package client

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func doRequest(client *http.Client, req *http.Request) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		log.WithError(err).Error("Request error:")
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			log.Errorf("error closing body: %v", err)
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("status: %d, Response body: %s", resp.StatusCode, body)
	}
	return body, nil
}
