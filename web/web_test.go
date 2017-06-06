package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/altwebplatform/core/storage"
	"github.com/facebookgo/ensure"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
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
	var services = make(map[string][]storage.Service)

	// first clean up using the API
	rr = request(t, "GET", "/api/v1/services", nil, nil)
	EnsureSuccess(t, rr)
	MustUnmarshall(&services, rr.Body.Bytes())

	for _, service := range services["services"] {
		rr = request(t, "DELETE", "/api/v1/services/"+strconv.FormatUint(service.ID, 10), nil, nil)
		EnsureSuccess(t, rr)
	}

	rr = request(t, "POST", "/api/v1/services", MustMarshall(storage.Service{Name: "inserted"}), nil)
	EnsureSuccess(t, rr)

	rr = request(t, "GET", "/api/v1/services", nil, nil)
	EnsureSuccess(t, rr)
	MustUnmarshall(&services, rr.Body.Bytes())
	ensure.True(t, len(services["services"]) == 1)
	ensure.DeepEqual(t, services["services"][0].Name, "inserted")

	for _, service := range services["services"] {
		rr = request(t, "DELETE", "/api/v1/services/"+strconv.FormatUint(service.ID, 10), nil, nil)
		EnsureSuccess(t, rr)
	}

	rr = request(t, "GET", "/api/v1/services", nil, nil)
	EnsureSuccess(t, rr)
	MustUnmarshall(&services, rr.Body.Bytes())
	ensure.True(t, len(services["services"]) == 0)
}

func MustMarshall(obj interface{}) []byte {
	b, err := json.Marshal(obj)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func MustUnmarshall(obj interface{}, data []byte) {
	err := json.Unmarshal(data, obj)
	if err != nil {
		log.Fatal(err)
	}
}
