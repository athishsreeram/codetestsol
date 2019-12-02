/**1. map number to true
2. current false
    left  loop false
    right loop false
3. check max with cur keep the max array

*/

package main

import "fmt"

func main() {


  a := []int{3,2,4,5,0,7}
  b := make(map[int]bool,len(a))
  var max []int


  for i := 0; i < len(a) ; i++ {
    b[a[i]] = true
  }

  fmt.Println(b)
  fmt.Println(a)

  for i :=0; i < len(a) ; i++{

  if b[a[i]]{

    var cur []int

    fmt.Println("middle",a[i])
    cur = append(cur,a[i])
    b[a[i]]=false

    // Left check
    for j := a[i]; b[j-1] ; j--{
      fmt.Println("left",j-1)
      cur = append(cur,j-1)
      b[j-1]=false
    }


    //right check
     for j := a[i]; b[j+1] ; j++ {
       fmt.Println("right",j+1)
       cur = append(cur,j+1)
       b[j+1]=false
    }

    if len(max)<len(cur){
      max = cur
    }

  }

  }

       fmt.Println(max)
}