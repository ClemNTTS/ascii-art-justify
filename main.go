package main

import (
	"fmt"
	"lib"
)

func main() {
	//check arguments
	wrong_input, _, opt_justify, opt_output, opt_color, sentence, asset_file_name := lib.CheckOptions()

	//Error management
	if wrong_input {
		fmt.Println("Error : Wrong input, check the README file to know how to use this project ")
		return
	}

	//Get the file content
	asset, thinkertoy := lib.GetFileContent("assets/" + asset_file_name)
	//create a slice with all the template
	table_asset := lib.CreateTable(asset, thinkertoy)
	//Print management
	lib.PrintAscii(opt_justify, opt_output, opt_color, table_asset, sentence)
}
