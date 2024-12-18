package main

import (
	"fmt"

	"github.com/shauru1118/webservice/internal/application"
)

func Main() {
	fmt.Println("main: Start")
	application.StartServer()
	fmt.Println("main: End")
}
