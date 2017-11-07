package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Monkey-Ops Service!\n")
}

//Rest json (POST) to obtain token and projects for a user
func OcLogin(w http.ResponseWriter, r *http.Request) {
	
	var loginInput LoginInput
	
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &loginInput); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	loginOutput := Login(&loginInput)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(loginOutput); err != nil {
		panic(err)
	}
}

//Rest json (POST) to launch chaos in a devops project
func DoChaos(w http.ResponseWriter, r *http.Request) {

	var chaosInput ChaosInput

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &chaosInput); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	
	go ExecuteChaos(&chaosInput, "rest")	

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

}

