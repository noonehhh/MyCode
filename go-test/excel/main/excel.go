package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func sli() []interface{} {
	sli1 := [10]string{"SRTB00001", "东海", "工作", "2019-09-18 15:38:33", "其他省行", "模板", "农行_横屏", "100.66.45", "未授权", "ch"}
	sli2 := [10]string{"SRTB00102", "南海", "工作", "2019-09-18 15:38:34", "其他省行", "模板", "农行_横屏", "100.66.46", "授权", "ch"}
	sli3 := [10]string{"SRTB00003", "北海", "工作", "2019-09-18 15:38:35", "其他省行", "模板", "农行_横屏", "100.66.47", "未授权", "ch"}
	sli4 := [10]string{"SRTB00103", "西海", "工作", "2019-09-18 15:38:36", "其他省行", "模板", "农行_横屏", "100.66.48", "授权", "ch"}
	sli := []interface{}{}
	sli = append(sli, sli1, sli2, sli3, sli4)
	return sli
}
func main() {
	ex := excelize.NewFile()
	sheet := ex.NewSheet("sheet1")
	ex.SetCellValue("sheet1", "A2", "helloworld")
	ex.SetCellValue("sheet2", "B2", 100)
	ex.SetActiveSheet(sheet)
	err := ex.SaveAs("C:/Users/97556/Desktop/Book1.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
