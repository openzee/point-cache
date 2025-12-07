package flow

import (
	"fmt"
	"testing"
	"time"

	xlsx "github.com/openzee/xlsx-loader"
)

func LoadExcel() []*xlsx.Point {
	rst, _ := xlsx.ParseExcel("test/test.xlsx")
	return rst
}

func TestA(t *testing.T) {
	ptList := LoadExcel()

	pt := Point{
		Original:        ptList[0],
		Value:           "11111111122222",
		ChangeTimestamp: time.Now(),
	}

	b, _ := pt.Marshal()
	fmt.Println(len(b))
}
