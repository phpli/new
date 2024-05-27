package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func main() {
	server := NewHttpServer("test-server")
	//server.Route("/", home)
	//server.Route("/user", user)
	//server.Route("/user/create", createUser)
	server.Route(http.MethodGet, "/user/signup", SignUp)
	//server.Route("/order", order)
	err := server.Start(":8080")
	if err != nil {
		panic(err)
	}
}

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User Page")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create User Page")
}

func order(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Order Page")
}
