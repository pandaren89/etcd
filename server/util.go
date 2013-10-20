package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/coreos/etcd/log"
)

func decodeJsonRequest(req *http.Request, data interface{}) error {
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		log.Warnf("Malformed json request: %v", err)
		return fmt.Errorf("Malformed json request: %v", err)
	}
	return nil
}

func redirect(hostname string, w http.ResponseWriter, req *http.Request) {
	originalURL, _ := url.Parse(req.URL.String())
	redirectURL, _ := url.Parse(hostname)

	// we need the original path and raw query
	redirectURL.Path = originalURL.Path
	redirectURL.RawQuery = originalURL.RawQuery

	log.Debugf("Redirect to %s", redirectURL.String())
	http.Redirect(w, req, redirectURL.String(), http.StatusTemporaryRedirect)
}
