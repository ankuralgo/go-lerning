package main

import (
	"fmt"
	"time"
	"sync"
)


func main(){
	// goRoutineWithoutWaitGrp()
	goRoutineWithWaitGrp()
}



func goRoutineWithoutWaitGrp(){
 
	go func() {
	 
		fmt.Println("Sleeping for 2 sec")
		time.Sleep(2 * time.Second)
		fmt.Println("Anonymous function called after a bit of sleep")
	}()

	go func() {
		fmt.Println("Anonymous function called without wait")	
	}()
}

func goRoutineWithWaitGrp(){
	var waitGrp sync.WaitGroup
	waitGrp.Add(2)
	go func() {
		defer waitGrp.Done()
		fmt.Println("Sleeping for 2 sec")
		time.Sleep(2 * time.Second)
		fmt.Println("Anonymous function called after a bit of sleep")
	}()

	go func() {
		defer waitGrp.Done()
		fmt.Println("Anonymous function called without wait")
	}()

	waitGrp.Wait()
}

