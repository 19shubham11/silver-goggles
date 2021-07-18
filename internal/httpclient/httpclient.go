package httpclient

import "net/http"

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var Client HTTPClient

func init() {
	Client = &http.Client{}
}

func Get(url string, headers http.Header, queryParams map[string]string) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	request.Header = headers

	// add query params
	q := request.URL.Query()
	for key, element := range queryParams {
		q.Add(key, element)
	}
	request.URL.RawQuery = q.Encode()

	return Client.Do(request)
}
