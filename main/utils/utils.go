package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strconv"

	// "fmt"
	"os"
	"strings"
)

const (
	blue   = "\x1b[38;2;0;0;255m%s\x1b[0m"
	red    = "\x1b[38;2;255;0;0m%s\x1b[0m"
	green  = "\x1b[38;2;0;255;0m%s\x1b[0m"
	yellow = "\x1b[38;2;255;255;0m%s\x1b[0m"
	pink   = "\x1b[38;2;255;0;255m%s\x1b[0m"
	white  = "\x1b[38;2;255;255;255m%s\x1b[0m"
	black  = "\x1b[38;2;0;0;0m%s\x1b[0m"
	orange = "\x1b[38;2;255;128;0m%s\x1b[0m"
)

var ASSETS_PATH = "./assets/"

var (
	fonts = []string{"standard", "shadow", "thinkertoy"}
)

func removeFont(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func removeWord(slice []string, s int) []string {
	return append(slice[:s])
}

func GetFontAndWords(args *[]string) (string, []string) {
	for i, arg := range *args {
		for _, v := range fonts {
			if v == arg {
				*args = removeFont(*args, i)
				return v, removeWord(*args, i)
			}
		}
	}
	return "standard", *args
}

func ReadLine(f *os.File, n int) string {
	f.Seek(0, 0) //Устанавливает файловый указатель в определенную позицию
	bf := bufio.NewReader(f)
	var line string

	for lnum := 0; lnum < n; lnum++ {
		line, _ = bf.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
	}
	return line
}

func Fsss(font string, words []string) {
	f, err := os.Open(ASSETS_PATH + font + ".txt")
	if err != nil {
		return
	}
	arg := words[0]
	words = strings.Split(arg, "\\n")
	for _, word := range words {
		for i := 0; i < 8; i++ {
			for _, letter := range word {
				fmt.Print(ReadLine(f, 2+((int(letter)-32)*9)+i))
			}
			fmt.Println()
		}
	}
}

func Output(font string, words []string, output string) {
	f, err := os.Open(ASSETS_PATH + font + ".txt")
	if err != nil {
		return
	}
	file, err := os.Create(output)
	if err != nil {
		return
	}
	arg := words[0]

	words = strings.Split(arg, "\\n")
	for _, word := range words {
		for i := 0; i < 8; i++ {
			for _, letter := range word {
				file.WriteString(ReadLine(f, 2+((int(letter)-32)*9)+i))
			}

			file.WriteString("\n")
		}
	}
	file.Sync()

}

func Justify(font string, words []string, align string) {
	f, err := os.Open(ASSETS_PATH + font + ".txt")
	if err != nil {
		return
	}
	file, err := os.Create("text.txt")
	if err != nil {
		return
	}

	arg := words[0]
	words = strings.Split(arg, "\\n")
	if align == "justify" {
		cmd := exec.Command("tput", "cols")
		var stdout bytes.Buffer
		cmd.Stdout = &stdout
		err = cmd.Run()
		if err != nil {
			fmt.Print(err)
		}
		cols := strings.TrimSuffix(stdout.String(), "\n")

		col, err := strconv.Atoi(cols)

		if err != nil {
			fmt.Print(err)
		}
		lenLine := 0
		SpaceCount := 0 //spaces
		var lenArr []int
		var SpaceArr []int
		for _, word := range words {
			lenLine = 0
			SpaceCount = 0
			for _, letter := range word {
				if letter != ' ' {

					line := len(ReadLine(f, 2+((int(letter)-32)*9)+0))

					lenLine = lenLine + line
				} else {
					SpaceCount++

				}
			}
			lenArr = append(lenArr, lenLine)
			SpaceArr = append(SpaceArr, SpaceCount)
			// fmt.Println("SpaceArr:", SpaceArr)
			// fmt.Println("lenArr:", lenArr)

			if col <= lenLine {
				fmt.Println("Too short terminal, that`s what she said")
				return
			}
		}

		counter := 0
		NewCounter := 0
		for j, word := range words {
			if SpaceArr[j] == 0 {
				//for _, word := range words {
				for i := 0; i < 8; i++ {
					for _, letter := range word {
						out := ReadLine(f, 2+((int(letter)-32)*9)+i)
						fmt.Print(out)
					}
					fmt.Println()
				}
				//}

			} else {
				spaces := (col - lenArr[j]) / SpaceArr[j]
				residual := (col - lenArr[j]) % SpaceArr[j]
				// fmt.Println("residual:", residual)
				// fmt.Println("Terminal:", col, ", Length of Spaces:", spaces, ", Length of letters:", lenArr[j])
				NewCounter = 0
				for i := 0; i < 8; i++ {
					NewCounter = 0
					for _, letter := range word {
						if letter != ' ' {
							row := ReadLine(f, 2+((int(letter)-32)*9)+i)
							fmt.Print(row)
						} else {
							// for {
							// 	fmt.Print(" ")
							// 	counter++
							// 	if counter%spaces == 0 {
							// 		break
							// 	}

							// }
							for counter != spaces {
								fmt.Print(" ")
								counter++
							}
							counter = 0
							if residual != 0 && NewCounter >= SpaceArr[j]-residual { //решение с остатком
								fmt.Print(" ")

							}
							NewCounter = NewCounter + 1

						}

					}
					fmt.Println()
				}
			}

		}
	}
	for _, word := range words {
		for i := 0; i < 8; i++ {
			for _, letter := range word {
				row := ReadLine(f, 2+((int(letter)-32)*9)+i)
				file.WriteString(row)
			}

			file.WriteString("\n")
		}
	}
	file.Sync()
	cmd := exec.Command("./pick.sh", align, "text.txt")
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err = cmd.Run()
	if err != nil {
		fmt.Print("err", err)
	}
	fmt.Print(stdout.String())

	cmd = exec.Command("rm", "text.txt")

	err = cmd.Run()
	if err != nil {
		fmt.Print("err", err)
	}

}

func Color(font string, words []string, color string, letters string) {
	f, err := os.Open(ASSETS_PATH + font + ".txt")
	if err != nil {
		return
	}
	arg := words[0]
	words = strings.Split(arg, "\\n")
	var slice []int
	if letters != "" {
		arrColor := letters
		index, err1 := strconv.Atoi(arrColor)
		var first, second int
		if err1 != nil {
			array := strings.Split(arrColor, ":")
			if len(array) != 2 {
				fmt.Println("Please, choose slice which you want to be colored, like so: --letter=<number>:<number>1")
				return
			} else {
				if array[0] == "" {
					first = 0
				} else {
					first, err = strconv.Atoi(array[0])
					if err != nil {
						fmt.Println("Please, choose slice which you want to be colored, like so: --letter=<number>:<number>2")
						return
					}
				}

				//first, err2 := strconv.Atoi(array[0])

				if array[1] == "" {
					second = len(arg)

				} else {
					second, err = strconv.Atoi(array[1])
					if err != nil {
						fmt.Println("Please, choose slice which you want to be colored, like so: --letter=<number>:<number>3")
						return
					}
					if first > second {
						fmt.Println("Please, don`t do stupid things, begging you!")
						return
					} else if first == second {
						slice = append(slice, first)
					}

				}
				for i := first; i < second; i++ {

					slice = append(slice, i)
				}
			}
		} else {
			slice = append(slice, index)

		}

	}
	colors := strings.Split(color, ",")
	if len(colors) == 1 {
		color, boolean := isValidColor(colors[0])
		if !boolean {
			fmt.Println("ERROR: if you want to colorize output use: --color=<color>1")
			return
		}

		Print(f, words, color, "%s", slice)
	} else if len(colors) == 2 {
		color1, boolean := isValidColor(colors[0])
		color2, boolean2 := isValidColor(colors[1])
		if !boolean {
			fmt.Println("ERROR: if you want to use two colors use: --color=<color>,<secondColor>")
			return

		}
		if !boolean2 {
			fmt.Println("ERROR: if you want to use two colors use: --color=<color>,<secondColor>")
			return
		}
		Print(f, words, color1, color2, slice)
	} else {
		fmt.Println("you are absolutely wrong, you can choose only 2 colors, that`s our RULE")
	}
}

func Print(f *os.File, words []string, color1 string, color2 string, slice []int) {

	for _, word := range words {
		for i := 0; i < 8; i++ {
			for j, letter := range word {
				if len(slice) == 0 || isColor(j, slice) {

					row := ReadLine(f, 2+((int(letter)-32)*9)+i)
					fmt.Printf(color1, row)
				} else {

					row := ReadLine(f, 2+((int(letter)-32)*9)+i)
					fmt.Printf(color2, row)
				}

			}
			fmt.Println()
		}
	}

}

func isColor(n int, arr []int) bool {
	for _, v := range arr {
		if n == v {
			return true
		}
	}
	return false
}

func isValidColor(str string) (string, bool) {
	color := ""
	flag := false
	if str == "blue" {
		color = blue
		flag = true
	} else if str == "red" {
		color = red
		flag = true
	} else if str == "green" {
		color = green
		flag = true
	} else if str == "orange" {
		color = orange
		flag = true
	} else if str == "white" {
		color = white
		flag = true
	} else if str == "black" {
		color = black
		flag = true
	} else if str == "yellow" {
		color = yellow
		flag = true
	} else if str == "pink" {
		color = pink
		flag = true
	}
	return color, flag

}
