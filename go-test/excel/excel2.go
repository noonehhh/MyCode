package main

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
)

func slice() [4]interface{} {
	sli1 := [10]string{"SRTB00001", "东海", "工作", "2019-09-18 15:38:33", "其他省行", "模板", "农行_横屏", "100.66.45", "未授权", "ch"}
	sli2 := [10]string{"SRTB00102", "南海", "工作", "2019-09-18 15:38:34", "其他省行", "模板", "农行_横屏", "100.66.46", "授权", "ch"}
	sli3 := [10]string{"SRTB00003", "北海", "工作", "2019-09-18 15:38:35", "其他省行", "模板", "农行_横屏", "100.66.47", "未授权", "ch"}
	sli4 := [10]string{"SRTB00103", "西海", "工作", "2019-09-18 15:38:36", "其他省行", "模板", "农行_横屏", "100.66.48", "授权", "ch"}
	sli := [4]interface{}{sli1, sli2, sli3, sli4}
	return sli
}
func changeIndexToAixs(indexX int, indexY int) string {
	var arr = [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	indexY = indexY + 1
	resultY := ""
	for true {
		if indexY <= 26 {
			resultY = resultY + arr[indexY-1]
			break
		}
		mo := indexY % 26
		resultY = arr[mo-1] + resultY
		shang := indexY / 26
		if shang <= 26 {
			resultY = arr[shang-1] + resultY
			break
		}
		indexY = shang
	}
	return resultY + strconv.Itoa(indexX+1)
}
func MoidfyExcelCell(xlsx *excelize.File, sheet string, axis string, value interface{}) int {
	xlsx.SetCellValue(sheet, axis, value)
	return 0
}

func main() {
	exNew := excelize.NewFile()
	sheet := "sheet1"
	exNew.NewSheet("sheet1")
	//result := slice()
	//result := [3][9]{{"SRTB00001","东海","工作","2019-09-18 15:38:33","其他省行","模板","农行_横屏","100.66.45","未授权","ch"},
	//                    {"SRTB00102","南海","工作","2019-09-18 15:38:34","其他省行","模板","农行_横屏","100.66.46","授权","ch"},
	//                    {"SRTB00003","北海","工作","2019-09-18 15:38:35","其他省行","模板","农行_横屏","100.66.47","未授权","ch"},
	//                    {"SRTB00103","西海","工作","2019-09-18 15:38:36","其他省行","模板","农行_横屏","100.66.48","授权","ch"}
	//    }
	result := [5][5]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}}
	for i, row := range result {
		for j, cell := range row {
			zuobiao := changeIndexToAixs(i, j)
			MoidfyExcelCell(exNew, sheet, zuobiao, cell)
		}
	}
	exNew.SaveAs("C:/Users/97556/Desktop/Book2.xlsx")
}
