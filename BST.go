// To execute Go code, please declare a func main() in a package "main"

package main

import "fmt"

type BST struct{
  val int
  left *BST
  right *BST
}

var maxSum int


func main() {

  t :=  &BST{val:5}
  t.AddBST(6)
  t.AddBST(2)
  t.AddBST(3)
  t.AddBST(4)
  t.AddBST(1)

  maxSum = 0

  fmt.Println(t.maxPathSum())
  fmt.Println(maxSum)


  fmt.Println(t.sumExist(11))

}


func (t *BST) maxPathSum()int{

  if t == nil{
    return 0
  }

  left := Max(0,t.left.maxPathSum())
  right := Max(0,t.right.maxPathSum())
  maxSum = Max(maxSum,left+right+t.val)

  return Max(left,right)+t.val


}

func (t *BST) sumExist(sum int)bool{

  if t == nil{
    return false
  }else if( t.left == nil && t.right == nil && sum - t.val ==0){
    return true
  }else{
    return t.left.sumExist(sum-t.val) || t.right.sumExist(sum-t.val)
  }


}


func Max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

// Min returns the smaller of x or y.
func Min(x, y int) int {
    if x > y {
        return y
    }
    return x
}



func (t *BST) AddBST(v int) *BST{
    if v < t.val {
      if t.left == nil{
        var left = &BST{val:v}
        t.left = left
      }else{

        t.left.AddBST(v)
      }
    }else if v > t.val{
      if t.right == nil{
        var right = &BST{val:v}
        t.right = right
      }else{
        t.right.AddBST(v)
      }
    }

  if t == nil{
    return &BST{val:v}
  }
   return t
}