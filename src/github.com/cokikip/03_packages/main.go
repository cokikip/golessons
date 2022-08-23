package main

import (
	"fmt"
	"math"
	"net/http"

	"cokikip/packages/src/github.com/cokikip/03_packages/strutil"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println(math.Floor(4.756))
	fmt.Println(math.Ceil(4.756))
	fmt.Println(math.Sqrt(9))
	fmt.Println(strutil.Reverse("Collins"))
	fmt.Print()
	r := mux.NewRouter()
	http.ListenAndServe(":9000", r)
}
