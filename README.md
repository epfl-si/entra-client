# Entra client

## Prerequisites

### Entra configuration (prerequisite)


[Inscrire une application](https://learn.microsoft.com/fr-fr/entra/identity-platform/quickstart-register-app#register-an-application)

Créer un secret (dans Certificats et secrets)

Copier sa valeur

Le tenant id est : b6cddbc1-2348-4644-af0a-2fdb55573e3b (Vue d'ensemble, id du client)

Utiliser MSAL pour Go (pour récupérer un token/configurer un client)

Utiliser http pour appeler l'API REST avec le token

### .env configuration

```
mv env.sample .env
```

Ensuite modifier .env en renseignant les bonnes valeurs


## Ressources
* [Entra API](https://learn.microsoft.com/en-us/graph/azuread-identity-access-management-concept-overview)
* [Authentication](https://learn.microsoft.com/en-us/graph/auth/auth-concepts)
* [Golang module for AzureAD MSAL](https://github.com/AzureAD/microsoft-authentication-library-for-go)
* [Microsoft documentation for Golang AzureAD MSAL module](https://github.com/AzureAD/microsoft-authentication-library-for-go)
* [Tutoriel Go d'appel de l'API Graph](https://github.com/microsoftgraph/msgraph-training-go)
* [http-client tutorial](https://www.sohamkamani.com/golang/http-client/)
* [Graph API explorer]()
