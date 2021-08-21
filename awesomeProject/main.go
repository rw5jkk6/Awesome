package main

import (
	"fmt"
   "./mylib/calc"
   car  "./mylib/calc_array"
	
)

func main(){
   a := calc.Add(2, 3)
   fmt.Println(a)

   b := calc.Sub(3, 1)
   fmt.Println(b)

   ar :=[]float64{1.2, 2.3, 3.0}
   r := car.Avg(ar)
   fmt.Println(r)
   
}