package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main()  {
	resp, _ := http.Get("http://openlibrary.org/works/OL27258W/editions.json?limit=")
	bs, _ := ioutil.ReadAll(resp.Body)
	var data map[string]interface{}
	err := json.Unmarshal(bs, &data)
	if err != nil {
		os.Exit(1)
	}
	entries := data["entries"].([]interface{})
	size := data["size"].(float64)
	for _,v := range entries {
		title := v.(map[string]interface{})["title"].(string)
		publishers := v.(map[string]interface{})["publishers"].([]interface{})
		lastModified := v.(map[string]interface{})["last_modified"].(map[string]interface{})["value"].(string)
		latestRevision := v.(map[string]interface{})["latest_revision"].(float64)
		key := v.(map[string]interface{})["key"].(string)
		publishDate := v.(map[string]interface{})["publish_date"].(string)
		fmt.Println("Title:", title)
		fmt.Println("Publishers:",publishers)
		fmt.Println("LastModified:",lastModified)
		fmt.Println("LatestRevision",latestRevision)
		fmt.Println("Key:",key)
		fmt.Println("PublishDate:", publishDate)
		fmt.Println("Size:", size)
		fmt.Println("---")
	}
}

