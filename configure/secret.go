package configure

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

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

// GetSecret returns SecretStruct
func GetSecret(file string) (SecretStruct, error) {
	secStr := SecretStruct{}
	data, err := readFile(file)
	if err != nil {
		return secStr, err
	}
	err = json.Unmarshal([]byte(data), &secStr)
	if err != nil {
		return secStr, err
	}
	return secStr, nil

}
