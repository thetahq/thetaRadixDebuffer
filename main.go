package main

import "C"
import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/thetahq/thetaRadixDebuffer/thetaradix"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

//go:generate protoc --go_out=plugins=grpc:thetaradix thetaradix.proto
func main() {
	log.Println("Starting Theta Radix® Debufffer")
	log.Println("Purpose of this software is connecting gRPC backend with REST based one")
	startGrpcServer()
}

func startGrpcServer() {
	lis, err := net.Listen("tcp", "localhost:8557") // TODO: Use config file
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	thetaradix.RegisterThetaRadixServer(grpcServer, DebufferServer{cfg: loadConfig()})

	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}

type DebufferServer struct {
	cfg Config
}

func (d DebufferServer) Login(ctx context.Context, request *thetaradix.LoginRequest) (*thetaradix.LoginReply, error) {
	log.Println("Login request")

	req, err := http.Post(d.cfg.ServerAddress, "application/json", nil)
	authHeader, err := generateAuthHeader(request.Mail, request.Password)
	if err != nil {
		log.Println("Failed to generate auth header", err) // TODO: Deduplicate code
		return nil, err
	}
	req.Header.Set("Authorization", authHeader)

	responseData, err := readJSONFromBody(req.Body)
	if err != nil {
		logError(err)
		return &thetaradix.LoginReply{}, errors.Wrap(err, "failed to read json from body")
	}

	isSuccess, err := successToBool(responseData["status"].(string))
	if err != nil {
		log.Println("Failed to parse status", err)
		return nil, err
	}

	jwt := ""
	msg := responseData["message"].(string)
	if isSuccess {
		jwt = responseData["message"].(string)
		msg = ""
	}

	return &thetaradix.LoginReply{
		Success: isSuccess,
		Jwt:     jwt,
		Message: msg,
	}, nil
}

func (d DebufferServer) Register(ctx context.Context, request *thetaradix.RegisterRequest) (*thetaradix.RegisterReply, error) {
	log.Println("Register request")

	bodyData := make(map[string]interface{})
	bodyData["username"] = request.Nickname
	bodyData["terms"] = true
	jsonBody, _ := json.Marshal(bodyData)

	req, err := http.Post(d.cfg.ServerAddress, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Println("Failed to send post request", err)
		return nil, err
	}

	authorizationString, err := generateAuthHeader(request.Mail, request.Password)
	if err != nil {
		log.Println("Failed to generate auth header", err)
		return nil, err
	}
	req.Header.Set("Authorization", authorizationString)

	responseData, err := readJSONFromBody(req.Body)
	if err != nil {
		logError(err)
		return &thetaradix.RegisterReply{}, errors.Wrap(err, "failed to read json from body")
	}

	isSuccess, err := successToBool(responseData["status"].(string))
	if err != nil {
		log.Println("Failed to parse status", err)
		return nil, err
	}

	return &thetaradix.RegisterReply{
		Success:              isSuccess,
		PleaseVerifyPassword: true,
		ErrorMessage:         responseData["message"].(string),
	}, nil
}
