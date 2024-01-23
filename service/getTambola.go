package service

import (
	"context"
	"fmt"
	"math/rand"
	"ticket_generator/utils"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func generateTambolaTicket() [][]int {
	ticket := make([][]int, 3)
	for i := range ticket {
		ticket[i] = make([]int, 9)
	}

	rand.Seed(time.Now().UnixNano())

	// Filling the ticket with numbers
	for col := 0; col < 9; col++ {
		colNumbers := rand.Perm(10)
		for i := 0; i < 3; i++ {
			ticket[i][col] = col*10 + colNumbers[i] + 1
		}
	}

	// Shuffle the numbers in each column
	for col := 0; col < 9; col++ {
		rand.Shuffle(3, func(i, j int) {
			ticket[i][col], ticket[j][col] = ticket[j][col], ticket[i][col]
		})
	}

	return ticket
}

// Function to generate N sets of Tambola tickets
func generateTambolaSets(numSets int) map[string][][][]int {
	tambolaSets := make(map[string][][][]int)

	for i := 11; i <= 10+numSets; i++ {
		setID := fmt.Sprintf("%d", i)
		tambolaSets[setID] = make([][][]int, 6)
		usedTickets := make(map[string]bool)

		for j := 0; j < 6; j++ {
			// Generate a unique ticket
			for {
				newTicket := generateTambolaTicket()
				ticketStr := fmt.Sprint(newTicket)
				if !usedTickets[ticketStr] {
					usedTickets[ticketStr] = true
					tambolaSets[setID][j] = newTicket
					break
				}
			}
		}
	}

	return tambolaSets
}

func insertTambolaSetsIntoDB(tambolaSets map[string][][][]int, insertConn *pgxpool.Conn) error {
	// Iterate over the sets and tickets, then insert into the database
	for setID, set := range tambolaSets {
		for i, ticket := range set {
			// Convert 2D array to 1D array for database insertion
			var flattenedTicket []int
			for _, row := range ticket {
				flattenedTicket = append(flattenedTicket, row...)
			}

			// Insert data into the database
			insCmd, err := insertConn.Exec(context.Background(), utils.Insert_data, "Tambola Set "+setID, i+1, flattenedTicket, time.Now().Format("2006-01-02 15:04:05"))
			if err != nil || insCmd.RowsAffected() == 0 {
				fmt.Println("error in inserting the data", err)
				return err
			}
		}

	}

	return nil
}
