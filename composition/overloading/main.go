package main
import (
  "fmt"
)

type Person struct {
  Name string
}

type Saiyan struct {
  *Person
  Power int
}

func (s *Saiyan) Introduce() {
  fmt.Printf("Hi, I'm %s. Ya!\n", s.Name)
}

func (p *Person) Introduce() {
  fmt.Printf("Hi, I'm %s\n", p.Name)
}

func main(){
  goku := &Saiyan{
    Person: &Person{"Goku"},
    Power: 9001,
  }
  goku.Introduce()
  goku.Person.Introduce()
}
