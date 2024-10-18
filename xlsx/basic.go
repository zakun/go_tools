package xlsx

import (
	"zk/tools/common"

	"github.com/xuri/excelize/v2"
)

type Xlsx struct {
	Name         string
	StartRow     int
	CurrentRowNo int
}

type RowHandler func([]string, *Xlsx, any) error

func NewXlsx(name string, startRow int) *Xlsx {
	return &Xlsx{
		Name:         name,
		StartRow:     startRow,
		CurrentRowNo: 0,
	}
}

func (x *Xlsx) Process(rh RowHandler, params any) {
	f, err := excelize.OpenFile(x.Name)
	common.CheckError(err)
	defer func() {
		err := f.Close()
		common.CheckError(err)
	}()

	sn := f.GetSheetName(0)
	rows, err := f.Rows(sn)
	common.CheckError(err)

	rn := 0
	for rows.Next() {
		rn++
		x.CurrentRowNo = rn
		if rn < x.StartRow {
			continue
		}

		cols, err := rows.Columns()
		common.CheckError(err)

		err = rh(cols, x, params)
		common.CheckError(err)
	}

	err = rows.Close()
	common.CheckError(err)
}
