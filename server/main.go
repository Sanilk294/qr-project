package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type Item struct {
	ID    int
	Title string
	Link  string
}

var items = []Item{
	{1, "Givi ST611 Tanklock Tank Bag", "https://www.amazon.in/s?k=Givi+ST611+Tanklock+Tank+Bag"},
	{2, "Balaclava", "https://www.amazon.in/s?k=balaclava"},
	{3, "Royal Enfield Black Adventure Rider Seat KXA00427", "https://www.amazon.in/s?k=Royal+Enfield+Black+Adventure+Rider+Seat+KXA00427"},
	{4, "Rynox Navigator Frame Bags 24L", "https://www.amazon.in/s?k=Rynox+Navigator+Frame+Bags+24L"},
	{5, "Amazon Voucher Card", "https://www.amazon.in/gp/gift-card"},
}

func main() {

	// Serve static files (QR images)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web"))))
	// Serve CSS from /template as /assets/
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("template"))))

	// Homepage → Only QR Code
	http.HandleFunc("/", QRPage)

	// Product List Page → Shown after scanning QR
	http.HandleFunc("/list", ListPage)

	log.Println("Server running at http://localhost:10043")
	log.Fatal(http.ListenAndServe(":10043", nil))
}

// Only shows QR
func QRPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(filepath.Join("web", "qrpage.html")))
	tmpl.Execute(w, nil)
}

// Shows full product list
func ListPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(filepath.Join("web", "list.html")))
	tmpl.Execute(w, items)
}
