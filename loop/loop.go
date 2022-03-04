package main

import (
	"fmt"
	"time"
)

func main(){
	finiteLoop()
	rangeLoop()
	infiniteLoop()
}

func rangeLoop(){
	fmt.Println("Range Loop")
	strings := []string{"hello", "world"}
	for i, s := range strings {
		fmt.Println(i, s)
	}
}

func infiniteLoop(){
	fmt.Println("Infinite Loop")
	sum := 0
	for {
		sum++ // repeated forever
		fmt.Println(sum)
	}
	fmt.Println(sum) // never reached
}

func finiteLoop(){
	fmt.Println("Finite Loop")
	for timer := 5; timer >=0; timer--{
		if timer == 0{
			fmt.Println("Boom..!!")
			break
		}
		fmt.Println(timer)
		time.Sleep(1*time.Second)
	}
}