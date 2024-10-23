package internal

import (
	"fmt"
	"io"
	"net/http"
)

func GetUrl(urlString string) ([]byte, int, error) {
	if data, ok := URLCache.Get(urlString); ok {
		return data, http.StatusInternalServerError, nil
	}

	res, err := http.Get(urlString)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if res.StatusCode > 299 && res.StatusCode != http.StatusNotFound {
		return nil, res.StatusCode, fmt.Errorf("response error: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	defer res.Body.Close()

	URLCache.Add(urlString, body)

	return body, res.StatusCode, nil
}
