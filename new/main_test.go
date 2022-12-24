package main

import (
	"testing"
)

func Benchmark_Match(b *testing.B) {
	for i := 0; i < b.N; i++ {
		users := Users{
			User{name: "John"},
			User{name: "Jane"},
			User{name: "Josh"},
		}
		_, err := users.Match("Josh")
		if err != nil {
			panic(err)
		}
		// b.Log(user.name)
	}
}
