# Benchmark - Django

Ce README.md présente tous les points forts/faibles du framework Django.

## Présentation du Django

Django est un framework haut niveau en python pour du développement web. L'objectif de Django est de pouvoir mettre en place des architectures/infrastructures webs assez rapidement.

Et en effet on se rend compte que c'est le cas.

## Points forts

Django est un framework Python qui a énormément de points forts rendant son utilisation crédible dans un domaine de production et voici ci-dessous les arguments à cela.

* Le temps de développement est rapide, avec la commande `django-admin` on peut très rapidement initialiser un projet, ainsi que des routes API et autre.
* Les intégrations d'autres frameworks/bibliothèques sont assez simple car Django est fait en Python. Par exemple, mettre en place des chats bots est simple avec l'ajout de dépendances comme Pytorch ou TensorFlow.
* Django a une intégration de Database intégrée, en SQL (sqlite). On peut très facilement mettre en place une DB surtout avec l'ORM Django intégrée, facilement manipulable car ce ne sont que des classes en Python (il ne faut que savoir coder en Python).
* La sécurité du framework est déjà gérée donc aucun travail d'utilisation doit être fait.

## Points faibles

Cependant certains points faibles du framework peut rendre compliquée son intégration dans un contexte de production.

* Le Python n'est pas aussi performant que d'autres langages comme le JS/TS ou bien le Go, un projet demandant des performance aura du mal à atteindre ses objectifs avec Django.
* Il faut beaucoup d'effort pour gérer beaucoup de traffic, beaucoup de clients, des bibliothèques tiers sont donc nécéssaires.
* L'architecture monolithique du framework est bonne pour des petits projets mais pas sur des gros projets, posant un problème sur la scalabilité et la maintenabilité de l'application.
* Django est fait en Python et le typage n'est pas natif mais optionnel et surtout pas assez restreignant (comme le Go par exemple), sur du long terme cela peut être un problème.

## Conclusion

Le Django peut être un bon framework pour des petits projets mais dans le contexte d'AREA qui est un projet demandant un très gros effort sur l'architecture, la scalabilité du software ou bien la stabilité de l'application.

Ainsi, je ne pense pas que Django pourrait être le Framework utilisé pour le backend dans ce projet.

## Commandes pour tester

```sh
python manage.py makemigrations djangoapp
python manage.py migrate
python manage.py runserver
```
