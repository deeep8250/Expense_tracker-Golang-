package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	var save name
	var currentTime = time.Now()
	formateDate := currentTime.Format("01/02/2006")
	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Println("\n\nCHOOSE YOUR DESIRE OPTION")
		fmt.Println("1. Add Expense")
		fmt.Println("2. View Expense")
		fmt.Println("3. Delete Expense")
		fmt.Println("4. Exit")
		fmt.Print("\n")
		fmt.Print("Enter your choice : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":

			var n interface{}
			for i := 0; ; i++ {

				fmt.Print("\n\t\t1. Add Your Expense\n")
				fmt.Print("\n\t\t2. back\n")
				fmt.Print("\n\t\tEnter your choice :")
				n = User_input()

				if n == 1 || n == 2 {
					break
				} else {
					fmt.Println("Enter valid input : ")
				}
			}

			switch n {
			case 1:
				fmt.Print("Enter Expense Nmae : ")
				expenseN := bufio.NewReader(os.Stdin)
				expenseN_input, _ := expenseN.ReadString('\n')
				expenseN_input = strings.TrimSpace(expenseN_input)

				fmt.Print("Enter Amount name : ")
				amountN := User_input()

				save = name{Expense_Name: append(save.Expense_Name, expenseN_input), Amount: append(save.Amount, float64(amountN)), Date: append(save.Date, formateDate)}

			case 2:
				break
			}

		case "2":

			fmt.Print("\n\t\tView your expense\n")

			if len(save.Expense_Name) == 0 {
				fmt.Println("You didn't add something Please add something")
			}

			for i := 0; ; i++ {

				if i >= len(save.Expense_Name) {
					fmt.Printf("Press E for back to the main menu : ")
					n := bufio.NewReader(os.Stdin)
					n_input, _ := n.ReadString('\n')
					n_input = strings.TrimSpace(n_input)
					if n_input == "E" {
						break
					} else {
						fmt.Println("Please enter valid input ")
						continue

					}
				}

				fmt.Printf("%d Expense name : %v\n", i+1, save.Expense_Name[i])
				fmt.Printf("  *Amount       : +%v  \n ", save.Amount[i])
				fmt.Printf("  *Date         : %v\n\n", save.Date[i])

			}

		case "3":

			for i := 0; ; i++ {

				if i == len(save.Expense_Name) {
					fmt.Print("Enter the expense number that you want to delete : ")
					nm := User_input() - 1

					if nm <= len(save.Expense_Name) {
						save.Expense_Name = append(save.Expense_Name[:nm], save.Expense_Name[nm+1:]...)
						save.Amount = append(save.Amount[:nm], save.Amount[nm+1:]...)
						save.Date = append(save.Date[:nm], save.Date[nm+1:]...)
						break

					} else {
						fmt.Print("invalid input")
					}

				}
				if len(save.Expense_Name) > 0 {
					fmt.Println("i = ", i)
					fmt.Printf("Expense name : %v\n", save.Expense_Name[i])
					fmt.Printf("Amount name : %v\n", save.Amount[i])
					fmt.Printf("Date name : %v\n", save.Date[i])
				}
			}

		case "4":

			return

		default:
			fmt.Println("Invalid input please try again")

		}

	}
}

func User_input() int {

	m := bufio.NewReader(os.Stdin)
	input, _ := m.ReadString('\n')
	input = strings.TrimSpace(input)
	conv_uin, _ := strconv.ParseInt(input, 10, 64)

	n := int(conv_uin)
	return n
}

type name struct {
	Amount []float64

	Expense_Name []string

	Date []string
}
