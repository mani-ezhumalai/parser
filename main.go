package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/net/html"
)

// const
// place for store strings,int values
const (
	ConstAnchorTag       = "a"
	ConstImgTag          = "img"
	Const2               = 2
	Const1               = 1
	ConstEmptyInputError = "please provide input as url on os args"
	Const200             = 200
	ConstHttpError       = "error occured while fething html content of"
	ConstTimeFormat      = "Mon, 02 Jan 2006 15:04 UTC"
	ConstSite            = "site:"
	ConstNumLinks        = "num_links:"
	ConstImages          = "images:"
	ConstLastFetch       = "last_fetch:"
	ConstURLSpliter      = "//"
	ConstSlash           = "/"
	Const0               = 0
	ConstHTMLExtention   = ".html"
)

// main ...
// entry point of script
func main() {
	// getting url value from os parameters
	if len(os.Args) < Const2 {
		fmt.Println(ConstEmptyInputError)
		os.Exit(Const1)
	}
	// iterating through the os arguments
	for index, value := range os.Args {
		var err error
		// skiping index zero due to excuteable path
		if index >= Const1 {
			// call for fetch html content from url
			var htmlContent []byte
			if htmlContent, err = fetchHTML(value); err != nil {
				fmt.Println(ConstHttpError, value)
				continue
			}
			// parsing html and getting element count
			anchorCount, imageCount := parseHTML(htmlContent)
			// call for print parsed data
			printResponse(value, anchorCount, imageCount)
			// archive html data into local file system
			archiveContent(value, htmlContent)
		}
	}
}

// fetchHTML ...
// fetches the provided URL and returns the response body or an error
func fetchHTML(url string) (htmlContent []byte, err error) {
	for {
		// call for get content using go http
		var resp *http.Response
		if resp, err = http.Get(url); err != nil {
			break
		}
		defer resp.Body.Close()
		// check for request is success or not
		if resp.StatusCode != Const200 {
			err = errors.New(ConstHttpError)
			break
		}
		// converting as byte
		htmlContent, err = ioutil.ReadAll(resp.Body)
		break
	}
	return
}

// parseHTML ...
// will parse the provided html content and return the count of tag elements
func parseHTML(inputData []byte) (anchorCount int, imageCount int) {
	// converting response as html tokenizer
	content := io.NopCloser(strings.NewReader(string(inputData)))
	htmlTokens := html.NewTokenizer(content)
	// itrating through html content and finding count of tag element
loop:
	for {
		tokenType := htmlTokens.Next()
		switch tokenType {
		case html.ErrorToken:
			break loop
		case html.StartTagToken:
			tag := htmlTokens.Token()
			// check for anchor
			if tag.Data == ConstAnchorTag {
				anchorCount++
			} else if tag.Data == ConstImgTag { // check for image
				imageCount++
			}
		}
	}
	return
}

// printResponse ...
// will print parsed and input data as a required format
func printResponse(url string, anchorCount int, imageCount int) {
	fmt.Println(ConstSite, url)
	fmt.Println(ConstNumLinks, anchorCount)
	fmt.Println(ConstImages, imageCount)
	fmt.Println(ConstLastFetch, time.Now().UTC().Format(ConstTimeFormat))
	fmt.Println("--------------------------------------------")
}

// archiveContent ...
// will store html content of url in local system
func archiveContent(url string, content []byte) {
	urlSplit := strings.Split(url, ConstURLSpliter)
	if len(urlSplit) > Const1 {
		nameArray := strings.Split(urlSplit[Const1], ConstSlash)
		if len(nameArray) > Const0 {
			// replacing dots
			replacedString := strings.ReplaceAll(nameArray[Const0], ".", "")
			ioutil.WriteFile(replacedString+ConstHTMLExtention, content, 0755)
		}
	}
}
