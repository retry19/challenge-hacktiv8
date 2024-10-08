package main

import (
	"fmt"
	"os"
	"strings"
)

type Person struct {
	Id      uint8
	Name    string
	Address string
	Job     string
	Reason  string
}

var peoples = []Person{
	{Id: 1, Name: "Reza", Address: "Kuningan", Job: "Backend Engineer", Reason: "Upgrade skill"},
	{Id: 2, Name: "Foo", Address: "Jakarta", Job: "Fullstack Engineer", Reason: "Learn new tech"},
	{Id: 3, Name: "Bar", Address: "Bandung", Job: "Frontend Engineer", Reason: "For work"},
	{Id: 4, Name: "Robert", Address: "Bali", Job: "Backend Engineer", Reason: "Interest in tech"},
	{Id: 5, Name: "Toiro", Address: "Cirebon", Job: "Devops Engineer", Reason: "Learning"},
}

func findPersonByName(name string) (person Person, isExists bool) {
	isExists = false

	for _, p := range peoples {
		if strings.EqualFold(p.Name, name) {
			person = p
			isExists = true
			break
		}
	}

	return
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("error: Missing person-name parameter\n\nUsage: %s <person-name>\n", os.Args[0])
		return
	}

	person, isExists := findPersonByName(os.Args[1])
	if !isExists {
		fmt.Println("error: person not found")
		return
	}

	fmt.Println("ID :", person.Id)
	fmt.Println("nama :", person.Name)
	fmt.Println("alamat :", person.Address)
	fmt.Println("pekerjaan :", person.Job)
	fmt.Println("alasan :", person.Reason)
}
