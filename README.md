# multi-agent-system-traffic-modelisation

The goal of this project is to simulate the traffic near ou school and see if we can optimise it


## Structure of the repo

Here is a quick view of the overall structure.
This is a starting point, don't hesitate to add new folder and add it in the Readme.

```
project/
├── go.mod
├── cmd/
│   ├── gui/main.go        # launch simulation + graphical window
│   └── headless/main.go   # (optionnal) simu without graphic : tests, stats, batch
├── internal/
│   ├── agent/             # interface Agent + type of agent (one file per type)
│   ├── env/               # state of the world, percept, application of act
│   ├── graph/             # Everything that is needed for the construction of a graph.
│   ├── simulation/        # loop, clock, start/pause/stop, snapshots
│   ├── gui/               # everything that is linked to Ebiten : Game, rendu, input, sprites
│   └── config/            # parameter (nb of agents, size of the map, speed)
└── assets/                # images, font
```

### Important Point

`internal/` : Go convention for code that will not be imported from outsite.

A `gui/` folder separated : `simu`, `agt` and `env` must never import Ebiten. Like this we can launch the simu without the graphical window, for test or even if we want to change the game engine.

Two main in `cmd/` : One version without graphical window, helps for debugging.


# Bonnes pratiques de gestion de projet et de développement

Voici les quelques règles qu'on essaie de suivre pour travailler ensemble sur le projet : gestion des issues, organisation du backlog, échanges autour du code et utilisation de Git.

L'objectif est simple : savoir qui travaille sur quoi, pourquoi on le fait, ce qui a été décidé et comment retrouver ces informations deux semaines plus tard.

On cherche surtout à éviter de perdre du temps à retrouver le contexte d'un problème ou à comprendre pourquoi un bout de code a été écrit d'une certaine manière.

## 1. Pourquoi prendre le temps de faire des issues ?

Une issue n'est pas simplement une tâche à cocher. C'est aussi un espace de réflexion et une documentation du projet.

Quelques raisons de prendre le temps de rédiger des tickets :
- **Garder une trace de la réflexion collective** : on rencontre un problème, on propose des solutions, on en discute. Deux semaines plus tard, on peut retrouver le cheminement de pensée du groupe sans devoir tout recommencer.
- **Savoir qui travaille sur quoi** : on identifie les tâches en cours, les personnes concernées et les éventuels blocages.
- **Documenter le développement au fur et à mesure** : plutôt que d'essayer de reconstituer la documentation après avoir codé, on écrit directement le contexte, le problème et les décisions prises.
- **Préparer l'implémentation** : une bonne issue permet de réfléchir à la solution avant de commencer à coder.

### La philosophie c'est : 90 % de réflexion, 10 % d'implémentation

Le commit est la dernière étape d'un travail de réflexion déjà effectué :
1. On identifie un problème ou un besoin.
2. On le décrit dans une issue.
3. On réfléchit collectivement aux solutions.
4. On se met d'accord sur ce qu'il faut faire.
5. On implémente la solution.
6. On vérifie et on merge.

## 2. Comment rédiger une issue ?

Une issue doit permettre à quelqu'un qui découvre le sujet de comprendre :
- Quel est le problème ?
- Dans quel contexte apparaît-il ?
- Qu'est-ce qui existe actuellement ?
- Qu'est-ce qu'on voudrait obtenir ?
- Quelles solutions ont déjà été envisagées ?

### Le titre

Le titre doit décrire une **action à réaliser** : une issue, c'est agir sur quelque chose, quelque part.

Par exemple :
- `Modifier la gestion des pointeurs dans la classe X`
- `Ajouter la vérification des paramètres dans le fichier Y`
- `Corriger le calcul des coordonnées dans la classe Z`

On essaie d'être suffisamment précis pour comprendre ce qui va être modifié sans avoir à ouvrir le ticket.

### L'objectif général

Décrire le besoin auquel répond le ticket. On explique ici la finalité de la modification : pourquoi cette fonctionnalité est nécessaire ou quel comportement on cherche à améliorer.

### Le problème rencontré

Décrire le problème à l'origine de l'issue. C'est ici qu'on raconte le scénario dans lequel le problème est apparu :
- Dans quelles circonstances l'a-t-on rencontré ?
- Quelles manipulations ont permis de le reproduire ?
- Qu'est-ce qui ne fonctionne pas actuellement ?
- Qu'est-ce qui nous gêne dans le comportement actuel ?

On peut aussi détailler le cheminement de réflexion qui nous a amenés à ouvrir cette issue.

### Ce qu'on a actuellement

On montre concrètement le fonctionnement actuel du code. C'est le moment de mettre les extraits de code concernés (snippets), avec éventuellement les noms des fichiers et des classes. L'objectif est que les autres puissent rapidement identifier la partie du code dont on parle.

### Ce qu'on souhaiterait avoir concrètement

Décrire le résultat attendu, de manière aussi précise que possible. Ici, on descend d'un niveau par rapport à l'objectif général : on décrit ce qu'on veut réellement modifier dans le code. Par exemple :
- Remplacer un pointeur par une référence.
- Modifier la signature d'une méthode.
- Ajouter un attribut à une classe.
- Changer le comportement d'une fonction dans un cas particulier.

Le pourquoi a normalement déjà été expliqué dans l'objectif général et la description du problème. On peut donc se concentrer sur le quoi.

### Les premières idées de solution

