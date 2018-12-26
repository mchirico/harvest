package configure

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/levigross/grequests"
	"log"
	"os/user"
	"strings"
)

type TODOStruct struct {
	Id     string `json:"client_id"`
	Secret string `json:"client_secret"`
}

type TODOtoken struct {
	Token string `json:"access_token"`
	Type  string `json:"token_type"`
}

type LiveStruct struct {
	Refresh string `json:"refresh_token"`
	Id      string `json:"client_id"`
	Secret  string `json:"client_secret"`
	Type    string `json:"grant_type"`
}

type LastAccess struct {
	Access  string `json:"access_token"`
	Expires string `json:"expires_in"`
	Start   string `json:time_write`
	Refresh LiveStruct
}

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

type RefreshStruct struct {
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

func GetLastAccess2() LastAccess {

	usr, _ := user.Current()
	file := usr.HomeDir + "/.harvestLive"
	datain, _ := readFile(file)
	l := LastAccess{}
	json.Unmarshal([]byte(datain), &l)
	log.Println(l.Start)
	log.Println(l.Refresh.Id)
	return l
}

func GetTODO() TODOStruct {
	usr, _ := user.Current()
	file := usr.HomeDir + "/.todoapi"
	jsonData, _ := readFile(file)
	t := TODOStruct{}
	json.Unmarshal([]byte(jsonData), &t)

	return t
}

func GetTODOtoken() TODOtoken {
	usr, _ := user.Current()
	file := usr.HomeDir + "/.todotoken"
	jsonData, _ := readFile(file)
	t := TODOtoken{}
	json.Unmarshal([]byte(jsonData), &t)

	return t

}

func Refresh() string {

	l := GetLastAccess2()
	r := LiveStruct{}
	Access := l.Access
	r = l.Refresh
	GrantType := "refresh_token"

	ro := grequests.RequestOptions{}
	headers := map[string]string{}
	headers["Content-Type"] = "application/json"
	headers["grant_type"] = GrantType
	headers["Authorization"] = fmt.Sprintf("Bearer %v", Access)

	ro.Headers = headers
	data, _ := json.Marshal(r)
	ro.JSON = data
	url := "https://id.getharvest.com/api/v2/oauth2/token"
	result, err := grequests.Post(url, &ro)

	log.Println(result.String())
	log.Println(err)

	return result.String()
}

func UnmarshelRefreshToken(str string) RefreshStruct {
	res := RefreshStruct{}
	json.Unmarshal([]byte(str), &res)
	return res
}

/*
curl "https://id.getharvest.com/api/v2/accounts" \
  -H "Authorization: Bearer 14530" \
  -H "User-Agent: AiPiggybot (mchirico@gmail.com)"
*/
func GetID(access_token string) (string, error) {
	ro := grequests.RequestOptions{}
	headers := map[string]string{}
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = fmt.Sprintf("Bearer %s", string(access_token))
	headers["User-Agent"] = "AiPiggybot (mchirico@gmail.com)"
	ro.Headers = headers
	url := "https://id.getharvest.com/api/v2/accounts"
	result, err := grequests.Get(url, &ro)
	if err != nil {
		return "", err
	}

	if strings.Contains(result.String(), "error") {
		return result.String(), errors.New("Result contains error")
	}
	return result.String(), err

}
