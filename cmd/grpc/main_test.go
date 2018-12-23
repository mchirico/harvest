package main

import (
	"bytes"
	"github.com/mchirico/harvest/pkg/rpkg"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a rpkg.App

func TestMain(m *testing.M) {

	a = rpkg.App{}
	a.Initilize()
	code := m.Run()

	os.Exit(code)
}

func TestEmptyProducts(t *testing.T) {

	body := []byte(`{
   "method": "JSONServer.GiveBookDetail",
   "params": [{
   "Id": "1234"
   }],
   "id": "1"
}`)

	req, _ := http.NewRequest("POST", "/rpc", bytes.NewBuffer(body))

	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	expectedResult := `{"result":{"Id":"1234","Name":"In the sunburned country","Author":"Bill Bryson"},"error":null,"id":"1"}
`

	if body := response.Body.String(); body != expectedResult {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
