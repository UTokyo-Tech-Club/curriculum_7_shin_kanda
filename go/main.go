package main
import (
	"fmt"
	"sort"
	"strings"
	"strconv"
)

func convertStringToIntArray(m string) []int {
	parts := strings.Split(m, ",")
	var result []int
	for _, part := range parts {
		num, err := strconv.Atoi(strings.TrimSpace(part))
		if err == nil {
			result = append(result, num)
		}
	}
	return result
}

func countNumberFrequency(a []int) map[int]int {
	freq := make(map[int]int)
	for _, num := range a {
		freq[num]++
	}
	return freq
}
	
func countCardCombinations(a []int, s int) int {
	dp := make([]int, s+1)
	dp[0] = 1
	for _, num := range a {
		for i := s; i >= num; i-- {
			dp[i] += dp[i-num]
		}
	}

	if dp[s] == 0 {
		return -1
	}
	return dp[s]
}

func printMapKeyAndValue(m map[int]int) {
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Println(k, m[k])
	}
}


func main(){
	var m string
	var s int

	fmt.Scan(&m)
	fmt.Scan(&s)

	a := convertStringToIntArray(m)
	frequencyCount := countNumberFrequency(a)
	combinationCount := countCardCombinations(a, s)

	printMapKeyAndValue(frequencyCount)
	fmt.Println(combinationCount)
}