package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var scanner *bufio.Scanner = readFileAndGetScanner()
	var sumOfCaloriesOfElves []int = sumOfCaloriesOfElves(scanner)
	mostSumOfCalories(sumOfCaloriesOfElves)
	SumOfCaloriesOfTop3(sumOfCaloriesOfElves)
}

func readFileAndGetScanner() *bufio.Scanner {
	file, err := os.Open("./input.txt")
	checkIfNilError(err)
	scanner := bufio.NewScanner(file)
	return scanner
}

func sumOfCaloriesOfElves(scanner *bufio.Scanner) []int {
	var arrayOfSum []int
	var sumOfCalories int = 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			arrayOfSum = append(arrayOfSum, sumOfCalories)
			sumOfCalories = 0
		} else {
			Calories, err := strconv.Atoi(scanner.Text())
			checkIfNilError(err)
			sumOfCalories += Calories
		}
	}
	arrayOfSum = append(arrayOfSum, sumOfCalories)
	return arrayOfSum
}

func mostSumOfCalories(arrayOfSum []int) {
	var largetSumOfCalories int = 0
	for _, sumOfCalories := range arrayOfSum {
		if sumOfCalories > largetSumOfCalories {
			largetSumOfCalories = sumOfCalories
		}
	}
	fmt.Println("most Calories Elf carrying : ", largetSumOfCalories)
}

func SumOfCaloriesOfTop3(arrayOfSum []int) {
	var sumOfTop3Calories int = 0
	sort.Ints(arrayOfSum)
	for i := len(arrayOfSum) - 1; i >= len(arrayOfSum)-3; i-- {
		sumOfTop3Calories += arrayOfSum[i]
	}
	fmt.Println("sum of top 3 Calories : ", sumOfTop3Calories)
}

func checkIfNilError(err error) {
	if err != nil {
		panic(err)
	}
}
