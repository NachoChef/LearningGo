package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	// parse()
	s3()
}

// parse arg to float, convert to C, display w/format
func parse() {
	f, e := strconv.ParseFloat(os.Args[1], 64)
	if e != nil {
		fmt.Println("error parsing float")
		os.Exit(1)
	}
	cValue := fToC(f)
	fmt.Printf("%.2fF = %.2fC\n", f, cValue)
}

// pointers
func s3() {
	i := pt()
	fmt.Println("> i := pt()")
	fmt.Printf("i is now a pointer to some address containing %v\n", *i)
	origI := *i
	j := &i
	fmt.Println("> j := &i")
	fmt.Println("j is now a pointer that points to i (aka points to another pointer), so *i == **j and &i == *&j")
	k := *i
	fmt.Println("> k := *i")
	fmt.Println("k takes the value of *i, but it is not a pointer, so modifying *i will not modify k")
	fmt.Println("address to i is now *&j, dereference j and then get the value")

	fmt.Printf("location of i: %v\n", &i)
	fmt.Printf("location of j: %v\n", &j)
	fmt.Printf("value of j: %v and address that value points to: %v\n", *j, *&j)
	fmt.Printf("location of k: %v\n", &k)
	fmt.Printf("data at i: %d\n", *i)
	fmt.Printf("data at j: %d\n", **j)
	fmt.Printf("data of k: %d\n", k)
	*i = 3
	fmt.Println("> *i = 3")
	fmt.Printf("value at i is now %d\n", *i)
	fmt.Printf("k is still %v: %v\n", origI, k == origI)
	fmt.Printf("**j is still equal to *i: %v\n", *i == **j)
	l := i
	fmt.Println("> l = i")
	i = pt()
	fmt.Println("> i = pt()")
	fmt.Printf("*i is now %v\n", *i)
	fmt.Println("i points to some other address from pt(), j will still point to the address of &i, and fully dereference to the same value since the location of i hasn't moved")
	fmt.Printf("*i == **j: %v\n", *i == **j)
	fmt.Printf("however, the old value is still at previous *&i (which we saved in l), *l = %v\n", *l)
	ptInc(*&i)
	fmt.Println("> ptInc(*&i)")
	fmt.Printf("passing by reference to pointer i, i is now %v\n", *i)

}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func pt() *int {
	i := rand.Intn(100)
	return &i
}

func ptInc(p *int) {
	*p++
}
