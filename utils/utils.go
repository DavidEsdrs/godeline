package utils

type MapFunction[T any, Y any] func(T) Y

func Map[T any, Y any](arr []T, f MapFunction[T, Y]) []Y {
	res := make([]Y, len(arr))
	for i, item := range arr {
		res[i] = f(item)
	}
	return res
}
