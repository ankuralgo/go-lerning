package main

import (
	"fmt"
	"strings"
)


func main(){

	author := "Sudha murthy"
	book := "Grandma's bag of stories"
	
	fmt.Println(  convert(author, book))

	fmt.Println("Top score: ", topScore(1,3,12,43,56,78,98,43,122));
	fmt.Println("Top score: ", topScore(123,43,129,876,2562,78,98,43,122));

	
}


// Method with definate input parameters an return type.
// This xgf pass by value
func convert(author, book string) (a,b string){
	
	author = strings.ToUpper(author)
	book = strings.Title(book)
	return author, book
}


//Variadic functions: which accepts variable number of inputs

func topScore(scores ... int) int{

	topScore := scores[0]
	for _, i := range scores{
		if i > topScore {
			topScore = i
		}
	}
	return topScore
}