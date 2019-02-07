package swaggerui

import (
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
)

func doHttp(method, path, body string) (string, *http.Response) {
	handler := Handler("/swagger", "MOOMOO")
	server := httptest.NewServer(handler)

	var bodyReader io.Reader
	if len(body) > 0 {
		bodyReader = strings.NewReader(body)
	}

	req, _ := http.NewRequest(method, server.URL + path, bodyReader)
	cli := &http.Client{}

	resp, _ := cli.Do(req)
	respbody, _ := ioutil.ReadAll(resp.Body)

	return string(respbody), resp
}

func TestHandler(t *testing.T) {
	body, resp := doHttp("GET", "/swagger/index.html", "")
	assert.Equal(t, 200, resp.StatusCode)
	assert.Regexp(t, regexp.MustCompile("^<!-- HTML"), body)

	body, resp = doHttp("GET", "/swagger", "")
	assert.Equal(t, 200, resp.StatusCode)
	assert.Regexp(t, regexp.MustCompile("^<!-- HTML"), body)

	body, resp = doHttp("GET", "/swagger", "")
	assert.Equal(t, 200, resp.StatusCode)
	assert.Regexp(t, regexp.MustCompile("^<!-- HTML"), body)
	assert.Regexp(t, regexp.MustCompile("MOOMOO"), body)

	body, resp = doHttp("GET", "/swagger/swagger-ui.css", "")
	assert.Equal(t, 200, resp.StatusCode)
	assert.Regexp(t, regexp.MustCompile("^\\.swagger-ui{"), body)
	assert.Regexp(t, regexp.MustCompile("^text/css"), resp.Header.Get("content-type"))
}
