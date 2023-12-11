package main

import ("fmt")

func main() {
	//Day1()
	//Day2()
	//Day3()
	//Day4()
	//Day5()
	//Day6()
	//Day7()
	//Day8()
	//Day9()
	//Day10()
	//Day11()
	//Day12()
	//Day13()
	//Day14()
	//Day15()
	//Day16()
	//Day17()
	//Day17other()
	//Day18()
	//Day19()
	//Day20()
	//Day21()
	//Day22()
	//Day22Other()
	//Day23()
	//Day24()
	//Day25()

	arr := []int{0,1,2,3,6,8,9,14,17}
	fmt.Println(bin_search(arr, 0))
	fmt.Println(bin_search(arr, 17))
	
	
}

func bin_search(arr []int, target int) bool {

	
	lo := 0
	hi := len(arr)
	
	for lo < hi {
		mid := lo + (hi - lo) / 2
	  	if arr[mid] == target {
			return true
		} else if arr[mid] > target {
			hi = mid		
		} else {
			lo = mid + 1
		}
	}

	return false

}
