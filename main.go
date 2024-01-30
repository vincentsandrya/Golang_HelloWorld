package main

import (
	"GOLANG_HELLOWORLD/config"
	"GOLANG_HELLOWORLD/controllers/homecontroller"
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("main start")
	config.GetConnect()

	http.HandleFunc("/", homecontroller.Welcome)
	err := http.ListenAndServe("localhost:8080", nil)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("main end")
}
