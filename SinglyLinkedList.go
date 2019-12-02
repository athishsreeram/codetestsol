package main

import "fmt"


type SinglyLinkedList struct{
  val int
  next *SinglyLinkedList
}

func main() {
  
  var lst SinglyLinkedList
 
  lst.addNode(10)
  lst.addNode(12)
  
  lst.printLst()
  
  
}


func (lst *SinglyLinkedList) addNode(v int) *SinglyLinkedList{
  
    fmt.Println(lst.val)
    fmt.Println(lst.next)
    fmt.Println(v)
    fmt.Println(" ")

    if lst.val == 0 && lst.next == nil {
       lst.val = v
       return lst
    }
   
  
    for lst.next != nil {
      lst = lst.next
    }
    if lst.val != 0 && lst.next == nil {
       lst.next = &SinglyLinkedList{val:v}
    }
   
    return lst
  
  
}

func (lst *SinglyLinkedList) printLst() {
    for lst.next != nil{
      fmt.Println(lst.val)
      lst = lst.next
    }
    fmt.Println(lst.val)
  
}
