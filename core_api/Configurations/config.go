package Configurations

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

type Configurations struct {
	Settings
	Operation
}

type (
	Settings struct {
		ServerPort         string   `json:"server_port"`
		AllowedOrigins     []string `json:"allowed_origins"`
		AllowedHeaders     []string `json:"allowed_headers"`
		AllowedMethods     []string `json:"allowed_methods"`
		AllowCredentials   bool     `json:"allow_credentials"`
		Debug              bool     `json:"debug"`
		OptionsPassthrough bool     `json:"options_passthrough"`
		RunTime            string   `json:"run_time"`
	}
)

type Operation map[string]gojsonschema.JSONLoader
var FileValid Operation
type Routes struct {
	RoutePath     string
	RouteFunction func(w http.ResponseWriter, req *http.Request)
	RouteMethods  string
}

var Configs Configurations

func (c *Configurations) Init(cf string) error {
	// reading configurations settings from file
	err := c.Settings.Init(cf)
	if err != nil {
		return err
	}
	ValidFile()
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

	return nil
}

func ValidFile() {

	var files []string

	root := Configs.RunTime
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".json") {
			files = append(files, path)
			return nil
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	m := make(map[string]gojsonschema.JSONLoader)
	for _, file := range files {
		
		s := strings.Split(file, "/")
		fmt.Println(s[len(s)-3] + s[len(s)-2] + s[len(s)-1])
		m[s[len(s)-3]+s[len(s)-2]+s[len(s)-1]] = gojsonschema.NewReferenceLoader("file://" + file)
        
		//fmt.Println(FileValid[s[len(s)-2]][s[len(s)-2]])

	}
	FileValid=m
}
