package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    
 moveToNumberMap := map[string]int{
         "A": 1 ,
		 "B": 2 ,
		 "C": 3 ,
         "X": 1 ,
		 "Y": 2 ,
		 "Z": 3 ,
    }

 var pointsMatrix = [3][3]int{
	{4,8,3},
	{1,5,9},
	{7,2,6},
}

var pointsMatrixPart2 = [3][3]int{
	{3,4,8},
	{1,5,9},
	{2,6,7},
}

	var totalScore int = 0
	var part2totalScore int = 0


	var scanner *bufio.Scanner = readFileAndGetScanner()
	for scanner.Scan() {
		var splitedString =strings.Split(scanner.Text()," ")
		totalScore += getRoundScore(moveToNumberMap[splitedString[0]],moveToNumberMap[splitedString[1]],pointsMatrix)
		part2totalScore += getRoundScore(moveToNumberMap[splitedString[0]],moveToNumberMap[splitedString[1]],pointsMatrixPart2)
	}
	fmt.Println("Your Total Score part 1 :",totalScore)
	fmt.Println("Your Total score part 2 :",part2totalScore)


}

func readFileAndGetScanner() *bufio.Scanner {
	file, err := os.Open("./input.txt")
	checkIfNilError(err)
	scanner := bufio.NewScanner(file)
	return scanner
}

func checkIfNilError(err error) {
	
	if err != nil {
		panic(err)
	}
}

func getRoundScore(opp int,you int,pointsMatrix [3][3]int) int{
	return pointsMatrix[opp-1][you-1]
}
func get2ndRoundScore(opp int,you int,pointsMatrix [3][3]int) int{
	return pointsMatrix[opp-1][you-1]
}

// TO-DO
//

 // . points Matrix 
//             X - R -  1      Y - P - 2     Z - S - 3

//  A - R - 1  (3 + 1) = 4 .  (6 + 2) = 8    (0 + 3) = 3

//  B - P - 2   (0 + 1) = 1   (3 + 2) = 5    (6 + 3) = 9

//  C - S - 3   (6 + 1) = 7   (0 + 2) = 2    (3 + 3) = 6


 // . points Matrix part 2
//             X - R -  1      Y - P - 2     Z - S - 3

//  A - R - 1   (0 + 3) = 3 .  (3 + 1) = 4    (6 + 2) = 8

//  B - P - 2   (0 + 1) = 1   (3 + 2) = 5    (6 + 3) = 9

//  C - S - 3   (0 + 2) = 2   (3+ 3) = 6    (6 + 1) = 7



// x -  loss 
// y -  draw 
// z -  win

// A Y  
// B X
// C Z