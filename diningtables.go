
//Worked with Ethan and Dylan on second and third seating
package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Person struct {
	Firstname string
	Lastname  string
	Table1    string
	Table2    string
	Table3    string
}

//var  declared globally to be used in func findTable()
var table = 1

func findTable() {
	//find new table assignment 
	table = rand.Intn(33)
	//take out table 0
	table++

}

func contains(s []int, e int) bool {
	//Checks to see if a table is already filled before assigning to a name
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {

	var assignment = ""
	//tablevar keeps track of how much each person changes tables between dinners
	var tablevar = 1
	//strings display table assignments at the end.
	var table1disp = ""
	var table2disp = ""
	var table3disp = ""

	//randomizes each time
	rand.Seed(time.Now().UnixNano())

	//tableFill keeps track of how full each table is
	tableFill := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	//usedTables keeps track of tables that are full
	usedTables := []int{}

	csvFile, _ := os.Open("Dinner.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		//picks a random number to assign the current person to a table
		findTable()

		for {
			//continues to find a new number until it finds a table or assignment that isn't filled up yet.
			if contains(usedTables, table) {
				findTable()
			} else {
				break
			}
		}
		//sets person's order of seating in the table. First person at tables gets tableVar value of 1, second person tableVar=2 etc...
		tablevar = tableFill[table] + 1
		//loopvar is set to manually run the for loop three times.
		var loopvar = 1

		for {
			if table < 32 {
				//numbers below 32 are regular tables.
				if loopvar == 1 {
					//continues to fill table until there are 8 people.
					if tableFill[table] < 8 {
						tableFill[table]++
					} else {
						tableFill[table]++
						//adds table to usedTables
						usedTables = append(usedTables, table)
					}
				}
				//converts to a string for assignment
				assignment = strconv.Itoa(table)
			} else if table == 32 {
				//32 will be kitchen crew
				if loopvar == 1 {
					if tableFill[table] < 6 {
						tableFill[table]++
					} else {
						tableFill[table]++
						//adds to usedTables after 6 people are added
						usedTables = append(usedTables, table)
					}
				}
				//set assignment for Kitchen Crew
				assignment = "Kitchen Crew"
			} else if table == 33 {
				//33 will be for waiters
				if loopvar == 1 {
					if tableFill[table] < 30 {
						tableFill[table]++
					} else {
						tableFill[table]++
						//add to usedTables once full
						usedTables = append(usedTables, table)
					}
				}
				assignment = "Waiter"
			}
			//1 loop run through = 1 night.
			if loopvar == 1 {
				table1disp = assignment
			} else if loopvar == 2 {
				table2disp = assignment
			} else if loopvar == 3 {
				table3disp = assignment
			}

			//move to the next table assignment and repeat loop.
			table = table + tablevar
			//avoids index out of range error.
			for {
				if table > 33 {
					table = table - 33
				} else {
					break
				}
			}

			//if statement continues the loop for three seatings
			if loopvar < 3 {
				loopvar++
			} else if loopvar == 3 {
				break
			}

		}
		//Once all three nights are assigned, add specific person to struct Person
		people := Person{
			Firstname: line[1],
			Lastname:  line[0],
			Table1:    table1disp,
			Table2:    table2disp,
			Table3:    table3disp,
		}
		fmt.Println(people)

	}
}
