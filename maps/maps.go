package main

import (
	"fmt"
)


func main(){

	// basicMap()
	// rangeLoopsMap()
	updateMap()
}

func updateMap(){
	treeMap := map[string]int{
		"B": 2,
		"A": 1,
		"D": 4,
		"C": 3,
	}

fmt.Printf("Before modifing, value of  Key A: [%d] \n", treeMap["A"])
treeMap["A"] = 100
fmt.Printf("After modifing, value of  Key A: [%d] \n", treeMap["A"])
fmt.Printf("Key Z value: [%d] \n", treeMap["Z"])

fmt.Println("Before deleting key D",treeMap)
delete(treeMap, "D")
fmt.Println("After deleting key D",treeMap)

}

func rangeLoopsMap(){
	treeMap := map[string]int{
		"B": 2,
		"A": 1,
		"D": 4,
		"C": 3,
	}

	fmt.Printf("\nTree map:  %v \n" , treeMap)

//rang loop, no guarantee of order

for key, val := range treeMap{
	fmt.Printf("Key is: %v Value is %v \n", key, val)
}


}
func basicMap(){
	leagueTitles := make(map[string] int)
	leagueTitles["Sunderland"] = 6
	leagueTitles["Newcastel"] = 4


	recentHead2headWins := map[string]int{
		"Sunderland": 2,
		"Newcastel": 2,
	}
	fmt.Printf("League title: %v \n Recent Head to head: %v \n", leagueTitles, recentHead2headWins)
}