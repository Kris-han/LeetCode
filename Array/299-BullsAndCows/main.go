package main

import "fmt"

func getHint(secret string, guess string) string {

	bucket := make([]int, 10)
	chS := []byte(secret)
	chG := []byte(guess)
	A := 0

	for i := 0; i < len(secret); i++ {
		if chS[i] == chG[i] {
			A++
		}
		bucket[chS[i]] += 1
		bucket[chG[i]] -= 1
	}

	B := 0
	for _, v := range bucket{
		B += v
	}

	B = len(secret) - B - A

	return fmt.Sprintf("%dA%dB", A, B)
}