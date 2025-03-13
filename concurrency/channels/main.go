package main

func main1() {
	//var intCh chan int = make(chan int)
	//strCh := make(chan string)
	//fmt.Println(intCh)
	//fmt.Println(strCh)

	//intCh := make(chan *int)
	//v := 5
	//go func(v int) {
	//	fmt.Println("Go routines starts")
	//	fmt.Scanln()
	//	intCh <- &v
	//	v += 1
	//}(v)
	//fmt.Scanln()
	//fmt.Println(*<-intCh)
	//fmt.Println("The end", "v =", v)
	//intCh := make(chan int)

	//go factorial(5, intCh)
	//fmt.Println(<-intCh)
	//fmt.Println("The end")
}


//func factorial(n int, ch chan int) {
//	result := 1
//	for i := 1; i <= n; i++ {
//		result *= i
//	}
//	fmt.Println(n, "-", result)
//	ch <- result
//}