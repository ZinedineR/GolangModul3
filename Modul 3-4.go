package main
 
import (
    "fmt"
    humanize "github.com/dustin/go-humanize"
    "strings"
    "math"
)
var idr, usd float64
//I'm using package humanize to use its Comma function, it's good to separate the thousands.
func FormatRupiah(amount float64) string {
	humanizeValue := humanize.CommafWithDigits(math.Round(amount),2) //I'm also using math.Round to give better calculation, this round the float to the nearest integer
	stringValue := strings.Replace(humanizeValue, ",", ".", -1) //This to change the strings of , to ., with unlimited search
  return stringValue}

func FormatDollar(amount float64) string {
	humanizeValue := humanize.CommafWithDigits(amount,2)
	stringValue := strings.Replace(humanizeValue, ",", ".", -1)
  return stringValue}

func main(){
  fmt.Println("Input your numeric : ")
  fmt.Scanf("%f",&idr)
  usd = idr * 0.000071
  fmt.Printf("%f\n",usd)
fmt.Print("IDR : Rp. ", FormatRupiah(idr),",00\n")
fmt.Println("USD : $ ", FormatDollar(usd))
}
