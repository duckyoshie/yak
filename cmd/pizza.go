package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// intializing function
func init() {
	// add to root
	rootCmd.AddCommand(pizzaCmd)
}

// pizza
type PizzaProduct struct {
	Name  string
	Price float32
}

// list to hold types of pizzas
var pizzaList []PizzaProduct

// types of pizza defined
var pepperoni = &PizzaProduct{
	Name:  "Pepperoni",
	Price: 10.95,
}

var meatFeast = &PizzaProduct{
	Name:  "Meat Feast",
	Price: 15.95,
}

var margherita = &PizzaProduct{
	Name:  "Margherita",
	Price: 9.12,
}

// function to hold pizza shop interaction
func PizzaShop(userItemsList []PizzaProduct) {
	// loop infintely
	for {
		// loop through list to display options
		for i := 0; i < len(pizzaList); i++ {
			fmt.Println(i, ": ", pizzaList[i].Name, " ", pizzaList[i].Price)
		}

		// get user's input
		var choice string
		fmt.Scanln(&choice)

		if choice == "0" {
			userItemsList = append(userItemsList, *pepperoni)
		} else if choice == "1" {
			userItemsList = append(userItemsList, *meatFeast)
		} else if choice == "2" {
			userItemsList = append(userItemsList, *margherita)
		} else {
			fmt.Println("Not specified")
			return
		}

		// checks if further input
		fmt.Println("Any more?")
		var wantMore string
		fmt.Scanln(&wantMore)

		if wantMore == "yes" {
			// loop
			PizzaShop(userItemsList)
		} else {
			// add payment
			var payment float32

			for i := 0; i < len(userItemsList); i++ {
				payment += userItemsList[i].Price
			}

			// return payment
			fmt.Println("You need to pay: £", payment)
			os.Exit(2)
		}
	}
}

// pizza command - this will mimic a pizza shop
var pizzaCmd = &cobra.Command{
	Use:        "pizza",
	Aliases:    []string{},
	SuggestFor: []string{},
	Short:      "A pizza shop",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Pizza Shop")
		var userItemsList []PizzaProduct // user's list of items
		pizzaList = append(pizzaList, *pepperoni)
		pizzaList = append(pizzaList, *meatFeast)
		pizzaList = append(pizzaList, *margherita)

		PizzaShop(userItemsList)
	},
}
