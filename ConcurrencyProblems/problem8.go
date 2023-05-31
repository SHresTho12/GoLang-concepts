package main

import (
	"fmt"
	"sync"
)

type DownloadedData struct {
	URL  string
	Data []byte
}

// though the problem statement does not have waitgroup as a parameter in the function but it is necessary to pass it as a parameter
// as we need to wait for the routines to finish executing
// waitgroup is better than time.sleep
func downloadURL(url string, ch chan DownloadedData, wg *sync.WaitGroup) {
	defer wg.Done()
	// download the url
	fmt.Println("Downloading", url)
	// send the downloaded data to the channel
	ch <- DownloadedData{url, []byte("some data")}

}

func Problem8() {
	fmt.Println("Problem8")
	fmt.Println()

	wg := sync.WaitGroup{}
	urls := []string{
		"https://example.com",
		"https://golang.org",
		"https://google.com",
	}
	// Make channel
	dataChanel := make(chan DownloadedData)
	// Read all the URLs using goroutine
	for _, url := range urls {
		wg.Add(1)
		go downloadURL(url, dataChanel, &wg)
	}

	// Get & store the data
	var datas []DownloadedData
	for i := 0; i < len(urls); i++ {
		data := <-dataChanel
		datas = append(datas, data)
	}

	wg.Wait()
	// Print the data
	for _, data := range datas {
		fmt.Println(data.URL, len(data.Data))
	}

}
