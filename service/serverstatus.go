package service

import (
	"fmt"
	"io"
	"net/http"
)

func GetServerStatus(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Server Status :: Active !!")
	fmt.Println("request coming from Host::", r.Host, "of the remoteAddress :: ", r.RemoteAddr)
}
