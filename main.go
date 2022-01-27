package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const hostname = "localhost"
const port = ":8080"

type Img struct {
	Id  int    `json:"id"`
	Src string `json:"src"`
}

func main() {
	http.HandleFunc("/load", func(rw http.ResponseWriter, r *http.Request) {
		dir, err := os.ReadDir("./db")
		if err != nil {
			panic(err)
		}
		imgs := make([]Img, 0, 10)
		for _, e := range dir {
			name := e.Name()
			li := strings.LastIndex(name, ".")
			if li == -1 {
				log.Println("warn: no extension: ", err)
				continue
			}
			id, err := strconv.Atoi(name[:li])
			if err != nil {
				log.Println(err)
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}
			imgs = append(imgs, Img{
				Id:  id,
				Src: "http://" + hostname + port + "/db/" + name,
			})
		}
		je := json.NewEncoder(rw)
		err = je.Encode(imgs)
		if err != nil {
			log.Println(err)
		}
	})
	http.Handle("/db/", http.StripPrefix("/db/", http.FileServer(http.Dir("./db"))))
	http.Handle("/", http.FileServer(http.Dir("./app/dist")))

	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}

}
