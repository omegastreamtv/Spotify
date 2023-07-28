package spotify

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
)

func testClient(code int, body []byte) (*Client, *httptest.Server) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		w.Write(body)
	}))

	mockServerURL, _ := url.Parse(ts.URL)
	client, _ := NewClient(WithClientID("test"),
		WithBaseURL(mockServerURL.String()),
		WithRedirectURI("http://localhost:8080"),
	)

	return client, ts
}

func testClientFile(code int, filename string) (*Client, *httptest.Server) {
	f, err := os.Open(fmt.Sprintf("test_json/%s", filename))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	bytes, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	return testClient(code, bytes)
}
