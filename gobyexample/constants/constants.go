package main

import "fmt"
import "math"

func main(){
const n = 500000000
const d=3e20/n
fmt.Println(int64(d))
fmt.Println(math.Sin(n))
}
