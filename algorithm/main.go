package main

func dfs(x int) {
	if x > 999 {
		return
	}
	dfs(x*10 + 1)
	dfs(x*10 + 2)
	dfs(x*10 + 3)
}

func main() {
	dfs(0)
}
