package cmd

import (
	"io"
	"net/http"
	"net/url"
	"os"
)

func openFileOrWeb(u string) (io.ReadCloser, error) {
	parsedURL, err := url.Parse(u)
	if err != nil {
		return nil, err
	}
	if parsedURL.Scheme == "http" ||
		parsedURL.Scheme == "https" {
		resp, err := http.Get(u)
		if err != nil {
			return nil, err
		}
		return resp.Body, nil
	}
	return os.Open(u)
}
