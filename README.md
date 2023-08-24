# Command line tools for [KubeBuild](https://github.com/EikenDram/kube-build) project

| NAME     | DESCRIPTION
|----------|-----------------------------
| keycloak | Tool for managing [Keycloak](https://www.keycloak.org/) server

## Keycloak

Command line tool for managing Keycloak server by processing text files

| COMMAND | DESCRIPTION
|---------|-------------------------------
| auth    | Authorize in Keycloak server
| add     | Add users to Keycloak from a text file
| update  | Update users in Keycloak from a text file
| delete  | Delete users from Keycloak from a text file

## Keycloak API

In order to provide admin-cli secret need to switch client auth on, allow service account and copy the token from credentials
```sh
curl --location --request POST 'http://localhost:8080/realms/master/protocol/openid-connect/token' --header 'Content-Type: application/x-www-form-urlencoded' --data 'grant_type=client_credentials' --data 'client_id=admin-cli' --data 'client_secret=<client secret>'
```

Duration is 1 minute, can be changed in settings

Get all users:
```sh
curl -k -X GET http://keycloak.k3s.local/admin/realms/master/users -H "Authorization: Bearer <access token>"
```

Get user information:
```sh
curl -k -X GET http://keycloak.k3s.local/admin/realms/master/users?exact=true&username=username -H "Authorization: Bearer <access token>"
```

Change password for user:
```sh
curl -k -X PUT http://keycloak.k3s.local/admin/realms/master/users/<id>/reset-password -H "Authorization: Bearer <access token>" -H "Content-type: application/json" --data '{"type":"password","value":"<new password>","temporary":false}"'
```

Add user:
```sh
curl -k -X POST http://keycloak.k3s.local/admin/realms/master/users/ -H "Content-Type: application/json" -H "Authorization: Bearer <access token>" --data '{"firstName":"xyz","lastName":"xyz", "username":"xyz123","email":"demo2@gmail.com", "enabled":"true","credentials":[{"type":"password","value":"test123","temporary":false}]}'
```