package main

import (
	"fmt"
	"os"
	"strconv"
)

type Student struct {
	Id         int
	Name       string
	Address    string
	Job        string
	JoinReason string
}

func main() {
	students := []Student{
		{
			Id:         1,
			Name:       "Budi",
			Address:    "Jl Budi",
			Job:        "Student",
			JoinReason: "Budi's reason",
		},
		{
			Id:         2,
			Name:       "Adam",
			Address:    "Jl Adam",
			Job:        "Student",
			JoinReason: "Adam's reason",
		},
		{
			Id:         3,
			Name:       "Yusuf",
			Address:    "Jl Yusuf",
			Job:        "Student",
			JoinReason: "Yusuf's reason",
		},
	}

	args := os.Args

	if len(args) < 2 {
		panic("Please enter the id")
	}

	id, err := strconv.Atoi(args[1])
	if err != nil {
		panic(err.Error())
	}

	for _, v := range students {
		if v.Id == id {
			fmt.Println("Name: ", v.Name)
			fmt.Println("Alamat: ", v.Address)
			fmt.Println("Pekerjaan: ", v.Job)
			fmt.Println("Alasan: ", v.JoinReason)
			return
		}
	}

	fmt.Printf("Student with Id %d not found\n", id)
}
