package portia

import (
	"fmt"
	"io"
	"net/http"
)

type WebPage struct {
	name    string
	URL     string
	Status  int
	Content string
}

func NewWebPage(name string, url string) (*WebPage, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making http request: %s", err)
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %s", err)
	}

	return &WebPage{
		name:    name,
		URL:     url,
		Status:  res.StatusCode,
		Content: string(resBody),
	}, nil
}
