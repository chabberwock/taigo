package main

/**
реализует скачивание файла, а затем запускает его парсинг. Вообще идея с запуском парсинга прямо отсюда не очень,
но я пока плохо умею в горутины, не понимаю как заставить их ждать друг друга
 */
import (
	"io"
	"os"
	"net/http"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"strconv"
)

type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	fmt.Printf("\r%s", strings.Repeat(" ", 35))
	homeData.DownloadProgress = strconv.FormatUint(wc.Total, 10)
	fmt.Printf("\rСкачивание... %s", wc.Total)
}

func downloadFile(filepath string, url string) error {

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Отслеживает процесс закачки файла, обновляя значение homedata.DownloadProgress
	counter := &WriteCounter{}
	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	if err != nil {
		return err
	}

	// The progress use the same line so print a new line once it's finished downloading
	fmt.Print("\n")

	go parseFile("passports.bz2")

	return nil
}

