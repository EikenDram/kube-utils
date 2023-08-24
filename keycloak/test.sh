#!/bin/bash

KEYCLOAK_HOST=http://keycloak.swk3s
TOKEN=$(curl -s -X POST $KEYCLOAK_HOST/realms/master/protocol/openid-connect/token -H 'Content-Type: application/x-www-form-urlencoded' --data 'grant_type=client_credentials' --data 'client_id=admin-cli' --data 'client_secret=jKYLLZusOC3L6TrbUSwmWGFRYsktuN2a'  | jq -r ".access_token" ;)

curl -s -X GET -H "Authorization: Bearer $TOKEN" -H "Content-type: application/json" $KEYCLOAK_HOST/admin/realms/master/users?realm=cluster&exact=true&username=k3sadmin ;
#ADMIN_USER=$(curl -s -X GET -H "Authorization: Bearer $TOKEN" -H "Content-type: application/json;charset=UTF-8" -H 'Accept: application/json' "$KEYCLOAK_HOST/auth/admin/realms/master/users" | jq -r '.[] | select(.username=="admin") | .id' )
#curl -s -X PUT -H "Authorization: Bearer $TOKEN" -H "Content-type: application/json;charset=UTF-8" -H 'Accept: application/json' "$KEYCLOAK_HOST/auth/admin/realms/master/users/$ADMIN_USER_ID/reset-password" -d "{\"type\":\"password\",\"value\":\"$ADMIN_USER_NEW_PASSWORD\",\"temporary\":false}"