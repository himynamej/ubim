package main

import (
	"fmt"
	"flag"
	"log"
	"core_api/Configurations"
	"core_api/Server"
)


func main() {

	config := flag.String("config", "please enter config path!!!!!", "url of configs")

	flag.Parse()

	value := *config
	fmt.Println("Value:", value)
	
	
	err := Configurations.Configs.Init(value)
	if err != nil {
		log.Fatalf("failed to config: %v", err)
	}

	fmt.Println("Starting API Core Service")
	Server.Service.Serve()
}
