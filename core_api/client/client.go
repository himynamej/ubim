package client

import(
	 pb "coreBase/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"fmt"
)
func client(){
	conn, err := grpc.Dial("192.168.43.169:50051", grpc.WithInsecure())
	if err != nil {
	log.Fatalf("did not connect: %v", err)
	}
//"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmUiOiIyMDE5LTEyLTEyVDE1OjU5OjM0LjU1NDU0OTAyMiswMzozMCIsInRpbWUiOiIyMDE5LTEyLTEyVDE1OjQ0OjM0LjU1NDU1MTY4NSswMzozMCIsInVzZXJuYW1lIjoiamF2YWRhaCJ9.dS7cbNKYmIr2jtcJs1U5xPQfeAt1r3VYGjl5AST0PUo"
	defer conn.Close()
	c := pb.NewServiceClient(conn)
	//byt := []byte(`{"username":"javadah","email":"javadah1376@gmail.com","mobile":"09186543321","password":"1234"}`)
	//req := &pb.NoAuth{Group:"register",Key:"registerUser",Payload:byt} 
	// initialize a pb.NoAuth
	//byt := []byte(`{"verify":"5da865b910f0682c422d9007","code":"7507"}`)
	//req := &pb.NoAuth{Group:"register",Key:"verifiction",Payload:byt} 
	byt := []byte(`{"username":"javadah","password":"1234"}`)
	req := &pb.NoAuth{Group:"register",Key:"login",Payload:byt} 
	//byt := []byte(`{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmUiOiIyMDE5LTEyLTEyVDE1OjU5OjM0LjU1NDU0OTAyMiswMzozMCIsInRpbWUiOiIyMDE5LTEyLTEyVDE1OjQ0OjM0LjU1NDU1MTY4NSswMzozMCIsInVzZXJuYW1lIjoiamF2YWRhaCJ9.dS7cbNKYmIr2jtcJs1U5xPQfeAt1r3VYGjl5AST0PUo"}`)
	//req := &pb.NoAuth{Group:"user",Key:"checkuser",Payload:byt} 
	str, err := c.NormRequest(context.Background(), req)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(str)


}


