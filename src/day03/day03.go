package main

import "fmt"

func main() {
	// 创建一个整型切片
	// 其长度和容量都是4个元素
	slice := []int{10, 20, 30, 40}

	// 迭代每一个元素，并显示其值
	for index, value := range slice {
		fmt.Printf("Index: %d　Value: %d\n", index, value)
	}

	// 通过声明映射创建一个nil映射
	// 创建一个映射，存储颜色以及颜色对应的十六进制代码
	colors := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
		"blue":        "#66ccff",
	}

	value, exists := colors["blue"]

	if exists {
		fmt.Printf("this colors is %s\n", value)
	}

	// 删除键为Coral的键值对
	removeColor(colors, "blue")

	// 显示映射里的所有颜色
	for key, value := range colors {
		fmt.Printf("Key: %s　Value: %s\n", key, value)
	}
}

// removeColor将指定映射里的键删除
func removeColor(colors map[string]string, key string) {
	delete(colors, key)
}
