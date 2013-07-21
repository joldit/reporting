package main

import (
  "os"
  "io"
  "log"
  "flag"
  "fmt"
  "runtime"
  "database/sql"
  "net/http"
  _ "github.com/go-sql-driver/mysql"
)

var (
  reportName = flag.String("name", "", "Name of the report to generate.")
  httpAddr = flag.String("http", "", "Serve reporting via http.")
)

type reportInstructions struct {
  Query string
  Scan []string
  Description string
}

var mapReport = make(map[string]reportInstructions)

func usage() {
  fmt.Fprintf(os.Stderr, "Usage of %s:\n\n Reports available:\n", os.Args[0])
  for report_name := range mapReport {
    fmt.Fprintf(os.Stderr, "  %s\n", report_name)
  }
  fmt.Fprintf(os.Stderr, "\n")
  flag.PrintDefaults()
}

func RegisterReport(name string, query string, scan []string, description string) reportInstructions {
  report := reportInstructions{
    Query: query,
    Scan: scan,
    Description: description,
  }
  mapReport[name] = report
  return report
}

func getReportData(w io.Writer, name string) {
  instructions := mapReport[name]
  if instructions.Query == "" {
    return
  }
  
  logger := log.New(w, "", 0)
  
  // Connect to database
  db, err := sql.Open("mysql", os.Getenv("REPORTING_DATABASE"))
  if err != nil {
    log.Fatal("Could not connect to database:", err)
  }
  // Get rows for query
  rows, err := db.Query(instructions.Query)
  if err != nil {
    logger.Println("Query Error:", err.Error())
    return
  }
  // Get columns from descriptor
  columns, err := rows.Columns()
  if err != nil {
    logger.Println("Columns Error:", err.Error())
    return
  }

  // Build pointers so we can dinamically scan results at runtime
  valueColumn := make(map[string]*interface{})
  pointers := make([]interface{}, len(columns))

  for i, column := range columns {
    var value interface{}
    valueColumn[column] = &value
    pointers[i] = &value
  }
  
  // Iterate through resultset and parse rows
  for rows.Next() {
    if err := rows.Scan(pointers...); err != nil {
      log.Fatal("Scan Error:", err)
    }
    for i, scan := range instructions.Scan {
      fmt.Fprintf(w, scan, *pointers[i].(*interface{}))
    }
    fmt.Fprintf(w, "\n")
  }
  if err := rows.Err(); err != nil {
    log.Fatal(err)
  }
}

func main() {
  flag.Usage = usage
  flag.Parse()
  
  if *reportName != "" {
    getReportData(os.Stderr, *reportName)
  }
  
  if *httpAddr != "" {
    log.Printf("Reporting Server")
    log.Printf("version = %s", runtime.Version())
    log.Printf("address = %s", *httpAddr)

    registerHandlers(http.DefaultServeMux)

    if err := http.ListenAndServe(*httpAddr, http.DefaultServeMux); err != nil {
      log.Fatal("http error:", err)
    }
  }
}
