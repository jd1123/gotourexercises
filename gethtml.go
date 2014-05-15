package main 

import "fmt"
import "net/http"
import "io/ioutil"
import "code.google.com/p/go.net/html"
import "strings"
import "log"

func parse_links(s string) ([]string, string){	
	var links []string
	var title string

	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
	    log.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data =="title"{
			title = n.FirstChild.Data

		}
	    if n.Type == html.ElementNode && n.Data == "a" {
	        for _, a := range n.Attr {
	            if a.Key == "href" {
	                links = append(links, a.Val)
	                break
	            }
	        }
	    }
	    for c := n.FirstChild; c != nil; c = c.NextSibling {
	        f(c)
	    }
	}
	f(doc)
	return links, title
}

func get_links(url string) ([]string, string){
	resp, err:=http.Get(url)
	if err != nil{
		return nil,""
	} 
	defer resp.Body.Close()
	body, err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		return nil,""
	}
	return parse_links(string(body))
}

func main(){
	found_links:=make(map[string]bool)

	links, title:=get_links("http://github.com")
	fmt.Println(title, ":")
	for _,v:=range links {
		fmt.Println(v)
		found_links[v]=true
	}
}