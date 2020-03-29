package main

import "fmt"

type field struct {
	firstname string
	lastname  string
	job       string
	salary    string
}

func (s *field) GetField(firstname string, lastname string, job string, salary string) {
	s.firstname = firstname
	s.lastname = lastname
	s.job = job
	s.salary = salary

}
func main() {
	var s1 = field{"Budi", "Pratama", "Marketing", "2.000.000"}
	s1.GetField("Budi", "Pratama", "Marketing", "2.000.000")
	fmt.Println("Firstname:", s1.firstname)
	fmt.Println("Lastname:", s1.lastname)
	fmt.Println("Job:", s1.job)
	fmt.Println("Salary:", s1.salary)

}
