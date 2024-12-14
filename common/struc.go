package common

import "fmt"

type GasEngine struct {
	Mpg       uint8
	Gallons   uint16
	OwnerInfo Owner
}

type electricEngine struct {
	MpKwh     uint8
	Kwh       uint16
	OwnerInfo Owner
}

type Engine interface {
	GetOwner() string
}

type Owner struct {
	Name string
}

func (T GasEngine) GetOwner() string {
	return T.OwnerInfo.Name
}

type article struct {
	title  string
	author string
}

func (a article) String() string {
	return fmt.Sprintf("%q by %s", a.title, a.author)
}

type book struct {
	title  string
	author string
	pages  int
}

func (b book) String() string {
	return fmt.Sprintf("%q by %s (%d pages)", b.title, b.author, b.pages)
}

type stringer interface {
	String() string
}

func Print(s stringer) {
	fmt.Println(s.String())
}

func StrucDemo() {
	fmt.Println("### Struc Demo ###")
	a := article{"My article", "Alex"}
	b := book{"Dear", "Mike", 100}
	fmt.Println(a.String())
	Print(a)
	Print(b)
}
