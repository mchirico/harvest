package configure

import (
	"encoding/json"
	"log"
	"os/user"
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
