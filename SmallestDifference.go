
package main

import "fmt"
import "math"

func main() {

  a := []int{10,2,3,4,5,0}
  b := []int{11,12,13,12,45,20}
  
  
  smallestDiff(a,b)
  
}

func smallestDiff(a []int,b []int){

  var minDiff int
  var myslice []int
        
  
  minDiff = math.MaxInt64 

  for  i := 0; i < len(a)-1; i++{
    for  j := 0; j < len(b)-1; j++{
      
      var absDiff int
       
      if(a[i] > b[j]){
        absDiff = a[i] - b[j]
      }else{
        absDiff = b[j] - a[i]
      }
      
      
      if minDiff > absDiff {
        myslice = []int{a[i],b[j]}
        minDiff = absDiff
      }
      
    }
  }
  
  
  
  fmt.Println(myslice);
  fmt.Println(minDiff);
  
  
  
}


