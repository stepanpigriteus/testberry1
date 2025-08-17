package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	result, err := delete(arr, 4)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(result)
}

func delete(arr []int, index int) ([]int, error) {
	if index < 0 || index >= len(arr) {
		return arr, fmt.Errorf("index %d out of range", index)
	}
	// Создаем копию слайса чтобы не мутировать массив
	result := make([]int, len(arr))
	copy(result, arr)
	copy(result[index:], result[index+1:])
	return result[:len(result)-1], nil // удаляем последний элемент, чтобы убрать ссылку на крайний элемент в исходном массиве
}
