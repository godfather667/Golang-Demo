// gov_data is a demostration program.
// It reads a file from "https://www.data.gov", process the data, and then
// displays to formated data on a webpage at localhost:8000.
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/buger/jsonparser"
	"github.com/go-resty/resty"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
)

// Error Variable.
var err error

// Database Handle.
var db *gorm.DB

// Template Handle.
var reportTemplate *template.Template

// Product Slice containing data read from www.data.gov.
var product []Product

// View Data Struct for template.
type ViewData struct {
	Collect [][]string
}

// Type definition for product structure.
type Product struct { // Product Type
	gorm.Model
	StrategyID        int64
	StrategyTitle     string
	DecisionDate      string
	OmbInitiative     string
	AmountType        string
	Fy2012_Amount     float32
	Fy2012_NetOrGross string
	Fy2013_Amount     float32
	Fy2013_NetOrGross string
	Fy2014_Amount     float32
	Fy2014_NetOrGross string
	Fy2015_Amount     float32
	Fy2015_NetOrGross string
	Fy2016_Amount     float32
	Fy2016_NetOrGross string
	Fy2017_Amount     float32
	Fy2017_NetOrGross string
	Fy2018_Amount     float32
	Fy2018_NetOrGross string
	Fy2019_Amount     float32
	Fy2019_NetOrGross string
}

// RestGov responds to endpoint localhost:8000/gov.
// It reads data from www.data.gov and builds a database with the information.
func RestGov(c echo.Context) error {
	client := resty.New()
	resp, _ := client.R().Get("https://www.transportation.gov/sites/dot.gov/files/docs/costsavings.json")
	data := resp.Body()

	jsonLoad(data)
	return nil
}

// JsonLoad parses the data read from www.data.gov and builds the gov.db Database.
func jsonLoad(data []byte) {
	jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		d_strategyID, err := jsonparser.GetInt(value, "strategyId")
		if err != nil && err.Error() != "Key path not found" {
			fmt.Printf("Parsing Error: %s\n", err)
		}
		d_strategyTitle, err := jsonparser.GetString(value, "strategyTitle")
		if err != nil && err.Error() != "Key path not found" {
			fmt.Printf("Parsing Error: %s\n", err)
		}
		d_decisionDate, err := jsonparser.GetString(value, "decisionDate")
		if err != nil && err.Error() != "Key path not found" {
			fmt.Printf("Parsing Error: %s\n", err)
		}
		d_ombInitiative, err := jsonparser.GetString(value, "ombInitiative")
		if err != nil && err.Error() != "Key path not found" {
			fmt.Printf("Parsing Error: %s\n", err)
		}
		d_amountType, err := jsonparser.GetString(value, "amountType")
		if err != nil && err.Error() != "Key path not found" {
			fmt.Printf("Parsing Error: %s\n", err)
		}
		d_Fy2012_Amount, err := jsonparser.GetFloat(value, "fy2012", "amount")
		if err != nil && err.Error() != "Key path not found" {
			fmt.Printf("Parsing Error: %s\n", err)
		}
		d_Fy2012_NetOrGross, err := jsonparser.GetString(value, "fy2012", "netOrGross")
		if err != nil && err.Error() != "Key path not found" {
			fmt.Printf("Parsing Error: %s\n", err)
		}
		d_Fy2013_Amount, err := jsonparser.GetFloat(value, "fy2013", "amount")
		if err != nil && err.Error() != "Key path not found" {
			fmt.Printf("Parsing Error: %s\n", err)
		}
		d_Fy2013_NetOrGross, err := jsonparser.GetString(value, "fy2013", "netOrGross")
		if err != nil && err.Error() != "Key path not found" {
			fmt.Printf("Parsing Error: %s\n", err)
		}
		d_Fy2014_Amount, err := jsonparser.GetFloat(value, "fy2014", "amount")
		if err != nil && err.Error() != "Key path not found" {
			fmt.Printf("Parsing Error: %s\n", err)
		}
		d_Fy2014_NetOrGross, err := jsonparser.GetString(value, "fy2014", "netOrGross")
		if err != nil && err.Error() != "Key path not found" {
			fmt.Printf("Parsing Error: %s\n", err)
		}
		d_Fy2015_Amount, err := jsonparser.GetFloat(value, "fy2015", "amount")
		if err != nil && err.Error() != "Key path not found" {
			fmt.Printf("Parsing Error: %s\n", err)
		}
		d_Fy2015_NetOrGross, err := jsonparser.GetString(value, "fy2015", "netOrGross")
		if err != nil && err.Error() != "Key path not found" {
			fmt.Printf("Parsing Error: %s\n", err)
		}
		d_Fy2016_Amount, err := jsonparser.GetFloat(value, "fy2016", "amount")
		if err != nil && err.Error() != "Key path not found" {
			fmt.Printf("Parsing Error: %s\n", err)
		}
		d_Fy2016_NetOrGross, err := jsonparser.GetString(value, "fy2016", "netOrGross")
		if err != nil && err.Error() != "Key path not found" {
			fmt.Printf("Parsing Error: %s\n", err)
		}
		d_Fy2017_Amount, err := jsonparser.GetFloat(value, "fy2017", "amount")
		if err != nil && err.Error() != "Key path not found" {
			fmt.Printf("Parsing Error: %s\n", err)
		}
		d_Fy2017_NetOrGross, err := jsonparser.GetString(value, "fy2017", "netOrGross")
		if err != nil && err.Error() != "Key path not found" {
			fmt.Printf("Parsing Error: %s\n", err)
		}
		d_Fy2018_Amount, err := jsonparser.GetFloat(value, "fy2018", "amount")
		if err != nil && err.Error() != "Key path not found" {
			fmt.Printf("Parsing Error: %s\n", err)
		}
		d_Fy2018_NetOrGross, err := jsonparser.GetString(value, "fy2018", "netOrGross")
		if err != nil && err.Error() != "Key path not found" {
			fmt.Printf("Parsing Error: %s\n", err)
		}
		d_Fy2019_Amount, err := jsonparser.GetFloat(value, "fy2019", "amount")
		if err != nil && err.Error() != "Key path not found" {
			fmt.Printf("Parsing Error: %s\n", err)
		}
		d_Fy2019_NetOrGross, err := jsonparser.GetString(value, "fy2019", "netOrGross")
		if err != nil && err.Error() != "Key path not found" {
			fmt.Printf("Parsing Error: %s\n", err)
		}

		db.Create(&Product{
			StrategyID:        d_strategyID,
			StrategyTitle:     d_strategyTitle,
			DecisionDate:      d_decisionDate,
			OmbInitiative:     d_ombInitiative,
			AmountType:        d_amountType,
			Fy2012_Amount:     float32(d_Fy2012_Amount),
			Fy2012_NetOrGross: d_Fy2012_NetOrGross,
			Fy2013_Amount:     float32(d_Fy2013_Amount),
			Fy2013_NetOrGross: d_Fy2013_NetOrGross,
			Fy2014_Amount:     float32(d_Fy2014_Amount),
			Fy2014_NetOrGross: d_Fy2014_NetOrGross,
			Fy2015_Amount:     float32(d_Fy2015_Amount),
			Fy2015_NetOrGross: d_Fy2015_NetOrGross,
			Fy2016_Amount:     float32(d_Fy2016_Amount),
			Fy2016_NetOrGross: d_Fy2016_NetOrGross,
			Fy2017_Amount:     float32(d_Fy2017_Amount),
			Fy2017_NetOrGross: d_Fy2017_NetOrGross,
			Fy2018_Amount:     float32(d_Fy2018_Amount),
			Fy2018_NetOrGross: d_Fy2018_NetOrGross,
			Fy2019_Amount:     float32(d_Fy2019_Amount),
			Fy2019_NetOrGross: d_Fy2019_NetOrGross,
		})
	}, "strategies")

}

