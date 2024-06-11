package lib

import (
	"fmt"
	"os"
)

func PrintAscii(opt_justify, opt_output, opt_color bool, table_asset [][]string, sentence string) {
	var l1, l2, l3, l4, l5, l6, l7, l8 string

	//check if --align=justify cause we need to run the same code but modified, we also save run time
	if opt_justify && os.Args[1] == "--align=justify" {
		PrintJustify(l1, l2, l3, l4, l5, l6, l7, l8, os.Args[1], table_asset)
		return
	}
	//iteration above the user string
	for i := 0; i < len(sentence); i++ {
		//add on variables each line of characters
		if sentence[i] == '\\' && i < len(sentence) && len(sentence) > 1 {
			//check if justify, print, and make all variables empty for new line
			if opt_justify && (i < len(sentence)-1 && sentence[i+1] == 'n') {
				PrintJustify(l1, l2, l3, l4, l5, l6, l7, l8, os.Args[1], table_asset)
				l1, l2, l3, l4, l5, l6, l7, l8 = "", "", "", "", "", "", "", ""
				i++
			} else if i < len(sentence)-1 && sentence[i+1] == 'n' {
				Print(l1, l2, l3, l4, l5, l6, l7, l8)
				l1, l2, l3, l4, l5, l6, l7, l8 = "", "", "", "", "", "", "", ""
				i++
			} else {
				//add the following string of the template for each letter
				l1 += table_asset[sentence[i]-32][0]
				l2 += table_asset[sentence[i]-32][1]
				l3 += table_asset[sentence[i]-32][2]
				l4 += table_asset[sentence[i]-32][3]
				l5 += table_asset[sentence[i]-32][4]
				l6 += table_asset[sentence[i]-32][5]
				l7 += table_asset[sentence[i]-32][6]
				l8 += table_asset[sentence[i]-32][7]
			}
		} else if sentence[i] >= 32 && sentence[i] <= 126 {
			l1 += table_asset[sentence[i]-32][0]
			l2 += table_asset[sentence[i]-32][1]
			l3 += table_asset[sentence[i]-32][2]
			l4 += table_asset[sentence[i]-32][3]
			l5 += table_asset[sentence[i]-32][4]
			l6 += table_asset[sentence[i]-32][5]
			l7 += table_asset[sentence[i]-32][6]
			l8 += table_asset[sentence[i]-32][7]

			//if \n print all if they arn't empty, else \n

			//case for invalid inpput
		} else {
			fmt.Println("")
			fmt.Println("Error : invalid input")
			return
		}
	}
	if opt_justify {
		PrintJustify(l1, l2, l3, l4, l5, l6, l7, l8, os.Args[1], table_asset)
		return
	}
	fmt.Println(l1)
	fmt.Println(l2)
	fmt.Println(l3)
	fmt.Println(l4)
	fmt.Println(l5)
	fmt.Println(l6)
	fmt.Println(l7)
	fmt.Println(l8)
}
