package main
 
import (
    "fmt"
)
const Kelvin = 273.15
var Celcius, Fahrenheit float64

func main(){
  fmt.Println("Input your celcius degree : ")
  fmt.Scanf("%f", &Celcius)
  Fahrenheit = (Celcius*9/5)+32
  CeltoKel := Celcius+Kelvin
  fmt.Println("Your celcius degree : ", Celcius, "°C")
  fmt.Println("In Fahrenheit degree : ", Fahrenheit, "°F")
  fmt.Println("In Kelvin degree : ", CeltoKel, "°K")
}