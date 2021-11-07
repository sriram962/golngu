package main

import (
	// "fmt"
	"encoding/json"
	// "fmt"
	"io/ioutil"

	// "os"
	"net/http"
	// "strconv"
	
)

type Pokemon struct {
	Pokemon []User `json:"pokemon"`
}

type User struct {
	Num        string   `json:"num"`
	Id         int      `json:"id"`
	Name       string   `json:"name"`
	Type       []string `json:"type"`
	Weaknesses []string `json:"weaknesses"`
}

func main() {

	file, _ := ioutil.ReadFile("users.json")

	data := Pokemon{}

	_= json.Unmarshal([]byte(file), &data)

	http.HandleFunc("/userlist", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println(pokdata)
		// for i := 0; i < len(data.Pokemon); i++ {

		// 	w.Write([]byte(strconv.Itoa(data.Pokemon[i].Id)))
		// 	// fmt.Fprint(w, data + "\n")
		// 	w.Write([]byte(data.Pokemon[i].Name + ","))
		// 	for j := 0; j < len(data.Pokemon[i].Type); j++ {
		// 		w.Write([]byte(data.Pokemon[i].Type[len(data.Pokemon[i].Type)-j-1] + ","))
		// 	}
		// 	for j := 0; j < len(data.Pokemon[i].Weaknesses); j++ {
		// 		w.Write([]byte(data.Pokemon[i].Weaknesses[len(data.Pokemon[i].Weaknesses)-j-1] + ","))
		// 	}
		// 	w.Write([]byte("\n"))
		// }
		file, _ := ioutil.ReadFile("users.json")

		//data := Pokemon{}
		var result interface{}
		json.Unmarshal([]byte(file), &result)
		// fmt.Println(result)
	})

	// http.HandleFunc("/pokemon/{name}", func(w http.ResponseWriter, r *http.Request) {
	// 	// fmt.Println(pokdata)
	// 	vars := mux.Vars(r)
	// 	key := vars["id"]
	// 	for i := 0; i < len(data.Pokemon); i++ {

	// 		w.Write([]byte(strconv.Itoa(data.Pokemon[i].Id)))
	// 		// fmt.Fprint(w, data + "\n")
	// 		w.Write([]byte(data.Pokemon[i].Name + ","))
	// 		for j := 0; j < len(data.Pokemon[i].Type); j++ {
	// 			w.Write([]byte(data.Pokemon[i].Type[len(data.Pokemon[i].Type)-j-1] + ","))
	// 		}
	// 		for j := 0; j < len(data.Pokemon[i].Weaknesses); j++ {
	// 			w.Write([]byte(data.Pokemon[i].Weaknesses[len(data.Pokemon[i].Weaknesses)-j-1] + ","))
	// 		}
	// 		w.Write([]byte("\n"))
	// 	}
		
	// })
	http.ListenAndServe(":8081", nil)

}
