package work_poster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
)


type Poster struct {
	Title 		string
	Country		string
	Poster		string
	Year 		string
}

const apikey = "837a1b8b"
const apiUrl = "http://www.omdbapi.com/"

func GetPoster(title string) {
	resp, err := http.Get(strings.Join([]string{apiUrl,"?t=",url.QueryEscape(title),"&apikey=",apikey},""))
	if err != nil {
		log.Fatal( err )
		return
	}
	defer resp.Body.Close()
	var p Poster
	json.NewDecoder(resp.Body).Decode(&p)
	//fmt.Println(p.Title)
	if p.Poster != "" {
		downloadPoster(p.Poster)
	}
}
func GetPosterSprintf(title string) {
	resp, err := http.Get(fmt.Sprintf("%s?t=%s&apikey=%s",apiUrl,url.QueryEscape(title),apikey))
	if err != nil {
		log.Fatal( err )
	}
	defer resp.Body.Close()
	var p Poster
	json.NewDecoder(resp.Body).Decode(&p)
	//fmt.Println(p.Title)
}
func GetPosterAdd(title string) {
	resp, err := http.Get(apiUrl+"?t="+url.QueryEscape(title)+"&apikey="+apikey)
	if err != nil {
		log.Fatal( err )
	}
	defer resp.Body.Close()
	var p Poster
	json.NewDecoder(resp.Body).Decode(&p)
	//fmt.Println(p.Title)
}

func downloadPoster(url string ) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	bcontent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	index := strings.LastIndex(url, "/")
	filename ,err := os.Create(url[index+1:])
	if err != nil {
		log.Fatal(err)
		return
	}
	defer filename.Close()

	_, err = filename.Write(bcontent)
	if err != nil {
		log.Fatal(err)
	}



}

func SearchByTitle(titles ...strings) {
	var wg sync.WaitGroup
	for _, title := range titles {
		wg.Add(1)
		go func() {
				GetPoster(title)
		}
	}
}