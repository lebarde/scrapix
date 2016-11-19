# Get updated and see what changed

Create a config file inside `$HOME/.config/scrapix` and put the urls you want to watch:

```
params: "coucou2"

urls:
  "http://www.topologix.fr/exercices/":
    refresh: 1d
    watch: ".wrapper"
  "http://www.topologix.fr/sujets/":
```

Then run `~$ scrapix`.

# TODO

En un seul binaire en ligne de commande, on récupère toutes les url, et on doit les comparer avec la dernière version.

Il faut pour cela :
- lire la configuration yaml. Celle-ci, pour chaque url, contient l'url, la périodicité et éventuellement un sélecteur à la JQuery.
- lancer chaque requête http le cas échéant.
- récupérer l'état des dernières requêtes (dans un cache) et comparer les résultats
- enregistrer l'état des nouvelles url.
