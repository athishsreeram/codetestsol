package main

import "fmt"

func main() {

  a := []int{10,-1,-3,4,5}
  t := 9


  twoNumberSum(a,t)

}

func twoNumberSum(a []int,t int)[]int{

  diffMap := make(map[int]int,len(a))
 
    
  for  i := 0; i <= len(a)-1; i++{
    
      diffVal := t - a[i]
      
      
       value,ok := diffMap[diffVal]
       if ok{
          v1 := []int{value,i}
           
          fmt.Println(v1)
          return v1
                
        } else {
                diffMap[a[i]] = i
        }
      
      fmt.Println(diffMap)
    
  
  }
    
   return  nil



}