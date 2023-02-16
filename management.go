package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func getStudent(){
	var filename string
	fmt.Scan(&filename)

	st, err := ioutil.ReadFile(filename)
	if err!= nil{
		// Option 1: log the error and return a call to newDeck()
		fmt.Println("Error: ", err)
		// Option 2: Log the error and entirely quit the program
		os.Exit(1)
	}
	s := strings.Split(string(st), ",") // become a slice of string

	student := populateStudent(s)
	actions(student)

	
}
func actions(student map[string]string){

	for{
		fmt.Println("Now, you can choose action to get information about the student")
		fmt.Print("\n 1): Get name\n 2): Get ID\n 3): Get status\n 4): Get job\n 5): Get Salary\n 6): Print statistic\n 7): Exit program\n")
		var choice int
		fmt.Scan(&choice)
		fmt.Println("Your action", choice)
	
		switch choice {
		case 1:
			getName(student)
			actions(student)
		case 2:
			getID(student)
			actions(student)
		case 3:
			getStatus(student)
			actions(student)
		case 4:
			getJob(student)
			actions(student)
		case 5:
			getSalary(student)
			actions(student)
		case 6:
			statistic(student)
			actions(student)
		case 7:
			exitProgram()
		default:
			fmt.Println("Choose an Int between 1 and 5")
			actions(student)
		}
	}
}

func populateStudent(s []string) map[string]string{
	student := make(map[string]string)
	student["name"] = s[0]
	student["id"] = s[1]
	student["status"] = s[2]
	student["job"] = s[3]
	student["baseSalary"] = s[4]

	return student
}

func getName(st map[string]string) {
	fmt.Println("Name: "+ st["name"])
}

func getID(st map[string]string) {
	fmt.Println("His/her ID is" + st["id"])
}

func getJob(st map[string]string){
	fmt.Println("His/her job is" + st["job"])
}

func getSalary(st map[string]string){

	fmt.Println("Enter the total hours:")
	var hours int
	fmt.Scan(&hours)
	salary, err := strconv.Atoi(strings.TrimSpace(st["baseSalary"]))
	if err != nil {
		fmt.Println("Error in salary")
		os.Exit(1)
	}
	totalSalary := hours*salary
	fmt.Println("Total salary is:", totalSalary, "euros")
}

func getStatus(st map[string]string){
	fmt.Println("Status:", st["status"], ", year ", time.Now().Year())

}

func statistic(st map[string]string){
	fmt.Println("Here is the statistic of the student:")
	getName(st)
	getID(st)
	getStatus(st)
	getJob(st)
	getSalary(st)
}

// func addStudent(){
// 	var newStudent string
// 	fmt.Println("Enter information for new student, separate by comma: <<name, id, status, job, baseSalary>>")
// 	fmt.Scan(&newStudent)
// 	fmt.Println("Please enter a file name:")
// 	var fName string
// 	fmt.Scan(&fName)
// 	populateStudent(strings.Split(newStudent, ","))
// 	ioutil.WriteFile(fName, []byte((newStudent)), 0666)
// }

func exitProgram(){
	fmt.Println("Bye!")
	os.Exit(1)
}

