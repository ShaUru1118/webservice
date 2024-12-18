package application

import (
	"fmt"
	"net/http"

	"github.com/shauru1118/webservice/pkg/simplecalc"
)

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	expretion := r.URL.Query().Get("expression")
	res, errcalc := simplecalc.Calc(expretion)
	if errcalc != nil {
		http.Error(w, "Expression is not valid", 422)
	}
	fmt.Fprint(w, res)
}

func StartServer() {
	fmt.Println("application: Starting Server")

	http.HandleFunc("/api/v1/calculate", CalcHandler)

	http.ListenAndServe(":8080", nil)

	fmt.Println("application: Stopping Server")
}
