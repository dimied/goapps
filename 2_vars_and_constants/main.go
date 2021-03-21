package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

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

func variadicFunction(nums ...int) int64 {
	var result int64
	for _, num := range nums {
		result += int64(num)
	}

	return result
}

type Val struct {
	a int
}

func (v Val) asString() string {
	return strconv.Itoa(v.a)
}

// using struct embedding
type ValSquared struct {
	Val
}

// "Overwrite"
func (v ValSquared) asString() string {
	return strconv.Itoa(v.a*v.a)
}

type Chain struct {
	Sum             int64
	numberOfEntries int
}

// can create function chains
func (chain *Chain) Add(val int) *Chain {
	chain.Sum += int64(val)
	chain.numberOfEntries++
	return chain
}

var str string

// Is called before main
func init() {
	str = "Assigned in init"
}

var (
	MyError = errors.New("My error")
)

type MyCustomError struct {
	code int
	msg string
}

// now it can be used for output as string
func (err MyCustomError) Error() string {
	return "Error: "+err.msg+"Code: "+strconv.Itoa(err.code)
}

func main() {

	fmt.Println(str)

	v := Val{2}
	vs := ValSquared{Val{2}}

	fmt.Println("Normal ", v.asString())
	fmt.Println("Embedded ", vs.asString())



	fmt.Println("Error: ", MyError)
	customErr := MyCustomError{}
	customErr.code = 1
	customErr.msg = "Bla bla"
	fmt.Println("Error: ", customErr)

	useMaps()
	
	printGlobals()

	printString()

	printValuesAndPointers()

	slicesAndArray()

	maps()

	r := variadicFunction(1, 2, 3, 4)
	fmt.Println("Sum is ", r)

	c := Chain{}

	c.Add(1).Add(2).Add(3)
	fmt.Println("Sum-chain: ", c.Sum, c.numberOfEntries)

	//Anonymous structs
	a := struct {
		b string
		c int
	}{"a", 1}
	fmt.Println("Anonymous structs: ", a)

	sum := 0

MyLabel:
	for i := 0; i < 10; i++ {
		if i > 6 {
			break MyLabel
		}
		sum += i
	}
	fmt.Println("Sum: ", sum)

	fmt.Println("Named return value", lastValueReturn(6))
	fmt.Println("Named return value", lastValueReturn(3))
}

func lastValueReturn(v int) (r int) {
	if r = v % 2; r > 0 {
		r = 1
	}
	return // here r is returned
}

func useMapsAsSets() {
	// Using maps as sets, no value, we are only interested in keys
	mySet := make(map[string]struct{})
	mySet["Hallo"] = struct{}{}
	mySet["Hallo du da"] = struct{}{}
	mySet["Hallo"] = struct{}{}
	fmt.Println("My first set", mySet)

	_, ok := mySet["Hallo"]
	fmt.Println("Element found ", ok)
}

func useMaps() {
	useMapsAsSets()

	// Use interface as key type
	outputMap := make(map[io.Writer]bool)
	byteBuffer := new(bytes.Buffer)

	outputMap[os.Stdout] = true
	outputMap[os.Stderr] = true
	outputMap[byteBuffer] = true

	for writer, value := range outputMap {
		if value {
			fmt.Fprintf(writer, "My message\n")
		}
	}

	fmt.Print("In buffer: ", byteBuffer.String())
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

	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("Rune: x=%v, type of %T\n", r, r)
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

func maps() {
	m1 := map[int]int{1: 1, 2: 4}
	fmt.Println("Map: ", m1)

	var lookup map[int]string
	lookup = make(map[int]string, 2)
	lookup[1] = "one"
	lookup[3] = "three"

	fmt.Println("Map2: ", lookup)

	delete(lookup, 3)
	delete(lookup, 2)
	fmt.Println("Map2:after delete: ", lookup)

	// comma-ok notation
	value, ok := lookup[100]
	if !ok {
		fmt.Println("Value not found for ", 100)
		fmt.Println("Value = ", value)
	}

	asJson, _ := json.Marshal(lookup)
	fmt.Println("Map2:json: ", string(asJson))
}
