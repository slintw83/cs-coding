package main

import "fmt"

type Dep struct {
	Prereq int
	Job    int
}

type Digraph struct {
	adj map[int][]int
}

func (G Digraph) V() []int {
	keys := make([]int, 0)

	for key := range G.adj {
		keys = append(keys, key)
	}

	return keys
}

func (G Digraph) Adj(v int) []int {
	return G.adj[v]
}

func (G Digraph) AddEdge(v, w int) {
	G.adj[v] = append(G.adj[v], w)
}

func (G Digraph) AddVertex(v int) {
	G.adj[v] = make([]int, 0)
}

func NewDigraph(jobs []int, deps []Dep) Digraph {
	G := Digraph{adj: make(map[int][]int)}
	for _, job := range jobs {
		G.AddVertex(job)
	}

	for _, dep := range deps {
		G.AddEdge(dep.Prereq, dep.Job)
	}

	return G
}

func ReverseDepthFirstOrder(G Digraph) []int {
	reversePost := []int{}
	visited := make(map[int]bool)

	for _, v := range G.V() {
		if _, ok := visited[v]; !ok {
			reversePost = dfs(G, v, reversePost, visited)
		}
	}

	for i, j := 0, len(reversePost)-1; i < j; i, j = i+1, j-1 {
		reversePost[i], reversePost[j] = reversePost[j], reversePost[i]
	}
	return reversePost
}

func dfs(G Digraph, v int, reversePost []int, visited map[int]bool) []int {
	visited[v] = true

	for _, w := range G.Adj(v) {
		if _, ok := visited[w]; !ok {
			reversePost = dfs(G, w, reversePost, visited)
		}
	}

	reversePost = append(reversePost, v)
	return reversePost
}

func DirectedCycle(G Digraph) bool {
	onStack := make(map[int]bool)
	visited := make(map[int]bool)
	cycle := false
	for _, v := range G.V() {
		if _, ok := visited[v]; !ok {
			cycle = dfsCycle(G, v, cycle, visited, onStack)
		}
	}

	return cycle
}

func dfsCycle(G Digraph, v int, cycle bool, visited, onStack map[int]bool) bool {
	onStack[v] = true
	visited[v] = true
	for _, w := range G.Adj(v) {
		if cycle {
			return true
		} else if _, ok := visited[w]; !ok {
			cycle = dfsCycle(G, w, cycle, visited, onStack)
		} else if alreadyOnStack, ok2 := onStack[w]; ok2 && alreadyOnStack {
			cycle = true
		}
	}

	onStack[v] = false
	return cycle
}

func TopologicalSort(jobs []int, deps []Dep) []int {
	G := NewDigraph(jobs, deps)
	fmt.Println(G)
	if DirectedCycle(G) {
		return []int{}
	}

	order := ReverseDepthFirstOrder(G)
	return order
}
