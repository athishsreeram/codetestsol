package main

import "fmt"
import "sort"

func main() {

  a := []int{10,2,-3,4,5,0}
  t := 9


  threeNumberSum(a,t)

}

func threeNumberSum(a []int,t int){


  var myslice [][]int

  for  i := 0; i < len(a)-2; i++{
    for  j := i+1; j < len(a)-1; j++{
      for  k := j+1; k < len(a); k++{

       fmt.Println( a[i],a[j],a[k]);
          if a[i]+a[j]+a[k] == t {
            v := make([]int,0,2)

            v = append(v,a[i])
            v = append(v,a[j])
            v = append(v,a[k])

            sort.Ints(v)
            myslice = append(myslice,v)

          }
      }
    }
  }



  fmt.Println(myslice);



}