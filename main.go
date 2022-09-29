package main

import (
	"fmt"
)

func main(){
	arr := []int{1,2,3,4,5,6,7,8,9,10}
	size := len(arr)
	result, error := binarySearch(arr, 0, size, 1)
	if error == nil{
		fmt.Printf("index of target: %d\n", result)
	}else{
		fmt.Println(error)
	}
}

func binarySearch(arr []int, start int, end int, target int) (int, error){
	if start < end{
		var mid int = (start+end)/2
		if arr[mid] == target{
			return mid, nil
		}else if target < arr[mid]{
			return binarySearch(arr, start, mid-1, target)
		}else if target > arr[mid]{
			return binarySearch(arr, mid+1, end, target)
		}
		
	}
	return -1, fmt.Errorf("target not found")
}
