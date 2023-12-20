package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

var sqlTmp = "INSERT INTO `dsj_mms`.`%s` (`user_id`, `channel_code`, `create_time`, `update_time`, `deadline`, `order_id`, `order_type`) VALUES ('%s', '%s', %d, %d, %d, '', 2);\n"

func main() {
	writeSql("30", 30, "huiyuanmianfei_30d", 1695005999)
	writeSql("365", 365, "huiyuanmianfei_1y", 1723949999)
	writeSql("730", 730, "huiyuanmianfei_2y", 1755485999)
}

func writeSql(fileName string, days int64, pCode string, deadline int64) {
	file, err := os.Open(fileName + ".csv")
	if err != nil {
		fmt.Println("open file fail: ", err.Error())
		return
	}
	defer file.Close()
	r := csv.NewReader(file)
	resultSqlFile, err := os.OpenFile(fileName+".sql", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open result file fail: ", err.Error())
		return
	}
	now := time.Now().Unix()
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if len(record) == 0 {
			continue
		}
		if record[0] == "user_id" {
			continue
		}
		//fmt.Printf("userid: %s, table: %s \n", record[0], GetTableName(record[0]))
		sql := fmt.Sprintf(sqlTmp, GetTableName(record[0]), record[0], pCode, now, now, deadline)
		resultSqlFile.WriteString(sql)
	}
}

// modUserId 对userId取模
func modUserId(userId string, tableNum int) int {
	if userId == "" {
		return 0
	}
	len := len(userId)
	num := 0
	for i := len; i > 0 && len-i < 2; i-- { // 取后2位
		if n, ok := s[userId[i-1:i]]; ok {
			num += n * base(len-i)
		}
	}
	return num % tableNum
}

func base(n int) int {
	switch n {
	case 0:
		return 1
	case 1:
		return len(s)
	}
	return 1
}

func GetTableName(userId string) string {
	return "user_pay_channel_" + strconv.Itoa(modUserId(userId, 16))
}

var s = map[string]int{
	"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "a": 10, "b": 11, "c": 12, "d": 13,
	"e": 14, "f": 15, "g": 16, "h": 17, "i": 18, "j": 19, "k": 20, "l": 21, "m": 22, "n": 23, "o": 24, "p": 25, "q": 26,
	"r": 27, "s": 28, "t": 29, "u": 30, "v": 31, "w": 32, "x": 33, "y": 34, "z": 35, "A": 36, "B": 37, "C": 38, "D": 49,
	"E": 40, "F": 41, "G": 42, "H": 43, "I": 44, "J": 45, "K": 46, "L": 47, "M": 48, "N": 49, "O": 50, "P": 51, "Q": 52,
	"R": 53, "S": 54, "T": 55, "U": 56, "V": 57, "W": 58, "X": 59, "Y": 60, "Z": 61,
}
