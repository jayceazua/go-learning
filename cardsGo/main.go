package main

import "fmt"

func main() {
  // var card string = "Ace of Spades"
  // only useful for assigning a new variable; use it once when initializing the variable
  cards := deck{"Ace of Spades", newCard()} // declaring a slice []type{}
  cards = append(cards, "Six of Spades") // it returns a new slice not add anything to it

  for i, card := range cards {
    fmt.Println(i, card)
  }
}

// array (fixed length of size) / slice(more features, grow or shrink)

func newCard() string { // return types
  return "Five of Diamonds"
}
