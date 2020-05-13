package main

import (
	"api"
	"common"
	"log"
	"net/http"
)

type config struct {
	Port string `yaml:"port"`
}

func main() {
	cfg := config{}
	common.ReadConfig(&cfg)

	log.Printf("http://localhost:" + cfg.Port)
	router := api.InitRouter()
	panic(http.ListenAndServe(":"+cfg.Port, router))
}
