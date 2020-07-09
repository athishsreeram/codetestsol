package main

import "fmt"

func main(){
  
  a := []int{ 3,8,2,1,5}
  
  
  for i:=0 ; i < len(a)-1;  i++{
    for j:=0; j < len(a)-i-1; j++{ 
      if(a[j] > a[j+1]){
        swap(a,j,j+1)
      }
    }
  
  }
  
  for _,v := range a{
    fmt.Println(v)
  }

}

func swap(a []int,i int,j int){
  t := a[i]
  a[i] = a[j]
  a[j] = t
}