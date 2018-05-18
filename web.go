package main

import (
	"net/http"
	"strings"
	"html/template"
)

type HomeData struct {
	DownloadProgress string
	ParseProgress string
}

type ResultData struct {
	Passports []string
}

var homeData HomeData

// Главная страница /
func homeHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, _ := template.ParseFiles("home.html")
	tpl.Execute(writer, homeData)
}

// Запуск закачки и парсинга файла /parse
func parseHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	go downloadFile("passports.bz2", "http://guvm.mvd.ru/upload/expired-passports/list_of_expired_passports.csv.bz2")
	tpl, _ := template.ParseFiles("started.html")
	tpl.Execute(writer, homeData)
}

// Проверка паспортов и вывод результата /check
func checkHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	request.ParseForm();
	passports := strings.Split(request.Form.Get("nums"), "\n")
	stm, _ := db.Prepare("select * from passport where pnum = ?")
	result := ResultData{}
	var res string
	var err error
	for i := range passports {
		err = stm.QueryRow(strings.TrimSpace(passports[i])).Scan(&res)
		if err == nil {
			result.Passports = append(result.Passports, res)
		}
	}
	tpl, _ := template.ParseFiles("result.html")
	tpl.Execute(writer, result)


}
