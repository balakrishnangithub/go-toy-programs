// credit: https://golangbot.com/interfaces-part-1/#practicaluseofaninterface

package main

import (
	"fmt"
)

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empID    int
	basicpay int
	pf       int
}

type Contract struct {
	empID    int
	basicpay int
}

type Freelancer struct {
	empID       int
	ratePerHour int
	totalHours  int
}

//salary of permanent employee is sum of basic pay and pf
func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

//salary of contract employee is the basic pay alone
func (c Contract) CalculateSalary() int {
	return c.basicpay
}

//salary of freelancer
func (f Freelancer) CalculateSalary() int {
	return f.ratePerHour * f.totalHours
}

/*
total expense is calculated by iterating through the SalaryCalculator slice and summing
the salaries of the individual employees
*/
func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month ₹%d\n", expense)
}

func salaryCalculator() {
	pemp1 := Permanent{
		empID:    1,
		basicpay: 5000,
		pf:       20,
	}
	pemp2 := Permanent{
		empID:    2,
		basicpay: 6000,
		pf:       30,
	}
	cemp1 := Contract{
		empID:    3,
		basicpay: 3000,
	}
	freelancer1 := Freelancer{
		empID:       4,
		ratePerHour: 70,
		totalHours:  120,
	}
	freelancer2 := Freelancer{
		empID:       5,
		ratePerHour: 100,
		totalHours:  100,
	}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1, freelancer1, freelancer2}
	totalExpense(employees)
}
