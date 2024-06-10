package lib

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func PrintJustify(l1, l2, l3, l4, l5, l6, l7, l8, full_opt string, table_asset [][]string) {
	option := full_opt[8:]
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdout
	out, _ := cmd.CombinedOutput()
	out_splited := strings.Fields(string(out))
	width, _ := strconv.Atoi(out_splited[1])
	space := width - (len(l1))
	switch option {
	case "right":
		if len(l1) != 0 {
			fmt.Printf("%*s%v\n", space, " ", l1)
			fmt.Printf("%*s%v\n", space, " ", l2)
			fmt.Printf("%*s%v\n", space, " ", l3)
			fmt.Printf("%*s%v\n", space, " ", l4)
			fmt.Printf("%*s%v\n", space, " ", l5)
			fmt.Printf("%*s%v\n", space, " ", l6)
			fmt.Printf("%*s%v\n", space, " ", l7)
			fmt.Printf("%*s%v\n", space, " ", l8)
		} else {
			fmt.Println("")
		}

	case "left":
		if len(l1) != 0 {
			fmt.Println(l1)
			fmt.Println(l2)
			fmt.Println(l3)
			fmt.Println(l4)
			fmt.Println(l5)
			fmt.Println(l6)
			fmt.Println(l7)
			fmt.Println(l8)
		} else {
			fmt.Println("")
		}
	case "center":
		if len(l1) != 0 {
			fmt.Printf("%*s%v%*s\n", space/2, " ", l1, space/2, " ")
			fmt.Printf("%*s%v%*s\n", space/2, " ", l2, space/2, " ")
			fmt.Printf("%*s%v%*s\n", space/2, " ", l3, space/2, " ")
			fmt.Printf("%*s%v%*s\n", space/2, " ", l4, space/2, " ")
			fmt.Printf("%*s%v%*s\n", space/2, " ", l5, space/2, " ")
			fmt.Printf("%*s%v%*s\n", space/2, " ", l6, space/2, " ")
			fmt.Printf("%*s%v%*s\n", space/2, " ", l7, space/2, " ")
			fmt.Printf("%*s%v%*s\n", space/2, " ", l8, space/2, " ")
		} else {
			fmt.Println("")
		}
	case "justify":
		sentence := os.Args[2]
		final_size := 0
		var t_txt []int
		//for var space
		for i := 0; i < len(sentence); i++ {
			//add on variables each line of characters
			if sentence[i] == '\\' && i < len(sentence) && len(sentence) > 1 {
				if i < len(sentence)-1 && sentence[i+1] == 'n' {
					if len(l1) > final_size {
						final_size = len(l1)
					}
					t_txt = append(t_txt, len(l1))
					l1, l2, l3, l4, l5, l6, l7, l8 = "", "", "", "", "", "", "", ""
					i++
				} else {
					l1 += table_asset[sentence[i]-32][0]

				}
			} else if sentence[i] >= 32 && sentence[i] <= 126 {
				l1 += table_asset[sentence[i]-32][0]
			} else {
				fmt.Println("")
				fmt.Println("Error : invalid input")
				return
			}
		}
		t_txt = append(t_txt, len(l1))
		if len(l1) > final_size {
			final_size = len(l1)
		}
		l1, l2, l3, l4, l5, l6, l7, l8 = "", "", "", "", "", "", "", ""
		space = width - final_size

		t := TablesSpaces(os.Args[2])
		cpt_spaces := 0
		for i := 0; i < len(sentence); i++ {
			//add on variables each line of characters
			if sentence[i] == '\\' && i < len(sentence) && len(sentence) > 1 {
				if i < len(sentence)-1 && sentence[i+1] == 'n' {
					Print(l1, l2, l3, l4, l5, l6, l7, l8)
					l1, l2, l3, l4, l5, l6, l7, l8 = "", "", "", "", "", "", "", ""
					cpt_spaces++
					space = width - t_txt[cpt_spaces]
					i++
				} else if sentence[i] == ' ' {
					l1 += table_asset[sentence[i]-32][0] + AddSpace(space/t[cpt_spaces])
					l2 += table_asset[sentence[i]-32][1] + AddSpace(space/t[cpt_spaces])
					l3 += table_asset[sentence[i]-32][2] + AddSpace(space/t[cpt_spaces])
					l4 += table_asset[sentence[i]-32][3] + AddSpace(space/t[cpt_spaces])
					l5 += table_asset[sentence[i]-32][4] + AddSpace(space/t[cpt_spaces])
					l6 += table_asset[sentence[i]-32][5] + AddSpace(space/t[cpt_spaces])
					l7 += table_asset[sentence[i]-32][6] + AddSpace(space/t[cpt_spaces])
					l8 += table_asset[sentence[i]-32][7] + AddSpace(space/t[cpt_spaces])
				} else {
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
				if sentence[i] == ' ' {
					l1 += table_asset[sentence[i]-32][0] + AddSpace(space/t[cpt_spaces])
					l2 += table_asset[sentence[i]-32][1] + AddSpace(space/t[cpt_spaces])
					l3 += table_asset[sentence[i]-32][2] + AddSpace(space/t[cpt_spaces])
					l4 += table_asset[sentence[i]-32][3] + AddSpace(space/t[cpt_spaces])
					l5 += table_asset[sentence[i]-32][4] + AddSpace(space/t[cpt_spaces])
					l6 += table_asset[sentence[i]-32][5] + AddSpace(space/t[cpt_spaces])
					l7 += table_asset[sentence[i]-32][6] + AddSpace(space/t[cpt_spaces])
					l8 += table_asset[sentence[i]-32][7] + AddSpace(space/t[cpt_spaces])
				} else {
					l1 += table_asset[sentence[i]-32][0]
					l2 += table_asset[sentence[i]-32][1]
					l3 += table_asset[sentence[i]-32][2]
					l4 += table_asset[sentence[i]-32][3]
					l5 += table_asset[sentence[i]-32][4]
					l6 += table_asset[sentence[i]-32][5]
					l7 += table_asset[sentence[i]-32][6]
					l8 += table_asset[sentence[i]-32][7]
				}
				//if \n print all if they arn't empty, else \n

				//case for invalid inpput
			} else {
				fmt.Println("")
				fmt.Println("Error : invalid input")
				return
			}
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

}

func AddSpace(n int) string {
	txt := ""
	for i := 0; i < n; i++ {
		txt += " "
	}
	return txt
}

func TablesSpaces(s string) []int {
	cpt := 0
	var t []int
	for i, ch := range s {
		if ch == ' ' {
			cpt++
		} else if i < len(s)-1 && ch == '\\' {
			if s[i+1] == 'n' {
				t = append(t, cpt)
				cpt = 0
				i++
			}
		}
	}
	t = append(t, cpt)
	return t
}
