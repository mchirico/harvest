package configure

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/levigross/grequests"
	"log"
	"os/user"
)

type SecretStruct struct {
	Id      string `json:"clientID"`
	Secret  string `json:"clientSecret"`
	Url     string `json:"url"`
	Code    string `json:"code"`
	Seconds string `json:"seconds"`
	Expire  string `json:"expire"`
}

type CodeToPassStruct struct {
	Code      string `json:"code"`
	Id        string `json:"client_id"`
	Secret    string `json:"client_secret"`
	GrantType string `json:"grant_type"`
}

type ResultData struct {
	Access  string `json:"access_token"`
	Refresh string `json:"refresh_token"`
	Type    string `json:"token_type"`
	Expires int    `json:"expires_in"`
}

func (c *CodeToPassStruct) Marshel() ([]byte, error) {
	return json.Marshal(*c)
}

func WriteResponseDataToFile(response *grequests.Response, url string) (int, error) {

	log.Printf("WriteResponseDataToFile")

	var f interface{}
	err := response.JSON(&f)
	if err != nil {
		return -1, errors.New("Error response.JSON(&f) ")
	}

	m := f.(map[string]interface{})
	if m == nil {
		return -1, errors.New("Error m := f.(map[string]interface{}) ")
	}

	r := m["json"].(map[string]interface{})

	if val, ok := r["access_token"]; ok {
		ro := grequests.RequestOptions{}
		headers := map[string]string{}
		headers["Content-Type"] = "application/json"
		headers["Authorization"] = fmt.Sprintf("Bearer %v", val)
		headers["User-Agent"] = "AiPiggybot (mchirico@gmail.com)"
		ro.Headers = headers
		result, err := grequests.Get(url, &ro)

		if err != nil {
			log.Printf("ERROR:  %v\n", err)
			return -1, err
		}

		type Account struct {
			Id      int    `json:"id"`
			Name    string `json:"name"`
			Product string `json:"product"`
			Google  bool   `json:"google_sign_in_required"`
		}

		type AccountStruct struct {
			User     interface{} `json:"user"`
			Accounts []Account   `json:"accounts"`
		}

		v := AccountStruct{}
		err = result.JSON(&v)

		if err != nil {
			return -1, err
		}

		if len(v.Accounts) == 0 {
			return -1, errors.New("Didn't get ID")
		}

		log.Println(v.Accounts[0].Id)
		usr, _ := user.Current()
		file := usr.HomeDir + fmt.Sprintf("/.harvestData_%v", v.Accounts[0].Id)

		type Store struct {
			Access string `json:"access"`
			Id     int    `json:"id"`
		}

		store := Store{}
		store.Access = val.(string)
		store.Id = v.Accounts[0].Id
		data, err := json.Marshal(store)
		n, err := writeFile(string(data), file)
		return n, err
	}
	log.Printf("We did not write")
	return -1, errors.New("Can't work with result")

}
