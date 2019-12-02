package main

import "fmt"
import "sort"

func main() {
  m := map[string]int{"Bob":3,"Alice":2,"Frank":1}


  fmt.Println("map:", m)

  keys := make([]string,0,len(m))

  for k := range m{
    keys = append(keys,k)
  }

  fmt.Println("Keys : ",keys)

  sort.Strings(keys)

  fmt.Println("Keys : ",keys)

  for _,k1 := range keys{
    fmt.Println("Val : ",m[k1])
  }

}