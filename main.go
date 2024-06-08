package main

import (
	"fmt"
	"lib"
)

func main() {
	//check arguments
	wrong_input, opt_template, opt_justify, opt_output, opt_color, sentence, asset_file_name := lib.CheckOptions()
	if wrong_input {
		fmt.Println("Error : Wrong input, check the README file to know how to use this project ")
		return
	}

	fmt.Printf("template : %v\n", opt_template)
	fmt.Printf("output : %v\n", opt_output)
	fmt.Printf("justify : %v\n", opt_justify)
	fmt.Printf("color : %v\n", opt_color)

	asset, thinkertoy := lib.GetFileContent("assets/" + asset_file_name)
	table_asset := lib.CreateTable(asset, thinkertoy)
	lib.PrintAscii(opt_justify, opt_output, opt_color, table_asset, sentence)
}
