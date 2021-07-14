package main

func main() {
	var a = 1
	for a < 15 {
		println(a)
		
		if a == 5 {
			a += a
			continue
		}
		a++

		if a > 10 {
			break
		}
	}

	if a == 11 {
		goto END
	}
	println(a)

END:
	println("End")
}
