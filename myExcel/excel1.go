package myExcel

import (
	"fmt"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func writexlsx() {
	xlsx := excelize.NewFile()
	// 创建一个工作表
	index := xlsx.NewSheet("Sheet2")
	// 设置单元格的值
	xlsx.SetCellValue("Sheet2", "A2", "Hello world.")
	xlsx.SetCellValue("Sheet1", "B2", 100)
	// 设置工作簿的默认工作表
	xlsx.SetActiveSheet(index)
	// 根据指定路径保存文件
	err := xlsx.SaveAs("./Book1.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

func readfile() {
	xlsx, err := excelize.OpenFile("./Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	sheets := xlsx.GetSheetMap()
	for key, value := range sheets {
		log.Println(key, value)
	}

	log.Printf("current work sheet=%s", xlsx.GetSheetName(xlsx.GetActiveSheetIndex()))

	// 获取工作表中指定单元格的值
	cell := xlsx.GetCellValue("Sheet1", "B2")
	fmt.Println(cell)
	// 获取 Sheet1 上所有单元格
	rows := xlsx.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
	xlsx.SetColVisible("Sheet1", "D", false)
	xlsx.SetCellValue("Sheet1", "A1", 123)
	v := xlsx.GetRows("Sheet1")

	for i, value := range v {
		for ii, value1 := range value {
			fmt.Printf("%d,%d %s ", i, ii, value1)
		}
		fmt.Println()
	}

	//log.Println(v)
	//xlsx.DeleteSheet("Sheet1")

	xlsx.Save()

}

func zhexiantu() {
	xlsx, err := excelize.OpenFile("Book3.xlsx")
	if err != nil {
		log.Println(err)
		return
	}

	err = xlsx.AddChart("Sheet1", "M2",
		`{
  "type": "line",
  "series": [
    {
      "name": "",
      "categories": "",
      "values": "Sheet1!$B$2:$D$2"
    },
    {
      "name": "kai",
      "categories": "",
      "values": "Sheet1!$B$3:$D$3"
    },
    {
      "name": "good",
      "categories": "",
      "values": "Sheet1!$B$4:$D$4"
    }
  ],
  "title": {
    "name": "Fruit Line Chart1"
  }
}`)
	if err != nil {
		log.Println(err)
		return
	}

	//	xlsx.AddChart("Sheet1", "E20",
	//		`
	//{
	//  "type": "line",
	//  "series": [
	//    {
	//      "name": "dingKai",
	//      "categories": "Sheet1!B17:B30",
	//      "values": "Sheet1!C17:C30"
	//    }
	//  ],
	//  "title": {
	//    "name": "Fruit Line Chart"
	//  }
	//}`)

	// 保存工作簿
	err = xlsx.Save()
	if err != nil {
		fmt.Println(err)
	}
}

func MyExcelMain1() {
	//writexlsx()
	//readfile()
	zhexiantu()
}
