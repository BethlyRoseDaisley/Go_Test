package main

import(
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
	"bytes"
	"github.com/ailumiyana/latency"
)

func main() {

	get  := latency.New("HTTP", "GET")
	post := latency.New("HTTP", "POST")
	put  := latency.New("HTTP", "PUT")
	del  := latency.New("HTTP", "DELETE")

	//--------get---------
	get.Start()

	res, err := http.Get("http://127.0.0.1:8089/get")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(get.End())

	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(res)	
	fmt.Printf("%s\n", robots)


	// -----post-----
	post.Start()

	//json := "{\"sample\":\"example\"}"
	body := bytes.NewBufferString("{\"sample\":\"example\"}")
	
	res, err =  http.Post("http://127.0.0.1:8089/post", "application/json",	body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(post.End())

	robots, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)	
	fmt.Printf("%s\n", robots)

	// ------Put-----------
	put.Start()

	req, err := http.NewRequest("PUT", "http://127.0.0.1:8089/put",  bytes.NewBufferString("{\"example\":\"sample\"}"))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	res, _ = http.DefaultClient.Do(req)
	
	fmt.Println(put.End())
 	
	robots, _ = ioutil.ReadAll(res.Body) 	
	res.Body.Close()

	fmt.Println(res)	
	fmt.Println(string(robots))

	// ------Delete------
	del.Start()

	req, err = http.NewRequest("PUT", "http://127.0.0.1:8089/del",  nil)
	if err != nil {
		log.Fatal(err)
	}

	res, _ = http.DefaultClient.Do(req)
	
	fmt.Println(del.End())
 	
	robots, _ = ioutil.ReadAll(res.Body) 	
	res.Body.Close()

	fmt.Println(res)	
	fmt.Println(string(robots))

}
