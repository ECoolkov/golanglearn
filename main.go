package main

// data model:
// 1. types
// 2. operations
// 3. constraints

// formal language toolset:
// 1. primitive ideas (atomic, non-dividable)
// 2. complex idead (composition of primitives or other complex ideas)
// 3. tools for composing ideas

//     * // complex idea
//    / \ // relationships of composition
//   *   * // more complex ideas
//  / \ / \ // relationships of composition
// *  * *  * // primitive ideas

// ideas in formal language
// 1. data
// 2. procedures

// object - composition of data and related procedures that operate on that data as a single idea
// state - object internal data
// methods - object procederus

// types of data
// 1. primitive data
// 2. complex data

// types of procedures
// 1. primitive
// 2. complex

// types of operations (expressions, return values OF THE OPERANDS' TYPE)
// 1. unary - operates on single operand
// 2. binary - operates on two operands OF THE SAME DATA TYPE

// application to golang

// 8 bit = 1 byte // octet
// 8   = 2^3
// 16  = 2^4
// 32  = 2^5
// 64  = 2^6
// 128 = 2^7
// 256 = 2^8 // from 0 to 255
// ? = 2^64 // from 0 ti 2^64-1 ~9_400_200_000

// uint - Unsigned integer, that is has no sign capability, that is capable to store only 0 and POSITIVE NUMBERS
// int - signed integer, possible to store negative numbers, -1, -17, etc.

// zero value - default value FOR PARTICULAR DATA TYPE

// golang primitive data types
// 1. numeric
// 	1.1 integer: int (default), int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr
//  	zero value: 0
//  1.2 decimal: float32, float64
//      zero value: 0.0
//  1.3 complex: complex64, complex128
// 2. bool: false, true
//    zero value: false
// 3. byte: alias for uint8
//    zero value: 0
// 4. rune: alias for int32
//    zero value 0
// 5. arrays: [integer_size]Type_of_array_elements
//    5.1 MUST define size
//    5.2 array literals: [...]Type{elem1, elem2, elem3}
// 6. nil - MARK DENOTING ABSENCE OF ANY VALUE
// 7. pointer: *Type
// 8. function

// golang complex data types
// 1. struct
// 2. interface
// 3. slice // built-in
// 4. string // built-in
// 5. map // built-in

// assignability
// variable = value

// type literals
// 1. primitive
// 2. composite

// assignability of literals to var and const, and type conversion
// myInt := 4 // type of myInt ? int
// variable = literal e.g. var myInt int64 = 4; // type of myInt and type of 4? will assign?
// variable2 = variable1 e.g. var variable1 int32 = 4; var variable2 int64 = variable1;

// func f(areaID int64) {
//  fmt.Println(areaID)
// }
//
// f(42) // type of literal value 42 ? untyped int (assignable to int, int8, uint, uint8, ...)
//
// myAreaID := 42
// f(myAreaID) // compile-time error
//
// how to fix?
// f(int64(myAreaID))

// func f(areaID int64) {}
//
// var myAreaID int64 = 42
// f(myAreaID)
//
// what really happens to all this parameter-argument-shmargument shit above?
//
// this:
// func myFunction(param int) { fmt.Println(param) }
// is equivalent to this:
// var param int
// func myFunction { fmt.Println(param) }

// func GetProductsCount(tag string, areaID int64) (int, error) {...}
//
// tag := "myTag"
// areaID := 42
// count, err := GetProductsCount(tag, areaID)

// var myInt = 4
//    * // =
//   / \
//  /   * // 4
// * //var myInt
