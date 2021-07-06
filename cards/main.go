package main

//import "fmt"
func main() {


//	cards := newDeck()
	// handCard, deckLeft := deal(cards, 5)
	// fmt.Println( handCard.toString())
	// fmt.Println (deckLeft.toString())
	//fmt.Println(cards.toString())
	x := newDeckFromFile("cardDeck")
	x.print()


}
