package main

import (
	"github.com/indigo-dc/chronosctl/cmd"
	"github.com/indigo-dc/chronosctl/config"
	"log"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Println("Could not load configuration file: ", err)
	}

	cmd.InitRootCmd()

	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
