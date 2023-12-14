package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	whatToSay := doctor.Intro()
	fmt.Println("AI:" + whatToSay)

	for{
		userInput, _ := reader.ReadString('\n')

		//Stripping /n from the userInput to allow our program terminate

		userInput = strings.Replace(userInput, "\r\n", "", -1) //For Windows only
		userInput = strings.Replace(userInput, "\n", "", -1) //For other operating systems

		if(userInput == "quit"){
			break
		}else{
			fmt.Println("AI: " + doctor.Response(userInput))
		}
	}
}