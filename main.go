package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Format struct {
	Internationll string `json:"international"`
	Local         string `json:"local"`
}

type Country struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Prefix string `json:"prefix"`
}

type phoneValid struct {
	Phone    string  `json:"phone"`
	Valid    bool    `json:"valid"`
	Format   Format  `json:"format"`
	Country  Country `json:"country"`
	Location string  `json:"location"`
	Type     string  `json:"type"`
	Carrier  string  `json:"carrier"`
}

func main() {
	var key = "d9266a3b857245e484cdc185ca1faca8"
	var phoneNumber string = ""
	for {
		fmt.Print("Введите номер телефона: ")
		_, err := fmt.Scan(&phoneNumber)
		if err != nil {
			fmt.Println("Ошибка ввода")
			continue
		}
		break
	}
	resp, err := http.Get("https://phonevalidation.abstractapi.com/v1/?api_key=" + key + "&phone=" + phoneNumber)
	if err != nil {
		fmt.Println("Ошибка запроса: ", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения тела ответа: ", err)
	}

	var res phoneValid
	json.Unmarshal(body, &res)

	fmt.Print(res.Format.Local)
	if !res.Valid {
		fmt.Println("\t invalid")
		return
	}
	fmt.Print("\tvalid")
	fmt.Println("\n\tCountry: " + res.Country.Name)
	fmt.Println("\tType: " + res.Type)
	fmt.Println("\tOperator: " + res.Carrier)
}
