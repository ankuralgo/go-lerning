package main

import (
	"fmt"
)

func main(){
	// sliceLengthCapacity()
	// fetchData()
	// playWithSlice()
	appendSlice()
}


func appendSlice() {
	mySlice := []int{1,2,3}
	fmt.Printf("Initial lenght of slice is %d and capacity is %d \n", len(mySlice), cap(mySlice))

	for i := 1; i<=6; i++ {
		mySlice = append(mySlice, i)
		fmt.Printf("Modified lenght of slice is %d and capacity is %d \n", len(mySlice), cap(mySlice))
	}

	mySlice = []int{1,2,3, 4, 5}
	newSlice := []int{11,22,33}
	mySlice = append(mySlice, newSlice...)
	fmt.Printf("Append new slice %d has new size %d and capacity %d \n", mySlice, len(mySlice), cap(mySlice) )

}
func playWithSlice() {
	mySlice := []int{1,2,3,4,5}
	fmt.Printf("Element at index 0 is [%d] \n", mySlice[0])
	fmt.Printf("Element at index 2 is [%d] \n", mySlice[2])

	fmt.Println("Before changing value in slice")
	fmt.Println(mySlice)
	fmt.Println("After changing value in slice")
	mySlice[2] = 0
	fmt.Println(mySlice)


	sliceOfSlice := mySlice[2:5]
	fmt.Println("Slice of slice mySlice[2:5] ",sliceOfSlice)

	sliceOfSlice = mySlice[:3]
	fmt.Println("Slice of slice mySlice[:3] ",sliceOfSlice)

	sliceOfSlice = mySlice[2:]
	fmt.Println("Slice of slice mySlice[2:] ",sliceOfSlice)
	
}

func fetchData(){
	mySlice := []int{10,11,12,13,14,15}
	fmt.Printf("Element at index 0 is [%d] \n", mySlice[0])
}

func sliceLengthCapacity(){
	courses := make([]string, 5, 10)

	courses[0] ="Go lang for beginners"
	courses[1] ="JAVA in action"
	courses[2] ="Docker in action"


	fmt.Printf("Lenght of slice is %d and capacity is %d \n", len(courses), cap(courses))
	fmt.Println("List of courses:")
	for _,course := range courses {
		fmt.Println(course)
	}


	newCourses := []string{"Go lang for beginners", "JAVA in action", "Docker in action"}
	fmt.Printf("Lenght of slice is %d and capacity is %d \n", len(newCourses), cap(newCourses))
	fmt.Println("List of newCourses:")
	for _,course := range newCourses {
		fmt.Println(course)
	}
}