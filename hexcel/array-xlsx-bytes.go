package hexcel

import (
	"github.com/xuri/excelize/v2"
	"strconv"
)

// ArrayToXlsx 二维数组转为Excel表格
// 文件过大时，请注意内存是否足够
func ArrayToXlsx(data [][]any) (xlsxBytes []byte, err error) {

	var lineHeight float64 = 26 // 行高
	var colWidth float64 = 22   // 列宽

	f := excelize.NewFile()
	var sheetNmae = f.GetSheetName(0)

	err = f.SetColWidth(sheetNmae, "A", COL[len(data[0])], colWidth)
	if err != nil {
		return
	}

	for i, row := range data {
		err = f.SetRowHeight(sheetNmae, i+1, lineHeight)
		if err != nil {
			return
		}
		var style int
		style, err = f.NewStyle(&excelize.Style{Alignment: &excelize.Alignment{Horizontal: "left", Vertical: "center"}})
		if err != nil {
			return
		}
		err = f.SetCellStyle(sheetNmae, "A"+strconv.Itoa(i+1), COL[len(row)]+strconv.Itoa(i+1), style)
		if err != nil {
			return
		}
		for j, col := range row {
			err = f.SetCellValue(sheetNmae, COL[j]+strconv.Itoa(i+1), col)
			if err != nil {
				return
			}
		}
	}

	buffer, err := f.WriteToBuffer()
	if err != nil {
		return
	}
	xlsxBytes = buffer.Bytes()
	return
}
