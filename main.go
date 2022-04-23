package main
import ("fmt")

type person struct {
  FirstName string `json:"firstName"`
  LastName  string `json:"lastName"`
  Age       int    `json:"age"`
}

var people = []person{
  {FirstName: "Gabriel", LastName: "Ford", Age: 5},
  {FirstName: "Yaya", LastName: "DaGiwlfwend", Age: 5},
  {FirstName: "Gabriela", LastName: "Ford", Age: 5},
  {FirstName: "Yayoid", LastName: "DaGiwlfwend", Age: 5},
}

func main() {
  for itemIndex, item := range people {
    fmt.Print(item.FirstName + " ")
    fmt.Println(item.LastName)
    fmt.Println(itemIndex)
    fmt.Println(item)
  }
}
