package main

// ISP - Break up interface into different parts so that people can use or need

import "fmt"

type Document struct {
}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type Fax interface {
	Fax(d Document)
}

//Photocopier (Print and Scan)
type Photocopier struct{}

func (p Photocopier) Print(d Document) {
	fmt.Println("Print")
}
func (p Photocopier) Scan(d Document) {
	fmt.Println("Scan")
}

//Multi Function Device (using split interface)
type MultiPurposeDevice struct {
	Scanner
	Fax
}

//Decorator

type MultiPurposeDevice2 struct {
	printer Printer
	scanner Scanner
}

func (m MultiPurposeDevice2) PrintScan(d Document) {
	m.printer.Print(d)
	m.scanner.Scan(d)
}

func main() {}
