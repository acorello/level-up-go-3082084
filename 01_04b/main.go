package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

type coin uint16

const (
	c100p coin = 100
	c50p  coin = 50
	c20p  coin = 20
	c10p  coin = 10
	c05p  coin = 5
	c02p  coin = 2
	c01p  coin = 1
)

const coinCount = 7

var coinSizesInDescendingOrd = [coinCount]coin{
	c100p,
	c50p,
	c20p,
	c10p,
	c05p,
	c02p,
	c01p,
}

var coinNames [coinCount]string

func init() {
	for i, c := range coinSizesInDescendingOrd {
		switch c {
		case 100:
			coinNames[i] = "1 pound"
		case 50, 20, 10, 5, 2:
			coinNames[i] = fmt.Sprintf("%d pence", c)
		case 1:
			coinNames[i] = "1 penny"
		default:
			log.Panicf("missing case for %d", c)
		}
	}
}

func (I coin) Name() string {
	for i, c := range coinSizesInDescendingOrd {
		if c == I {
			return coinNames[i]
		}
	}
	log.Panicf("missing name for %d", I)
	return ""
}

// calculateChange returns the coins required to calculate the
func calculateChange(amount uint64) change {
	change := NewChange()
	var quantity uint64
	for _, nextCoin := range coinSizesInDescendingOrd {
		quantity, amount = nextCoin.DivMod(amount)
		change.Add(nextCoin, quantity)
		if amount == 0 {
			break
		}
	}
	return change
}

type change map[coin]uint64

func NewChange() change {
	return make(change)
}

func (I coin) DivMod(amount uint64) (q uint64, r uint64) {
	i := uint64(I)
	return amount / i, amount % i
}

func (I change) Add(c coin, q uint64) {
	I[c] += q
}

// printCoins prints all the coins in the slice to the terminal.
func (I change) printCoins() {
	if len(I) == 0 {
		fmt.Println("No change found.")
		return
	}
	fmt.Println("Change has been calculated.")
	for _, nextCoin := range coinSizesInDescendingOrd {
		count := I[nextCoin]
		if count > 0 {
			fmt.Printf("%d x %s \n", count, nextCoin.Name())
		}
	}
}

var amountRegex = regexp.MustCompile(`^\s*(\d+)(?:\.(\d{1,2}))?\s*$`)

func main() {
	amountFlag := flag.String("amount", "", "The amount you want to make change for")
	flag.Parse()
	parts := amountRegex.FindStringSubmatch(*amountFlag)
	if parts == nil {
		log.Fatal("Invalid format:", *amountFlag)
	}
	amountUnit, _ := strconv.ParseUint(parts[1], 10, 64)
	var amountDec uint64
	if len(parts) > 2 {
		amountDec, _ = strconv.ParseUint(parts[2], 10, 64)
	}
	amount := amountUnit*100 + amountDec
	change := calculateChange(amount)
	change.printCoins()
}