On note les premières pistes auxquelles on a pensé. Elles ne sont pas forcément abouties, et c'est ok ! Le but est de donner une première direction à la discussion. On peut notamment préciser :
- Les classes ou fonctions qu'on pense devoir modifier.
- Les approches envisagées.
- Les éventuelles difficultés anticipées.
- Les alternatives auxquelles on a pensé.

### Éventuellement les personnes concernées

Si certaines personnes sont directement concernées par le sujet, on peut les mentionner avec le `@`. Cela permet de solliciter les bonnes personnes dès l'ouverture du ticket.

## Un exemple de structure d'une issue

On peut reprendre directement cette structure lors de la création d'un ticket :

````
# Objectif général

Décrire le besoin auquel répond cette issue.

## Problème rencontré

Décrire le scénario dans lequel le problème apparaît,
les difficultés rencontrées et le cheminement de réflexion.

## Ce qu'on a actuellement

Décrire le fonctionnement actuel.

```Go
// Extrait de code concerné
```

## Ce qu'on souhaiterait avoir concrètement

Décrire précisément le comportement ou la modification attendue.

## Premières idées de solution

- Première piste envisagée
- Deuxième piste éventuelle

---
Avis @personne ?
````

## 3. Comment utiliser les commentaires d'une issue ?

L'issue, c'est aussi surtout les commentaires, qui servent à faire avancer la réflexion collective. On ne cherche pas simplement à recueillir des avis du type "je suis d'accord" ou "je ne suis pas d'accord".

=> **On essaie d'apporter quelque chose de directement exploitable pour la suite.**

Quelques exemples de contributions utiles :
- Proposer un snippet de code qui illustre une solution.
- Partager une documentation ou une source technique pertinente.
- Expliquer pourquoi une approche risque de poser problème.
- Proposer une alternative concrète.
- Pointer une dépendance avec une autre partie du projet.
- Poser une question qui permet de préciser le besoin.

L'idée est que chaque commentaire apporte une brique supplémentaire à la résolution du problème. Une discussion doit progressivement nous permettre de converger vers une solution sur laquelle on s'est mis d'accord.


## 4. Workflow Git : de l'issue à la MR

On essaie de garder un workflow commun pour éviter les modifications qui arrivent directement dans `main` sans discussion ni relecture.

### La branche main est protégée

**On ne travaille jamais directement sur `main`.** La branche `main` doit toujours contenir une version du projet intégrée et relue. Toutes les modifications passent par une Merge Request (MR). Cela permet de :
- Relire les modifications avant leur intégration.
- Discuter des choix d'implémentation.
- Vérifier que le code correspond à l'issue.
- Garder une trace des changements apportés.

### Créer une branche feature

Une fois qu'on s'est mis d'accord sur une issue et sur la solution à implémenter, on crée une branche dédiée à partir de `main`. On utilise une branche par fonctionnalité ou par tâche. Par exemple :

```bash
git checkout main
git pull
git checkout -b feature/42-modification-pointeurs
````

Quelques exemples de nomenclatures :

```
feature/42-modification-pointeurs
fix/51-correction-calcul
docs/63-mise-a-jour-readme
```

Attention : le numéro de l'issue doit apparaître dans le nom de la branche pour retrouver facilement le ticket associé.

### Faire ses commits

Les commits doivent référencer l'issue concernée. On ajoute `#numeroDuTicket` dans le message de commit (ou directement l'URL de l'issue). Exemples :

```
git commit -m "Modification de la gestion des pointeurs #42"
git commit -m "Correction du calcul des coordonnées #51"
```

Les commits seront alors référencés directement sur le ticket.

### Créer une MR

Une fois le travail terminé, on ouvre une MR vers `main`. Dans la description de la MR, on ajoute :

```
close #42
```

Ou directement l'URL de l'issue.

Ça permet de lier la MR à son issue et de fermer automatiquement cette dernière lors du merge.

### Reviewer et assignee

Chaque MR doit avoir au minimum :
- **1 reviewer** : la personne qui relit le code et vérifie les modifications.
- **1 assignee** : la personne qui prend en charge la MR, répond aux commentaires et s'occupe de son intégration.

Le reviewer et l'assignee peuvent être deux personnes différentes.

La relecture doit permettre de vérifier que :
- Le code répond bien à l'objectif de l'issue.
- Les modifications restent cohérentes avec le reste du projet.
- Les éventuels cas particuliers ont été pris en compte.
- Les décisions prises dans l'issue ont bien été respectées.

On ne merge pas directement une MR sans avoir effectué la relecture.

## Récapitulatif du workflow

Pour résumer, voici le cheminement qu'on essaie de suivre pour chaque nouvelle tâche :

1. **Créer une issue** : décrire le problème, le contexte et le résultat attendu.
2. **Ajouter l'issue au Backlog SR08** : renseigner les tags et les personnes concernées.
3. **Discuter de la solution** : utiliser les commentaires pour proposer des pistes concrètes et se mettre d'accord.
4. **S'assigner l'issue** : la personne qui prend en charge la tâche devient assignee.
5. **Créer une branche feature** : partir de `main` avec le numéro de l'issue dans le nom de la branche.
6. **Implémenter la solution** : faire des commits qui référencent l'issue.
7. **Ouvrir une MR** : ajouter `close #numeroDuTicket` dans sa description.
8. **Faire relire le code** : renseigner au moins un reviewer et un assignee.
9. **Merger dans main** : une fois la relecture effectuée et les éventuelles remarques traitées.

**En résumé : on réfléchit dans les issues, on implémente dans les branches et on valide dans les MR.**