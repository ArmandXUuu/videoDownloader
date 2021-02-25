package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

type video struct {
	fileName string
	url      string
}

// Attention :
// Pas du tout de gestion d'exceptions ! (i'm lazzzy)
// Donc faut faire bien attention.
func main() {
	videos := ini()
	captureVideosViaFFMPEG(videos)
	fmt.Println("Au revoir.")
	// Si jamais le programme crashe, il faut vérifier que 720p.m3u8 est supporté pour cette vidéo :
	// Vas dans la route que tu viens de copier : https://pod.utt.fr/media/videos/c0b729d711ff70f8ad707c4046415d963b4ee50e52172303bbb4714831081bd5/2082/playlist.m3u8 pour voir
	// si 720p est bien dans la liste ; s'il y a un 1080p supporté tu peux bien modifier le code.
}

func ini() []video {
	var videos []video
	mode := os.Args[1]
	fmt.Println("Bienvenu !")
	if mode == "--normal" || mode == "-n" {
		url := os.Args[3]
		fmt.Println("URL :", url, "Téléchargement commence.")
		fmt.Println("Veuillez patienter...")
		videos = append(videos, video{
			fileName: os.Args[2],
			url:      url,
		})
	} else if mode == "--batch" || mode == "-b" {
		fmt.Println("Trying to get urls from the file ./urls.txt")
		videos = getFileContents()
	} else {
		log.Fatal("SYNTAX ERROR.")
	}
	return videos
}

func getFileContents() []video {
	var videos []video
	b, err := ioutil.ReadFile("./url.txt")
	if err != nil {
		log.Fatal("Error reading file url.txt :", err)
	}
	s := string(b)
	for _, ligne := range strings.Split(s, "\n") {
		if ligne == "" {
			continue
		}
		videos = append(videos, video{
			fileName: strings.Split(ligne, " ")[0],
			url:      strings.Split(ligne, " ")[1],
		})
	}
	return videos
}

func captureVideosViaFFMPEG(videos []video) {
	/*	numCPUs := runtime.NumCPU()
		numVideos := len(videos)
		println(numCPUs)
		println(numVideos)*/
	println(" - Number of videos :", len(videos))
	for _, video := range videos {
		cmdArguments := []string{"-i", strings.Replace(video.url, "playlist.m3u8", "720p.m3u8", -1), video.fileName, "-y"}
		execFfmpeg(cmdArguments...)
	}
}

func execFfmpeg(arg ...string) {
	fmt.Println("\t - Start Downloading", arg[2], "... Veuillez patienter...")
	ctxt, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
	defer cancel()

	cmd := exec.CommandContext(ctxt, "ffmpeg", arg...)
	if _, err := cmd.Output(); err != nil {
		if ctxt.Err() != nil && ctxt.Err() == context.DeadlineExceeded {
			log.Fatal("timeout")
		}
		log.Fatalf("FAILED - %s : %s", arg[1], err.Error())
	}
	fmt.Println("\t - SUCCEEDED !")
}
