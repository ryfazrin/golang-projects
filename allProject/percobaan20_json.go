package main

import (
	"fmt"
	"encoding/json"
)

type UserInfo struct {
	Username    string `json:"user"`
	Email    	string `json:"email"`
	Password 	string `json:"pass"`
}

func main() {
	// contoh 1
	u :=UserInfo{}
	u.Username = "test"
	u.Email = "test@test.com"
	u.Password = "123456"

	fmt.Println(u)

	data_json, err:=json.Marshal(&u)
	if err == nil{
		fmt.Println(string(data_json))
	}


	// contoh 2
	external_data:= []byte("{\"user\":\"test\",\"email\":\"test@test.com\",\"pass\":\"123456\"}")
	u2 :=UserInfo{}
	fmt.Println(u2)

	json.Unmarshal(external_data, &u2)
	fmt.Println(u2)
}
