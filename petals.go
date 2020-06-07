package main

import (
	"fmt"
	"math/rand"
	"time"
)

type dice struct {
	petals int
	face   [3]string
}

func main() {
	d := [6]dice{
		{
			0,
			[3]string{
				" |   | ",
				" | * | ",
				" |   | ",
			},
		}, {
			0,
			[3]string{
				" |  *| ",
				" |   | ",
				" |*  | ",
			},
		}, {
			2,
			[3]string{
				" |  *| ",
				" | * | ",
				" |*  | ",
			},
		}, {
			0,
			[3]string{
				" |* *| ",
				" |   | ",
				" |* *| ",
			},
		}, {
			4,
			[3]string{
				" |* *| ",
				" | * | ",
				" |* *| ",
			},
		}, {
			0,
			[3]string{
				" |* *| ",
				" |* *| ",
				" |* *| ",
			},
		},
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rolls := [5]int{0, 0, 0, 0, 0}
	petals := 0
	guess := 0
	tries := 0
	correct := 0
	streak := 0
	quit := ""

	for {
		fmt.Println("The name of the game is Petals Around the Rose, and the name is significant.")
		fmt.Printf("The answer is zero or an even number.\n\n")
		petals = 0
		for n, _ := range rolls {
			rolls[n] = r.Intn(6)
			petals = petals + d[rolls[n]].petals
		}
		for n := 0; n < 3; n++ {
			for _, roll := range rolls {
				fmt.Print(d[roll].face[n])
			}
			fmt.Print("\n")
		}
		fmt.Println("\nWhat is your answer?")
		if n, err := fmt.Scanln(&guess); n == 1 && err == nil {
			if guess == petals {
				fmt.Println("Your answer was correct!")
				correct++
				streak++
			} else {
				streak = 0
				fmt.Println("Your answer was wrong!")
				fmt.Printf("You answered %d.\n", guess)
			}
		} else {
			streak = 0
		}
		tries++
		fmt.Printf("There are %d petals.\n", petals)
		fmt.Printf("\n\nYour stats:\nTries:\t%d\nCorect:\t%d\nStreak:\t%d\n\n", tries, correct, streak)
		fmt.Println("Input any keys to quit, enter to continue.")
		if n, _ := fmt.Scanln(&quit); n != 0 {
			return
		}
		fmt.Printf("\n\n")
	}
}
