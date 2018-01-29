package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "This is a first string"
	fmt.Println("s =", s)
	s = s + ". This is another string"
	fmt.Println("s =", s)
	fmt.Println("substring(0,5) of s =", s[:5])
	s2 := s
	s3 := s[:1]
	fmt.Printf("Address of s is = %p\n", &s)
	fmt.Printf("Address of s2 is = %p\n", &s2)
	fmt.Printf("Address of s3 is = %p\n", &s3)

	fmt.Printf("Index of ('is a') in s = %d\n", strings.Index(s, "is a"))
	fmt.Printf("Len of s is %d\n", len(s))

	s = "field1;field2;field3;field4;field5;;;;"
	fmt.Println("Nos s is", s)
	fields := strings.Split(s, ";")
	fmt.Printf("s contains %d fields\n", len(fields))
	for i := 0; i < len(fields); i++ {
		fmt.Printf("fields[%d] = %s\n", i, fields[i])
	}
	fmt.Printf("cap of fields is %d\n", cap(fields))
	fmt.Println("join of fields with , is =", strings.Join(fields, ","))
	var ss = strings.Repeat(s, 10)
	fmt.Println(ss)
	ss = strings.Replace(ss, "field", "@", -1)
	fmt.Println(ss)

	c := s[2]
	fmt.Printf("type of c is %T\n", c)
	s = s + string(c)
	fmt.Println(s)
	i := 45423573453
	s = s + string(i)
	fmt.Println(s)
}
