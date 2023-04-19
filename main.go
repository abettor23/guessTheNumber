package main

import (
	"fmt"
)

func main() {
	fmt.Println("Это прозвучит странно, но загадайте число от 1 до 5, а я угадаю его!")
	var number int
	fmt.Scan(&number)

	if number > 5 {
		fmt.Println("Написали чушь! Давайте заново...")
	} else if number < 1 {
		fmt.Println("Написали чушь! Давайте заново...")
	} else {
		fmt.Println("Ваше число меньше 5? (отвечайте так - y/n)")
		var answer string
		fmt.Scan(&answer)

		if answer == "n" {
			fmt.Println("Ваше число 5!")
		} else if answer == "y" {
			fmt.Println("Ваше число меньше 4?")
			fmt.Scan(&answer)

			if answer == "n" {
				fmt.Println("Ваше число 4!")
			} else if answer == "y" {
				fmt.Println("Ваше число меньше 3?")
				fmt.Scan(&answer)

				if answer == "n" {
					fmt.Println("Ваше число 3!")
				} else if answer == "y" {
					fmt.Println("Ваше число меньше 2?")
					fmt.Scan(&answer)

					if answer == "n" {
						fmt.Println("Ваше число 2!")
					} else if answer == "y" {
						fmt.Println("Ваше число 1!")
					} else {
						fmt.Println("Несёте чушь! Давайте заново...")
					}
				} else {
					fmt.Println("Несёте чушь! Давайте заново...")
				}
			} else {
				fmt.Println("Несёте чушь! Давайте заново...")
			}
		} else {
			fmt.Println("Несёте чушь! Давайте заново...")
		}
	}
}
