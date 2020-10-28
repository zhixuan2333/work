package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("HELLO WORLD")
	fileurl := "http://youku.cdn-163.com/20180507/6910_65bfcd86/1000k/hls/"

	resp, err := http.Get(fileurl + "index.m3u8")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("index.m3u81", data, 0644)

}
