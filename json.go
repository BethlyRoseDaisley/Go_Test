package main 

import(
	"encoding/json"
	"fmt"
	"os"
)

func json_encode() {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	os.Stdout.Write(b)
}

func json_decode() {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}

	r := ColorGroup{}
	err = json.Unmarshal(b, &r)
	if err != nil {
		fmt.Println("json.Unmarshal error : ", err)
		os.Exit(1)
	}

	fmt.Println(r)

}

func json_struct_and_map() {
		
	type People struct {
		Name string `json:"name"`
		Age int `json:"age"`
		Feature map[string]interface{} `json:feature`
	}

	var feature map[string]interface{} = make(map[string]interface{})
	feature["personality"] = "cute"
	eature["Inner"] = "strength"

	p := People {
		Name : "ailumiyana",
		Age  : 18,
		Feature : feature,
	}

	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	fmt.Println("encode : ")
	os.Stdout.Write(b)

	r := People{}
	err = json.Unmarshal(b, &r)
	if err != nil {
		fmt.Println("json.Unmarshal error : ", err)
		os.Exit(1)
	}

	fmt.Println("\ndecode : ")
	fmt.Println(r)

	//update
	if r.Feature == nil {
		r.Feature = make(map[string]interface{})
	}
	r.Feature["fea"] = "warmth"
	
	b, err = json.Marshal(r)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	fmt.Println("update : ")
	os.Stdout.Write(b)
}

func main(){	
	//json_encode()
	//json_decode()

	json_struct_and_map()

}