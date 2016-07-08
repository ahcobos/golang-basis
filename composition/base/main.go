package main
import (
  "fmt"
)

type Person struct {
  Name string
}

type Stupid struct {
  Profession string
}

func (s *Stupid) Introduce() {
  fmt.Printf("Hi, I'm a %s\n", s.Profession)
}

func (p *Person) Introduce() {
  fmt.Printf("Hi, I'm %s\n", p.Name)
}
type Saiyan struct {
  *Person
  Power int//cant add a Stupid structure it's an "ambiguous selector goku.Introduce"
}

func main(){
  goku := &Saiyan{
    Person: &Person{"Goku"},
    Power: 9001,
  }
  goku.Introduce()
  goku.Introduce()
}
