package main

import "fmt"

//각 팀별로 경기

func main() {
	//팀 개수 N
	//리그경기의 결과 입력 팀명1 세트승수 팀명2 세트승수로 입력
	//경기 수 N * N-1
	var n int
	var input string
	fmt.Scan(&n)
	for i := 0; i < 4*n*(n-1); i++ {
		fmt.Scan(&input)
	}
}
