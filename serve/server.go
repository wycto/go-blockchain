package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"wycto/core"
)

var blockchain *core.Blockchain

func main() {
	blockchain = core.NewBlockchain()
	run()
}

func run() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/api/list", List)
	http.HandleFunc("/api/add", Add)
	fmt.Println("启动服务  http://127.0.0.1:9099")
	http.ListenAndServe("127.0.0.1:9099", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "列表  <a href='http://127.0.0.1:9099/api/list'>http://127.0.0.1:9099/api/list</a>")
}

func List(w http.ResponseWriter, r *http.Request) {
	bytes, error := json.Marshal(blockchain)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
	}

	io.WriteString(w, string(bytes))
}

func Add(w http.ResponseWriter, r *http.Request) {
	blockdata := r.URL.Query().Get("data")
	blockchain.SendData(blockdata)
	List(w, r)
}
