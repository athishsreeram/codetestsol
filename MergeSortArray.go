package main

import "fmt"


// Wont work for 0 in the array

func main(){
  
  a := [][]int{ 
               {3,8,2,1,5},
               {3,2},
              }
  
  size := 0
  
  for i:=0 ; i < len(a);  i++{
    size = size + len(a[i])
  }
  
  
  o := make([]int,size)
 
  for i:=0 ; i < len(a) ;  i++{
    for j :=0 ; j < len(a[i]) ; j++ {
      fmt.Println(a[i][j])
      for k :=0; k < len(o); j++{    
        
        v := a[i][j]
        if  v <= o[k]{
          fmt.Println("match");
          break;
        }
      }
      
    }
    
  }  

}
