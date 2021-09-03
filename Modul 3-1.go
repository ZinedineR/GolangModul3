package main
 
import (
    "fmt"
    "bufio"
    "os"
    "strings"
)
//this function to detect duplicates. Basically using two loops, the first loop is to pick the character, then using the inner loop to loop the others characters, if detected. Count = 2, and if count > 1, it appends and array to return as function 
func detect(input string, n int) []string {
  a := []string{}
    for i := 0; i<n; i++{
      count := 1
      for j := i+1;j<n; j++{
      if input [i] == input[j]{
          count += 1
          }
      if count > 1{
        a = append(a, string(input [i]))
        j=n
      }
      }  
}
return a
}
//I'm using bufio and strings package to read " " or space. Because fmt package doesn't recognize space character. 
func main() {
    consoleReader := bufio.NewReader(os.Stdin)
    fmt.Println("Input some string :  ")
    input, _ := consoleReader.ReadString('\n')
    input = strings.TrimSuffix(input, "\n")
    var n int = len(input)
  fmt.Println("Sum of counted strings : ", n)
  dups := detect(input, n)
  fmt.Println("List(s) of duplicate characters detected : ",strings.Join(dups,", "))
              
}