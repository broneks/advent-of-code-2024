package day5

import "fmt"

func getCorrectedPageUpdateMiddleValue() {

}

func SolutionPart2() {
	var result int

	pageUpdates, pageOrderingRules := scanInput()

	for _, pageUpdate := range pageUpdates {
		result += getCorrectedPageUpdateMiddleValue(pageUpdate, pageOrderingRules)
	}

	fmt.Println(result)
}
