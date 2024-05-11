package main

import (
	"fmt"
	"github.com/xiezerozero/mypackage/excel/excel"
	"net/http"
)

func process(w http.ResponseWriter, req *http.Request) {
	file, _, err := req.FormFile("file")
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	defer file.Close()
	allRows, err := excel.ExcelSheetRows(file, "Sheet1")
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	fmt.Fprint(w, allRows)
}

func main() {
	http.HandleFunc("/process", process)
	http.ListenAndServe(":8090", nil)
}
