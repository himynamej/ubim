package Server

import (
	"core_api/Configurations"
	"core_api/DataHandler"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var Service *Server

type Server struct {
	Router *mux.Router
	Routes *[]Configurations.Routes
}

var Routes = []Configurations.Routes{
	//------------------------------------------------------------------------------------
	{"/ubim/request/v1", DataHandler.Request, "POST"},
	
}

func (s *Server) Serve() {

	smux := mux.NewRouter().StrictSlash(true)
	for _, route := range Routes {
		smux.HandleFunc(route.RoutePath, route.RouteFunction)
	}

	c := cors.New(cors.Options{
		AllowedMethods:     Configurations.Configs.AllowedMethods,
		AllowedOrigins:     Configurations.Configs.AllowedOrigins[:],
		AllowCredentials:   Configurations.Configs.AllowCredentials,
		AllowedHeaders:     Configurations.Configs.AllowedHeaders,
		OptionsPassthrough: Configurations.Configs.OptionsPassthrough,
		Debug:              Configurations.Configs.Debug,
	})
	fmt.Println(Configurations.Configs.AllowedOrigins)
	handler := c.Handler(smux)

	err := http.ListenAndServe(":"+Configurations.Configs.ServerPort, handler)
	fmt.Println("", err)
}
