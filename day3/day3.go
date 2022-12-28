package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var priorities = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
	
		"A": 27,
		"B": 28,
		"C": 29,
		"D": 30,
		"E": 31,
		"F": 32,
		"G": 33,
		"H": 34,
		"I": 35,
		"J": 36,
		"K": 37,
		"L": 38,
		"M": 39,
		"N": 40,
		"O": 41,
		"P": 42,
		"Q": 43,
		"R": 44,
		"S": 45,
		"T": 46,
		"U": 47,
		"V": 48,
		"W": 49,
		"X": 50,
		"Y": 51,
		"Z": 52,
	}

    
//  moveToNumberMap := map[string]int{
//          "A": 1 ,
// 		 "B": 2 ,
// 		 "C": 3 ,
//          "X": 1 ,
// 		 "Y": 2 ,
// 		 "Z": 3 ,
//     }

//  var pointsMatrix = [3][3]int{
// 	{4,8,3},
// 	{1,5,9},
// 	{7,2,6},
// }

// var pointsMatrixPart2 = [3][3]int{
// 	{3,4,8},
// 	{1,5,9},
// 	{2,6,7},
// }

// 	var totalScore int = 0
// 	var part2totalScore int = 
	var scanner *bufio.Scanner = readFileAndGetScanner()
	for scanner.Scan() {
		var splitedString =strings.Split(scanner.Text()," ")
		fmt.Println(splitedString)
	}
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