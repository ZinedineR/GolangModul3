package main
 
import (
    "fmt"
)
var num int

//this function to create a series of the sequence, or sum the sequence. It works basically by looping every element, then adding it each by
func sum(array []int) int{
  res := 0
  for _, i := range array{res+=i}
  return res
}

//this function generates a sequence by int input. Using looping, I also use if to determine if the number modulo by 2 = 0, then that's an even number
func even(num int) []int{
  evenseq := []int{}
  for i:=0; i<num;i++{
    result := i % 2
    if result == 0{
      evenseq = append(evenseq,int(i))
    }
  }
return evenseq
}

//this function generates a sequence by int input. Using looping, I also use if to determine if the number modulo by 2 != 0, then that's an uneven number
func uneven(num int) []int{
  unevenseq := []int{}
  for i:=0; i<num;i++{
    result := i % 2
    if result != 0{
      unevenseq = append(unevenseq,int(i))
    }
  }
return unevenseq
}

func main(){
  fmt.Println("Input how many sequence that you wish : ")
  fmt.Scanf("%d", &num)
  fmt.Println("Your even sequence = ",even(num))
  fmt.Println("Sum : ", sum(even(num)))
  fmt.Println("Your uneven sequence = ",uneven(num))
  fmt.Println("Sum : ", sum(uneven(num)))
}