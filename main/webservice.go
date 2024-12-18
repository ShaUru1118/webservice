package main

import (
	"fmt"

	"github.com/shauru1118/webservice/internal/application"
)

func main() {
	fmt.Println("main: Start")
	application.StartServer()
	fmt.Println("main: End")
}
