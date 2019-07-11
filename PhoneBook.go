package main

import (
	_ "bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type PhoneBook struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

var Numbers = make(map[int]PhoneBook)

func main() {

	Numbers[0] = PhoneBook{"Ivanov", "+79001234567"}
	Numbers[1] = PhoneBook{"Petrov", "+79001234568"}
	Numbers[2] = PhoneBook{"Kirkorov", "+79001234560"}
	Numbers[3] = PhoneBook{"Galkin", "+79001234561"}

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
		fmt.Println("Вы используете неверный метод запроса")
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

		fmt.Println(len(Numbers))
	}
}

func getPhoneNumber(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Println("Вы используете неверный метод запроса")
	} else {

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			return
		}

		var str = string(b)
		i, err := strconv.Atoi(str)
		if i > len(Numbers) {
			fmt.Println("записи с таким id, не существует...")
		} else {
			fmt.Println(Numbers[i])
		}

	}
}

func deletePhoneNumber(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		fmt.Println("Вы используете неверный метод запроса")
	} else {

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			return
		}

		var str = string(b)
		i, err := strconv.Atoi(str)
		if i > len(Numbers) {
			fmt.Println("записи с таким id, не существует...")
		} else {
			delete(Numbers, i)
			fmt.Println("record delete...")
		}

		fmt.Println(len(Numbers))

	}
}

func postPhoneNumber(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Println("Вы используете неверный метод запроса")
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
		fmt.Println("record add ...")

		fmt.Println(len(Numbers))

	}
}

func updatePhoneNumber(w http.ResponseWriter, r *http.Request) {
	if r.Method != "UPDATE" {
		fmt.Println("Вы используете неверный метод запроса")
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
		fmt.Println("record update ...")

		fmt.Println(Numbers[i])
	}
}
