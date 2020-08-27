package main

import "fmt"

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

func (p *Product) String() string {
	return fmt.Sprintf("%s:%d", p.name, p.color)
}

type Filter struct {
}

type BetterFilter struct{}

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

//去實作Specification interface
func (spec ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == spec.color
}

func (f *BetterFilter) Filter(products []*Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for _, v := range products {
		if spec.IsSatisfied(v) {
			result = append(result, v)
		}
	}
	return result
}
func (f *Filter) filterByColor(products []*Product, color Color) []*Product {
	result := make([]*Product, 0)
	for _, v := range products {
		if v.color == color {
			fmt.Println("p:", v.name)
			result = append(result, v)
		}
	}
	return result
}

func main() {
	apple := &Product{"Apple", green, small}
	tree := &Product{"Tree", green, large}
	house := &Product{"House", blue, large}
	prpducts := []*Product{apple, tree, house}
	//舊程式碼沒有使用介面　所以有變動就要修改程式碼主體
	f := Filter{}
	for _, v := range f.filterByColor(prpducts, green) {
		fmt.Printf("%s is green\n", v.name)
	}
	//新程式碼有使用介面　所以有變動不須修改程式碼主體
	greenSpec := ColorSpecification{green}
	bf := BetterFilter{}
	for _, v := range bf.Filter(prpducts, greenSpec) {
		fmt.Println("value:", v)
	}
}
