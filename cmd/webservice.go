package cmd

import (
	"fmt"

	"../pkg"
)

func main() {
	fmt.Println("main: Start")
	pkg.StartServer()
	fmt.Println("main: End")
}
