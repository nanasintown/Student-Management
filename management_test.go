package main

import "testing"

func TestNewStudent(t *testing.T) {
	newStudent := []string{"Gabrielle", "21313", "Active", "PhD", "20"}
	student := populateStudent(newStudent)
	if len(student) != 5{
		t.Errorf("Expect 5 values, but got %d", len(student))
	}
}