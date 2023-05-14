package main

import "fmt"

func main() {

	var b bool = true
	fmt.Println(b)

	var s string = "hello"
	fmt.Println(s)

	var i int = 42
	fmt.Println(i)

	var f float64 = 12.25
	fmt.Println(f)

	{
		var one, two int = 1, 2
		fmt.Println(one, two)
	}
	{
		var sunny = true
		fmt.Printf("%v is %T\n", sunny, sunny)
	}
	{
		var num int
		var str string
		var ok bool

		fmt.Printf("%#v %#v %#v\n", num, str, ok)
	}
	{
		food := "apple"
		fmt.Println(food)
	}
	{
		const s string = "constant"
		fmt.Println(s)

	}
	{
		const n = 233
		const d = 334 / n
		fmt.Println(d)
	}
	// For loops
	{
		i := 1
		for i <= 3 {
			fmt.Println(i)
			i = i + 1
		}
	}
	{
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}
	{
		for {
			fmt.Println("loop")
			break
		}
	}
	{
		for n := 0; n < 5; n++ {
			if n%2 == 0 {
				continue
			}
			fmt.Println(n)
		}
	}

	// if-else
	{
		if 7%2 == 0 {
			fmt.Println("7 is even")
		} else {
			fmt.Println("7 is odd")
		}

		if 8%4 == 0 {
			fmt.Println("8 is divisible by 4")
		}
	}
}
