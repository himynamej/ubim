package Configurations

import (
	"coreBase/mongo"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Configurations struct {
	Settings
}

type (
	Settings struct {
		MongoURL     string `json:"mongo_url"`
		DatabaseName string `json:"database_name"`
		ServerPort   string `json:"server_port"`
	}
)

var Mongodb mongo.MongoConnection
var Configs Configurations

func (c *Configurations) Init(cf string) error {
	// reading configurations settings from file
	err := c.Settings.Init(cf)
	if err != nil {
		return err
	}
	return nil
}

func (s *Settings) Init(cf string) error {
	data, err := ioutil.ReadFile(cf)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	err = json.Unmarshal([]byte(string(data)), &s)
	if err != nil {
		return err
	}
	
	err=MongoConnect()
	if err != nil {
		return err
	}
	fmt.Println("url:",s.MongoURL)
	return nil
}
func MongoConnect() error {
	Mongodb.MongoUrl = Configs.MongoURL
	err := Mongodb.ConnectMongo()
	if err != nil {
		fmt.Printf("failed to connect mongo: %v", err)
		return err
	}
	return nil
}
