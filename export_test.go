
package main

import (
	"bytes"
	"io"
  "fmt"
	"testing"
	"encoding/csv"
  "github.com/ryokdy/go-kintone"
)

func makeTestData(app *kintone.App) error {
	err := deleteRecords(app, "")
	if err != nil {
		return err
	}
	records := make([]*kintone.Record, 0)

	record := make(map[string]interface{})
	record["single_line_text"] = kintone.SingleLineTextField("single line1")
	record["multi_line_text"] = kintone.SingleLineTextField("multi line1\nmulti line")
	record["number"] = kintone.DecimalField("12345")
	table := make([]*kintone.Record, 0)
	sub := make(map[string]interface{})
	sub["table_single_line_text"] = kintone.SingleLineTextField("table single line1")
	sub["table_multi_line_text"] = kintone.SingleLineTextField("table multi line1\nmulti line")
	table = append(table, kintone.NewRecord(sub))
	sub = make(map[string]interface{})
	sub["table_single_line_text"] = kintone.SingleLineTextField("table single line2")
	sub["table_multi_line_text"] = kintone.SingleLineTextField("table multi line2\nmulti line")
	table = append(table, kintone.NewRecord(sub))
	record["table"] = kintone.SubTableField(table)

	records = append(records, kintone.NewRecord(record))

	record = make(map[string]interface{})
	record["single_line_text"] = kintone.SingleLineTextField("single line2")
	record["multi_line_text"] = kintone.SingleLineTextField("multi line2\nmulti line")
	record["number"] = kintone.DecimalField("12345")
	records = append(records, kintone.NewRecord(record))

	_, err = app.AddRecords(records)

	return err
}

func TestExport1(t *testing.T) {
  buf := &bytes.Buffer{}

	app := newApp()
	makeTestData(app)

	config.fields = []string{"single_line_text", "multi_line_text", "number"}
	config.query = "order by record_number asc"
  err := writeCsv(app, buf)
	if err != nil {
		t.Error(err)
	}

	//output := buf.String()
  //fmt.Printf(output)
	fmt.Printf("\n")

	reader := csv.NewReader(buf)

	row, err := reader.Read()
	if err != nil {
		t.Error(err)
	}
	//fmt.Printf(row[0])
	if row[0] != "single_line_text" {
		t.Error("Invalid field code")
	}
	if row[1] != "multi_line_text" {
		t.Error("Invalid field code")
	}
	if row[2] != "number" {
		t.Error("Invalid field code")
	}

	row, err = reader.Read()
	if err != nil {
		t.Error(err)
	}
	if row[0] != "single line1" {
		t.Error("Invalid 1st field value of row 1")
	}
	if row[1] != "multi line1\nmulti line" {
		t.Error("Invalid 2nd field value of row 1")
	}
	if row[2] != "12345" {
		t.Error("Invalid 3rd field value of row 1")
	}

	row, err = reader.Read()
	if err != nil {
		t.Error(err)
	}
	if row[0] != "single line2" {
		t.Error("Invalid 1st field value of row 2")
	}
	if row[1] != "multi line2\nmulti line" {
		t.Error("Invalid 2nd field value of row 2")
	}
	if row[2] != "12345" {
		t.Error("Invalid 3rd field value of row 2")
	}

	row, err = reader.Read()
	if err != io.EOF {
		t.Error("Invalid record count")
	}
}

func TestExport2(t *testing.T) {
  buf := &bytes.Buffer{}

	app := newApp()
	makeTestData(app)

	config.fields = []string{"single_line_text", "multi_line_text", "number", "table"}
	config.query = "order by record_number asc"
  err := writeCsv(app, buf)
	if err != nil {
		t.Error(err)
	}

	//output := buf.String()
  //fmt.Printf(output)

	reader := csv.NewReader(buf)

	row, err := reader.Read()
	if err != nil {
		t.Error(err)
	}
	//fmt.Printf(row[0])
	if row[0] != "*" {
		t.Error("Invalid field code")
	}
	if row[1] != "single_line_text" {
		t.Error("Invalid field code")
	}
	if row[2] != "multi_line_text" {
		t.Error("Invalid field code")
	}
	if row[3] != "number" {
		t.Error("Invalid field code")
	}
	if row[4] != "table_single_line_text" {
		t.Error("Invalid field code")
	}
	if row[5] != "table_multi_line_text" {
		t.Error("Invalid field code")
	}

	row, err = reader.Read()
	if err != nil {
		t.Error(err)
	}
	if row[0] != "*" {
		t.Error("Invalid 1st field value of row 1")
	}
	if row[1] != "single line1" {
		t.Error("Invalid 2nd field value of row 1")
	}
	if row[2] != "multi line1\nmulti line" {
		t.Error("Invalid 3rd field value of row 1")
	}
	if row[3] != "12345" {
		t.Error("Invalid 4th field value of row 1")
	}
	if row[4] != "table single line1" {
		t.Error("Invalid 5th field value of row 1")
	}
	if row[5] != "table multi line1\nmulti line" {
		t.Error("Invalid 6th field value of row 1")
	}

	row, err = reader.Read()
	if err != nil {
		t.Error(err)
	}
	if row[0] != "" {
		t.Error("Invalid 1st field value of row 2")
	}
	if row[1] != "single line1" {
		t.Error("Invalid 2nd field value of row 2")
	}
	if row[2] != "multi line1\nmulti line" {
		t.Error("Invalid 3rd field value of row 2")
	}
	if row[3] != "12345" {
		t.Error("Invalid 4th field value of row 2")
	}
	if row[4] != "table single line2" {
		t.Error("Invalid 5th field value of row 2")
	}
	if row[5] != "table multi line2\nmulti line" {
		t.Error("Invalid 6th field value of row 2")
	}

	row, err = reader.Read()
	if err != nil {
		t.Error(err)
	}
	if row[0] != "*" {
		t.Error("Invalid 1st field value of row 3")
	}
	if row[1] != "single line2" {
		t.Error("Invalid 2nd field value of row 3")
	}
	if row[2] != "multi line2\nmulti line" {
		t.Error("Invalid 3rd field value of row 3")
	}
	if row[3] != "12345" {
		t.Error("Invalid 4th field value of row 3")
	}
	if row[4] != "" {
		t.Error("Invalid 5th field value of row 3")
	}
	if row[5] != "" {
		t.Error("Invalid 6th field value of row 3")
	}

	row, err = reader.Read()
	if err != io.EOF {
		t.Error("Invalid record count")
	}
}
