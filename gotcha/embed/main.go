package main


import "fmt"

type parent struct {
	name string
}

func (p *parent) Info() string {
	return fmt.Sprintf("Name: %s", p.name)
}

type child struct {
	parent
	//other fileds
}

func getInfo(p *parent) string {
	return p.Info()
}

func main() {
	c := &child {
		parent : parent {
			name: "TheChild",
		},
		// other fields
	}

	fmt.Println(c.Info()) //child using parent method Info
	fmt.Println(c.name) //child using parent field name

	fmt.Println(getInfo(c))
}
