/*Package restapi RCG horodatage

RCG horodatage est un service qui permet l'horodatage numérique via
sur la blockchain Ethereum.
Le principe est d'envoyer des fichiers qui sont ensuite passer dans
une fonction hachage SHA3-256. Les « hash » sont ensuite intégrés
dans un arbre de Merkle dont la racine est inséré dans une
transaction blockchain, l'(es) adresse(s) signant la transaction
identifie le Registre du Commerce, c'est une information qui doit
être publique.



    Schemes:
      http
    Host: localhost
    BasePath: /
    Version: 0.1.0

    Consumes:
    - application/json


    Produces:
    - application/json

    - application/octet-stream


swagger:meta
*/
package restapi
