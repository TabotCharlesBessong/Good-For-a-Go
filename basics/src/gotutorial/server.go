package main

import (
	"net/http"
	"fmt"
	"log"
	"errors"
	"encoding/json"
)

// const time string = "Hello Unchagable"

func main(){
	// time = "Really?"
	// fmt.Println(time)
	http.HandleFunc("/create-payment-intend", handleCreatePaymentIntend)

	// doesChange := "can be changed"
	// fmt.Println(doesChange)

	// doesChange = "has been changed"
	// fmt.Println(doesChange)

	http.HandleFunc("/health-check", handleHealth)

	log.Println("Listening on localhost:4242.....")

	var err error = http.ListenAndServe("localhost:4242",nil)
	if err != nil {
		log.Fatal(err)
	}

	// var names []string = []string{"Charles","Joseph","Junior"}
	// fmt.Println(names)

	// var me string = returnValue("How are you doing my nigga?")
	// fmt.Println(me)

	// var err error = returnsError("wrongpassword")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// functionOne(anotherOne)
}

func functionOne(functionTwo func()){
	functionTwo()
}

func anotherOne(){
	fmt.Println("Hello I am function 2")
}

func returnsError(password string) error {
	var secretPassword string = "supersecret"
	if password == secretPassword{
		return nil
	}else{
		return errors.New("Invalid password")
	}
}

func handleCreatePaymentIntend(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST" {
		http.Error(w,http.StatusText(http.StatusMethodNotAllowed),http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		ProductId string `json:"product_id"`
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
		Address1 string `json:"address_1"`
		Address2 string `json:"address_2"`
		City string `json:"city"`
		Region string `json:"region"`
		Zip string `json:"zip"`
		Country string `json:"country"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	fmt.Println(req.FirstName)
	fmt.Println(req.LastName)
	fmt.Println(req.Address1)
	fmt.Println(req.Address2)
	fmt.Println(req.City)
	fmt.Println(req.Region)
	fmt.Println(req.Country)

	// w.Write(req)
}

func returnValue (hey string) string {
	return hey
}

func handleHealth(w http.ResponseWriter, r *http.Request){
	// fmt.Println("OK!")
	response := []byte("The Server is up and running")

	someString, someInt, someBool := returnsMultiple()
	fmt.Println(someString,"\n",someInt,"\n",someBool)
	
	w.Write(response)
}

func returnsMultiple() (string,int,bool){
	return "hello", 1, true
}
