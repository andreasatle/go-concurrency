package main

import "fmt"

func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range in {
			out <- num * num
		}
		close(out)
	}()
	return out

}

func main() {
	nums := []int{1, 2, 3, 5, 8, 13, 21}
	squareV1(nums...)
	squareV2(nums...)
	squareSquare(nums...)
}

func squareV1(nums ...int) {
	ch := generator(nums...)
	out := square(ch)
	for num := range out {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}

func squareV2(nums ...int) {
	for num := range square(generator(nums...)) {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}

func squareSquare(nums ...int) {
	for num := range square(square(generator(nums...))) {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}
