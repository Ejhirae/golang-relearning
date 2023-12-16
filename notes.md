# GoLang Basic Structure

    package main

    import "fmt"

    func main() {

    fmt.Println("Hello World!") 
    
    }

# GoLang Variables Declarations
variables called in GoLang must be used

Basic syntax for declaring variables

Longhand variable declaration
    var characterName string = "Victor"

Shorthand variable declaration
    characterAge := 21

# Inputting in GoLang

    reader := bufio.NewReader(os.Stdin)
    userInput , _ := reader.ReadString('\n')


# Stripping "/n" from the userInput to allow our program terminate

		userInput = strings.Replace(userInput, "\r\n", "", -1) //For Windows only
		userInput = strings.Replace(userInput, "\n", "", -1) //For other operating systems

# Forever For Loop
    for{
        fmt.Println(true)
    }

# Variables in GoLang