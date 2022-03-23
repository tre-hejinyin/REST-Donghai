package main

import (
	"log"
	"net/http"
	"strconv"

	"server/infra"
	"server/middleware/cache"
	"server/middleware/config"
)

func main() {
	log.Println("service start")
	// config init
	config.Init("./config.toml")

	defer func() {
		if err := recover(); err != nil {
			log.Printf("caught panic: %v ", err)
		}
	}()
	// connect redis
	err := cache.InitClient()
	if err != nil {
		log.Fatal(err)
	}
	// postgres
	conn, err := infra.Open()
	if err != nil {
		log.Print(err)
	}
	defer func() {
		db, _ := conn.DB()
		_ = db.Close()
	}()
	conn = conn.Debug()
	r, err := infra.SetupServer(conn)
	if err != nil {
		log.Print(err)
	}
	serveMux := http.NewServeMux()
	serveMux.Handle("/", r)

	err = http.ListenAndServe(":"+strconv.Itoa(config.ServicePort()), serveMux)
	if err != nil {
		log.Print("an error occurred when starting the server", err)
	}
	log.Println("service close")
}
