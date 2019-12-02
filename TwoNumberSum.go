package main

import "fmt"
import "sort"

func main() {

  a := []int{10,2,-3,4,5}
  t := 9


  twoNumberSum(a,t)

}

func twoNumberSum(a []int,t int){

  v := make([]int,0,2)

  for  i := 0; i < len(a)-1; i++{
    for  j := i+1; j < len(a); j++{

       fmt.Println( a[i],a[j]);
      if a[i]+a[j] == t {

        v = append(v,a[i])
        v = append(v,a[j])
        break

      }

    }
  }

  sort.Ints(v)

  fmt.Println(v);



}