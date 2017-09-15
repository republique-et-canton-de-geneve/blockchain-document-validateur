Prototype Proof of Concept Registre du Commerce
===============================================

---------

I. Backlog
--------------
Interface de génération d'extrait
--------------------------------
 - Interface de téléversement des extraits.
 => Cette interface permet le téléversement d'extrait et téléchargement de reçu au format PDF ainsi que la signature et l'horodatage des extraits.
 
 Interface de validation d'extrait
--------------------------------
 - Interface de téléversement d'extrait et reçu.
 => Cette interface permet le téléversement d'extrait et reçu au format PDF et indique en retour la validité.
 La validité est défini par le fait que le reçu corresponde à l'extrait et que le signataire de la transaction Ethereum est bien celle défini dans le backend de validation (ici, le registre du commerce).

II. Documentation technique
--------------------------------

Interfaces
------------
Dans le cadre du PoC les interfaces sont des écrans web accessible via un navigateur de type Chrome ou Firefox.

Horodateur:

L'envoi des extraits[^extrait] se fait via un téléversement par l'officier public sur l'interface, une fois les extraits téléversés, il est possible de télécharger le reçu de la signature+horodatage[^trustedtimestamping] des-dit extraits.

[^extrait]: Il a été déterminé que que les extraits ne peuvent venir uniquement de la base de donnée interne du registre du commerce car la base de donnée publique n'est pas mis à jour en temps réel.

[^trustedtimestamping]: https://en.wikipedia.org/wiki/Trusted_timestamping

Validateur:

L'envoi des extraits et reçus se fait via un téléversement par un utilisateur sur l'interface, en retour un message clair de couleur vert indique si l'horodatage est confirmé ou par un message sur fond rouge si la validation de l'horodatage a échoué.

Backend
-------
Les backends sont séparés en deux micro-service, un horodateur qui possède la clef privé et le validateur qui est verrouillé sur l'adresse issu de la clef privé de l'horodateur.
Ces deux paramètres sont néanmoins distincts et configurables.
Les fonctionnalités de l'interfaces sont volontairement minimal afin d'être conçu comme une brique simple se reposant sur les backends.
Les backends s'interface avec la blockchain et les différentes transactions qui sont émis sur celle-ci et aura la gestion des clefs associés à ceux-ci (émission/signature de transaction, gestion des identités blockchain, émission de reçu, ...).
Une Application Programming Interface (API) expose les différentes fonctionnalités, une documentation peut-être fournis au format html.

ChainPoint 2.1[^chainpoint] est la technologie implémenté car plus rapide pour les besoin PoC, il est néanmoins à noter qu'il existe des alternatives en cours de développement de type ChainPoint v3 ou OpenTimestamps[^opentimestamps].

[^chainpoint]: https://github.com/chainpoint/whitepaper/blob/master/chainpoint_white_paper.pdf

[^opentimestamps]: https://opentimestamps.org

How to run the prototype ?
----------

1. Install Docker[^docker] and Docker Compose[^dockercompose] (Window 10, macOS, Linux, ...)
2. Install git[^git]
3. Clone the repository with sub-modules
``` git clone --recursive https://github.com/Magicking/rc-ge-ch-pdf.git ```
or
``` git clone --recursive https://github.com/Magicking/rc-ge-validator.git ```
3. Place yourself within the directory containing this document
4. Build the containers using docker-compose
``` docker-compose build ```
5. Edit environnements variables (see below) according to your needs[^dockercomposespec] in the docker-compose.yml
6. Run the prototype
```docker-compose up```
7. Access interfaces at 
 - http://127.0.0.1:8001 for the timestamping service
 - http://127.0.0.1:8002 for the validator service
[^docker]: https://docs.docker.com/engine/installation/#server

[^dockercompose]: https://docs.docker.com/compose/install/

[^git]: https://git-scm.com/book/en/v2/Getting-Started-Installing-Git

[^dockercomposespec]: https://docs.docker.com/compose/compose-file/

Environment variables
--------------------

 - WS_URI is a URI pointing to an Ethereum RPC endpoint[^RPC] (e.g: http://localhost:8545)
 => The Ethereum node must be fully sync prior to use.
 - LOCKED_ADDR is an Ethereum address used by the validate service to verify the transaction signer of the receipt (e.g: 0x533a245f03a1a46cacb933a3beef752fd8ff45c3)
 - PRIVATE_KEY is an Ethereum private key used by the timestamping service to sign the transaction used to insert the merkle root. (e.g: 18030537dbdd38d0764947d40bed98fc4d2a21af82765a7de7b13d2e4076773c)
 => The account related to the private key must be funded prior to use.
[^RPC]: https://github.com/ethereum/wiki/wiki/JSON-RPC

Frameworks & softwares
------------

 - [Git](https://git-scm.com)
 - [OpenAPI](https://www.openapis.org/)
 - [Golang](https://golang.org/)
 - [Go-Ethereum](https://geth.ethereum.org/)
 - [Twitter Bootstrap](http://getbootstrap.com/)
 - [Docker](https://www.docker.com/)
 - [ChainPoint](https://chainpoint.org)
 - [Caddy](https://caddyserver.com)
