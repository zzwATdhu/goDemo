package main

import(
"./memo1"
"fmt"
"os"
"time"
"log"
"io/ioutil"
"net/http"
)


func main(){
	m := memo1.New(HttpGetBody)
	for _,url:=range os.Args[1:]{
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("%s, %s, %d bytes\n",url, time.Since(start), len(value.([]byte)))
	}
}



func HttpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}