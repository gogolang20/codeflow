package Command

import "fmt"

//命令模式
type Person struct {
	name string
	cmd  Command
}
type Command struct {
	person *Person
	method func()
}

func NewCommand(p *Person, method func()) Command {
	return Command{
		person: p,
		method: method,
	}
}
func (c *Command) Execute() {
	c.method()
}

func NewPerson(name string, cmd Command) Person {
	return Person{
		name: name,
		cmd:  cmd,
	}
}

func (p *Person) Buy() {
	fmt.Printf("%s is buying \n", p.name)
	p.cmd.Execute()
}

func (p *Person) Cook() {
	fmt.Printf("%s is cooking \n", p.name)
	p.cmd.Execute()
}

func (p *Person) Wash() {
	fmt.Printf("%s is washing \n", p.name)
	p.cmd.Execute()
}

func (p *Person) Listen() {
	fmt.Printf("%s is Listening \n", p.name)
}

func (p *Person) Talk() {
	fmt.Printf("%s is Talking \n", p.name)
	p.cmd.Execute()
}
