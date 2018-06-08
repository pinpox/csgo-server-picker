package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// serverList is a CS:GO server list, as specified by github.com/SteamDatabase/SteamTracking.
type serverList struct {
	Revision int               `json:"revision"`
	Servers  map[string]server `json:"pops"`
}

// server is a CS:GO server.
type server struct {
	Description          string   `json:"desc"`
	Geo                  geo      `json:"geo"`
	Partners             int      `json:"partners"`
	RelayAddresses       []string `json:"relay_addresses"`
	ServiceAddressRanges []string `json:"service_address_ranges"`
}

// geo is a set of geographical coordinates.
type geo struct {
	Lat, Long float64
}

func (g *geo) UnmarshalJSON(data []byte) error {
	var tmp []float64
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	if len(tmp) != 2 {
		return fmt.Errorf("geo: unexpected JSON array of length %d, want 2", len(tmp))
	}
	g.Lat = tmp[1]
	g.Long = tmp[0]
	return nil
}

const defaultServerListURL = "https://raw.githubusercontent.com/SteamDatabase/SteamTracking/master/Random/NetworkDatagramConfig.json"

type updater struct {
	c             *http.Client
	serverListURL string
}

func (u *updater) fetchServerList(ctx context.Context) (*serverList, error) {
	req, err := http.NewRequest(http.MethodGet, u.serverListURL, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	resp, err := u.c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch server list: %v", err)
	}
	// We have a response. No matter what we do with it, we should drain the
	// body and close it, just in case we want to continue using the HTTP
	// client connection. readServerList uses a json.Decoder, which might not
	// read until EOF, which is what the copy to ioutil.Discard is for.
	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()
	// We may have gotten a nil error, but a non-200 status code (for example, we may
	// be rate, limited, or the server may be down). In that case, we need to process
	// the response body differently. Try to tell the user something meaningful.
	if resp.StatusCode != http.StatusOK {
		buf := new(bytes.Buffer)
		io.Copy(buf, resp.Body)
		return nil, fmt.Errorf("failed to fetch server list: %d (%s): %s", resp.StatusCode, resp.Status, buf.String())
	}

	list, err := readServerList(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to decode server list: %v", err)
	}
	return list, nil
}

func readServerList(r io.Reader) (*serverList, error) {
	dec := json.NewDecoder(r)
	list := new(serverList)
	if err := dec.Decode(list); err != nil {
		return nil, err
	}
	return list, nil
}
