// TASK: Calculate the biggest users market for a list of users records.
// Interpretations: high-level market size is a function of a set of users.
// Eg. M(fn U -> sum (map (fn u -> 1) U)) count of users in a market
// Eg. M(fn U -> sum (map (fn u -> u.Budget) U)) sum of the budget of each user

// which interpretation is correct is constrained by the input-dataset (its attributes)
// the structure is User{id int; name string; country string}
// so I map country to a market
package main

import (
	"encoding/json"
	"log"
	"os"
)

// User represents a user record.
type User struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

const path = "users.json"

// getBiggestMarket takes in the slice of users and
// returns the biggest market.
func getBiggestMarket(users []User) (market string, size int) {
	markets := map[string]int{}
	for _, u := range users {
		s := markets[u.Country]
		s += 1
		markets[u.Country] = s
		if s > size {
			market = u.Country
			size = s
		}
	}
	return
}

// I'd like to have tests setup…

func main() {
	users := importData()
	country, count := getBiggestMarket(users)
	log.Printf("The biggest user market is %s with %d users.\n",
		country, count)
}

// importData reads the raffle entries from file and
// creates the entries slice.
func importData() []User {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var data []User
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
