# videoDownloader

pour pod.utt.fr

@Ziyi XU, le 19 février 2021

## But
Pour télécharger les vidéos sous format `.m3u8` :

Si jamais vous souhaitez télécharger une vidéo mais que vous constatez qu'elle n'est pas au format mp4, utilisez-le.

## Dépendances
- [Go](https://golang.org/dl/) / python
- [ffmpeg](https://ffmpeg.org/download.html)

## Go
1. Allez dans la page de moodle, cliquez droit pour copier l'adresse de la vidéo.
2. Et puis dans terminal :
`go run main.go FILEPATH/FILENAME.mp4 URLCOPIÉ`
La durée d'exécution dépend de la configuration de l'ordinateur et de l'environnement réseau.

### Mode 1 : normal

Par exemple :
``` bash
go run main.go -n test.mp4 https://pod.utt.fr/media/videos/a_demo/xxxx/playlist.m3u8
go run main.go --normal test.mp4 https://pod.utt.fr/media/videos/a_demo/xxxx/playlist.m3u8
```

### Mode 2 : batch
``` bash
go run main.go -b
go run main.go --batch
```

    
Et il faut dans `./url.txt` :
```
file1.mp4 https://url1
file2.mp4 https://url2
...
```

## Python
Utilisez jupyter-lab ou command line pytho.

![image](https://github.com/ArmandXUuu/videoDownloader/blob/develop/img/pythonOutput.png?raw=true)

## Des fonctionnalités à rajouter :
- [x] Télécharger en vrac (en utilisant `-b` ou `--batch`) - 25.02.2021

- [ ] Téléchargement simultané multitâche

        en fait en utilisant le `ffmpeg` ça fait très lentement : une vidéo de 15 mins prendra 4 mins pour télécharger et exporter pour moi
        -- car ma machine est vieille.
        -- faut implementer le merger à la main
        
- [ ] Simplifier la commande

***Si vous trouvez que ce programme n'est pas "légal", n'hésitez pas à me contacter.***
et créez un pullRequest si vous trouverez un problème
