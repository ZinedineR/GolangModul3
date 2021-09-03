package main
 
import (
    "fmt"
)
var width, length, area, perimeter float32

func main(){
  fmt.Println("Input your length : ")
  fmt.Scanf("%f",&length)
  fmt.Println("Input your width : ")
  fmt.Scanf("%f",&width)
  area = length * width
  perimeter = 2*(length+width)
  fmt.Printf("Rectangle's area is  %.2f \n", area)
  fmt.Printf("Rectangle's perimeter is %.2f", perimeter)
}
