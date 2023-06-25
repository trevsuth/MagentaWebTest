package main

import (
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

func main() {
	go func() {
		// Start the server
		http.Handle("/", http.FileServer(http.Dir("./")))
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Open the browser
	openBrowser("http://localhost:8080")

	// Block forever
	select {}
}

func openBrowser(url string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = exec.Command("xdg-open", url)
	}

	err := cmd.Start()
	if err != nil {
		log.Fatalf("Failed to open browser: %v", err)
	}
}

