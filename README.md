liens du canva: https://canva.link/h1mmpq2kl8o786j

# Go Reloaded

Go Reloaded est un projet écrit en Go. Il lit le contenu d'un fichier texte, applique cinq niveaux de correction, puis affiche le résultat dans le terminal.

Le programme est conçu pour traiter du texte en anglais.

## Prérequis

- Go installé sur l'ordinateur
- Un terminal

Le projet utilise uniquement la bibliothèque standard de Go.

## Utilisation

Depuis la racine du projet, exécutez :

```bash
go run main.go
```

Le programme lit le fichier `test.txt`, applique les corrections dans l'ordre prévu, puis affiche le texte modifié.

## Fonctionnement général

- Chaque ligne est traitée séparément.
- Le nombre et l'ordre des lignes sont conservés, y compris les lignes vides.
- Les espaces et tabulations successifs sont remplacés par un seul espace.
- Les espaces au début et à la fin de chaque ligne sont supprimés.
- Les corrections sont appliquées dans cet ordre :
  1. Conversion des nombres
  2. Modification de la casse
  3. Ponctuation
  4. Guillemets
  5. Articles anglais

## Niveau 1 — Conversion des nombres

Le programme reconnaît les modificateurs placés après un nombre :

| Modificateur | Action |
| --- | --- |
| `(hex)` | Convertit un nombre hexadécimal en nombre décimal |
| `(bin)` | Convertit un nombre binaire en nombre décimal |

Le modificateur reconnu disparaît dans tous les cas. Si le nombre n'est pas valide, il reste inchangé.

### Exemples

```text
Entrée : 1E (hex) files were added
Sortie : 30 files were added
```

```text
Entrée : It has been 10 (bin) years
Sortie : It has been 2 years
```

## Niveau 2 — Modification de la casse

Les modificateurs agissent sur les mots précédents :

| Modificateur | Action |
| --- | --- |
| `(up)` | Transforme les lettres en majuscules |
| `(low)` | Transforme les lettres en minuscules |
| `(cap)` | Met la première lettre en majuscule |
| `(up, N)` | Met les `N` mots précédents en majuscules |
| `(low, N)` | Met les `N` mots précédents en minuscules |
| `(cap, N)` | Capitalise les `N` mots précédents |

### Exemples

```text
Entrée : Ready, set, go (up) !
Sortie : Ready, set, GO!
```

```text
Entrée : This is so exciting (up, 2)
Sortie : This is SO EXCITING
```

## Niveau 3 — Ponctuation

Les signes concernés sont : `.`, `,`, `!`, `?`, `:` et `;`.

- Une ponctuation séparée du mot précédent est collée à celui-ci.
- Plusieurs signes consécutifs restent groupés.
- Un espace est ajouté après le groupe de ponctuation lorsqu'un autre élément suit.
- La ponctuation déjà intégrée à un mot, comme dans `3.14`, n'est pas modifiée.

### Exemples

```text
Entrée : I was thinking ... You were right
Sortie : I was thinking... You were right
```

```text
Entrée : Really ! ? No way
Sortie : Really!? No way
```

## Niveau 4 — Guillemets simples

Les guillemets simples `'` sont associés par paires, de gauche à droite, sur chaque ligne :

1. Le premier guillemet ouvre.
2. Le deuxième guillemet ferme.
3. Le troisième guillemet ouvre.
4. Et ainsi de suite.

- Les espaces situés juste après un guillemet ouvrant sont supprimés.
- Les espaces situés juste avant un guillemet fermant sont supprimés.
- Une apostrophe située entre deux lettres, comme dans `don't`, n'est pas modifiée.
- Un guillemet sans partenaire reste inchangé.

### Exemple

```text
Entrée : He said ' hello ' !
Sortie : He said 'hello'!
```

## Niveau 5 — Articles anglais

Le programme corrige les articles `a`, `A`, `an` et `An` selon la première lettre du mot suivant :

- Devant une voyelle (`a`, `e`, `i`, `o`, `u`) : utilisez `an` ou `An`.
- Devant une autre lettre : utilisez `a` ou `A`.
- Si le mot suivant commence par un chiffre ou un symbole, ou si l'article n'a pas de mot suivant : l'article reste inchangé.
- La casse initiale de l'article est conservée.
- Les guillemets simples placés au début du mot suivant sont ignorés pour cette vérification.

### Exemples

```text
Entrée : She ate an banana and a orange
Sortie : She ate a banana and an orange
```

```text
Entrée : It was a ' epic ' fail
Sortie : It was an 'epic' fail
```
