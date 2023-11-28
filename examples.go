package main

import (
	"concurrency-runner/concurrencyrunner"
	"fmt"
)

func example1() {
	functions := []func() (int, error){
		func() (int, error) { return 1, nil },
		func() (int, error) { return 2, nil },
		func() (int, error) { return 3, nil },
	}

	results := concurrencyrunner.Run(functions)

	for i, result := range results.Results {
		if result.Error == nil {
			fmt.Printf("Result %d: %v\n", i+1, result.Result)
		}
	}

	err := results.CombineErrors()
	if err != nil {
		fmt.Printf("Combined Errors:\n%s\n", err)
	}
}

func example2() {
	testValues := []int{2, 3, 5, 7, 11, 15}
	var functions []func() (*Test, error)

	for _, value := range testValues {
		val := value
		functions = append(functions, func() (*Test, error) {
			if val < 10 {
				fmt.Printf("val:%d\n", val)
				return &Test{Id: "test", Value: val}, nil
			}
			panic("bigger than ten")
		})
	}

	results := concurrencyrunner.Run(functions)

	for i, result := range results.Results {
		if result.Error == nil {
			fmt.Printf("Result %d: %v\n", i+1, result.Result)
		}
	}

	err := results.CombineErrors()
	if err != nil {
		fmt.Printf("Combined Errors:\n%s\n", err)
	}
}

type Test struct {
	Id    string
	Value int
}
