package main

import "fmt"

// AntIncome interface.
type AntIncome interface {
	calculate() float64
	source() string
}

// Stock struct.
type Stock struct {
	code  string
	total int
	price float64
}

// Fund struct.
type Fund struct {
	vendor string
	total  int
	price  float64
}

// Advertisement struct.
type Advertisement struct {
	name          string
	clickCounter  int
	perClickPrice float64
}

func (s Stock) calculate() float64 {
	return float64(s.total) * s.price
}

func (s Stock) source() string {
	return s.code
}

func (f Fund) calculate() float64 {
	return float64(f.total) * f.price
}

func (f Fund) source() string {
	return f.vendor
}

func (a Advertisement) calculate() float64 {
	return float64(a.clickCounter) * a.perClickPrice
}

func (a Advertisement) source() string {
	return a.name
}

func main() {
	s1 := Stock{"SH019008", 1000, 89.89}
	s2 := Stock{"USBABA", 1000, 120.05}
	f1 := Fund{"000013", 1000, 12.88}
	f2 := Fund{"100292", 1000, 11.55}
	a1 := Advertisement{"Baidu Search Ad", 1000000, 0.03}
	a2 := Advertisement{"Google Search Ad", 1000000, 0.04}
	income := []AntIncome{s1, s2, f1, f2, a1, a2}
	fmt.Printf("Total amount: %.2f\n", calculateTotalIncome(income))
}

func calculateTotalIncome(incomes []AntIncome) float64 {
	amount := 0.0
	for _, income := range incomes {
		fmt.Printf("Source: %s, Amount: %.2f\n",
			income.source(), income.calculate())
		amount += income.calculate()
	}
	return amount
}
