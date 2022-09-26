package main

// From https://zetcode.com/golang/json/
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type User struct {
	Id         int
	Name       string
	Occupation string
}

func main() {
	{
		u1 := User{1, "John Doe", "gardener"}
		json_data, err := json.Marshal(u1)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(json_data))
		users := []User{
			{Id: 2, Name: "Roger Roe", Occupation: "driver"},
			{Id: 3, Name: "Lucy Smith", Occupation: "teacher"},
			{Id: 4, Name: "David Brown", Occupation: "programmer"},
		}
		json_data2, err := json.Marshal(users)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(json_data2))
	}
	{
		var u1 User
		data := []byte(`{
        "Id" : 1,
        "Name": "John Doe",
        "Occupation": "gardener"
    }`)
		err := json.Unmarshal(data, &u1)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Struct is:", u1)
		fmt.Printf("%s is a %s.\n", u1.Name, u1.Occupation)
		var u2 []User
		data2 := []byte(`
    [
        {"Id":2,"Name":"Roger Roe","Occupation":"driver"},
        {"Id":3,"Name":"Lucy Smith","Occupation":"teacher"},
        {"Id":4,"Name":"David Brown","Occupation":"programmer"}
    ]`)
		err2 := json.Unmarshal(data2, &u2)
		if err2 != nil {
			log.Fatal(err2)
		}
		for i := range u2 {
			fmt.Println(u2[i])
		}
	}
	{
		birds := map[string]interface{}{
			"sounds": map[string]string{
				"pigeon":  "coo",
				"eagle":   "squak",
				"owl":     "hoot",
				"duck":    "quack",
				"cuckoo":  "ku-ku",
				"raven":   "cruck-cruck",
				"chicken": "cluck",
				"rooster": "cock-a-doodle-do",
			},
		}
		data, err := json.MarshalIndent(birds, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(data))
	}
	{
		type User struct {
			Name       string
			Occupation string
			Born       string
		}
		filename, err := os.Open("data.json")
		if err != nil {
			log.Fatal(err)
		}
		defer filename.Close()
		data, err := ioutil.ReadAll(filename)
		if err != nil {
			log.Fatal(err)
		}
		var result []User
		jsonErr := json.Unmarshal(data, &result)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		fmt.Println(result)
	}
	{
		type Astronaut struct {
			Name  string
			Craft string
		}
		type people struct {
			Number  int
			People  []Astronaut
			Message string
		}
		url := "http://api.open-notify.org/astros.json"
		var netClient = http.Client{
			Timeout: time.Second * 10,
		}
		res, err := netClient.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		astros := people{}
		jsonErr := json.Unmarshal(body, &astros)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		fmt.Println(astros)
	}
}
