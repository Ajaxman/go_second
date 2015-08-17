package main

import "fmt"

func main() {

//Estyo es una declaracion
  var arrayInt [5]int
  slice := make([]string,3)

  fmt.Println("***** My Second programa with Go ******")
  fmt.Println("This is an empty array:", arrayInt)

//Esto es una asignacion
  arrayIntShort := [2]int{1,2}
  fmt.Println("This is an array with short notation", arrayIntShort)
  println("========")


  fmt.Println("This is a slice", slice)


  slice[0] = "test"
  slice[1] = "test"
  slice[2] = "test"
  fmt.Println("This is a slice changed", slice)

  for i := 0; i<3; i++ {
    fmt.Println("2d: ", slice[i])
    if i == 2 {
        fmt.Println("The programas has been finished ")
    }
  }

}
