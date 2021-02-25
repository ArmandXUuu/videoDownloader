import datetime
import os
import re
import threading
import time
import requests
from queue import Queue

def download(ts_queue, headers, filename="720p.ts"):
    url = ts_queue.get()
    try:
        requests.packages.urllib3.disable_warnings()
        r = requests.get(url, stream=True, headers=headers, verify=False)
        with open('cache/' + filename, 'wb') as fp:
            for chunk in r.iter_content(5242):
                if chunk:
                    fp.write(chunk)
        print('    Téléchargement de cache', filename, 'a réussi.')
    except:
        print('    Téléchargement de cache', filename, 'a échoué.')

def merge(name):
    try:
        path = 'cache/' + name
        command = 'ffmpeg -y -i cache/720p.ts -c:v libx264 -c:a copy -bsf:a aac_adtstoasc %s' % (path)
        res = os.popen(command)
        print(res.read())
        print('2/2 Terminé.')
    except:
        print('2/2 ÉCHOUÉ !')

def remove():
    try:
        os.remove('cache/720p.ts')
        print('Fichiers temporaires supprimés avec succès.')
    except:
        print('La suppression du dossier temporaire a échoué.')

if __name__ == '__main__':
    print("Bienvenu !")
    name = input('Veuillez saisir le nom du fichier `xxx.mp4`:')
    url = input('Veuillez entrer le lien de la vidéo :').strip()
    base_url = url.rsplit('/', 1)[0] + "/"
    headers = {
        'user-agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36'
    }

    start = datetime.datetime.now().replace(microsecond=0)

    s = Queue(10)
    s.put(base_url + '720p.ts')
    print("1/2 Commencer à télécharger des données...")
    download(s, headers)
    print("1/2 Fin du téléchargement.")
    end = datetime.datetime.now().replace(microsecond=0)
    print('    Temps passé à télécharger les fichiers : ' + str(end - start))
    print('2/2 Commencer à fusionner les données...')
    merge(name)

    remove()

    over = datetime.datetime.now().replace(microsecond=0)
    print('La fusion et la suppression de fichiers prend ' + str(over - end))
    print("Tout est bon")