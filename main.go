package main

func main() {
	// sliceの先頭に要素を加えたい
	hoge := []string{
		"hoge1",
		"hoge2",
	}
	hoge, hoge[0] = append(hoge[:1], hoge[0:]...), "hoge0"
	println("動かない:", hoge)

	hoge = []string{
		"hoge1",
		"hoge2",
	}
	hoge = append(hoge[:1], hoge[0:]...)
	hoge[0] = "hoge0"
	println("動く:", hoge)
}
