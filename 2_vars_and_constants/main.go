package main

import "fmt"

var x int

var px *int

// Declaring many variables at once
var (
	b    bool
	s    string
	f    float32
	d    float64
	i    int
	i8   int8
	i16  int16
	i32  int32
	i64  int64
	ui   uint
	ui8  uint8
	ui6  uint16
	ui32 uint32
	ui64 uint64
	r    rune = '*'
)

const my_const = 3.14

// Defining many constants at once
const (
	a  = 1
	a2 = 4
)

func variadicFunction (nums ...int) int64 {
	var result int64
	for _, num := range nums {
		result += int64(num)
	}

	return result;
}

type Chain struct {
	Sum int64
	numberOfEntries int
}

// can create function chains
func (chain *Chain) Add(val int) *Chain {
	chain.Sum += int64(val)
	chain.numberOfEntries++
	return chain;
}

func main() {

	printGlobals()

	printString()

	printValuesAndPointers()

	slicesAndArray()

	r := variadicFunction(1,2,3,4)
	fmt.Println("Sum is ", r)

	c := Chain{}

	c.Add(1).Add(2).Add(3)
	fmt.Println("Sum-chain: ", c.Sum, c.numberOfEntries)
}

func printGlobals() {
	s = "some string"
	fmt.Println("Global variables: ", x, s, f)
	fmt.Println("Global constants: ", my_const, a, b)
	// Will fail
	//my_const = 2
	declare_and_assign := 5 * 6
	fmt.Println("Declared", declare_and_assign)
	fmt.Println("Rune :", r)
}

func printString() {
	// UTF-8 encoded string for smiley (0xE2 0x98 0xBA)
	var s string = "☺"
	fmt.Println("Smiley (and byte size): ", s, len(s))
	var asBytes = []byte(s)
	fmt.Println("Bytes: ", asBytes, len(asBytes))
}

func printValuesAndPointers() {
	px = &x
	fmt.Println("Value and Pointer: ", x, px)
	*px = 10
	fmt.Println("Value and Pointer: ", x, px)
}

func slicesAndArray() {
	a := [5]int{}
	// Length and capacity always equal
	fmt.Println("Array (has fixed length): ", a,
		"length = ", len(a),
		"capacity = ", cap(a))

	// defining a slice
	slice := []int{1, 1, 0}
	fmt.Println("Slice: ", slice)
	slice[2] = 2
	fmt.Println("Slice: ", slice)

	// Initialize a slice
	var slice2 []int
	length := 5
	capacity := 10
	slice2 = make([]int, length, capacity)
	fmt.Println("Slice2: ", slice2,
		"length = ", len(slice2),
		"capacity = ", cap(slice2))
	// Add something to the slice
	slice2 = append(slice2, 3)
	fmt.Println("Slice2: ", slice2)
	// Concatenate slices
	slice2 = append(slice2, slice...)
	fmt.Println("Slice2: ", slice2)

	// Creating new slice
	slice3 := slice2[1:4]
	fmt.Println("Slice3: ", slice3, "len = ", len(slice3), "cap = ", cap(slice3))
	// Since go 1.2 can also define capacity
	slice3 = slice2[1:4:6]
	fmt.Println("Slice3: ", slice3, "len = ", len(slice3), "cap = ", cap(slice3))

	// "Reslicing"
	slice3 = slice3[1:3]
	fmt.Println("Slice3: ", slice3, "len = ", len(slice3), "cap = ", cap(slice3))

	// Copy from slices, smaller destination slice won't be resized
	copy(slice3, slice2[4:8])
	fmt.Println("Slice3: ", slice3, "len = ", len(slice3), "cap = ", cap(slice3))
}
