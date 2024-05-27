package main

import (
	"fmt"
	"log"
	"net/http"
	"unicode/utf8"
)

func main() {
	println("Hello World")
	println(len("你好"))                     //len 算的是字节长度，字节长度与编码无关
	println(utf8.RuneCountInString("早上好")) //字符数量与编码无关，用编码库来取下·
	http.HandleFunc("/", home)
	http.HandleFunc("/user", user)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love 2 %s!", r.URL.Path[1:])
}
