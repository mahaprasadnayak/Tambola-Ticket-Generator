package main

import (
	"fmt"
	"log"
	"net/http"
	"ticket_generator/global"
	"ticket_generator/service"
)

func main() {
	fmt.Println("Setting up the server...")
	//defer global.PostgresConn.Close()
	http.HandleFunc("/", service.GetServerStatus) //health check of the server with default endpoint
	serverErr := http.ListenAndServe(global.Port, nil)
	if serverErr != nil {
		fmt.Println("Error starting the server ", serverErr, " from the port::", global.Port)
		log.Fatal("Error starting the server ", serverErr, " from the port::", global.Port)
	}
}
