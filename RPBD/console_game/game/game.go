package game

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Horsen121/TBD/RPBD/console_game/wombat/wombat"
)

func StartGame() {
	w := wombat.New()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		var input string = ""
		for {
			fmt.Println("Select an action:")
			fmt.Print("1. Dig\n2. Eat\n3. Fight\n4. Sleep\n")
			scanner.Scan()
			input := scanner.Text()

			if input == "1" || input == "2" || input == "3" || input == "4" {
				break
			}
		}

		switch input {
		case "1":
			for {
				fmt.Print("1. Intense\n2. Lazy\n")
				scanner.Scan()
				input = scanner.Text()

				if input == "1" {
					w.Dig(1)
					break
				} else if input == "2" {
					w.Dig(2)
					break
				}
			}
		case "2":
			for {
				fmt.Print("1. Yellow\n2. Green\n")
				scanner.Scan()
				input = scanner.Text()

				if input == "1" {
					w.Eat(1)
					break
				} else if input == "2" {
					w.Eat(2)
					break
				}
			}
		case "3":
			for {
				fmt.Print("1. Weak(30)\n2. Middle(50)\n3. Strong(70)")
				scanner.Scan()
				input = scanner.Text()

				if input == "1" {
					fmt.Println(w.Fight(30))
					break
				} else if input == "2" {
					fmt.Println(w.Fight(50))
					break
				} else if input == "3" {
					fmt.Println(w.Fight(70))
					break
				}
			}
		case "4":
			w.Sleep()
			fmt.Println("Good night! :)")
		}
		stats := w.Stats()
		if stats == "You loose" || stats == "You win!!!" {
			fmt.Println(stats)
			os.Exit(1)
		}
		fmt.Println(stats + "\n")

		fmt.Println("Night is coming")
		w.Sleep()
		fmt.Println("New day")
	}
}
