package main

/*

A binary gap within a positive integer N is any maximal sequence of
consecutive zeros that is surrounded by ones at both ends in the binary representation of N.

For example, number 9 has binary representation 1001 and contains a binary gap of length 2.
The number 529 has binary representation 1000010001 and contains two binary gaps:
one of length 4 and one of length 3. The number 20 has binary representation 10100
and contains one binary gap of length 1. The number 15 has binary representation 1111
and has no binary gaps. The number 32 has binary representation 100000 and has no binary gaps.

Write a function:

func Solution(N int) int

that, given a positive integer N, returns the length of its longest binary gap.
The function should return 0 if N doesn't contain a binary gap.

For example, given N = 1041 the function should return 5,
because N has binary representation 10000010001 and so its
longest binary gap is of length 5. Given N = 32 the function
should return 0, because N has binary representation '100000' and thus no binary gaps.

Write an efficient algorithm for the following assumptions:
N is an integer within the range [1..2,147,483,647].

Copyright 2009–2024 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.

*/

import (
	"strconv"
)

func Solution(N int) int {
	binaryRepresentation := strconv.FormatInt(int64(N), 2)

	maxGap := 0
	currentGap := 0
	startCounting := false

	for _, digit := range binaryRepresentation {
		if digit == '1' {
			if startCounting && currentGap > maxGap {
				maxGap = currentGap
			}
			startCounting = true
			currentGap = 0
		} else if startCounting {
			currentGap++
		}
	}

	return maxGap

}

/*
func main() {
	fmt.Println(Solution(1041)) // Output: 5
	fmt.Println(Solution(32))   // Output: 0
}

### Task Description:

The task is to find the length of the longest binary gap
in the binary representation of a positive integer `N`.
A binary gap is defined as a maximal sequence of consecutive
zeros surrounded by ones at both ends in the binary representation.

For example:
- The binary representation of 9 is 1001, and it
contains a binary gap of length 2.
- The binary representation of 529 is 1000010001,
containing two binary gaps of lengths 4 and 3.

The function `Solution(N)` needs to return
the length of the longest
binary gap for a given positive integer `N`.
If there are no binary gaps, it should return 0.

### Approach:

1. **Convert to Binary:**
   - Start by converting the given integer `N` into
   its binary representation. This can be done using standard
   functions provided by the programming language you are using.
   In the Go solution, `strconv.FormatInt` is used.

2. **Iterate Through Binary Digits:**
   - Once you have the binary representation,
   iterate through each binary digit.

3. **Track Gaps:**
   - Keep track of the current binary gap length and
   the maximum gap length encountered so far.

4. **Counting Zeros:**
   - When you encounter a '1', check if you were already
   counting zeros. If yes, update the maximum gap length if
   the current gap is greater. Reset the current gap length to zero.

5. **Continue Counting:**
   - If the digit is '0' and you are already counting zeros,
   increment the current gap length.

6. **Return Result:**
   - After iterating through all digits,
   return the length of the longest binary gap.

### Advice:

- **Understand the Binary Representation:**
  Ensure you understand how to convert an integer to its binary
  representation and vice versa. Familiarize yourself with
  bitwise operations if needed.

- **Iterative Approach:**
  Consider an iterative approach to go through each binary digit one by one.
  Track the state of whether you are currently counting zeros or not.

- **Edge Cases:**
  Handle edge cases carefully. For instance,
  consider what happens when there are no binary gaps.

- **Efficiency:**
  Aim for an efficient solution. The given constraints indicate that
  the integer `N` can be as large as 2,147,483,647, so the solution
  should be able to handle large inputs efficiently.

- **Test Cases:**
  Test your solution with various test cases,
  including cases with multiple binary gaps, no gaps,
  and the largest possible input.

*/
