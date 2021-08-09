package main

import (
	"io"
	"net/http"
	"os"
	"strings"
)


func Wget(url, fileName string) {
	resp := getResponse(url)
	if fileName == "" {
		urlSplit := strings.Split(url, "/")
		fileName = urlSplit[len(urlSplit)-1]
	}
	writeToFile(fileName, resp)
}



func getResponse(url string) *http.Response {
	tr := new(http.Transport)
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	errorChecker(err)
	return resp
}



func writeToFile(fileName string, resp *http.Response) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)
	errorChecker(err)
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	errorChecker(err)
}

func errorChecker(err error) {
	if err != nil {
		panic(err)
	}
}
func main(){
	Wget("https://top4man.ru/inews/aHR0cHM6Ly9jZG5pbWcucmcucnUvaW1nL2NvbnRlbnQvMjAyLzAwLzQ0LzEwMTBfZF84NTAuanBn","1.jpg")
}
