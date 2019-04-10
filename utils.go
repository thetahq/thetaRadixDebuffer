package main

import "errors"

func successToBool(data string) (success bool, err error) {
	switch data {
	case "ok":
		return true, nil
	case "error":
		return false, nil
	}
	return false, errors.New("unknown string value: " + data)
}
