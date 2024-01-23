package service

import (
	"context"
	"fmt"
	"net/http"
	"ticket_generator/global"
)

// Function to generate a single Tambola ticket

func SetTicket(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	tambolaSets := generateTambolaSets(6)
	insertConn, err := global.PostgresConn.Acquire(context.Background())
	if err != nil {
		fmt.Println("Error acquiring connection ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	inserterr := insertTambolaSetsIntoDB(tambolaSets, insertConn)
	insertConn.Release()
	if inserterr != nil {
		http.Error(w, inserterr.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Tambola sets inserted successfully!!!")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Tambola sets inserted successfully"  ))

}
