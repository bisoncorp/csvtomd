package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"flag"
)

var semicolon = flag.Bool("s", false, "if the file use semicolon instead comma")

func main() {
	flag.Parse()
	
	file, err := os.Open(flag.Arg(0))
	checkErr(err)

	reader:= csv.NewReader(file)
	if *semicolon {
		reader.Comma = ';'
	}

	records, err := reader.ReadAll()
	checkErr(err)

	builder := strings.Builder{}
	for i, line := range records {
		for _, data := range line {
			builder.WriteString("|" + data)
		}
		builder.WriteString("|\n")
		if i == 0 {
			for i := 0; i < len(line); i++ {
				builder.WriteString("|---")
			}
			builder.WriteString("|\n")
		}
	}
	fmt.Println(builder.String())
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}