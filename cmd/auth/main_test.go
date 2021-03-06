package main

import (
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	. "github.com/mchirico/harvest/configure"
	"github.com/mchirico/harvest/pkg"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"os/user"
	"testing"
)

var a pkg.App

func TestMain(m *testing.M) {
	a = pkg.App{}

	a.Initilize()
	code := m.Run()

	os.Exit(code)
}

func TestAuth(t *testing.T) {

	oSecretStruct := SecretStruct{}

	oSecretStruct.Id = "01223"
	oSecretStruct.Secret = "password"
	oSecretStruct.Url = "http://httpbin.org/post"

	a.InitSS(&oSecretStruct)

	req, _ := http.NewRequest("GET", "/auth?code=3.43", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "code" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

// Ref: https://play.golang.org/p/UGeNKd-cw34
func TestResponseCode(t *testing.T) {
	type SendData struct {
		Num  float32  `json:"num"`
		Strs []string `json:"strs"`
	}
	ro := grequests.RequestOptions{}
	headers := map[string]string{}
	headers["Content-Type"] = "application/json"
	ro.Headers = headers
	s := SendData{Num: 6.23, Strs: []string{"one", "two"}}
	ro.JSON = s
	url := "http://httpbin.org/post"
	result, _ := grequests.Post(url, &ro)

	log.Printf("\n\n result:%v\n", result)

	var f interface{}
	b := result.String()
	err := result.JSON(&f)
	if err != nil {
		t.Fail()
	}

	m := f.(map[string]interface{})
	log.Printf("json: %s\n", b)

	log.Printf("m: %v\n type: %T\n\n", m["json"], m["json"])

	for key, value := range m["json"].(map[string]interface{}) {
		fmt.Println("Key:", key, "Value:", value)
	}

	r := m["json"].(map[string]interface{})

	fmt.Println("num", r["num"].(float64))
	fmt.Println("strs[0]", r["strs"].([]interface{})[0])
	fmt.Println("strs[1]", r["strs"].([]interface{})[1])

	if r["strs"].([]interface{})[0] != "one" {
		t.Fail()
	}
	if r["strs"].([]interface{})[1] != "two" {
		t.Fail()
	}
}

func TestGettingSecret(t *testing.T) {

	usr, _ := user.Current()

	file := usr.HomeDir + "/.secretHarvest"

	oSecretStruct := SecretStruct{}

	oSecretStruct.Id = "01223"
	oSecretStruct.Secret = "password"
	oSecretStruct.Url = "http://httpbin.org/post"

	odata, err := json.Marshal(oSecretStruct)

	n, err := writeFile(string(odata),
		file)
	if err != nil {
		log.Printf("error: %v, %v\n", n, err)
		t.Fail()
	}

	data, err := readFile(file)
	if err != nil {
		t.Fail()
	}

	res := SecretStruct{}
	err = json.Unmarshal([]byte(odata), &res)
	if err != nil {
		t.Fail()
	}
	fmt.Println("res.Id: ", res.Id)
	fmt.Println("res.Secret: ", res.Secret)
	fmt.Println("res.Url: ", res.Url)
	fmt.Println("data: ", data)

}

func TestSecret(t *testing.T) {
	type SendData struct {
		Num  float32  `json:"num"`
		Strs []string `json:"strs"`
	}
	ro := grequests.RequestOptions{}
	headers := map[string]string{}
	headers["Content-Type"] = "application/json"
	ro.Headers = headers
	s := SendData{Num: 6.23, Strs: []string{"one", "two"}}
	ro.JSON = s
	url := "http://httpbin.org/post"
	result, _ := grequests.Post(url, &ro)

	log.Printf("\n\n result:%v\n", result)

	var f interface{}
	b := result.String()
	result.JSON(&f)

	m := f.(map[string]interface{})
	log.Printf("json: %s\n", b)

	log.Printf("m: %v\n type: %T\n\n", m["json"], m["json"])

	for key, value := range m["json"].(map[string]interface{}) {
		fmt.Println("Key:", key, "Value:", value)
	}

	r := m["json"].(map[string]interface{})

	fmt.Println("num", r["num"].(float64))
	fmt.Println("strs[0]", r["strs"].([]interface{})[0])
	fmt.Println("strs[1]", r["strs"].([]interface{})[1])

	if r["strs"].([]interface{})[0] != "one" {
		t.Fail()
	}
	if r["strs"].([]interface{})[1] != "two" {
		t.Fail()
	}

}

func readFile(file string) (string, error) {
	data, err := ioutil.ReadFile(file)
	return string(data), err
}

func writeFile(data string, file string) (int, error) {
	f, err := os.Create(file)
	defer f.Close()

	if err != nil {
		return -1, err
	}

	n, err := f.WriteString(data)

	return n, err
}

func TestEmptyProducts(t *testing.T) {

	req, _ := http.NewRequest("GET", "/products", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestRoot(t *testing.T) {

	req, _ := http.NewRequest("GET", "/", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body !=
		`[{"page":1,"fruits":["pear","orange"]},{"page":2,"fruits":["pear","orange"]}]` {
		t.Errorf("Expected an array. Got %s", body)
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
