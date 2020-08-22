package main

import (
	"exercise/dataStruct"
	"fmt"
)

func main(){
	fmt.Println("abcde = ", dataStruct.Hash("abcde"))
	fmt.Println("abcde = ", dataStruct.Hash("abcde"))
	fmt.Println("abcdf = ", dataStruct.Hash("abcdf"))
	fmt.Println("tbcde = ", dataStruct.Hash("tbcde"))
	fmt.Println("tbcdasdadsade = ", dataStruct.Hash("tbcdasdadsade"))
}
