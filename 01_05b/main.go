package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"sort"
)

const path = "items.json"

// SaleItem represents the item part of the big sale.
type SaleItem struct {
	Name          string  `json:"name"`
	OriginalPrice float64 `json:"originalPrice"`
	ReducedPrice  float64 `json:"reducedPrice"`
}

type Item struct {
	SaleItem
	SalePercentage float64
}

// Here is a list:
//
//  1. Item 1.
//
//     Subitem 1.
//
//     Subitem 2.
//
//  2. Item 2.
//
//  3. Item 3.
func NewItem(i SaleItem) Item {
	pcd := ((i.OriginalPrice / i.ReducedPrice) - 1) * 100
	return Item{i, pcd}
}

// matchSales adds the sales procentage of the item and sorts the array accordingly.

// I'm assuming the match can have at lest the two following objectives
//
// N. OBJECTIVE ->
//
//	   SORTING-STRATEGY
//
//	1. maximize the number of items ->
//	   by ReducedPrice ascending
//
//	2. present the most discounted first ->
//	   by SalePercentage descending
//
// either way, I'm assuming I should not mutate the input slice
func matchSales(budget float64, items []Item) []Item {
	if budget <= 0 {
		return nil
	}
	var res []Item
	for _, item := range items {
		if item.ReducedPrice <= budget {
			res = append(res, item)
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].SalePercentage > res[j].SalePercentage
	})
	return res
}

func main() {
	budget := flag.Float64("budget", 0.0,
		"The max budget you want to shop with.")
	flag.Parse()
	items := importData()
	matchedItems := matchSales(*budget, items)
	printItems(matchedItems)
}

// printItems prints the items and their sales.
func printItems(items []Item) {
	log.Println("The BIG sale has started with our amazing offers!")
	if len(items) == 0 {
		log.Println("No items found. Try increasing your budget.")
	}
	for i, r := range items {
		log.Printf("[%d]:%s is %.0f%% OFF! Get it now for JUST %.2f!\n",
			i, r.Name, r.SalePercentage, r.ReducedPrice)
	}
}

// importData reads the raffle entries from file and
// creates the entries slice.
func importData() (res []Item) {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var data []SaleItem
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}
	res = make([]Item, len(data))
	for i := range data {
		res[i] = NewItem(data[i])
	}
	return
}
