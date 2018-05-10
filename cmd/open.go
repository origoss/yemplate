package cmd

import (
	"io"
	"net/http"
	"net/url"
	"os"
)

func openFileOrWeb(u string) (io.ReadCloser, error) {
	_, err := url.Parse(u)
	if err != nil {
		return os.Open(u)
	}
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}
