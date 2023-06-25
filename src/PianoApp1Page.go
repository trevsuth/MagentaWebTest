package main

import (
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := `<!DOCTYPE html>
<html>
<head>
    <title>Magenta.js Melody Generation</title>
    <script src="https://cdn.jsdelivr.net/npm/@magenta/music@^1.23.1/dist/magentamusic.js"></script>
</head>
<body>
    <h1>Magenta.js Melody Generation</h1>
    <button id="play">Generate Melody</button>
    <script>
	// Initialize the model.
music_vae = new mm.MusicVAE('https://storage.googleapis.com/magentadata/js/checkpoints/music_vae/mel_4bar_small_q2');
music_vae.initialize();

// Create a player to play the sampled sequence.
vaePlayer = new mm.Player();
vae_temperature = 1.0;

function playVAE() {
  if (vaePlayer.isPlaying()) {
    vaePlayer.stop();
    return;
  }
  music_vae
  .sample(1, vae_temperature)
  .then((sample) => vaePlayer.start(sample[0]));
}

	//Add event listner
	document.getElementById('play').addEventListener('click',playVAE);
    </script>
</body>
</html>`
		w.Write([]byte(html))
	})

	http.ListenAndServe(":8080", nil)

	//open the browser
	openBrowser("http://localhost:8080/")

	//Block forever
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

