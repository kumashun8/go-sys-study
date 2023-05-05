package chan2

import (
	"net"
	"net/http"
)

// 終了した順に書き出し
// チャネルに結果が投入された順に処理される
func writeToConnLIFO(reposonses chan *http.Response, conn net.Conn) {
	defer conn.Close()
	// 順番に取り出す
	for reposponse := range reposonses {
		reposponse.Write(conn)
	}
}

// 開始した順に書き出し
// チャネルにチャネルを入れた(開始した)順に処理される
func writeToConnFIFO(sessionReposonses chan chan *http.Response, conn net.Conn) {
	defer conn.Close()
	for essionReposons := range sessionReposonses {
		// 選択した仕事が終わるまで待つ
		reposponse := <-essionReposons
		reposponse.Write(conn)
	}
}
