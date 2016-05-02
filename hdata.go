package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func hdata(w http.ResponseWriter, r *http.Request) {
	from, errx := strconv.Atoi(r.FormValue("from"))
	if errx != nil {
		log.Println(err.Error())
		from = 0
	}
	d, errs := historyData.ToJSON(from)
	if errs != nil {
		log.Println(errs.Error())
		http.Error(w, errs.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(d))
}
