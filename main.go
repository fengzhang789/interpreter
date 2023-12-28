package main

import "fmt"

func main() {
	testing := []struct {
		expectedType int
		expectedLiteral string
	}{
		{1, "one"},
		{2, "two"},
	}

	for _, j := range testing {
		fmt.Println(j.expectedType, j.expectedLiteral)
	}
}
