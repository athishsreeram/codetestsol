package main

func main() {
	println(fibc(6))
}

func fibc(n int) int {

	lastTwo := []int{0, 1}

	counterVal := 3

	for counterVal <= n {
		fibNoc := lastTwo[0] + lastTwo[1]
		lastTwo = []int{lastTwo[1], fibNoc}
		counterVal++
	}
	for i, a := range lastTwo {
		println(i, a)
	}

	if n < 1 {
		return lastTwo[0]
	} else {
		return lastTwo[1]
	}

}
