package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/atotto/skyway-webrtc-gateway-go-client/client"
	"github.com/atotto/skyway-webrtc-gateway-go-client/client/data"
	"github.com/atotto/skyway-webrtc-gateway-go-client/client/media"
	"github.com/atotto/skyway-webrtc-gateway-go-client/client/peers"
	"github.com/atotto/skyway-webrtc-gateway-go-client/models"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

var apikey = os.Getenv("SKYWAY_API_KEY")

var peerID = flag.String("peer", "caller", "peer id")
var targetID = flag.String("target", "callee", "target id")
var useData = flag.Bool("data", false, "data")
var useMedia = flag.Bool("media", true, "media")
var domain = flag.String("domain", "localhost", "domain name")
var address = flag.String("addr", "localhost:8000", "gateway address")

var (
	DATA_RECV_ADDR  = "127.0.0.1"
	DATA_RECV_PORT  = uint16(10001)
	RECV_ADDR       = "127.0.0.1"
	VIDEO_RECV_PORT = uint16(20002)
	AUDIO_RECV_PORT = uint16(20003)
)

func main() {
	flag.Parse()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		select {
		case <-sig:
			log.Print("stopping..")
			cancel()
			return
		}
	}()

	c := client.NewHTTPClientWithConfig(strfmt.Default, &client.TransportConfig{
		Host:     *address,
		BasePath: "/",
		Schemes:  []string{"http"},
	})

	// create_peer
	peer, err := c.Peers.Peer(&peers.PeerParams{
		Context: ctx,
		Body: &models.PeerOptions{
			Domain: domain,
			Key:    &apikey,
			PeerID: *peerID,
			Turn:   true,
		}})
	if err != nil {
		log.Printf("failed to connect peer: %s", err)
		return
	}
	peerToken := peer.Payload.Params.Token
	peerID = peer.Payload.Params.PeerID
	fmt.Printf("peer_id: %s peer_token: %s\n", *peerID, *peerToken)

	defer c.Peers.PeerDestroy(&peers.PeerDestroyParams{
		Context: context.Background(),
		PeerID:  *peerID,
		Token:   *peerToken,
	})

	if *useMedia {
		vMedia, err := c.Media.Media(&media.MediaParams{
			Context: ctx,
			Body: &models.MediaOptions{
				IsVideo: swag.Bool(true),
			},
		})
		if err != nil {
			log.Printf("failed to create video media: %s", err)
			return
		}
		log.Printf("media video: port %d", *vMedia.Payload.Port)

		aMedia, err := c.Media.Media(&media.MediaParams{
			Context: ctx,
			Body: &models.MediaOptions{
				IsVideo: swag.Bool(false),
			},
		})
		if err != nil {
			log.Printf("failed to create audio media: %s", err)
			return
		}
		log.Printf("media audio: port %d", *aMedia.Payload.Port)

		// call
		constraints := &models.PeerCallConstraints{}
		if err := constraints.UnmarshalBinary([]byte(defaultConstraints)); err != nil {
			log.Printf("failed to create constraints: %s", err)
			return
		}
		constraints.VideoParams.MediaID = vMedia.Payload.MediaID
		constraints.AudioParams.MediaID = aMedia.Payload.MediaID

		conn, err := c.Media.MediaConnectionCreate(&media.MediaConnectionCreateParams{
			Context: ctx,
			Body: &models.PeerCallOptions{
				Constraints: constraints,
				PeerID:      peerID,
				Token:       peerToken,
				TargetID:    targetID,
				RedirectParams: &models.MediaRedirectOptions{
					Video: &models.MediaRedirectOptionsVideo{
						IPV4: RECV_ADDR,
						Port: VIDEO_RECV_PORT,
					},
					Audio: &models.MediaRedirectOptionsAudio{
						IPV4: RECV_ADDR,
						Port: AUDIO_RECV_PORT,
					},
				},
			},
		})
		if err != nil {
			log.Printf("failed to create media connection: %s", err)
			return
		}
		mediaConnID := conn.Payload.Params.MediaConnectionID

		defer c.Media.MediaConnectionClose(&media.MediaConnectionCloseParams{
			Context:           context.Background(),
			MediaConnectionID: mediaConnID,
		})

		// wait_stream
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}
			ev, err := c.Media.MediaConnectionEvent(&media.MediaConnectionEventParams{
				Context:           ctx,
				MediaConnectionID: mediaConnID,
			})
			if err != nil {
				log.Printf("failed to : %s", err)
				time.Sleep(time.Second)
				continue
			}
			if *ev.Payload.Event == "STREAM" {
				break
			}
		}

		log.Println("called")
	}

	if *useData {
		// create_data
		d, err := c.Data.Data(&data.DataParams{
			Body:    struct{}{},
			Context: ctx,
		})
		if err != nil {
			log.Printf("failed to data: %s", err)
			return
		}

		dataID := d.Payload.DataID
		log.Printf("data channel: port %d", *d.Payload.Port)

		// connect
		dc, err := c.Data.DataConnectionsCreate(&data.DataConnectionsCreateParams{
			Context: ctx,
			Body: &models.PeerConnectOptions{
				PeerID:   peerID,
				Token:    peerToken,
				TargetID: targetID,
				Options: &models.PeerConnectOptionsOptions{
					Serialization: "BINARY",
				},
				Params: &models.PeerConnectOptionsParams{
					DataID: *dataID,
				},
				RedirectParams: &models.DataConnectionRedirectOptions{
					IPV4: DATA_RECV_ADDR,
					Port: &DATA_RECV_PORT,
				},
			},
		})
		if err != nil {
			log.Printf("failed to : %s", err)
			return
		}

		dataConnID := dc.Payload.Params.DataConnectionID

		// wait_open
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}
			ev, err := c.Data.DataConnectionEvents(&data.DataConnectionEventsParams{
				Context:          ctx,
				DataConnectionID: *dataConnID,
			})
			if err != nil {
				log.Printf("failed to : %s", err)
				return
			}
			if *ev.Payload.Event == "OPEN" {
				break
			}
		}

		log.Println("opened data channel")

		defer c.Data.DataConnectionClose(&data.DataConnectionCloseParams{
			Context:          context.Background(),
			DataConnectionID: *dataConnID,
		})
	}

	<-ctx.Done()
}

var defaultConstraints = `
{
    "video": true,
    "videoReceiveEnabled": true,
    "video_params": {
        "band_width": 1500,
        "codec": "H264",
        "media_id": "video_id",
        "payload_type": 100
    },
    "audio": true,
    "audioReceiveEnabled": true,
    "audio_params": {
        "band_width": 1500,
        "codec": "opus",
        "media_id": "audio_id",
        "payload_type": 111
    }
}`
