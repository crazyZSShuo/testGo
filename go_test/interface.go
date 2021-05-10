package main

import "fmt"

//type VowelsFinder interface {
//	FindVowels() []rune
//}
//
//type MyString string
//
//func (ms MyString) FindVowels() []rune {
//	var vowels []rune
//	for _, rune := range ms{
//		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u'{
//			vowels = append(vowels, rune)
//		}
//	}
//	return vowels
//}
//
//func main()  {
//	name := MyString("az")
//	fmt.Println(name)
//	var v VowelsFinder
//	v = name
//	fmt.Println(v.FindVowels())
//}

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

type Contract struct {
	empId    int
	basicpay int
}

type Freelancer struct {
	empId int
	ratePerHour int
	totalHours int
}


func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

func (c Contract) CalculateSalary() int {
	return c.basicpay
}

func (f Freelancer) CalculateSalary() int{
	return f.ratePerHour * f.totalHours
}

func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense += v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d", expense)
}

func main() {
	pemp1 := Permanent{1, 5000, 20}
	pemp2 := Permanent{2, 6000, 30}
	cemp1 := Contract{3, 3000}
	free1 := Freelancer{
		empId:       4,
		ratePerHour: 70,
		totalHours:  120,
	}
	free2 := Freelancer{
		empId:       3,
		ratePerHour: 100,
		totalHours:  100,
	}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1, free1, free2}
	totalExpense(employees)
}
