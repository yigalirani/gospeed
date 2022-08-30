package main

import (
	"fmt"
	"runtime"
	"time"
)

var run_counter uint64

func PrintMemUsage(title string) uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	var ans = m.HeapAlloc
	fmt.Println(title, ans)
	return ans

}

type node struct {
	counter     uint64
	left, right *node
}

func build(depth int) *node {
	if depth == 0 {
		return nil
	}
	run_counter++
	ans := node{counter: run_counter, left: build(depth - 1), right: build(depth - 1)}
	return &ans
}
func count(x *node) uint64 {
	if x == nil {
		return 0
	}
	return 1 + count(x.left) + count(x.right)
}
func main() {
	var start_time = time.Now()
	var start_mem = PrintMemUsage("start")
	var tree = build(25)
	var mid_mem = PrintMemUsage("mid")

	var diff = mid_mem - start_mem
	var ans = count(tree)
	time_diff := float32(time.Since(start_time).Microseconds())

	//var end_mem = PrintMemUsage("end")
	var per_node = diff / ans
	oper := float32(ans) / time_diff
	fmt.Println("num_nodes", ans, "num_mem", diff, "pernode", per_node, "time_diff", time_diff, "oper", oper)
}
