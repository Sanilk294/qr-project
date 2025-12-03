package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	qrcode "github.com/skip2/go-qrcode"
)

// Automatically find the directory that contains go.mod
func findProjectRoot() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("Cannot get current directory:", err)
	}

	for {
		// Look for go.mod in this folder
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir // FOUND
		}

		// Move one level UP
		parent := filepath.Dir(dir)
		if parent == dir {
			log.Fatal("ERROR: go.mod NOT FOUND — run inside a valid project!")
		}

		dir = parent
	}
}

func main() {

	// Detect the correct project root
	projectRoot := findProjectRoot()
	fmt.Println("DETECTED PROJECT ROOT:", projectRoot)

	// Correct QR folder = {project root}/web/qrcodes
	outDir := filepath.Join(projectRoot, "web", "qrcodes")
	fmt.Println("QR WILL BE CREATED IN:", outDir)

	// Create folder
	err := os.MkdirAll(outDir, 0755)
	if err != nil {
		log.Fatal("Cannot create qrcodes folder:", err)
	}

	//siteURL := "http://localhost:10043/"
	siteURL := "http://localhost:10043/list"
	// siteURL := "https://YOUR_USERNAME.github.io/qr-project/list"
	outputFile := filepath.Join(outDir, "site_qr.png")

	err = qrcode.WriteFile(siteURL, qrcode.Medium, 512, outputFile)
	if err != nil {
		log.Fatal("Error generating QR:", err)
	}

	fmt.Println("QR CREATED:", outputFile)
}
