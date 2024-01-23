package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"ticket_generator/global"
	"ticket_generator/utils"
)

func GetTicket(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	getConn, err := global.PostgresConn.Acquire(context.Background())
	if err != nil {
		log.Printf("Error acquiring connection: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer getConn.Release()

	rows, err := getConn.Query(context.Background(), utils.Select_ticket)
	if err != nil {
		log.Printf("Error querying data: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var currentSet interface{}
	var currentTicket interface{}
	var currentNumbers []interface{}

	for rows.Next() {
		var set interface{}
		var ticketNumber interface{}
		var numberArray interface{}

		if err := rows.Scan(&set, &ticketNumber, &numberArray); err != nil {
			log.Printf("Error scanning row: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Check if we have moved to a new set or ticket
		if set != currentSet || ticketNumber != currentTicket {
			// Print the previous ticket
			if currentSet != "" && currentTicket != 0 {
				printTicket(w, currentSet, currentTicket, currentNumbers)
			}

			// Update current set and ticket
			currentSet = set
			currentTicket = ticketNumber
			currentNumbers = append(currentNumbers, numberArray)
		} else {
			// Concatenate the numbers if we are still in the same ticket
			currentNumbers = append(currentNumbers, numberArray)
		}
	}

	// Print the last ticket
	if currentSet != "" && currentTicket != 0 {
		printTicket(w, currentSet, currentTicket, currentNumbers)
	}

	if err := rows.Err(); err != nil {
		fmt.Printf("Error iterating over rows: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func printTicket(w http.ResponseWriter, set interface{}, ticketNumber interface{}, numbers []interface{}) {
	fmt.Fprintf(w, "\n%s:\n", set)
	fmt.Fprintf(w, "Ticket %d:\n", ticketNumber)
	for _, row := range numbers {
		fmt.Fprintln(w, row)
	}
}
