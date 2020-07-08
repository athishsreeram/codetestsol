// Simple match

// loop small string
// each small string loop through big string
// check on lenght on big and small string

package main


import "fmt"

func main(){
  
  big := "This is abt big"
  small := []string {"abt","is"}
  
  out := make([]bool,len(small))

  
  for smallIndex,v := range small{
    out[smallIndex]=isSmallinBig(big,v)
  }
  
  fmt.Println(out)
    
}

func isSmallinBig(big string,small string) bool{
  
  match := true
  
  for indexBig,_  := range big{
    
    
    if indexBig+len(small) > len(big){
      break
    }
    
    match = isBigSmallMatch(big,small,indexBig)
    
    if match{
      return true
    }
   
  }
  
  return match

}

func isBigSmallMatch(big string,small string,i int) bool{
  
  fmt.Println(small,big,i)
  
  leftS := 0
  leftB := i
  
  rightS := len(small)-1
  rightB := i + len(small)-1
  
  match := true
  
  for leftB <= rightB {
    
    fmt.Println(small[leftS],big[leftB],small[rightS],big[rightB],leftB,rightB,leftS,rightS)
    if(small[leftS] != big[leftB] || small[rightS] != big[rightB]){
      match = false
    }else{
      match=true
      if(leftB == rightB){
        return match
      }
    }
    
    leftS += 1;
    leftB += 1;
    rightB  -= 1;
    rightS  -= 1;
    
    
  }
  
  
  
  return match
}



