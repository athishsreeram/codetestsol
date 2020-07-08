/**

KMP 

**/

package main

import "fmt"

func main(){

a  := []string{"a","b","c","a","b","y"}
size := len(a)-1
var pat [size]string

pat[0] = 0

for i,v := range a {

	j := i+1
	if ( v == a[j]){
		pat[j] = i + 1
	}else{
		pat[j] = 0
	}

} 


fmt.Println(pat)



}
