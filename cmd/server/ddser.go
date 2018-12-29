package main

import (
	"github.com/freeddser/ddser/config"
	"github.com/freeddser/ddser/controllers"
	R "github.com/freeddser/ddser/router"

	"flag"
	"fmt"
	"github.com/freeddser/ddser/models"
	"github.com/freeddser/ddser/util"
	"log"
	"net/http"
	"os"
)

func main() {
	configFile := flag.String("c", "", "Configuration File")
	flag.Parse()

	if *configFile == "" {
		fmt.Println("\n\nUse -h to get more information on command line options\n")
		fmt.Println("You must specify a configuration file")
		os.Exit(1)
	}

	err := config.Initialize(*configFile)
	if err != nil {
		fmt.Printf("Error reading configuration: %s\n", err.Error())
		os.Exit(1)
	}
	//init template
	util.RecursionDir(config.MustGetString("server.TEMPLATE_DIR"))

	//init session
	controllers.InitSession()
	models.InitDB()

	//start http
	log.Printf("**************\tListenAndServe: %s\t**************\n", config.MustGetString("server.Run_Port"))
	log.Printf("**************\tServer Run Mode: %s\t**************\n", config.MustGetString("server.Run_Mode"))
	err = http.ListenAndServe(config.MustGetString("server.Run_Port"), R.NewRouter())
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
