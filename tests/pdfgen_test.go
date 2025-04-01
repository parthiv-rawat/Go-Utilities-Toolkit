package tests

import (
	"os"
	"testing"
	"go_utilities_toolkit/utils/pdfgen"
)

func TestGeneratePDF(t *testing.T) {
	filename := "test_output.pdf"
	err := pdfgen.GeneratePDF(filename, "Hello, PDF!")
	if err != nil {
		t.Errorf("PDF generation failed: %v", err)
	}
	defer os.Remove(filename)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Errorf("PDF file not created")
	}
}