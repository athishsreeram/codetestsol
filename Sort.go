package main

import "fmt"

func main() {

  arr := []int{10,2,6,5}


  fmt.Println(bubbleSort(arr))

  arr1 := []int{10,2,6,5}

  fmt.Println(insertSort(arr1))
}

// bubble  ---->

func bubbleSort(a []int)[]int{

  for  i:= 0  ; i < len(a)-2 ; i++{
    for j:= i  ; j < len(a)-1 ; j++{
      if(a[j] > a[j+1]){
        var temp = a[j+1]
        a[j+1] = a [j]
        a[j] = temp
      }

    }
  }

  return a

}

// sorted array | unsorted array

func insertSort(a []int)[]int{

  for  i:= 0  ; i < len(a) ; i++{
    var key = a[i]
    var j = i-1
    for  ; j>=0 && a[j] > key ;j-- {
      a[j + 1] = a[j];
    }
    a[j + 1] = key;
  }

  return a

}

// Selection sort