package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
)

var waiterTable = 1

//checks if each element is present within the array
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	tableFill := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	usedTables := []int{}

	csvFile, _ := os.Open("Dining.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, error := reader.Read()

		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		var table = rand.Intn(33)
		//add one to eliminate 0
		table++

		for {
			//find a new table assignment number
			if contains(usedTables, table) {
				table = rand.Intn(33)
				table++
			} else {
				break
			}
		}

		// Tables < 32 are real tables, table 32 is Waiters, table 33 is Kitchen Crew
		if table < 32 {
			fill table up to nine spots, stop fillinf after that
			if tableFill[table] < 9 {
				tableFill[table]++
			} else {

				//add tabkes to usedTables if greater than 32
				usedTables = append(usedTables, table)
			}
		}

		//anyone who is assigned 32 is a waiter
		if table == 32 {
			fmt.Println(line[0], line[1], "Waiter", waiterTable)
			waiterTable++
		//33 is not a seated table, this is kitchen crew
		} else if table == 33 {
			fmt.Println(line[0], line[1], "Kitchen Crew")
		} else {
		//if not 32 or 33, print just table number along with student name
			fmt.Println(line[0], line[1], table)
		}

	}
}
