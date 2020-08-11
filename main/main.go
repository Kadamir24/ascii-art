package main

import (
	"flag"
	"os"

	utils "./utils"
)

func main() {
	os.Args = os.Args[1:]
	font, words := utils.GetFontAndWords(&os.Args)
	align := flag.String("align", "", "")
	output := flag.String("output", "", "")
	color := flag.String("color", "", "")
	letters := flag.String("letters", "", "")
	flag.Parse()
	// fmt.Println("align:", *align)
	// fmt.Println("output:", *output)
	// fmt.Println("color:", *color)
	// fmt.Println("font:", font)
	// fmt.Println("words:", words)
	//utils.Fsss(font, words)
	if *output != "" {
		utils.Output(font, words, *output)
	} else if *align != "" {
		utils.Justify(font, words, *align)
	} else if *color != "" {
		utils.Color(font, words, *color, *letters)
	} else {
		utils.Fsss(font, words)
	}
}
