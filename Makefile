build:
	go build .

test:
	go test ./...

example_1: build
	./scalc [ LE 2 a.txt [ GR 1 b.txt c.txt ] ]

example_2: build
	./scalc [ GR 1 c.txt [ EQ 3 a.txt a.txt b.txt ] ]

