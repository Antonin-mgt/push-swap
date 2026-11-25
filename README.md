# Push-Swap

Projet d'algorithmie en Go.

Le projet contient deux programmes :

- **push-swap** : reçoit une liste d'entiers, calcule une suite d'instructions qui trie la pile `a` (le plus petit nombre au sommet) et l'affiche, une instruction par ligne.
- **checker** : reçoit la même liste d'entiers en argument, lit des instructions sur l'entrée standard, les exécute, puis affiche `OK` si la pile `a` est triée et la pile `b` vide, `KO` sinon.

Seuls les packages de la bibliothèque standard de Go sont utilisés.

## Les instructions

| Instruction | Effet |
|---|---|
| `sa` / `sb` / `ss` | échange les deux premiers éléments de `a` / `b` / des deux piles |
| `pa` | prend le premier élément de `b` et le pose sur `a` |
| `pb` | prend le premier élément de `a` et le pose sur `b` |
| `ra` / `rb` / `rr` | décale vers le haut : le premier élément devient le dernier |
| `rra` / `rrb` / `rrr` | décale vers le bas : le dernier élément devient le premier |

## Structure du projet

```
push-swap/
├── go.mod
├── piles/        les deux piles, les 11 instructions, EstTriee, et le tri (Trier)
├── parsing/      lecture et validation des arguments
├── checker/      programme checker
└── push-swap/    programme push-swap
```

- `piles/piles.go` : le type `Piles` (piles `A` et `B`, et la liste `Operations` des instructions exécutées) et toutes les instructions.
- `piles/tri.go` : la fonction `Trier`, qui génère la suite d'instructions.
- `parsing/parsing.go` : la fonction `Parser`, qui convertit le texte reçu en entiers et détecte les erreurs.
- Les dossiers `piles` et `parsing` contiennent leurs fichiers de tests (`*_test.go`).

## Compilation

Depuis la racine du projet :

```
go build -o bin/checker ./checker
go build -o bin/push-swap ./push-swap
```

## Utilisation

```
./bin/push-swap "2 1 3 6 5 8"
```

Afficher le nombre d'instructions :

```
ARG="4 67 3 87 23"; ./bin/push-swap "$ARG" | wc -l
```

Vérifier le résultat avec le checker :

```
ARG="4 67 3 87 23"; ./bin/push-swap "$ARG" | ./bin/checker "$ARG"
OK
```

Le checker seul, avec des instructions écrites à la main :

```
printf 'rra\npb\nsa\nrra\npa\n' | ./bin/checker "3 2 1 0"
OK
```

## Gestion des erreurs

Dans les deux programmes, `Error` est écrit sur la sortie d'erreur dans les cas suivants :

- un argument n'est pas un entier ;
- il y a des doublons ;
- (checker) une instruction est inconnue.

Sans argument, les programmes n'affichent rien. Si la liste est déjà triée, `push-swap` n'affiche aucune instruction.

## Algorithme de tri

La méthode est simple :

1. Tant que la pile `a` n'est pas triée, on cherche le plus petit élément, on l'amène au sommet avec `ra` ou `rra` (le sens le plus court), puis on le pousse sur `b` avec `pb`.
2. Quand `a` est triée, on remonte tous les éléments de `b` sur `a` avec `pa`.

Les éléments poussés sur `b` sont les plus petits, dans l'ordre croissant : `b` se retrouve donc rangée du plus grand au plus petit, et la remontée avec `pa` remet `a` dans le bon ordre.

Cette méthode donne toujours un résultat correct, mais le nombre d'instructions grandit vite avec la taille de la liste (environ 1 400 instructions pour 100 nombres). Une piste d'amélioration serait de trier directement les listes de 3 éléments, et d'utiliser `rr` et `rrr` quand cela évite des instructions.

## Tests

```
go test ./...
```
