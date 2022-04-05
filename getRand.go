package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func gen(c chan int, maxValue int) {
	id := rand.Intn(maxValue)
	for i := 0; ; i++ {
		el := 1 + rand.Intn(maxValue)
		fmt.Println(el, id)
		c <- el
	}
}

func counter(c chan int, n int) []int {
	counter := make(map[int]int)
	arr := make([]int, 0, n)

	for i := 0; i < n; {
		el := <-c
		if counter[el] == 0 {
			counter[el] = 1
			arr = append(arr, el)
			fmt.Println(el)
			i++
		}
	}
	return arr
}

func main() {
	var n, maxValue, genN int
	flag.IntVar(&n, "n", 10, "number of random numbers")
	flag.IntVar(&genN, "gen", 2, "number of generators random numbers")
	flag.IntVar(&maxValue, "max", 10, "maximum value of random numbers")
	flag.Parse()

	if n <= 0 || maxValue <= 0 {
		fmt.Println("параметры должны быть > 0")
		return
	}

	if n > maxValue {
		fmt.Println("неврные параметры ввода")
		return
	}

	rand.Seed(time.Now().UnixNano())
	var c chan int = make(chan int)
	for i := 0; i < genN; i++ {
		go gen(c, maxValue)
	}
	arr := counter(c, n)
	sort.Ints(arr)
	fmt.Println(arr)
}

func getRandSingle() {
	var n, maxValue int
	flag.IntVar(&n, "n", 10, "number of random numbers")
	flag.IntVar(&maxValue, "max", 10, "maximum value of random numbers")
	flag.Parse()

	if n <= 0 || maxValue <= 0 {
		fmt.Println("параметры должны быть > 0")
		return
	}

	if n > maxValue {
		fmt.Println("неврные параметры ввода")
		return
	}

	rand.Seed(time.Now().UnixNano())

	counter := make(map[int]int)
	arr := make([]int, 0, n)

	for i := 0; i < n; i++ {
		f := true
		var el int

		for f {
			el = 1 + rand.Intn(maxValue)
			_, f = counter[el]
		}

		counter[el] = 1
		arr = append(arr, el)
		fmt.Println(el)
	}

	sort.Ints(arr)
	fmt.Println(arr)
}

func getRandOs() {

	n, max := 1, 10
	for i, arg := range os.Args[1:] {
		switch arg {
		case "-n":
			n, _ = strconv.Atoi(os.Args[i+2])
		case "-max":
			max, _ = strconv.Atoi(os.Args[i+2])
		}
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		fmt.Println(1 + rand.Intn(max))
	}
}
