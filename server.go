package main

import (
  "time"
  "bytes"
  "net/http"
  "html/template"
)

func renderTemplate(w http.ResponseWriter, tmpl string, d interface{}) {
  t, _ := template.ParseFiles("templates/" + tmpl + ".html")
  t.Execute(w, d)
}

func home(w http.ResponseWriter, r *http.Request) {
  var name = r.FormValue("report")
  start := time.Now()
  reportWriter := new(bytes.Buffer)
  if name != "" {
    getReportData(reportWriter, name)
  }
  renderTemplate(w, "home", map[string]interface{}{
    "Report": mapReport,
    "ReportData": reportWriter.String(),
    "Took": time.Since(start).Seconds(),
  })
}

func registerHandlers(mux *http.ServeMux) {
  mux.HandleFunc("/", home)
}