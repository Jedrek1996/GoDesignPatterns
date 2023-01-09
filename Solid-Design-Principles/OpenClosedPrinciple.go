package main

import "fmt"

//OCP, open for extension closed for modification

// The interface type is opened for extension but closed for modification (unlikely need to modify the filter)
// (Once you design and tested the api you should not modifiy it, you already got clients working so try implement interface and making new types)

// Specification allows us to modify the inputs to specify
type Size int

type Color int

const (
	small Size = iota
	medium
	large
)

const (
	red Color = iota
	green
	blue
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

type Specification interface {
	IsSatisfied(p *Product) bool
}

// Color specification
type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

// Size specification
type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == s.size
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

type AndSpecification struct {
	first, second Specification
}

func (a AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}

func main() {

	Apple := Product{"Apple", red, large}
	Pear := Product{"Pear", green, large}
	Grape := Product{"Grape", blue, large}

	products := []Product{Apple, Pear, Grape}
	fmt.Printf("Green products (old):\n")
	f := Filter{}
	for _, v := range f.FilterByColor(products, green) {
		fmt.Printf("Product %v is green\n", v.name)
	}

	fmt.Printf("Green prodcut new:!!S \n")
	greenSpec := ColorSpecification{green}
	bf := BetterFilter{}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf("Product %v is green\n", v.name)
	}

	largeSpec := SizeSpecification{large}
	lgSpec := AndSpecification{greenSpec, largeSpec}
	fmt.Println("Large green products")
	for _, v := range bf.Filter(products, lgSpec) {
		fmt.Printf("%s Is large and green\n", v.name)
	}
}
