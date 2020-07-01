package countWordsAndImages

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

//func main() {
//	words, images, _ := CountWordsAndImages(os.Args[1])
//	fmt.Printf("文字：%d,图片：%d \n", words, images)
//}

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	//bare return
	return
}
func countWordsAndImages(n *html.Node) (words, images int) {
	//这里只是统计words总数,如果需要看每个单词可以通过map
	//counts := make(map[string]int)
	texts, images := visit(nil, 0, n)

	for _, v := range texts {
		v = strings.Trim(strings.TrimSpace(v), "\r\n")
		if v == "" {
			continue
		}
		scanner := bufio.NewScanner(strings.NewReader(v))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			words++
			//counts[scanner.Text()] ++
		}

	}
	//for _,v := range counts {
	//	words += v
	//}
	return
}

//递归循环html
func visit(texts []string, imgs int, n *html.Node) ([]string, int) {
	//文本
	if n.Type == html.TextNode {
		texts = append(texts, n.Data)
	}
	//图片
	if n.Type == html.ElementNode && (n.Data == "img") {
		imgs++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "script" || c.Data == "style" {
			continue
		}

		texts, imgs = visit(texts, imgs, c)
	}
	//多返回值
	return texts, imgs
}
