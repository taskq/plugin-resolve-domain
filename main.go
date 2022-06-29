package main

import (
	// "bufio"
	// "bytes"
	"context"
	// "encoding/json"
	"fmt"
	// "io"
	"net"
	// "strings"
	"time"

	"github.com/rs/zerolog/log"
)

var PluginName string = "ResolveDomain"
var PluginDescription string = "DNS Resolve Domain plugin for TaskQ Subscriber"
var BuildVersion string = "0.0.0"

type PayloadStruct struct {
}

type OutputPayloadStruct struct {
}

func ExecCommand(payload []byte, configurationRaw interface{}) (result []byte, err error) {
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, network, "r-led-1.nameserver-provider.com:53")
			// return d.DialContext(ctx, network, "ns42.nameserver-provider.com:53")
		},
	}
	ip, err := r.LookupHost(context.Background(), "dwww.google.com")

	if err != nil {
		fmt.Printf("Couldn't do the resolve: %v\n", err)
	} else {
		fmt.Printf("Response: %+v\n", ip[0])

	}

	return []byte{}, nil
}

func main() {

	ResolveDomain()

	returned, err := ExecCommand([]byte(`{"domain": "godaddy.com", "nameserver": "ns1.godaddy.com"}`), nil)

	log.Info().
		Err(err).
		Str("returned", string(returned)).
		Msgf("Done")

}
