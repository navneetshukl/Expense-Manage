package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	"github.com/navneetshukl/models"
)

func ShowPDF(c *gin.Context, data []models.PDFDetails) {
	// Create a new PDF instance
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)

	// Add a title
	pdf.Cell(40, 10, "Your PDF content goes here")

	// Add a line break
	pdf.Ln(10)

	// Create a table header
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(40, 10, "Date")
	pdf.Cell(40, 10, "Expense")
	pdf.Ln(10) // Move to the next line for data

	// Set font for data
	pdf.SetFont("Arial", "", 12)

	// Iterate through the data and add to the table
	for _, val := range data {
		date := val.Date.Format("02-01-2006")
		expense := val.Expense
		fmt.Println("Date is ", date)
		fmt.Println("Expense is ", expense)
		pdf.Cell(40, 10, date)
		pdf.Cell(40, 10, expense)
		pdf.Ln(10)
	}

	// Set headers for the HTTP response
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=my_expense.pdf")

	// Write the PDF to the response writer
	err := pdf.Output(c.Writer)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
}
