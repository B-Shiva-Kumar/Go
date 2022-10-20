package main

import (
	"fmt"
	"strconv"
)

func main() {

	// 1. int -> float conversion
	var var1 int = 2
	fmt.Printf("%v %T \n", var1, var1)

	var var2 float32
	var2 = float32(var1)
	fmt.Printf("%v %T \n", var2, var2)

	// 2. float -> int
	var var3 float32 = 999.99
	fmt.Printf("%v %T \n", var3, var3)

	var var4 int
	var4 = int(var3)
	fmt.Printf("%v %T \n\n", var4, var4)

	// 3. int to string
	// to convert to/from string we need Strconv pckg
	// Atoi (string to int) and Itoa (int to string).
	fmt.Printf("%v %T \n", var4, var4)

	var var5 string
	var5 = strconv.Itoa(var4)
	fmt.Printf("%v %T \n", var5, var5)

	// 4. string to int
	var var6 string = "Gogopher"
	fmt.Printf("%v %T \n", var6, var6)

	// var var7 int
	if var7, err := strconv.Atoi(var6); err == nil {
		fmt.Printf("%v %T \n", var7, var7)
	} else {
		fmt.Println("conversion failed :", err)
	}

	var var8 string = "10"
	if var9, err := strconv.Atoi(var8); err == nil {
		fmt.Printf("%v %T \n", var9, var9)
	} else {
		fmt.Println("conversion failed :", err)
	}
}
