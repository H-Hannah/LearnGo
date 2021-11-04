package main

var h string

func main() {
	h = "G"
	print(h)
	f1()
}

func f1() {
	h := "O"
	print(h)
	f2()
}

func f2() {
	print(h)
}

// G 0 G
