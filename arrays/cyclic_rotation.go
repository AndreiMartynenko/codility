package main

/*
ARRAYS

An array A consisting of N integers is given.
Rotation of the array means that each element is shifted right by one index,
and the last element of the array is moved to the first place.
For example, the rotation of array A = [3, 8, 9, 7, 6] is [6, 3, 8, 9, 7]
(elements are shifted right by one index and 6 is moved to the first place).

The goal is to rotate array A K times; that is, each element
of A will be shifted to the right K times.

Write a function:

func Solution(A []int, K int) []int

that, given an array A consisting of N integers
and an integer K, returns the array A rotated K times.

For example, given

    A = [3, 8, 9, 7, 6]
    K = 3
the function should return [9, 7, 6, 3, 8].
Three rotations were made:

    [3, 8, 9, 7, 6] -> [6, 3, 8, 9, 7]
    [6, 3, 8, 9, 7] -> [7, 6, 3, 8, 9]
    [7, 6, 3, 8, 9] -> [9, 7, 6, 3, 8]

	For another example, given

    A = [0, 0, 0]
    K = 1
the function should return [0, 0, 0]

Given

    A = [1, 2, 3, 4]
    K = 4
the function should return [1, 2, 3, 4]

Assume that:

N and K are integers within the range [0..100];
each element of array A is an integer within the range [−1,000..1,000].

Copyright 2009–2024 by Codility Limited. All Rights Reserved.
Unauthorized copying, publication or disclosure prohibited.
*/

import (
	"fmt"
)

func Solution(A []int, K int) []int {
	N := len(A)

	// Handle the case where the array is empty or K is 0
	if N == 0 || K == 0 {
		return A
	}

	// Adjust K to avoid unnecessary rotations
	K = K % N

	// Rotate the array by shifting elements to the right K times
	result := make([]int, N)
	for i := 0; i < N; i++ {
		result[(i+K)%N] = A[i]
	}

	return result
}

func main() {
	fmt.Println(Solution([]int{3, 8, 9, 7, 6}, 3)) // Output: [9 7 6 3 8]
	fmt.Println(Solution([]int{0, 0, 0}, 1))       // Output: [0 0 0]
	fmt.Println(Solution([]int{1, 2, 3, 4}, 4))    // Output: [1 2 3 4]
}
