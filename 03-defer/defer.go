package main

func main() {
	defer print("first line\n")
	print("second line\n")
	print("third line\n")
}
