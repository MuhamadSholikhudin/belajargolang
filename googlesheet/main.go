package main

// BEFORE RUNNING:
// ---------------
// 1. If not already done, enable the Google Sheets API
//    and check the quota for your project at
//    https://console.developers.google.com/apis/api/sheets
// 2. Install and update the Go dependencies by running `go get -u` in the
//    project directory.

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/api/sheets/v4"
)

func main() {
	ctx := context.Background()

	c, err := getClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	sheetsService, err := sheets.New(c)
	if err != nil {
		log.Fatal(err)
	}

	// The ID of the spreadsheet to retrieve data from.
	spreadsheetId := "1G-Xd2K9bsvgqsINmqU5s4awwK4EKqh8l4wFMzcsMiG8" // TODO: Update placeholder value.

	// The A1 notation of the values to retrieve.
	range2 := "Form Responses 1" // TODO: Update placeholder value.

	// How values should be represented in the output.
	// The default render option is ValueRenderOption.FORMATTED_VALUE.
	valueRenderOption := "FORMATTED_VALUE" // TODO: Update placeholder value.

	// How dates, times, and durations should be represented in the output.
	// This is ignored if value_render_option is
	// FORMATTED_VALUE.
	// The default dateTime render option is [DateTimeRenderOption.SERIAL_NUMBER].
	dateTimeRenderOption := "FORMATTED_STRING" // TODO: Update placeholder value.

	resp, err := sheetsService.Spreadsheets.Values.Get(spreadsheetId, range2).ValueRenderOption(valueRenderOption).DateTimeRenderOption(dateTimeRenderOption).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Change code below to process the `resp` object:
	fmt.Printf("%#v\n", resp)
}

func getClient(ctx context.Context) (*http.Client, error) {
	// TODO: Change placeholder below to get authentication credentials. See
	// https://developers.google.com/sheets/quickstart/go#step_3_set_up_the_sample
	//
	// Authorize using the following scopes:
	//     sheets.DriveScope
	//     sheets.DriveFileScope
	//     sheets.DriveReadonlyScope
	//     sheets.SpreadsheetsScope
	//     sheets.SpreadsheetsReadonlyScope
	return nil, errors.New("not implemented")
}
