// To execute Go code, please declare a func main() in a package "main"

package main

import "fmt"
import "sort"



func main() {

  b := []int{1,2,4,7,12,13,6,5,14,17}
  fmt.Println(b)
  subarraySort(b)

}


func subarraySort(a []int){

  var s []int
  var in []int



  for i :=0 ; i < len(a)-1 ; i++{
    if minMaxCheck(a,i)  {
      s = append(s,a[i])
    }
  }

  sort.Ints(s)

  fmt.Println(s)

  for i :=0 ; i < len(a)-1 ; i++{
    if minIndex(a,i,s[0])  {
      in = append(in,i)
      break
    }
  }

  for j := len(a)-1 ; j >= 0   ; j--{
    if maxIndex(a,j,s[len(s)-1])  {
      in = append(in,j)
      break
    }
  }


  fmt.Println(in)

}

func minMaxCheck(a []int,i int)bool{
  if i==0 {
     if a[i] > a[i+1]{
      return true
     }else{
      return false
     }
  }else if i==len(a)-1{
     if a[i-1] > a[i]{
      return true
     }else{
      return false
     }
  }else{
    if a[i-1] < a[i] && a[i] < a[i+1] {
      return false
    }else{
      return true
    }
  }

}

func minIndex(a []int,i int,s int)bool{
  if s < a[i]{
    return true
  }else {
    return false
  }
}
func maxIndex(a []int,j int,s int)bool{
  if s > a[j]{
    return true
  }else{
    return false
  }
}