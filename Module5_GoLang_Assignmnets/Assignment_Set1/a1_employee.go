// Employee Management System in Go
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Constants for departments
const (
	HR_DEPT  = "HR"
	IT_DEPT  = "IT"
	FIN_DEPT = "FINANCE"
)

type Employee struct {
	ID         int
	Name       string
	Age        int
	Department string
}

var employees []Employee

// AddEmployee : adds a new employee after validating input
func AddEmployee(id int, name string, age int, department string) error {
	for _, emp := range employees {
		if emp.ID == id {
			return errors.New("employee ID must be unique")
		}
	}
	if age <= 18 {
		return errors.New("employee age must be greater than 18")
	}

	employees = append(employees, Employee{ID: id, Name: name, Age: age, Department: department})
	return nil
}

// SearchEmployee: searches for an employee by ID or name
func SearchEmployee(query string) (*Employee, error) {
	for _, emp := range employees {
		if fmt.Sprintf("%d", emp.ID) == query || strings.EqualFold(emp.Name, query) {
			return &emp, nil
		}
	}
	return nil, errors.New("employee not found")
}

// ListEmployeesByDepartment : lists employees in a given department
func ListEmployeesByDepartment(department string) []Employee {
	var deptEmployees []Employee
	for _, emp := range employees {
		if strings.EqualFold(emp.Department, department) {
			deptEmployees = append(deptEmployees, emp)
		}
	}
	return deptEmployees
}

// CountEmployeesByDepartment: counts the employees in a given department
func CountEmployeesByDepartment(department string) int {
	count := 0
	for _, emp := range employees {
		if strings.EqualFold(emp.Department, department) {
			count++
		}
	}
	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nOptions:")
		fmt.Println("1: Add Employee")
		fmt.Println("2: Search Employee")
		fmt.Println("3: List Employees by Department")
		fmt.Println("4: Count Employees by Department")
		fmt.Println("5: Exit")
		fmt.Print("Choose an option: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		// we can also use switch case instead of taking input
		if input == "1" {
			fmt.Print("Enter Employee ID: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Invalid ID format.")
				continue
			}

			fmt.Print("Enter Employee Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Enter Employee Age: ")
			ageStr, _ := reader.ReadString('\n')
			ageStr = strings.TrimSpace(ageStr)
			age, err := strconv.Atoi(ageStr)
			if err != nil {
				fmt.Println("Invalid Age format.")
				continue
			}

			fmt.Print("Enter Employee Department: ")
			department, _ := reader.ReadString('\n')
			department = strings.TrimSpace(department)

			if err := AddEmployee(id, name, age, department); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Employee added successfully.")
			}
		} else if input == "2" {
			fmt.Print("Enter Employee ID or Name to search: ")
			query, _ := reader.ReadString('\n')
			query = strings.TrimSpace(query)

			if emp, err := SearchEmployee(query); err == nil {
				fmt.Printf("Employee found: %+v\n", *emp)
			} else {
				fmt.Println("Error:", err)
			}
		} else if input == "3" {
			fmt.Print("Enter Department to list employees: ")
			department, _ := reader.ReadString('\n')
			department = strings.TrimSpace(department)

			employees := ListEmployeesByDepartment(department)
			if len(employees) == 0 {
				fmt.Println("No employees found in this department.")
			} else {
				fmt.Println("Employees in", department, "Department:")
				for _, emp := range employees {
					fmt.Printf("%+v\n", emp)
				}
			}
		} else if input == "4" {
			fmt.Print("Enter Department to count employees: ")
			department, _ := reader.ReadString('\n')
			department = strings.TrimSpace(department)

			count := CountEmployeesByDepartment(department)
			fmt.Printf("Number of employees in %s: %d\n", department, count)
		} else if input == "5" {
			fmt.Println("Exiting Employee Management System.")
			break
		} else {
			fmt.Println("Invalid option, try again.")
		}
	}
}
