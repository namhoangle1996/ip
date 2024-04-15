package main

import "fmt"

func main() {
	arr := []string{"banana1", "banana", "banana2", "orange", "pineapple", "strawberry"} // Mảng các chuỗi chưa được sắp xếp
	target := "banana2"                                                                  // Chuỗi cần tìm

	// Thực hiện tìm kiếm nhị phân
	result := bsearch(arr, target)
	if result != -1 {
		fmt.Printf("Chuỗi %s được tìm thấy tại vị trí %d trong mảng.\n", target, result)
	} else {
		fmt.Printf("Chuỗi %s không được tìm thấy trong mảng.\n", target)
	}
}

func bsearch(arr []string, target string) int {

	var x string
	var y string
	x = "yrak"
	y = "yran"

	fmt.Println("get equal", x > y)

	left, right := 0, len(arr)-1
	for left <= right {
		mid := (right + left) / 2
		fmt.Println("get mid", mid)
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			fmt.Println("<<<<")
			fmt.Println("mid value", arr[mid])
			fmt.Println("target", target)
			left = mid + 1
		} else {
			fmt.Println(">>>>>")
			fmt.Println("mid value", arr[mid])
			fmt.Println("target", target)
			right = mid - 1
		}
	}
	return -1
}

func calculateArea(low, high, first, second int) int {
	if second > first {
		return first * (high - low)
	} else {
		return second * (high - low)
	}
}

func maxArea(height []int) int {
	low_pointer := 0
	high_pointer := len(height) - 1
	max_area := 0
	for low_pointer < high_pointer {
		newSum := calculateArea(low_pointer, high_pointer, height[low_pointer], height[high_pointer])
		fmt.Println("newSum", newSum)
		if newSum > max_area {
			max_area = newSum
		}
		if height[low_pointer] < height[high_pointer] {
			low_pointer++
		} else {
			high_pointer--
		}

	}
	return max_area
}

func removeKdigits(num string, k int) string {
	stack := make([]int32, 0, len(num))

	for _, v := range num {
		for k > 0 && len(stack) > 0 && v < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		if len(stack) > 0 || v != '0' {
			stack = append(stack, v)
		}

	}

	if stack = stack[:len(stack)-min(k, len(stack))]; len(stack) == 0 {
		return "0"
	}

	fmt.Println("get stack", stack)
	return string(stack)
}
