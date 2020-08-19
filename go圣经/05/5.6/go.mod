module fibtest

go 1.13

replace fib => ./fib

require (
	fib v0.0.0
	golang.org/x/net v0.0.0-20200813134508-3edf25e44fcc // indirect
)
