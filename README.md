# Reporting Service

This is a simple reporting service written in go that suits the basic need of getting live stats from a production database in a safe and secure way.

## Install

```
git clone https://github.com/joldit/reporting.git
cd reporting
go build
# Then you can expose your reports via http
REPORTING_DATABASE="a_dsn" ./reporting --http="an_internal_ip:8000"
# Or using the cli
REPORTING_DATABASE="a_dsn" ./reporting --name=name_of_report
```

## Writing new reports

Reports are written in Go using the following spec:

```
RegisterReport(name string,
  sql_query string,
  format []string,
  description string)
```

You can put reports in any file inside the main package, but be sure to init them by using `func init()`.

## Future work

* Filters
* Creation of reports via web
  * SQL Validation
  * SQL Wizard (?)
* Cursors
* Async Cursors
* (Schedule) Export via Spreadsheet/CSV/Email
* Viz
