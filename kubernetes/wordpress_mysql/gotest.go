package main

import (
        "strings"
        "log"
        "url"
)
var uuid_count = 5
var urlString = strings.Join("https://www.uuidtools.com/api/generate/v1/count/" + uuid_count)
func urlGenerate()
{	
	u, err := url.Parse(urlString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)
}