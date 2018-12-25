package configure

import (
	"encoding/json"
	"github.com/levigross/grequests"
	"log"
	"os/user"
	"strings"
	"testing"
)

func TestReadFile(t *testing.T) {
	usr, _ := user.Current()
	file := usr.HomeDir + "/.secretHarvest"

	data, err := readFile(file)
	if err != nil {
		t.Fail()
	}
	log.Printf("data: %v\n", data)
}

func createFile() string {
	usr, _ := user.Current()
	file := usr.HomeDir + "/.secretHarvest"
	secStruct := SecretStruct{}

	secStruct.Id = "01223"
	secStruct.Secret = "password"
	secStruct.Url = "http://httpbin.org/post"

	odata, err := json.Marshal(secStruct)

	n, err := writeFile(string(odata),
		file)
	if err != nil {
		log.Fatalf("error: %v, %v\n", n, err)
	}
	return file
}

func TestGetSecret(t *testing.T) {
	file := createFile()
	secStruct, err := GetSecret(file)
	if err != nil {
		t.FailNow()
	}
	if secStruct.Url != "http://httpbin.org/post" {
		t.FailNow()
	}
	if secStruct.Secret != "password" {
		t.FailNow()
	}

}

func TestCodeToPassStruct(t *testing.T) {
	c := CodeToPassStruct{}
	c.Secret = "secret"
	c.Id = "id"
	c.Code = "code"
	c.GrantType = "grant type"

	result, err := c.Marshel()
	if err != nil {
		t.FailNow()
	}
	expectedResult := `{"code":"code","client_id":"id","client_secret":"secret","grant_type":"grant type"}`
	if string(result) != expectedResult {
		t.FailNow()
	}

}

func TestInvalidGetAccessToken(t *testing.T) {
	c := CodeToPassStruct{}
	c.Secret = "t-xqP"
	c.Id = "sqXYl52"
	c.Code = "1453"
	c.GrantType = "authorization_code"

	result, err := c.Marshel()
	if err != nil {
		t.FailNow()
	}

	ro := grequests.RequestOptions{}
	ro.JSON = result
	url := "https://id.getharvest.com/api/v2/oauth2/token"
	r, err := grequests.Post(url, &ro)

	if strings.Contains(r.String(), "error") {
		type E struct {
			Error     string `json:"error"`
			ErrorDesc string `json:"error_description"`
		}
		e := E{}
		r.JSON(&e)
		if e.Error != "" {
			log.Printf("We have a problem: %v", e.ErrorDesc)
		}
	}

}

/*

{"access_token":"14IlpeGg","refresh_token":"14jzGD_rwi",
"token_type":"bearer","expires_in":1209599}

*/

func TestWriteResponseDataToFile(t *testing.T) {

	access := "1453046"
	ro := grequests.RequestOptions{}
	headers := map[string]string{}
	headers["Content-Type"] = "application/json"
	ro.Headers = headers

	s := ResultData{Access: access,
		Refresh: "14jzGD_rwi", Type: "bearer",
		Expires: 1209599}
	ro.JSON = s
	url := "http://httpbin.org/post"

	result, _ := grequests.Post(url, &ro)
	url = "http://httpbin.org/get"

	//url = "https://id.getharvest.com/api/v2/accounts"
	WriteResponseDataToFile(result, url)
}
