// To execute Go code, please declare a func main() in a package "main"

package main

import "fmt"

type BST struct{
 val int
 left *BST
 right *BST
}

func main() {
  
  t := AddBST(nil,1)
  AddBST(t,2)
  AddBST(t,3)
  AddBST(t,4)
  
  
  
  preOrder(t)
  
}

func AddBST(t *BST,v int) *BST{

  if(t == nil){
    return &BST{val:v}
  }
  
  if(t.val < v){
    if(t.left == nil){
      t.left = &BST{val:v}
    }else{
      AddBST(t.left,v)
    }
  }else if (t.val > v){
    if(t.right == nil){
       t.right = &BST{val:v}
    }else{
      AddBST(t.right,v)
    }
  }
  
  return t  

}

func preOrder(t *BST){
 
  if(t == nil){
    return 
  }
  
    preOrder(t.left);
    preOrder(t.right);
  fmt.Println(t.val);  
  
  
  
}