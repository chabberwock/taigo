package main

import (
	_ "github.com/go-sql-driver/mysql"
	"os"
	"bufio"
	"compress/bzip2"
	"strings"
	"strconv"
	"fmt"
)

func parseFile(filepath string) {
	f, err := os.OpenFile(filepath, 0, 0)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// create a reader
	br := bufio.NewReader(f)
	// create a bzip2.reader, using the reader we just created
	cr := bzip2.NewReader(br)
	// lets do this. parse the file.
	// create a reader, using the bzip2.reader we were passed
	d := bufio.NewReader(cr)
	// create a scanner
	s := bufio.NewScanner(d)
	var currentLine int64
	currentLine = 0

	// scan the file! until Scan() returns "EOF", print out each line
	db.Exec("truncate passport")
	stm, err := db.Prepare("INSERT INTO passport (pnum) VALUES(?)")
	var sn []string
	for s.Scan() {
		currentLine++
		sn = strings.Split(s.Text(), ",")
		stm.Exec(strings.Join(sn, " "))
		homeData.ParseProgress = strconv.FormatInt(currentLine, 10)
		fmt.Printf("\r%s", strings.Repeat(" ", 35))
		fmt.Printf("\rИмпорт... %s", homeData.ParseProgress)
	}
	// we're done. return.
	return
}


