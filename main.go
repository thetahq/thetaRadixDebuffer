package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/thetahq/thetaRadixDebuffer/thetaradix"
	"google.golang.org/grpc"
	"io/ioutil"
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

func (d DebufferServer) Register(ctx context.Context, request *thetaradix.RegisterRequest) (*thetaradix.RegisterReply, error) {
	log.Println("Register request")

	bodyData := make(map[string]interface{})
	bodyData["username"] = request.Nickname
	bodyData["terms"] = true
	jsonBody, _ := json.Marshal(bodyData)

	req, err := http.Post(d.cfg.ServerAddress, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Println("Failed to send post request")
		return nil, err
	}
	authorizationString := fmt.Sprintf("%s:%s:%s", request.Mail, request.Password, request.Password)
	// Encode authorization string into base64
	base64output := bytes.Buffer{}
	base64Encoder := base64.NewEncoder(base64.StdEncoding, &base64output)
	_, err = base64Encoder.Write([]byte(authorizationString))
	err = base64Encoder.Close()
	if err != nil {
		log.Println("Failed to encode authorization header")
		return nil, err
	}

	req.Header.Set("Authorization", base64output.String())

	responseData := make(map[string]interface{})
	responseBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("Failed to read response")
		return nil, err
	}

	err = json.Unmarshal(responseBytes, responseData)
	if err != nil {
		log.Println("Failed to unmarchal json", err)
		return nil, err
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
