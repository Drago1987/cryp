package main

import (
	"log"

	"drago.nik/polygon/utilities"
)

func init() {
	// Enable line numbers in logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	log.Print("Starting rlpdecoder")
	//https://codechain-io.github.io/rlp-debugger/
	list1 := "ea8d4976616e20426a656c616a61638e4d616c697361205075736f6e6a618c536c61766b6f204a656e6963"
	utilities.RlpDecoder(list1)
	list2 := "e5922034342e38313538393735343033373334319132302e3435343733343334343535353435"
	utilities.RlpDecoder(list2)

	log.Print("--- exit")
}
