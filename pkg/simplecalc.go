package pkg

import (
	"fmt"
	"strconv"
	"strings"
)

func count(str string) (string, error) {
	isdab := false
	var r interface{}
	defer func() {
		if r = recover(); r != nil {
			isdab = true
		}
	}()

	if isdab {
		return "", fmt.Errorf("%s", r)
	}

	var errend error
	for strings.Contains(str, "(") {
		ind1 := strings.Index(str, "(")
		ind2 := strings.Index(str, ")")
		strtemp, err := count(str[ind1+1 : ind2])
		str, errend = strings.Replace(str, str[ind1:ind2+1], strtemp, 1), err
	}
	if strings.Contains(str, "**") {
		return "-1", fmt.Errorf("** is not good")
	}
	for strings.Contains(str, "*") {
		ind1 := strings.Index(str, "*") - 1
		ind2 := strings.Index(str, "*") + 1
		num1, err1 := strconv.ParseFloat(string(str[ind1]), 64)
		num2, err2 := strconv.ParseFloat(string(str[ind2]), 64)
		var err error
		if err1 != nil {
			err = err1
		} else {
			err = err2
		}
		str, errend = strings.Replace(str, str[ind1:ind2+1], fmt.Sprint(num1*num2), 1), err
	}
	for strings.Contains(str, "/") {
		ind1 := strings.Index(str, "/") - 1
		ind2 := strings.Index(str, "/") + 1
		num1, err1 := strconv.ParseFloat(string(str[ind1]), 64)
		num2, err2 := strconv.ParseFloat(string(str[ind2]), 64)
		var err error
		if err1 != nil {
			err = err1
		} else {
			err = err2
		}
		str, errend = strings.Replace(str, str[ind1:ind2+1], fmt.Sprint(num1/num2), 1), err
	}
	for strings.Contains(str, "+") {
		ind1 := strings.Index(str, "+") - 1
		ind2 := strings.Index(str, "+") + 1
		num1, err1 := strconv.ParseFloat(string(str[ind1]), 64)
		num2, err2 := strconv.ParseFloat(string(str[ind2]), 64)
		var err error
		if err1 != nil {
			err = err1
		} else {
			err = err2
		}
		str, errend = strings.Replace(str, str[ind1:ind2+1], fmt.Sprint(num1+num2), 1), err
	}
	for strings.Contains(str, "-") {
		ind1 := strings.Index(str, "-") - 1
		ind2 := strings.Index(str, "-") + 1
		num1, err1 := strconv.ParseFloat(string(str[ind1]), 64)
		num2, err2 := strconv.ParseFloat(string(str[ind2]), 64)
		var err error
		if err1 != nil {
			err = err1
		} else {
			err = err2
		}
		str, errend = strings.Replace(str, str[ind1:ind2+1], fmt.Sprint(num1-num2), 1), err
	}
	return str, errend
}

func Calc(expression string) (float64, error) {
	expression, errcount := count(expression)
	res, errres := strconv.ParseFloat(expression, 64)
	var errend error

	if errcount != nil {
		errend = errcount
	} else {
		errend = errres
	}
	return res, errend

}

// func main() {
// 	var x string
// 	fmt.Scanln(&x)
// 	res, err := Calc(x)
// 	fmt.Println(res, err)
// 	// fmt.Println(strings.Replace(x, x[strings.Index(x, "3")-1:strings.Index(x, "3")+1], "w", 1))
// }
