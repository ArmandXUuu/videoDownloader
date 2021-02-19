package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

var (
	path = os.Args[1]
	url = os.Args[2]
)

// Attention :
// Pas du tout de gestion d'exceptions ! (i'm lazzzy)
// Donc faut faire bien attention.
func main() {
	fmt.Println("Bienvenu !")
	fmt.Println("URL :", url, "Téléchargement commence.")
	fmt.Println("Veuillez patienter...")

	cmdArguments := []string{"-i", strings.Replace(url, "playlist.m3u8", "720p.m3u8", -1), path}
	// Si jamais le programme crashe, il faut vérifier que 720p.m3u8 est supporté pour cette vidéo :
		// Vas dans la route que tu viens de copier : https://pod.utt.fr/media/videos/c0b729d711ff70f8ad707c4046415d963b4ee50e52172303bbb4714831081bd5/2082/playlist.m3u8 pour voir
		// si 720p est bien dans la liste ; s'il y a un 1080p supporté tu peux bien modifier le code.
	execFfmpeg(cmdArguments...)
}

func execFfmpeg(arg ...string) {
	ctxt, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
	defer cancel()

	cmd := exec.CommandContext(ctxt, "ffmpeg", arg...)
	if _, err := cmd.Output(); err != nil {
		if ctxt.Err() != nil && ctxt.Err() == context.DeadlineExceeded {
			log.Fatal("timeout")
		}
		log.Fatalf("FAILED - %s : %s", arg[1], err.Error())
	}
	fmt.Println("SUCCEEDED !")
}