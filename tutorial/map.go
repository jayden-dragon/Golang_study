package main

func main() {
	tickers := map[string]string{
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"FB":   "FaceBook",
		"AMZN": "Amazon",
	}

	// map 키 체크
	val, exists := tickers["MSFT"] // 2개의 return값이 존재. key & key 유무 bool
	if !exists {
		println("No MSFT Ticker")
	} else {
		println(val, exists)
	}
}
