package main
import (
  "fmt"
)

func main(){
  scores := [4]int{9001, 9333, 212, 33}
  for index, value := range scores {
    fmt.Printf("value %d at pos %d\n",value,index)
  }
}
