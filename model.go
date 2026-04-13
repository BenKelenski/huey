package main

type model struct {
	rooms    []Room
	cursor   int
	selected map[int]struct{}
}
