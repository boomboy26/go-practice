package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)


type deck []string

func (d deck) print(){
	for i , card := range d {
		fmt.Println(i, card)
	}
}

func newDeck()(deck){
	d := deck{}
	symbol := []string{"Clubs","Diamonds" , "Hearts","Spades"}
	number := []string{"Ace","Two","Three","Four"}
	for _, sym := range symbol {
		for _, num := range number {
		 d = append(d,num + " of " + sym)
		}
	}
	return d
}

func deal(d deck,handSize int)(deck,deck){
	hand := d[:handSize]
	deckSize := d[handSize:]
	return hand , deckSize
}

func (d deck)toString () string{
	
	return strings.Join([]string(d),",")
}

func (d deck)saveToFile(filename string)error{
	return ioutil.WriteFile(filename, []byte(d.toString()),0666)
}

func newDeckFromFile (filename string) deck {
	    bs ,err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println("Error:",err) 
			os.Exit(1)
			
		}
		s := strings.Split(string(bs),",")
	return deck(s)
}