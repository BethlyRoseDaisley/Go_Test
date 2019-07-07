package main

import (
    "fmt"
    "sort"
)
type Person struct {
    Name string
    Age  int
}
func (p Person) String() string {
    return fmt.Sprintf("%s: %d", p.Name, p.Age)
}
// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []Person
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func Example2() {
    people := []Person{
        {"Bob", 31},
        {"John", 42},
        {"Michael", 17},
        {"Jenny", 26},
    }
    fmt.Println(people)
    sort.Sort(ByAge(people))
    fmt.Println(people)
    // Output:
    // [Bob: 31 John: 42 Michael: 17 Jenny: 26]
    // [Michael: 17 Jenny: 26 Bob: 31 John: 42]
}


func Example1() {
	arry := []int{5,8,3,1,4,2,7,6}

	fmt.Println(arry)
	sort.Ints(arry)
	fmt.Println(arry)
	// Output:
	// [5 8 3 1 4 2 7 6]
	// [1 2 3 4 5 6 7 8]
}


type stature float32
type weight float32

type Pepole struct {
    name     string
    h        stature
    w 		 weight
}

// By is the type of a "less" function that defines the ordering of its Pepole arguments.
type By func(p1, p2 *Pepole) bool
// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(p []Pepole) {
    ps := &pepoleSorter{
        pepole:  pepole,
        by:      by, // The Sort method's receiver is the function (closure) that defines the sort order.
    }
    sort.Sort(ps)
}
// pepoleSorter joins a By function and a slice of Pepole to be sorted.
type pepoleSorter struct {
    pepole   []Pepole
    by       func(p1, p2 *Pepole) bool // Closure used in the Less method.
}
// Len is part of sort.Interface.
func (s *pepoleSorter) Len() int {
    return len(s.pepole)
}
// Swap is part of sort.Interface.
func (s *pepoleSorter) Swap(i, j int) {
    s.pepole[i], s.pepole[j] = s.pepole[j], s.pepole[i]
}
// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *pepoleSorter) Less(i, j int) bool {
    return s.by(&s.pepole[i], &s.pepole[j])
}

var pepole = []Pepole{
    {"Rose",    1.58, 66.6},
    {"Daisley", 1.78, 58.4},
    {"Lumiya",  1.65, 57.9},
    {"Sola",    1.68, 55.77},
}

// ExampleSortKeys demonstrates a technique for sorting a struct type using programmable sort criteria.
func Example_sortKeys() {
    // Closures that order the Planet structure.
    name := func(p1, p2 *Pepole) bool {
        return p1.name < p2.name
    }
    stature := func(p1, p2 *Pepole) bool {
        return p1.h < p2.h
    }
    weight  := func(p1, p2 *Pepole) bool {
        return p1.w < p2.w
    }
    decreasingWeight := func(p1, p2 *Pepole) bool {
        return !weight(p1, p2)
    }
    
    // Sort the pepole by the various criteria.
    By(name).Sort(pepole)
    fmt.Println("By name:", pepole)
    By(stature).Sort(pepole)
    fmt.Println("By stature:", pepole)
    By(weight).Sort(pepole)
    fmt.Println("By weight:", pepole)
    By(decreasingWeight).Sort(pepole)
    fmt.Println("By decreasing weight:", pepole)
    
    // Output:
	// By name: [{Daisley 1.78 58.4} {Lumiya 1.65 57.9} {Rose 1.58 66.6} {Sola 1.68 55.77}]
	// By stature: [{Rose 1.58 66.6} {Lumiya 1.65 57.9} {Sola 1.68 55.77} {Daisley 1.78 58.4}]
	// By weight: [{Sola 1.68 55.77} {Lumiya 1.65 57.9} {Daisley 1.78 58.4} {Rose 1.58 66.6}]
	// By decreasing weight: [{Rose 1.58 66.6} {Daisley 1.78 58.4} {Lumiya 1.65 57.9} {Sola 1.68 55.77}]
}



func main() {
	fmt.Println("--------Example1---------")
	Example1()
	fmt.Println("--------Example2---------")
	Example2()
	fmt.Println("--------Example3---------")
	Example_sortKeys()
}