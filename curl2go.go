// Open devtool in chrome, go to network tab, right click on any network requset, select "copy as cURL"
// use this curl url in curl-to-go website to generate go code

package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	req, err := http.NewRequest("GET", "https://33.media.tumblr.com/avatar_402b1a99819c_16.png", nil)
	if err != nil {
		println(err.Error())
	}
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Dnt", "1")
	req.Header.Set("Accept-Encoding", "gzip, deflate, sdch")
	req.Header.Set("Accept-Language", "en-US,en;q=0.8,th;q=0.6,zh;q=0.4")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.109 Safari/537.36")
	req.Header.Set("Accept", "image/webp,image/*,*/*;q=0.8")
	req.Header.Set("Referer", "https://defaultnamehere.tumblr.com/post/139351766005/graphing-when-your-facebook-friends-are-awake")
	req.Header.Set("Cookie", "rxx=1fqqvyimpzl.1x5b4f6&v=1; __cfduid=d3ce3806304a2b27c5ea8664aa5b9e2381450246863; _ga=GA1.2.317602803.1437461388; __utma=189990958.317602803.1437461388.1455160201.1455165988.21; __utmz=189990958.1455165988.21.21.utmcsr=boonmeelab.com|utmccn=(referral)|utmcmd=referral|utmcct=/")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cache-Control", "no-cache")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		println(err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println(err.Error())
	}

	err = ioutil.WriteFile("data.png", body, 0644)
	if err != nil {
		println(err.Error())
	}
}
