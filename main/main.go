/*
 * @Author: p_hanxichen
 * @Date: 2023-07-31 21:19:09
 * @LastEditors: p_hanxichen
 * @LastEditTime: 2023-07-31 22:28:21
 * @FilePath: /src/classthreehomework/main/main.go
 * @Description:
 *
 * Copyright (c) 2023 by gdtengnan, All Rights Reserved.
 */
package main

import (
	slicegeneric "classthreehomework/slice_generic"
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4}
	fmt.Printf("a = %v\n", a)
	res := slicegeneric.Del[int](a, 100)
	fmt.Printf("res = %v\n", res)
	res = slicegeneric.Del[int](a, 0)
	fmt.Printf("res = %v\n", res)
}
