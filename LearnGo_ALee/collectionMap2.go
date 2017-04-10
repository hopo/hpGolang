package main

func main() {
    tickers := map[string]string{
        "GOOG": "Google Inc",
        "MSFT": "Microsoft",
        "FB":   "FaceBook",
        "AMZN": "Amazon",
    }

    // map 키 유무 체크 
    val, exists := tickers["MSFT"] // key value, bool value
    if !exists {
        println("No MSFT ticker")
    } else {
		println(val)
	}
}
