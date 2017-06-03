package web

import (
	"net/http/httptest"
	"fmt"
	"github.com/facebookgo/ensure"
	"testing"
	"io"
	"bytes"
	"strconv"
	"net/http"
	"encoding/json"
)

func EnsureSuccess(t *testing.T, rr *httptest.ResponseRecorder) *httptest.ResponseRecorder {
	if rr.Code != 200 {
		fmt.Println("ERROR ", rr.Code, ": ", rr.Body.String())
	}
	ensure.DeepEqual(t, rr.Code, 200)
	return rr
}

func request(t *testing.T, method string, url string, body []byte, headers map[string]string) *httptest.ResponseRecorder {
	var bReader io.Reader
	if body != nil {
		bReader = bytes.NewReader(body)
	}
	req, err := http.NewRequest(method, url, bReader)

	if err != nil {
		if t != nil {
			t.Fatal(err)
		} else {
			panic(err)
		}
	}
	if body != nil {
		req.Header.Add("Content-Length", strconv.Itoa(len(body)))
	}
	for header, value := range headers {
		req.Header.Add(header, value)
	}
	rr := httptest.NewRecorder()
	CreateRouter().ServeHTTP(rr, req)
	return rr
}

func TestServicesAPI(t *testing.T) {
	var rr *httptest.ResponseRecorder
	uri := "/api/v1/service/list"

	rr = request(t, "GET", uri, nil, nil)
	EnsureSuccess(t, rr)
	var result map[string]interface{}
	fmt.Println(rr.Body.String())
	err := json.Unmarshal(rr.Body.Bytes(), &result)
	ensure.Nil(t, err)
}