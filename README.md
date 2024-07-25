# Entra client

## Objectifs

Le but de ce client entra (ecli) est multiple :
* Explorer les API Entra
* Fournir des modules go permettant de manipuler les objets IAM (utilisateurs, applications, groupes...)
  Avec pour objectif de pouvoir faire du provisionning complexe d'application, ou des opérations de synchronisations.
* Fournir un outil ligne de commande éventuellement utilisable pour les opérations simples


## Organisation du code

Le client s'appuie sur cobra (github.com/spf13/cobra) et donc main.go ne fait qu'appeler les commandes définies dans /cmd

Les commandes utilisent elles les packages httpClient et sdkClient qui sont des implémentations de client.Service.
Ces 2 packages httpClient et sdkClient sont les deux "moteurs" (--engine) que l'application peut utiliser pour accéder aux API Entra.
* httpClient est un moteur bas niveau qui s'appuie sur net/http pour faire des requêtes REST.
* sdkClient s'appuie sur le sdk microsoft qui est de plus haut niveau

Ces deux moteurs utilisent les models qui sont dans internal/models
Et httpClient utilise une fine surcouche de net/http qui est définie dans pkg/rest


## Build

```
make build
```

Si les make tools ne sont pas installés:
```
go build -o ecli
```

## Tests

```
make test
```

Si les make tools ne sont pas installés:
```
go test  ./...
```

## Utilisation

```
mv env.sample .env
```

Ensuite modifier .env en renseignant les bonnes valeurs

L'aide intégrée :

```
./ecli --help
```

constitue un bon point d'entrée, mais les commandes son assez simples.

Quelques utilisations possibles :

Afficher l'id et le displayName des 50 premiers groupes
```
./ecli group list --select id,displayname --top 50
```

Afficher tous les utilisateurs (en contournant la pagination server side d'Entra). Qui sont nombreux...
```
./ecli user list
```

Afficher les informations d'une application spécifique sélectionnée par son ID
```
./ecli application get --id 4338fbfb-83b6-44be-ab56-7bb5e1f91b86
```

Afficher les informations des applicatoins matchant certains critères
```
./ecli application list --search displayname:Portal
```

Créer une application (entité de base uniquemenent)
```
./ecli application create --engine sdk --post '{"displayName": "test API POST AA"}'
```

Supprimer une application par son id
```
./ecli application delete --id 5128baa5-03b7-49f8-9f06-d1f1464eff1e
```


## Ressources
* [Entra API](https://learn.microsoft.com/en-us/graph/azuread-identity-access-management-concept-overview)
* [Authentication](https://learn.microsoft.com/en-us/graph/auth/auth-concepts)
* [Tutoriel Go d'appel de l'API Graph](https://github.com/microsoftgraph/msgraph-training-go)
* [http-client tutorial](https://www.sohamkamani.com/golang/http-client/)
* [Graph API explorer](https://developer.microsoft.com/en-us/graph/graph-explorer)
* [Configuring SAML application through API](https://learn.microsoft.com/en-us/graph/application-saml-sso-configure-api?tabs=http%2Cpowershell-script)

Golang Modules

* [Microsoft graph API module](https://github.com/microsoftgraph/msgraph-sdk-go)
* Azure SDK for go 
  * github.com/Azure/azure-sdk-for-go/sdk/azidentity
  * github.com/Azure/azure-sdk-for-go/sdk/azcore/policy
  * [Microsoft documentation for Golang AzureAD MSAL module](https://github.com/AzureAD/microsoft-authentication-library-for-go)

Liste des autorisations 

* https://entra.microsoft.com/#view/Microsoft_AAD_RegisteredApps/ApplicationMenuBlade/~/CallAnAPI/appId/ce306f4f-63ea-4ae3-98ce-1dba7572e990/isMSAApp~/false

