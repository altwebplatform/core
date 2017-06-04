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
	"log"
	"github.com/altwebplatform/core/storage"
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

	rr = request(t, "PUT", "/api/v1/service", MustMarshall(storage.Service{Name: "inserted"}), nil)
	EnsureSuccess(t, rr)

	rr = request(t, "GET", "/api/v1/service/list", nil, nil)
	EnsureSuccess(t, rr)
	var result map[string]interface{}
	fmt.Println(rr.Body.String())
	err := json.Unmarshal(rr.Body.Bytes(), &result)
	ensure.Nil(t, err)
}

func MustMarshall(obj interface{}) []byte {
	b, err := json.Marshal(obj)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func MustUnmarshall(obj interface{}, data []byte) interface{} {
	err := json.Unmarshal(data, obj)
	if err != nil {
		log.Fatal(err)
	}
	return obj
}
