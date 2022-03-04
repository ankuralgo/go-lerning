package main

import (
	"fmt"
	"os"
)

func main(){
	
	fmt.Println("1------------->")
	passengerCatogry("Ram",21)
	passengerCatogry("Vivek",15)
	passengerCatogry("Dev",4)
	passengerCatogry("Sam",2)
	fmt.Println("2------------->")
	handLuggageScan(2)
	handLuggageScan(20)
	fmt.Println("3------------->")
	identifyFlightByCode("AI")
	identifyFlightByCode("G8")
	identifyFlightByCode("69")
	fmt.Println("4------------->")
	codeTranslation("FAIL")
	codeTranslation("FAILED")
	codeTranslation("PASS")
	fmt.Println("5------------->")
	errorHandling()
	fmt.Println("-----end-------")
	 
}

func passengerCatogry(name string, age int) {
 

	if age >= 18 {
		fmt.Println(name," is adult")
	} else if age <18 && age >= 12 {
		fmt.Println(name, " is teen")
	} else if age <12 && age >= 5 {
		fmt.Println(name, " is gradeschooler")
	} else if age <5 && age >= 3 {
			fmt.Println(name, " is preschool")
	} else if age <3 && age >= 1 {
			fmt.Println(name, " is toddler")
	} else {
		fmt.Println(name, " is baby")
	}
}

// Example of simple initializtion
func handLuggageScan(luggageWeight int){

	if maxHandLuggageWeight := 10; luggageWeight > maxHandLuggageWeight{
		fmt.Println("Hand luggage failed weight scan")
	} else{
		fmt.Println("Hand luggage weight scan pass")
	}
}

// Example of switch case 
func identifyFlightByCode(flightCode string){
	switch flightCode {
	case "AI" :
		fmt.Println("Air India")
	case "I5" :
		fmt.Println("AirAsia India")
	case "G8" :
		fmt.Println("Go First")
	case "6E" :
		fmt.Println("IndiGo")		
	default :
		fmt.Println("Invalid code")

	}
}

// Example of switch case with fallthrough
func codeTranslation(code string){
	
	switch code {
	case "FAIL" :
		fallthrough
	case "FAILED":
		fallthrough
	case "UNSUCCESSFUL" :
		fallthrough
	case "FATAL":
		fmt.Println("ERROR")
	case "PASS": 
		fallthrough
	case "SUCCESSFUL":
		fmt.Println("Successful")
	default :
		fmt.Println("Invalid code")

	}
// Example of error handling
}
func errorHandling() {
	_, error := os.Open(".tmp.txt")
	if error != nil {
		fmt.Println("Error occured while file reading ", error)
	}
}

