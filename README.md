Go Reloaded est un projet développé en Go. Il consiste à lire un fichier texte, à corriger son contenu selon plusieurs règles, puis à écrire le résultat dans le terminal.

Le programme traite les messages en anglais et applique cinq niveaux de correction dans un ordre précis : conversion des nombres, modification de la casse, ponctuation, guillemets et articles anglais.
Prérequis
- Go installé sur l'ordinateur.
- Un terminal pour exécuter les commandes.
Le projet utilise uniquement la bibliothèque standard de Go.
Lancer le programme
Depuis le dossier du projet, exécuter :
go run main.go
Exemple :
Le programme lit test.txt, effectue les corrections, puis écrit le texte corrigé dans le terminal.
Fonctionnement :
Chaque ligne est traitée séparément, sans modifier les autres lignes. Le nombre et l'ordre des lignes sont conservés, y compris les lignes vides. Les espaces et tabulations successifs sont remplacés par un espace, et les espaces en début ou fin de ligne sont supprimés.

Les corrections sont appliquées dans cet ordre : niveaux 1 et 2, puis niveau 3, niveau 4 et niveau 5.
Niveau 1 — Conversion des nombres
Le programme reconnaît deux modificateurs placés après un nombre :
- (hex) : convertit un nombre hexadécimal (base 16) en décimal (base 10).
- (bin) : convertit un nombre binaire (base 2) en décimal (base 10).
Si le nombre n'est pas valide, il reste inchangé. Le modificateur reconnu disparaît dans tous les cas.
Entrée : 1E (hex) files were added
Sortie : 30 files were added

Entrée : It has been 10 (bin) years
Sortie : It has been 2 years
Niveau 2 — Modification des majuscules et minuscules
Les modificateurs suivants agissent sur les mots précédents :
- (up) : transforme les lettres en majuscules.
- (low) : transforme les lettres en minuscules.
- (cap) : met la première lettre en majuscule et les suivantes en minuscules.
Entrée : Ready, set, go (up) !
Sortie : Ready, set, GO!

Entrée : This is so exciting (up, 2)
Sortie : This is SO EXCITING
Niveau 3 — Ponctuation
Les signes concernés sont : . , ! ? : ;.

Les signes de ponctuation séparés sont collés à l'élément précédent, sans espace. Plusieurs signes consécutifs restent groupés. Un espace est placé après le groupe de ponctuation s'il est suivi d'un autre élément. Une ponctuation déjà intégrée à un mot, comme dans 3.14, n'est pas modifiée.
Entrée : I was thinking ... You were right
Sortie : I was thinking... You were right

Entrée : Really ! ? No way
Sortie : Really!? No way
Niveau 4 — Guillemets simples
Les guillemets simples ' sont associés par paires, de gauche à droite, sur chaque ligne. Les espaces entre les guillemets et les mots qu'ils entourent sont supprimés.

Une apostrophe située entre deux lettres, comme dans don't, ne doit pas être modifiée. Un guillemet sans partenaire reste inchangé.
Entrée : He said ' hello ' !
Sortie : He said 'hello'!
Niveau 5 — Articles anglais
Le programme corrige les articles a, A, an et An selon la première lettre du mot suivant :
- Devant une voyelle (a, e, i, o, u), il utilise an ou An.
- Devant une autre lettre, il utilise a ou A.
- Si le mot suivant commence par un chiffre ou un symbole, ou s'il n'existe pas, l'article ne change pas.
La casse initiale de l'article est conservée. Les guillemets simples situés au début du mot suivant sont ignorés pour cette vérification.
Entrée : She ate an banana and a orange
Sortie : She ate a banana and an orange

Entrée : It was a ' epic ' fail
Sortie : It was an 'epic' fail
