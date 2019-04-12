package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"log"
)

func successToBool(data string) (success bool, err error) {
	switch data {
	case "ok":
		return true, nil
	case "error":
		return false, nil
	}
	return false, errors.New("unknown string value: " + data)
}

func generateAuthHeader(mail string, password string) (authHeader string, err error) {
	authorizationString := fmt.Sprintf("%s:%s:%s", mail, password, password)
	// Encode authorization string into base64
	base64output := bytes.Buffer{}
	base64Encoder := base64.NewEncoder(base64.StdEncoding, &base64output)
	_, err = base64Encoder.Write([]byte(authorizationString))
	err = base64Encoder.Close()
	if err != nil {
		log.Println("Failed to encode authorization header")
		return "", errors.New("failed to generate auth header: " + err.Error())
	}
	return base64output.String(), nil
}

func readJSONFromBody(body io.Reader) (jsonData map[string]interface{}, err error) {
	responseData := make(map[string]interface{})
	responseBytes, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response")
	}

	err = json.Unmarshal(responseBytes, responseData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarchal json")
	}

	return responseData, nil
}

func logError(err error) {
	fmt.Printf("%+s:%d\n", err, err)
}
