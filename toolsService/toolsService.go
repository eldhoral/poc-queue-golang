package tools

import (
	"fmt"
	"runtime"
	"time"
)

type PerformanceParam struct {
	Name  string
	Value any
}

func Performance(name string, params ...*PerformanceParam) func() {
	var m1, m2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m1)

	start := time.Now()
	return func() {
		fmt.Printf("========= %s =========\n", name)
		fmt.Println("time:", time.Since(start))

		runtime.ReadMemStats(&m2)
		fmt.Println("total memory allocated:", m2.TotalAlloc-m1.TotalAlloc)
		fmt.Println("mallocs (cumulative count of heap objects allocated):", m2.Mallocs-m1.Mallocs)

		for _, param := range params {
			fmt.Printf("%s: %v\n", param.Name, param.Value)
		}
	}
}
