package main

import "fmt"

func main() {
	{
		var arr [5]int
		fmt.Println(arr)

	}
	{
		arr := [...]int{1, 2, 3, 4}
		fmt.Println(arr)
	}
	{
		s := make([]string, 3)
		fmt.Println(s)
		s[0] = "Hello"
		s[1] = "Hello"
		s[2] = "Hello"
		fmt.Println(s)
		s = append(s, "Hey")
		fmt.Println(s)

		src := []string{"a", "b", "c"}
		dst := make([]string, len(src))
		copy(dst, src)
		fmt.Println(dst)

		sl1 := src[:2]
		fmt.Println(sl1)

	}
	{
		str := "rooo!"
		bytes := []byte(str)
		fmt.Println(bytes)
		fmt.Println(str == string(bytes))

		runes := []rune(str)
		fmt.Println(runes)
		fmt.Println(str == string(runes))

	}

	// Ranges
	{
		nums := []int{1, 2, 3, 4, 5}
		sum := 0
		for _, num := range nums {
			sum += num
		}
		fmt.Println("sum: ", sum)
	}
	{
		nums := []int{2, 3, 3, 4, 5}
		for idx, num := range nums {
			if num == 5 {
				fmt.Println("idx : ", idx)
			}
		}
	}
	{
		m := map[string]string{"A": "apple", "b": "bred"}
		for key, val := range m {
			fmt.Printf("%s -> %s\n", key, val)
		}

		for key := range m {
			fmt.Println("key: ", key)
		}
	}
	{
		for idx, char := range "oro" {
			fmt.Println(idx, char, string(char))
		}
	}
	{
		str := "ВУЗ"
		bytes := []byte(str)
		fmt.Println(bytes)
		fmt.Println(str == string(bytes))

		runes := []rune(str)
		fmt.Println(runes)
		fmt.Println(str == string(runes))

	}
}
