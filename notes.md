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
    var Publicvariable = "public access or exported"
    var privateVariable = "private access"

* Quick Note

    -defer keyword delays whatever that follows it until the current function it is placed in ends.

    -anonymous fucntions are functions without a name. They can be depicted as:

        func(){
        fmt.Println("I am an anonymous function")
        }()

# Using a 3rd party function to read keyboard presses

    package main

    import (
    	"fmt"
    	"log"

    	"github.com/eiannone/keyboard"
    )

    //Using the keyboard 3rd party package

    func main() {
    	err := keyboard.Open()

    	checkError(err)

    	// using defer to close the keyboard package
    	defer func() {
    		_ = keyboard.Close()
    	}()

    fmt.Println("Press any key to start. ESC to quit.")
    
    for{
    	char, key, err := keyboard.GetSingleKey()
    	checkError(err)
    	if key != 0{
    		fmt.Println("You pressed", char, key)
    	}else{
    		fmt.Println("You pressed",char)
    	}	
    	if key == keyboard.KeyEsc{
    		break
        }
    }

    fmt.Println("End of program. Goodbye!")