// Report Page loads the database information and creates the web page display.
func ReportPage(c echo.Context) (err error) {
	w := c.Response()
	w.Header().Set("Content-Type", "text/html")

	// Read Database - Load Data
	db.Find(&product)

	collect := make([][]string, len(product))
	for i := range collect {
		collect[i] = make([]string, 21)
	}

	for i, _ := range product {
		collect[i][0] = fmt.Sprintf("%d", product[i].StrategyID)
		collect[i][1] = product[i].StrategyTitle
		collect[i][2] = product[i].DecisionDate
		collect[i][3] = product[i].OmbInitiative
		collect[i][4] = product[i].AmountType
		collect[i][5] = fmt.Sprintf("%f", product[i].Fy2012_Amount)
		collect[i][6] = product[i].Fy2012_NetOrGross
		collect[i][7] = fmt.Sprintf("%f", product[i].Fy2013_Amount)
		collect[i][8] = product[i].Fy2013_NetOrGross
		collect[i][9] = fmt.Sprintf("%f", product[i].Fy2014_Amount)
		collect[i][10] = product[i].Fy2014_NetOrGross
		collect[i][11] = fmt.Sprintf("%f", product[i].Fy2015_Amount)
		collect[i][12] = product[i].Fy2015_NetOrGross
		collect[i][13] = fmt.Sprintf("%f", product[i].Fy2016_Amount)
		collect[i][14] = product[i].Fy2016_NetOrGross
		collect[i][15] = fmt.Sprintf("%f", product[i].Fy2017_Amount)
		collect[i][16] = product[i].Fy2017_NetOrGross
		collect[i][17] = fmt.Sprintf("%f", product[i].Fy2018_Amount)
		collect[i][18] = product[i].Fy2018_NetOrGross
		collect[i][19] = fmt.Sprintf("%f", product[i].Fy2019_Amount)
		collect[i][20] = product[i].Fy2019_NetOrGross
	}

	vd := ViewData{collect}
	err = reportTemplate.Execute(w, vd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return err
}

// Main provide the primary executive function for
func main() {
	fmt.Println("Welcome to the data.gov server")
	DatabaseFile := "gov.db"

	reportTemplate, err = template.ParseFiles("hawthorne/index.html")
	if err != nil {
		fmt.Printf("Template Parsing Error - %s\n", err)
		os.Exit(-1)
	}

	// Remove Old Database if it exists
	_, err := os.Stat(DatabaseFile)
	if !os.IsNotExist(err) {
		os.Remove(DatabaseFile)
	}

	// Open New Database
	db, err = gorm.Open("sqlite3", DatabaseFile)
	if err != nil {
		fmt.Printf("failed to connect database: %s\n", err)
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Setup Endpoints
	e := echo.New()
	e.GET("/gov", RestGov)
	e.GET("/rep", ReportPage)

	// Start Server
	e.Start(":8000")
}
