package main

import "fmt"



   /** 3 variable loops k - i - j
    i=1;i<n-1;
    j=i+1;j<n
        diff = fournumsum - sumofIJ
        if(pairMap.has(diff))
            newfour = k i i j
    k=0 ; k<i ; pairMap<pairsumIK,[][]sumofIK)
   */

func main(){
    input := {1,2,3,5,6,7,0}
    fournumsum := 11

    fmt.Println(fourNumSum(input, fournumsum))
}


func fourNumSum(input []int, fournumsum int) [][]int {
	allPairSums := map[int][][]int{}
	fournum := [][]int{}
	for i := 1; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			currentSum := input[i] + input[j]
			difference := fournumsum - currentSum
			if pairs, found := allPairSums[difference]; found {
				for _, pair := range pairs {
					newfournum := append(pair, input[i], input[j])
					fournum = append(fournum, newfournum)
				}
			}
		}
		for k := 0; k < i; k++ {
			currentSum := input[i] + input[k]
			allPairSums[currentSum] = append(allPairSums[currentSum], []int{input[k], input[i]})
		}
	}
	return fournum
}