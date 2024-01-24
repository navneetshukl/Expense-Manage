package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	"github.com/navneetshukl/models"
)

func ShowPDF(c *gin.Context, data []models.PDFDetails, name, category, month string)error {
	// Create a new PDF instance
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Set font for the title
	pdf.SetFont("Arial", "B", 16)

	// Add a title centered on the page
	pdf.CellFormat(190, 10, fmt.Sprintf("Hello %s. This is your Expense on %s for month %s", name, category, month), "", 0, "C", false, 0, "")

	// Add a line break
	pdf.Ln(20)

	// Create a two-column layout
	colWidth := 95.0
	colHeight := 10.0

	// Set font for headers
	pdf.SetFont("Arial", "B", 12)

	// Create the headers for the two columns
	pdf.CellFormat(colWidth, colHeight, "Date", "1", 0, "C", false, 0, "")
	pdf.CellFormat(colWidth, colHeight, "Expense", "1", 0, "C", false, 0, "")
	pdf.Ln(colHeight) // Move to the next line for data

	// Set font for data
	pdf.SetFont("Arial", "", 12)

	// Iterate through the data and add to the two columns
	for _, val := range data {
		date := val.Date.Format("02-01-2006")
		expense := val.Expense

		// Add data to the first column
		pdf.CellFormat(colWidth, colHeight, date, "1", 0, "C", false, 0, "")
		// Add data to the second column
		pdf.CellFormat(colWidth, colHeight, expense, "1", 0, "C", false, 0, "")
		pdf.Ln(colHeight)
	}

	// Set headers for the HTTP response
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=my_expense.pdf")

	// Write the PDF to the response writer
	err := pdf.Output(c.Writer)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return err
	}
	return nil
}
