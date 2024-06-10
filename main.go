package main

import (
	"fmt"
	"lib"
)

func main() {
	//check arguments
	wrong_input, _, opt_justify, opt_output, opt_color, sentence, asset_file_name := lib.CheckOptions()
	if wrong_input {
		fmt.Println("Error : Wrong input, check the README file to know how to use this project ")
		return
	}

	asset, thinkertoy := lib.GetFileContent("assets/" + asset_file_name)
	table_asset := lib.CreateTable(asset, thinkertoy)
	lib.PrintAscii(opt_justify, opt_output, opt_color, table_asset, sentence)
}
