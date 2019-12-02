// To execute Go code, please declare a func main() in a package "main"
//remove element 7 preserve order
  
package main


func main() {
  
  a := []int{8,2,7,4,12}
  
  arrayRemoveElePreserve(a)


}

func arrayRemoveElePreserve(a []int){


  ele := 7
  remove := false
  for i := 0 ; i < len(a) - 1; i++ {
    if a[i] == ele {
      remove = true
    }
     // println(a[i])
    if remove == true{
      a[i]=a[i+1]
    }

  }

  a = a[:len(a)-1]

  for _,ak := range a{
    println(ak)
  }

}

