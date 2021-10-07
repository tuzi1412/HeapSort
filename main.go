package main

import "fmt"

func main() {
	var input []int = []int{100000, 5, 7, 9, 44, 33, 4, 6, 882222226, 453, 244, 234}
	HeapSort(input, true)
	fmt.Println(input)
}

func HeapSort(input []int, flag bool) {
	//flag true从小到大排序,false从大到小排序
	if flag {
		//从小到大排序创建大顶堆
		CreatBigHeap(input)
		length := len(input)
		for length > 0 {
			temp := input[0]
			input[0] = input[length-1]
			input[length-1] = temp
			siftDown_Big(0, input[:length-1])
			length--
		}
	} else {
		//从大到小排序创建小顶堆
		CreatLittleHeap(input)
		length := len(input)
		for length > 0 {
			temp := input[0]
			input[0] = input[length-1]
			input[length-1] = temp
			siftDown_Little(0, input[:length-1])
			length--
		}
	}
}

func CreatLittleHeap(input []int) {
	for i := (len(input) - 1) / 2; i >= 0; i-- {
		siftDown_Little(i, input)
	}
}

func CreatBigHeap(input []int) {
	for i := (len(input) - 1) / 2; i >= 0; i-- {
		siftDown_Big(i, input)
	}
}

func siftUp_Little(i int, input []int) {
	var flag bool //标记是否需要继续上移
	var temp int
	for !flag && i > 0 {
		if input[i] < input[(i-1)/2] {
			temp = input[i]
			input[i] = input[(i-1)/2]
			input[(i-1)/2] = temp
			i = (i - 1) / 2
		} else {
			flag = true
		}
	}
}

func siftDown_Little(i int, input []int) {
	var flag bool     //标记是否需要继续下移
	var leftright int //标记左孩子与右孩子哪个比较小 1-左孩子 2-右孩子
	var temp int
	length := len(input)
	for !flag && i < length {
		leftright = 0
		//input[i*2+1] 左孩子
		//input[(i+1)*2] 右孩子
		//判断左右孩子是否存在
		if i*2+1 >= length { //没有左孩子
			flag = true
		} else if (i+1)*2 >= length { //没有右孩子
			leftright = 1
		} else {
			//判断左孩子与右孩子的大小
			if input[i*2+1] >= input[(i+1)*2] {
				leftright = 2
			} else {
				leftright = 1
			}
		}
		if leftright == 1 {
			if input[i] > input[i*2+1] {
				temp = input[i]
				input[i] = input[i*2+1]
				input[i*2+1] = temp
				i = i*2 + 1
			} else {
				flag = true
			}
		} else if leftright == 2 {
			if input[i] > input[(i+1)*2] {
				temp = input[i]
				input[i] = input[(i+1)*2]
				input[(i+1)*2] = temp
				i = (i + 1) * 2
			} else {
				flag = true
			}
		}
	}

}

func siftUp_Big(i int, input []int) {
	var flag bool //标记是否需要继续上移
	var temp int
	for !flag && i > 0 {
		if input[i] > input[(i-1)/2] {
			temp = input[i]
			input[i] = input[(i-1)/2]
			input[(i-1)/2] = temp
			i = (i - 1) / 2
		} else {
			flag = true
		}
	}
}

func siftDown_Big(i int, input []int) {
	var flag bool     //标记是否需要继续下移
	var leftright int //标记左孩子与右孩子哪个比较大 1-左孩子 2-右孩子
	var temp int
	length := len(input)
	for !flag && i < length {
		leftright = 0
		//input[i*2+1] 左孩子
		//input[(i+1)*2] 右孩子true
		//判断左右孩子是否存在
		if i*2+1 >= length { //没有左孩子
			flag = true
		} else if (i+1)*2 >= length { //没有右孩子
			leftright = 1
		} else {
			//判断左孩子与右孩子的大小
			if input[i*2+1] >= input[(i+1)*2] {
				leftright = 1
			} else {
				leftright = 2
			}
		}

		if leftright == 1 {
			if input[i] < input[i*2+1] {
				temp = input[i]
				input[i] = input[i*2+1]
				input[i*2+1] = temp
				i = i*2 + 1
			} else {
				flag = true
			}
		} else if leftright == 2 {
			if input[i] < input[(i+1)*2] {
				temp = input[i]
				input[i] = input[(i+1)*2]
				input[(i+1)*2] = temp
				i = (i + 1) * 2
			} else {
				flag = true
			}
		}
	}

}
