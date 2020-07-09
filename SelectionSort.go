package main

import "fmt"

func main(){
  
  a := []int{ 3,8,2,1,5}
  
  
  for i:=0 ; i < len(a);  i++{
    small := i
    for j:=i+1; j < len(a) ; j++{ 
      if(a[small] > a[j]){
        small = j
      }
    }
    fmt.Println(i,small)
    swap(a,i,small)
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