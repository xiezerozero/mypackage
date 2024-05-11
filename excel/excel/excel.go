package excel

import (
	"github.com/xuri/excelize/v2"
	"io"
)

// 解析数据流,获取 excel 文件的所有行数据
func ExcelSheetRows(fileReader io.Reader, sheetName string) (rows [][]string, err error) {
	f, err := excelize.OpenReader(fileReader)
	if err != nil {
		return
	}
	rows, err = f.GetRows(sheetName)
	return
}
