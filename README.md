# videoDownloader
pour pod.utt.fr
@Ziyi XU, le 19 février 2021

## But
Pour télécharger les vidéos sous format `.m3u8`

## Dépendances
- Go
- [ffmpeg](https://ffmpeg.org/download.html)

## Comment faire ?
1. Allez dans la page de moodle, cliquez droit pour copier l'adresse de la vidéo.
2. Et puis dans terminal :
`go run videoDownloader.go FILEPATH/FILENAME.mp4 URLCOPIÉ`
La durée d'exécution dépend de la configuration de l'ordinateur et de l'environnement réseau.

Par exemple :
`go run videoDownloader.go test.mp4 https://pod.utt.fr/media/videos/a_demo/xxxx/playlist.m3u8`


***Si vous trouvez que ce programme n'est pas "légal", n'hésitez pas à me contacter.***
et créez un pullRequest si vous trouverez un problème
