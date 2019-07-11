package main

import (
	_ "bytes"
	"encoding/json"
	"fmt"

	"io/ioutil"
	"log"
	"net/http"
)

type PhoneBook struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

var Numbers = make(map[int]PhoneBook)

func main() {

	Numbers[0] = PhoneBook{"Ivanov", "+79001234567"}
	Numbers[1] = PhoneBook{"Petrov", "+79001234568"}

	http.HandleFunc("/getall", getAllPhoneNumbers)
	http.HandleFunc("/get", getPhoneNumber)
	http.HandleFunc("/post", postPhoneNumber)
	http.HandleFunc("/delete", deletePhoneNumber)
	http.HandleFunc("/update", updatePhoneNumber)

	fmt.Println("starting ...")
	err := http.ListenAndServe("localhost:80", nil)
	if err != nil {
		log.Println(err)
		return
	}
}

func getAllPhoneNumbers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Println("Вы используете неверный запрос")
	} else {

		_, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			return
		}

		i := 0
		for i < len(Numbers) {
			fmt.Print(i, " ")
			fmt.Println(Numbers[i])
			i = i + 1
		}
	}
}

func getPhoneNumber(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Println("Вы используете неверный запрос")
	} else {

		_, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			return
		}

		i := 0
		for i < len(Numbers) {
			fmt.Print(i, " ")
			fmt.Println(Numbers[i])
			i = i + 1
		}
	}
}

func deletePhoneNumber(w http.ResponseWriter, r *http.Request) {
	//
}

func postPhoneNumber(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Println("Вы используете неверный запрос")
	} else {

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			return
		}

		var add PhoneBook
		err = json.Unmarshal(b, &add)
		if err != nil {
			log.Println(err)
			return
		}

		i := len(Numbers)
		Numbers[i] = PhoneBook{add.Name, add.Phone}
	}
}

func updatePhoneNumber(w http.ResponseWriter, r *http.Request) {
	//
}
