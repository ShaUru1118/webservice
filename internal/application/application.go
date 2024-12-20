package application

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/shauru1118/webservice/pkg/simplecalc"
)

type postQwes struct {
	Expression string `json:"expression"`
}

type postError struct {
	Error string `json:"error"`
}

type postAnser struct {
	Result string `json:"result"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {

	body, errbody := ioutil.ReadAll(r.Body)
	body = []byte(strings.ReplaceAll(strings.ReplaceAll(string(body), "\\", ""), "'", ""))

	if errbody != nil {
		anserr := postError{Error: "Internal server error : r.Body.Read(data)"}
		jsonstr, _ := json.Marshal(anserr)
		http.Error(w, string(jsonstr), 500)
		fmt.Println("Error: ", anserr.Error)
		return
	}

	fmt.Println("r.body.read: |", string(body), "|")

	var datago postQwes

	errunmar := json.Unmarshal(body, &datago)

	if errunmar != nil {
		anserr := postError{Error: "Internal server error : json.Unmarshal(body, &datago)"}
		jsonstr, _ := json.Marshal(anserr)
		http.Error(w, string(jsonstr), 500)
		fmt.Println("Error: ", anserr.Error)
		return
	}

	expretion := datago.Expression

	fmt.Println("expretion: |", expretion, "|")

	res, errcalc := simplecalc.Calc(expretion)

	if errcalc != nil {
		anserr := postError{Error: "Expression is not valid : simplecalc.Calc(expretion)"}
		jsonstr, _ := json.Marshal(anserr)
		http.Error(w, string(jsonstr), 500)
		fmt.Println("Error: ", anserr.Error)
		return
	}

	ans := postAnser{Result: fmt.Sprint(res)}
	jsonstr, _ := json.Marshal(ans)
	fmt.Fprintf(w, "%s", string(jsonstr))
	fmt.Println("Result: ", ans.Result)

	// fmt.Fprint(w, res)
}

func StartServer() {
	http.HandleFunc("/api/v1/calculate", CalcHandler)

	fmt.Println("application: Start Server")

	if http.ListenAndServe(":8080", nil) == nil {
		fmt.Println("application: Started Server")
	}

	fmt.Println("application: Stopping Server")
}
