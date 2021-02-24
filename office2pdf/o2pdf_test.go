package office2pdf

import (
	"testing"
)

func TestExcel2Pdf(t *testing.T) {
	excel := &Excel{}
	// excel.open("template.xlsx")

	excel.Export("D:\\Projects\\go-home\\src\\github.com\\annlumia\\hawa-cems\\pkg\\office2pdf\\template")
}
