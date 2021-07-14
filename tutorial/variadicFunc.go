package main

func main() {
	say("This", "is", "sparta")
	say("Hi")
}

func say(msg ...string) {
	for _, s := range msg {
		println(s)
	}
}
