# videoDownloader
pour pod.utt.fr
@Ziyi XU, le 19 février 2021

## But
Pour télécharger les vidéos sous format `.m3u8` :

Si jamais vous souhaitez télécharger une vidéo mais que vous constatez qu'elle n'est pas au format mp4, utilisez-le.

## Dépendances
- [Go](https://golang.org/dl/)
- [ffmpeg](https://ffmpeg.org/download.html)

## Comment faire ?
1. Allez dans la page de moodle, cliquez droit pour copier l'adresse de la vidéo.
2. Et puis dans terminal :
`go run main.go FILEPATH/FILENAME.mp4 URLCOPIÉ`
La durée d'exécution dépend de la configuration de l'ordinateur et de l'environnement réseau.

Par exemple :
`go run main.go test.mp4 https://pod.utt.fr/media/videos/a_demo/xxxx/playlist.m3u8`

## Des fonctionnalités à rajouter :
- Télécharger en vrac
- Téléchargement simultané multitâche
- Simplifier la commande

***Si vous trouvez que ce programme n'est pas "légal", n'hésitez pas à me contacter.***
et créez un pullRequest si vous trouverez un problème

*Si je suis d'humeur, j'irai de l'avant et je ferai une version pour Python :D*
