package main

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"sort"
	"time"
)

const (
	LIST_NUMS  = 100000 // 序列大小
	TEST_TIMES = 3      // 测试次数
	DEBUG      = true   // 调试
)

func main() {
	// 定义测试使用的算法
	algorithms := [8]string{"insert", "shell", "bubble", "quick", "select", "heap", "merge", "radix"}
	// 生成测试用数据
	list := generalList()
	// 定义正确结果
	var listCal []int
	var listCorrect []int
	var timeStart, timeEnd time.Time

	// 执行测试
	for _, algorithm := range algorithms {
		fmt.Println(algorithm)
		for t := 0; t < TEST_TIMES; t++ {
			timeStart = time.Now()
			listTmp := make([]int, len(list))
			copy(listTmp, list)
			switch algorithm {
			case "insert":
				listCal = insert_sort(listTmp)
			case "shell":
				listCal = shell_sort(listTmp)
			case "bubble":
				listCal = bubble_sort(listTmp)
			case "quick":
				listCal = quick_sort(listTmp, 0, len(listTmp)-1)
			case "select":
				listCal = select_sort(listTmp)
			case "heap":
				listCal = heap_sort(listTmp)
			case "merge":
				listCal = merge_sort(listTmp)
			case "radix":
				listCal = radix_sort(listTmp)
			default:
				sort.Ints(listTmp)
				listCal = listTmp
			}
			timeEnd = time.Now()
			if listCorrect == nil {
				listCorrect = listCal[:]
			}
			if reflect.DeepEqual(listCorrect, listCal) == false && DEBUG == true {
				fmt.Println("Origin:  ", list)
				fmt.Println("Error:   ", listCal)
				fmt.Println("Correct: ", listCorrect)
			}
			fmt.Printf("time %3d  %s \n", t, timeEnd.Sub(timeStart).String())
		}
	}
	fmt.Println("done.")
}

// 生成测试序列
func generalList() []int {
	list := make([]int, LIST_NUMS)
	listMap := make(map[int]int)
	for len(listMap) < LIST_NUMS {
		listMap[rand.Intn(LIST_NUMS*2)] = 1
	}
	i := 0
	for k := range listMap {
		list[i] = k
		i = i + 1
	}
	return list
}

// 插入排序
func insert_sort(list []int) []int {
	n := len(list)
	for i := 1; i < n; i++ {
		key := list[i]
		for j := i - 1; j >= 0; j-- {
			if list[j] > key {
				list[j+1], list[j] = list[j], key
			}
		}
	}
	return list
}

// 希尔排序
func shell_sort(list []int) []int {
	step := 2
	n := len(list)
	gap := n / step
	for gap > 0 {
		for k := 0; k < gap; k++ {
			for i := k; i < n; i += gap {
				key := list[i]
				for j := i - gap; j >= 0; j -= gap {
					if list[j] > key {
						list[j+gap], list[j] = list[j], key
					}
				}
			}
		}
		gap /= step
	}
	return list
}

// 冒泡排序
func bubble_sort(list []int) []int {
	n := len(list)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if list[j] < list[i] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
	return list
}

// 快速排序
func quick_sort(list []int, left int, right int) []int {
	if left >= right {
		return list
	}
	key := list[left]
	min := left
	max := right
	for left < right {
		for left < right && list[right] > key {
			right--
		}
		list[left] = list[right]
		for left < right && list[left] < key {
			left++
		}
		list[right] = list[left]
	}
	list[right] = key
	quick_sort(list, min, left)
	quick_sort(list, left+1, max)
	return list
}

// 选择排序
func select_sort(list []int) []int {
	n := len(list)
	for i := 0; i < n; i++ {
		key := i
		for j := i + 1; j < n; j++ {
			if list[j] < list[key] {
				key = j
			}
		}
		if key != i {
			list[i], list[key] = list[key], list[i]
		}
	}
	return list
}

// 堆排序
func heap_sort(list []int) []int {
	n := len(list) - 1
	initHeap(list, n+1)
	for i := 0; i <= n; i++ {
		list[0], list[n-i] = list[n-i], list[0]
		adjustHeap(list, 0, n-i)
	}
	return list
}

func initHeap(list []int, num int) {
	for i := (num - 1 - 1) / 2; i >= 0; i-- {
		adjustHeap(list, i, num)
	}
}

func adjustHeap(list []int, i, num int) {
	leftChild := i*2 + 1
	rightChild := leftChild + 1
	key := i
	if leftChild < num && list[leftChild] > list[key] {
		key = leftChild
	}
	if rightChild < num && list[rightChild] > list[key] {
		key = rightChild
	}
	if key != i {
		list[i], list[key] = list[key], list[i]
		adjustHeap(list, key, num)
	}

}

// 归并排序
func merge_sort(list []int) []int {
	n := len(list)
	if n <= 1 {
		return list
	}
	key := n / 2
	left := merge_sort(list[:key])
	right := merge_sort(list[key:])
	list = merge(left, right)
	return list
}

func merge(left, right []int) []int {
	leftN := len(left)
	rightN := len(right)
	listNew := make([]int, leftN+rightN)
	i, j, k := 0, 0, 0
	for i < leftN && j < rightN {
		if left[i] < right[j] {
			listNew[k] = left[i]
			i++
		} else {
			listNew[k] = right[j]
			j++
		}
		k++
	}
	for i < leftN {
		listNew[k] = left[i]
		i++
		k++
	}
	for j < rightN {
		listNew[k] = right[j]
		j++
		k++
	}
	return listNew
}

// 基数排序
func radix_sort(list []int) []int {
	n := len(list)
	k := int(math.Ceil(math.Log10(max(list))))
	var key int
	var listNew [][]int
	for i := 1; i <= k; i++ {
		listNew = make([][]int, 10)
		for j := 0; j < n; j++ {
			key = list[j] / int(math.Pow10(i-1)) % 10
			listNew[key] = append(listNew[key], list[j])
		}
		m := 0
		for ii := range listNew {
			for iii := range listNew[ii] {
				list[m] = listNew[ii][iii]
				m++
			}
		}
	}
	return list
}

func max(list []int) float64 {
	key := 0
	n := len(list)
	for i := 1; i < n; i++ {
		if list[i] > list[key] {
			key = i
		}
	}
	return float64(list[key])
}
