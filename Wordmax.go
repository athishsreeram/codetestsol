package main

import "fmt"

func main() {
	fmt.Println("Max char " + wordMax("Hello"))
}

func wordMax(input string) string {

	split := []rune(input)

	var wordsChar = make(map[string]int)

	for i, a := range split {
		var val = 0
		fmt.Printf("%d %c ", i, a)
		if wordsChar[string(a)] != 0 {
			val = wordsChar[string(a)]
			wordsChar[string(a)] = val + 1
		} else {
			wordsChar[string(a)] = 1
		}
	}
	println("")
	var max = 0
	var maxKey string
	for k, v := range wordsChar {
		if max < v {
			maxKey = k
			max = v
		}
		fmt.Println(k, v)
	}

	return maxKey
}
