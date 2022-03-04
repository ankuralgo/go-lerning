package main
import (
	"fmt"
)

func main(){

	type courseMeta struct{

		author string
		level string
		rating float64
	}

// Only declare struct variable
	// var gettingStartedWithK8s courseMeta
	// gettingStartedWithK8s := new (courseMeta)

	gettingStartedWithK8s := courseMeta{
		author : "Nigel Poulton",
		level : "Beginner",
		rating : 5,
	}

// Declare and initialie struct
	fmt.Println(gettingStartedWithK8s)
	fmt.Printf("Author of Getting started with K8s is %s and rating is %.1f \n", gettingStartedWithK8s.author, gettingStartedWithK8s.rating)
	gettingStartedWithK8s.rating = 7.5
	fmt.Printf("Author of Getting started with K8s is %s and modified rating is %.1f \n", gettingStartedWithK8s.author, gettingStartedWithK8s.rating)


}