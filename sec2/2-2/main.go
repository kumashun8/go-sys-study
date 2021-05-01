package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// file, err := os.Create("test.csv")
	// if err != nil {
	// 	panic(err)
	// }
	// writer := csv.NewWriter(file)
	// writer.Write([]string{"Hoge", "20"})
	// writer.Write([]string{"Fuga", "30"})
	// writer.Flush()

	file, err := os.Open("test.csv")
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(file)

	records, err2 := reader.ReadAll()
	if err2 != nil {
		panic(err2)
	}
	fmt.Fprintf(os.Stdout, "%v", records)
}
