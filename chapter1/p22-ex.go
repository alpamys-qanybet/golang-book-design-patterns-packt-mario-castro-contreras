package main

import "fmt"

/*
func main() {

	// addN := func(m int) func(int) int {
	// 	return func(n int) int {
	// 		return m + n
	// 	}
	// }
	// addFive := addN(5)
	// fmt.Println(addFive)
	// result := addN(6)
	// //5 + 6 must print 7
	// fmt.Println(result)

	// addN := func(m int) {
	// 	return func(n int) {
	// 		return m + n
	// 	}
	// }
	// addFive := addN(5)
	// fmt.Println(addFive)
	// result := addN(6)
	// //5 + 6 must print 7
	// fmt.Println(result)
}
*/

// This function intSeq returns another function, which we define anonymously in the body of intSeq. The returned function closes over the variable i to form a closure.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// We call intSeq, assigning the result (a function) to nextInt. This function value captures its own i value, which will be updated each time we call nextInt.
	nextInt := intSeq()

	// See the effect of the closure by calling nextInt a few times.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// To confirm that the state is unique to that particular function, create and test a new one.
	newInts := intSeq()
	fmt.Println(newInts())
	fmt.Println(newInts())

	fmt.Println(nextInt())
	fmt.Println(nextInt())

	fmt.Println(newInts())
	fmt.Println(newInts())

}

/*

1
2
3
1
2
4
5
3
4

*/
