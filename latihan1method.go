package main

import "fmt"

type Field struct {
	Firstname string
	Lastname  string
	Job       string
	Salary    string
}

func (s Field) GetFirstname() string {
	return s.Firstname
}
func (s *Field) SetFirstname(firstname string) {
	s.Firstname = firstname

}
func (s Field) GetLastname() string {
	return s.Lastname
}
func (s *Field) SetLastname(lastname string) {
	s.Lastname = lastname

}
func (s Field) GetJob() string {
	return s.Job
}
func (s *Field) SetJob(job string) {
	s.Job = job

}
func (s Field) GetSalary() string {
	return s.Salary
}
func (s *Field) SetSalary(salary string) {
	s.Salary = salary

}

func main() {
	var s1 Field

	s1.GetFirstname()
	s1.SetFirstname("Budi")
	fmt.Println("Firstname:", s1.Firstname)

	s1.GetLastname()
	s1.SetLastname("Pratama")
	fmt.Println("Lastname:", s1.Lastname)

	s1.GetJob()
	s1.SetJob("Marketing")
	fmt.Println("Job:", s1.Job)

	s1.GetSalary()
	s1.SetSalary("2.000.000")
	fmt.Println("Salary:", s1.Salary)

}
