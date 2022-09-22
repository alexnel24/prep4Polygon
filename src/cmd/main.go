package main

import (
	"fmt"
	"net/http"

	"github.com/alexnel24/prep4Polygon/src/app"
)

func main(){
	fmt.Println("Hello Polygon")

	server := app.NewServer()

	http.ListenAndServe(":8080", server)

}

