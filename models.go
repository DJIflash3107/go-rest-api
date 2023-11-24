package main

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

var users = []User{{
	Name:  "Rangga Yuda",
	Age:   17,
}, {
	Name:  "Hafidz Fachrur Viamsyah",
	Age:   17,
}}
