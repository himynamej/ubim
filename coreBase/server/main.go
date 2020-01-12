package main

import (
	Configurations "coreBase/Configurations"
	functions "coreBase/Func"
	DataTypes "coreBase/Type"
	pb "coreBase/proto"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) BroadCastRequest(ctx context.Context, chunk *pb.Auth) (*pb.Response, error) {

	// No feature was found, return an unnamed feature
	var req DataTypes.Request
	req.Token = chunk.Token

	functions.RequestHandler(req)
	return &pb.Response{}, nil
}
func (s *server) NormRequest(ctx context.Context, chunk *pb.NoAuth) (*pb.Response, error) {
	//messages := make(chan []byte)
	//payload := make(chan DataTypes.Response)
	//errors := make(chan error)
	// No feature was found, return an unnamed feature
	fmt.Println("f:", time.Now())
	var req DataTypes.Request

	req.Payload = chunk.Payload
	req.Group = chunk.Group
	req.Key = chunk.Key
	/*
		go func() {
			res := functions.RequestHandler(req)
			b, err := json.Marshal(res.Payload)
			if err != nil {
				fmt.Println(err)
				errors <- err
			}
			messages <- b
			payload <- res
			errors <- nil

		}()

		msg := <-messages
		pay := <-payload
		err := <-errors
		if err != nil {
			return &pb.Response{
				Message: err.Error(),
				Result:  "error",
				Payload: nil,
			}, nil
		}
	*/
	res := functions.RequestHandler(req)
	b, err := json.Marshal(res.Payload)
	if err != nil {
		fmt.Println(err)
		return &pb.Response{
			Message: err.Error(),
			Result:  "error",
			Payload: nil,
		}, nil
	}
	fmt.Println("l:", time.Now())
	fmt.Print(res.Message,"   result:",res.Result,"  payload:",b)
	return &pb.Response{
		Message: res.Message,
		Result:  res.Result,
		Payload: b,
	}, nil
	/*
		return &pb.Response{
			Message: pay.Message,
			Result:  pay.Result,
			Payload: msg,
		}, nil
	*/
}

func main() {

	config := flag.String("config", "please enter config path!!!!!", "url of configs")

	flag.Parse()

	value := *config
	fmt.Println("Value", value)

	err := Configurations.Configs.Init(value)
	if err != nil {
		log.Fatalf("failed to config: %v", err)
	}
	//connect to mongo

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
