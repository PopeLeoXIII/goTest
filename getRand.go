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

// gen генерирует случайные числа
func gen(c chan<- int, maxValue int) {
	//id := rand.Intn(maxValue)
	for {
		randomNumber := 1 + rand.Intn(maxValue)
		//fmt.Println(el, id)
		c <- randomNumber
	}
}

// counter следит за уникальность сгенерированных чисел
func counter(c <-chan int, n int) []int {
	counter := make(map[int]int)
	result := make([]int, 0, n)

	for i := 0; i < n; {
		randomNumber := <-c
		if counter[randomNumber] == 0 {
			counter[randomNumber] = 1
			result = append(result, randomNumber)
			//fmt.Println(el)
			i++
		}
	}
	return result
}

type appContext struct {
	n, maxValue, genN, cb int
}

func parseFlags() appContext {
	//var n, maxValue, genN, cb int
	context := appContext{}

	flag.IntVar(&context.n, "n", 10, "number of random numbers")
	flag.IntVar(&context.cb, "cb", 1, "chenal buf")
	flag.IntVar(&context.genN, "gen", 1, "number of generators random numbers")
	flag.IntVar(&context.maxValue, "max", 10, "maximum value of random numbers")
	flag.Parse()

	if context.n <= 0 || context.maxValue <= 0 ||
		context.cb <= 0 || context.genN <= 0 {
		fmt.Println("параметры должны быть > 0")
		os.Exit(1)
	}

	if context.n > context.maxValue {
		fmt.Println("неврные параметры ввода")
		os.Exit(1)
	}
	return context
}

func main() {

	context := parseFlags()

	rand.Seed(time.Now().UnixNano())
	var c chan int = make(chan int, context.cb)

	t := time.Now().UnixNano()

	for i := 0; i < context.genN; i++ {
		go gen(c, context.maxValue)
	}

	arr := counter(c, context.n)
	sort.Ints(arr)
	t = time.Now().UnixNano() - t
	fmt.Println("вермя выполнения ", t/100000)
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
