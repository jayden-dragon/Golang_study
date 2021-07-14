package main

func main() {
	msg := "hello"
	say(&msg)
	println(msg)
}

func say(msg *string) {
	println(*msg)
	*msg = "Changed" // Dereference
}
