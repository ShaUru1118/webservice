package cmd

import (
	"fmt"

	"github.com/ShaUru1118/webservice/internal/application"
)

func main() {
	fmt.Println("main: Start")
	application.StartServer()
	fmt.Println("main: End")
}
