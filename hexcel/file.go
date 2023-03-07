package hexcel

import "github.com/xuri/excelize/v2"

// TplSet 替换Excel模板变量
func TplSet(f *excelize.File, sheet string, key string, value any) (err error) {
	cntLocs, err := f.SearchSheet(sheet, key)
	if err != nil {
		return
	}
	for _, loc := range cntLocs {
		_ = f.SetCellValue(sheet, loc, value)
	}
	return
}
