package cmd

import (
	"fmt"

	"../pkg/application"
)

func main() {
	fmt.Println("main: Start")
	application.StartServer()
	fmt.Println("main: End")
}
