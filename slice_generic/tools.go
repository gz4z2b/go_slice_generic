/*
 * @Author: p_hanxichen
 * @Date: 2023-07-31 21:21:42
 * @LastEditors: p_hanxichen
 * @LastEditTime: 2023-07-31 22:18:23
 * @FilePath: /src/classthreehomework/slice_generic/tools.go
 * @Description:
 *
 * Copyright (c) 2023 by gdtengnan, All Rights Reserved.
 */
package slicegeneric

import "fmt"

func Del[T any](slice []T, index int) []T {
	if len(slice) < index {
		panic("超限")
	}
	result := slice[0:index]
	suffix := slice[index+1:]
	fmt.Printf("slice=%v\n", slice)
	fmt.Printf("suffix=%v\n", suffix)
	for _, value := range suffix {
		result = append(result, value)
	}
	return result
}
