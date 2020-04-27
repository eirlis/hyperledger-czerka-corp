# hyperledger-czerka-corp
Kafka-based Hyperledger Fabric network built on top of balance-transfer sample.

## Disclaimer
This is app was created in demo & training purposes to discover different features of Hyperledger Fabric. *Use in production at your own risk.*

## Use Case
Czerka is a galaxy-spanning business corporation ranging from consumer food products to military weapons. 
It is one of the wealthiest and most successful economic enterprises in operation, conducting commerce on virtually every civilized planet and–as an owner of multiple star systems and employer of several billion individuals–has representation in the Republic Senate.
It's also being known as kind of a criminal entity during the Jedi Civil Wars because of the cooperation with Sith and slavery business.

## Purpose
This is a common example of Hyperledger Fabric usage which covers almost every field of ledger technology usage from supply chain to property or identity certification.
Why Czerka? Probably because this repo belongs to a Star Wars fan who was lazy enough to create own use case. 

## Getting started 
I've modified existing scripts as well as added new ones, which I found pretty helpful during the development. Here is a step-by-step guide.

1. Before you start the network, generate all the neccessary artifacts. I've updated the way of generating `docker-compose.yaml` file. Now it doesn't contain hard-coded CA private keys and is generated from `docker-compose-template.yaml`. 
```
generateArtifactsUpdated.sh
```

2. After the artifacts was generated, run the Fabric network and NodeJS app. This will launch the network locally on port 4000.
```
runApp.sh
```

3. To use API endpoints (especially if you run the network several times), you'll need to check whether your admin keys are up-to-date. If not, replace them using the respective script.
```
replaceAdminKeystore.sh
```

4. After you shut the network down don't forget to clean up credentials.
```
credentialsCleanup.sh
```
