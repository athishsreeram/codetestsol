package main

import "fmt"

type Item struct{
  ArrayIdx int
  Num int
}

func main(){
  
  a := [][]int{ 
               {2,5,7},
               {1,2,3},
              }
  
	MergeSortedArray(a)

}

func MergeSortedArray(arrays [][]int){ 

  sortedList := []int{}
  elementIdxs := make([]int,len(arrays))

for {
	smallestItems := []Item{}
	for i :=0 ; i < len(arrays); i++{
    relevantArray := arrays[i]
    elementIdx := elementIdxs[i]
    if elementIdx == len(relevantArray){
      continue
    }

    smallestItems = append(smallestItems,Item{
      ArrayIdx: i,
      Num: relevantArray[elementIdx],
    })

    if len(smallestItems) == 0 {
      break
    }

		fmt.Println(relevantArray,elementIdx,sortedList,smallestItems)
	}

break;

}
  

}
