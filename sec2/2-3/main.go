package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")

	// json 化する元データ
	source := map[string]string{
		"Hello": "World",
	}

	// g に書き込まれた内容が w = web のレスポンスとして表示される
	g := gzip.NewWriter(w)
	// out に書き込んだものが, g と標準出力の両方に書き込まれる
	out := io.MultiWriter(g, os.Stdout)
	// json encode したものを out に書き込む
	encoder := json.NewEncoder(out)
	encoder.SetIndent("", "  ")

	// Encode した結果が g と標準出力の両方に書き込まれる
	if err := encoder.Encode(source); err != nil {
		log.Print(err)
	}
	if err := g.Close(); err != nil {
		log.Print(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
