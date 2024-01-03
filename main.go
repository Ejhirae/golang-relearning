package main

import (
	"fmt"
	"log"
	"strconv"

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

	//creating a map in golang

	programmingLanguages := make(map[int]string)

	programmingLanguages[1] = "Golang"
	programmingLanguages[2] = "Java"
	programmingLanguages[3] = "C#"
	programmingLanguages[4] = "Dart"
	programmingLanguages[5] = "PHP"
	programmingLanguages[6] = "Python"

	fmt.Println("Press any key to start. ESC to quit.")

	fmt.Println("Pick any programming language to learn about it")

	fmt.Println("1. Golang")
	fmt.Println("2. Java")
	fmt.Println("3. C#")
	fmt.Println("4. Dart")
	fmt.Println("5. PHP")
	fmt.Println("6. Python")
	fmt.Println("7. Quit")

	for{
		char, _, err := keyboard.GetSingleKey()

		checkError(err)

		
		if char == 'q' || char == 'Q'{
			break
		}

		runeToInteger, _ := strconv.Atoi(string(char))

		fmt.Println(fmt.Sprintf("You chose %s", programmingLanguages[runeToInteger]))

		fmt.Println()

		fmt.Println(responseAi(runeToInteger))

		fmt.Println()


	}

	fmt.Println("End of program. Goodbye!")
}

func checkError(err error){
	if err != nil{
		log.Fatal(err)
	}
}

func golangInfo()string{
	return `Go, commonly known as Golang, is a statically typed, compiled programming language developed by Google. Released in 2009, Go is designed to be efficient, simple, and easy to read, making it well-suited for building scalable and robust software systems. Go's syntax is straightforward, and it features built-in support for concurrent programming through goroutines and channels, enabling the development of highly concurrent and parallel applications. It emphasizes performance without sacrificing simplicity, offering features like garbage collection and a strong standard library. Go is particularly favored for backend development, cloud computing, and microservices architecture. Its static typing catches errors early in the development process, contributing to the creation of reliable and maintainable code. With its focus on speed, simplicity, and concurrency, Go has gained popularity for building scalable and efficient software solutions, especially in the context of distributed systems.`
}

func javaInfo() string{
	return `Java is a highly versatile and widely adopted programming language known for its platform independence, making it suitable for a variety of applications. It follows an object-oriented paradigm, emphasizing modularity and code reusability through classes and objects. Java's syntax, reminiscent of C++, is designed to be straightforward and less error-prone. Central to its portability is the Java Virtual Machine (JVM), which enables Java applications to run on diverse devices. The language boasts a rich standard library, facilitating efficient development with a wide range of built-in functionalities. Multithreading support allows for concurrent execution, and its robust security features, including a sandbox model, make it suitable for secure environments. With an active community and a vast ecosystem of libraries and frameworks, Java remains a preferred choice for developing enterprise applications, web services, mobile applications (Android), and more. Frequent updates and a commitment to backward compatibility contribute to Java's enduring popularity since its introduction in the 1990s.`
}

func csharpInfo() string{
	return `C# (pronounced "C sharp") is a modern, object-oriented programming language developed by Microsoft. Known for its simplicity, readability, and versatility, C# is widely used for developing a diverse range of applications, including web applications, desktop software, and games. C# integrates seamlessly with the .NET framework, providing a robust set of libraries and tools for rapid application development. Its syntax is influenced by C and C++, making it familiar to developers of these languages. C# supports features like automatic memory management through garbage collection, strong typing, and advanced language constructs like LINQ for querying data. With a strong emphasis on security, C# enables the development of secure and scalable applications. It is a key language for building applications on the Microsoft platform, including Windows applications and services, as well as cross-platform solutions using technologies like .NET Core.`
}

func dartInfo() string{
	return `Dart is a modern, open-source programming language developed by Google that is primarily associated with building web and mobile applications. Known for its simplicity and flexibility, Dart combines a concise syntax with strong typing, making it accessible to developers of various backgrounds. Dart is the language used for developing Flutter, a popular UI toolkit for building natively compiled applications for mobile, web, and desktop from a single codebase. Dart supports both just-in-time (JIT) and ahead-of-time (AOT) compilation, providing flexibility in terms of development and runtime performance. It incorporates features such as asynchronous programming for efficient handling of concurrent tasks and a rich standard library. Dart's ability to seamlessly integrate with web browsers and its strong focus on creating responsive and visually appealing user interfaces contribute to its increasing adoption in the development community.`
}

func phpInfo() string{
	return `PHP, originally an acronym for "Personal Home Page," is a widely used server-side scripting language designed for web development. Known for its simplicity and flexibility, PHP is embedded within HTML code and executed on the server, generating dynamic content that is sent to the client's browser. It is particularly adept at handling tasks like form processing, database connectivity, and session management. PHP has evolved into a full-fledged programming language and is widely employed for building dynamic web applications, content management systems (CMS), and e-commerce platforms. With a vast ecosystem of libraries and frameworks, including Laravel and Symfony, PHP facilitates rapid development and scalability. Its open-source nature and compatibility with various web servers and databases make PHP a popular choice for web developers seeking a pragmatic and efficient solution for building interactive and dynamic websites.`
}

func pythonInfo() string{
	return `Python is a versatile, high-level programming language renowned for its readability, simplicity, and extensive community support. Created by Guido van Rossum and first released in the early 1990s, Python emphasizes code readability and a clean syntax, making it an ideal choice for beginners and experienced developers alike. It supports multiple programming paradigms, including procedural, object-oriented, and functional programming, providing flexibility for diverse applications. Python is widely used in web development, data analysis, artificial intelligence, scientific computing, and automation. Its extensive standard library and third-party packages, such as NumPy, Pandas, and Django, contribute to its popularity and efficiency in diverse domains. Python's interpretive nature allows for easy prototyping and testing, and its cross-platform compatibility ensures broad deployment across various operating systems. The language's dynamic typing and automatic memory management simplify development, while its active community continually contributes to its growth and evolution.`
}

func responseAi(option int ) string{
	switch option {
	case 1:
		return golangInfo()
	case 2:
		return javaInfo()
	case 3:
		return csharpInfo()
	case 4:
		return dartInfo()
	case 5:
		return phpInfo()
	case 6:
		return pythonInfo()
	default:
		return `Invalid option chosen`
	}
	
}