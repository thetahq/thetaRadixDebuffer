package main

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
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
