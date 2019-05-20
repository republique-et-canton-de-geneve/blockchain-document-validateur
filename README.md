
# Prototype Registre du Commerce



## I. Backlog

## Extract Generation Interface

This interface allows upload an extract and proceed with a trusted timestamping of these extract and provide a PDF receipt as a proof. Trusted timestamping, is the process of securely keeping track of the creation and modification time of a document. Security here means that no one—not even the owner of the document—should be able to change it once it has been recorded provided that the timestamper's integrity is never compromised.

## II. Technical Documentation

## Interfaces

As part of the PoC, the front is a web screens accessible via a browser type Chrome or Firefox.

-   **Validator**:

The sending of the extracts and receipts is done via an upload by a user on the interface, in return a clear message of green color indicates if the timestamp is confirmed or by a message on a red background if the validation of the timestamp has failed.

## Backend

The backends communicate with the blockchain, broadcast the various transactions (issue / signature of a transaction, issue receipts ...) and manage the associated keys by checking the validity of the signer and his key.

## How to run the prototype ?

1.  Install Docker[^docker](https://docs.docker.com/engine/installation/#server)  and Docker Compose[^dockercompose](https://docs.docker.com/compose/install/)  (Window 10, macOS, Linux, ...)
2. Edit environments variables (see below) according to your needs[^dockercomposespec] (https://docs.docker.com/compose/compose-file/)  in the docker-compose-prod.yml 
3.  Build the set of containers by running `docker-compose -f docker-compose-prod.yml up -d` 
7.  Access interface at

-   [http://127.0.0.1:8001/ctivalidator/](http://127.0.0.1:8001/)  for the timestamping service

## How to check node disponibility ?

1.   Access API interface at  [http://127.0.0.1:8001/ctivalidator/api/sonde](http://127.0.0.1:8001/api/sonde)
## Validator API Environment variables

Mandatory :

-   WS_URI is a URI pointing to an Ethereum RPC endpoint (e.g:  [http://localhost:8545](http://localhost:8545/)) => The Ethereum node must be fully sync prior to use.
-   LOCKED_ADDR is an Ethereum address used by the validate service to verify the transaction signer of the receipt (e.g: 0x533a245f03a1a46cacb933a3beef752fd8ff45c3)
-   TIMEOUT is a UNIX timestamp to manage the expiration time of a private key. To generate one you could use the website [http://www.timestamp.fr/](http://www.timestamp.fr/)

Optional :

-   HTTP(S)_PROXY are environment variables used to specified a forward proxy for connection to pass through.

## HTTPS support

HTTPS support is provided via the docker images `jwilder/nginx-proxy`. The `nginx-proxy` image faces Internet and dispatches requests to the
concerned service. Services that are reached from the Internet must have the following environment variables :  
   
  - `VIRTUAL_HOST` : The domain name associated to the service. 
  
Administrators must add an A record to their DNS configuration that points to the IP of the machine that hosts
`nginx/proxy`. 
   

## Webapp Environment variables

-   `KEY_NAME` is the name given to the cert & key files used by the Service Provider (e.g:  myservice ). When updated, the names in the *volumes* tag of the *docker-compose-prod.yml* need to be updated too.
-   `IDP_METADATA` is the public url where the SAML package gets the Identity Provider metadata.
-   `SP_URL` is the root url of the Service Provider
-   `API_HOST` is the hostname of the API. It is based on the docker image name.
-   `MAIN_URI` is used to specify the required prefix for the webapp. Default is ctivalidator.
-   `VIRTUAL_HOST` is used by `nginx-proxy` to identify the domain name associated to the Webapp's Docker.

## Frameworks & softwares

-   [Git](https://git-scm.com/)
-   [OpenAPI](https://www.openapis.org/)
-   [Golang](https://golang.org/)
-   [Go-Ethereum](https://geth.ethereum.org/)
-   [Twitter Bootstrap](http://getbootstrap.com/)
-   [Docker](https://www.docker.com/)
-   [ChainPoint](https://chainpoint.org/)

## Disclaimer

The State of Geneva disclaims all liability for any use of all or part of the code, in particular due to programming defects.
