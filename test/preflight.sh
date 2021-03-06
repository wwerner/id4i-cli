#!/usr/bin/env bash
set -o pipefail

# User Name Generation
FIRST=$(sort --random-sort test-user-names.1.txt | head -n 1)
SECOND=$(sort --random-sort test-user-names.2.txt | head -n 1)
USER="${FIRST}${SECOND}"
EMAIL="${USER}@bluerain.de"
PASSWORD=$(head /dev/urandom | LC_ALL=C tr -dc 'A-Za-z0-9' | head -c20)

echo "USER=$USER" > .preflightData
echo "PASSWORD=$PASSWORD" >> .preflightData


# Register User
REGISTRATION_DATA='{"username":"'${USER}'","password":"'${PASSWORD}'","email":"'${EMAIL}'"}'
curl -s 'https://id4i-develop.herokuapp.com/account/registration' \
    -H 'Pragma: no-cache' \
    -H 'Origin: https://id4i-develop.herokuapp.com' \
    -H 'Accept-Encoding: gzip, deflate, br' \
    -H 'Accept-Language: en' \
    -H 'Content-Type: application/json' \
    -H 'Accept: application/json' \
    -H 'Cache-Control: no-cache' \
    -H 'X-ID4i-Client: ID4i CLI Test' \
    --data-binary ${REGISTRATION_DATA}

echo
echo Registered ${USER}

# Login
LOGIN_DATA='{"login":"'${USER}'","password":"'${PASSWORD}'"}'
AUTHORIZATION=$(curl -s -i 'https://id4i-develop.herokuapp.com/login' \
    -H 'Pragma: no-cache' \
    -H 'Origin: https://id4i-develop.herokuapp.com' \
    -H 'Accept-Encoding: gzip, deflate, br' \
    -H 'Accept-Language: en' \
    -H 'Content-Type: application/json' \
    -H 'Accept: application/json' \
    -H 'Cache-Control: no-cache' \
    -H 'X-ID4i-Client: ID4i CLI Test' \
    --data-binary ${LOGIN_DATA} | grep -Fi Authorization | sed -e 's/[[:cntrl:]]//')

echo Logged in as ${USER}

echo "AUTHORIZATION=\"$AUTHORIZATION\"" >> .preflightData

# Get default organization
ORGANIZATION=$(curl -s 'https://id4i-develop.herokuapp.com/api/v1/user/organizations?offset=0&limit=1' \
    -H 'Pragma: no-cache' \
    -H 'Accept-Encoding: gzip, deflate, br' \
    -H 'Accept-Language: en' \
    -H "${AUTHORIZATION}" \
    -H 'Content-Type: application/json' \
    -H 'Accept: application/json' \
    -H 'Cache-Control: no-cache' \
    -H 'X-ID4i-Client: ID4i CLI Test' \
    -H 'Referer: https://id4i-develop.herokuapp.com/' | sed -e 's/.*namespace":"//' | sed -e 's/","logoURL.*//')

echo "Retrieved organization ${ORGANIZATION} for ${USER}"
echo "ORGANIZATION=$ORGANIZATION" >> .preflightData

# Set organization address
curl -s "https://id4i-develop.herokuapp.com/api/v1/organizations/${ORGANIZATION}/addresses/default" -X PUT \
    -H 'Accept-Encoding: gzip, deflate, br' \
    -H 'Accept-Language: en' \
    -H "${AUTHORIZATION}" \
    -H 'Content-Type: application/json' \
    -H 'Accept: application/json' \
    -H 'Cache-Control: no-cache' \
    -H 'X-ID4i-Client: ID4i CLI Test' \
    --data-binary '{"street":"Roslyndale Avenue","postCode":"9303","city":"Arleta","countryCode":"US"}'

echo
echo Set address of ${ORGANIZATION}

# Create API Key
CREATE_KEY_DATA='{"label":"CLI Test API key","organizationId":"'${ORGANIZATION}'","secret":"'${PASSWORD}'"}'

APIKEY_ID=$(curl -s 'https://id4i-develop.herokuapp.com/api/v1/apikeys' \
    -H 'Accept-Encoding: gzip, deflate, br' \
    -H 'Accept-Language: en' \
    -H "${AUTHORIZATION}" \
    -H 'Connection: keep-alive' \
    -H 'Pragma: no-cache' \
    -H 'Content-Type: application/json' \
    -H 'Accept: application/json' \
    -H 'Cache-Control: no-cache' \
    -H 'X-ID4i-Client: ID4i CLI Test' \
    --data-binary "$CREATE_KEY_DATA" | sed -e 's/.*key":"//' | sed -e 's/","createdBy.*//')

echo Created API key ${APIKEY_ID}

echo "APIKEY_ID=$APIKEY_ID" >> .preflightData

# Add API Key Permissions
PERMISSIONS="CREATE_GUID CREATE_HISTORY READ_HISTORY WRITE_HISTORY"
PERMISSIONS="${PERMISSIONS} CREATE_DOCUMENTS LIST_DOCUMENTS READ_DOCUMENTS WRITE_DOCUMENTS"
PERMISSIONS="${PERMISSIONS} WRITE_DATA READ_DATA"
PERMISSIONS="${PERMISSIONS} ID4N_WRITE_OUTGOING_TRANSFER_INFO ID4N_RECEIVE_TRANSFER ID4N_READ_OUTGOING_TRANSFER_INFO"
PERMISSIONS="${PERMISSIONS} CREATE_LOGISTIC_COLLECTION CREATE_LABELLED_COLLECTION CREATE_ROUTING_COLLECTION"
PERMISSIONS="${PERMISSIONS} RENAME_LOGISTIC_COLLECTION RENAME_LABELLED_COLLECTION RENAME_ROUTING_COLLECTION"
PERMISSIONS="${PERMISSIONS} LIST_LOGISTIC_COLLECTION_CONTENT LIST_LABELLED_COLLECTION_CONTENT LIST_ROUTING_COLLECTION_CONTENT"
PERMISSIONS="${PERMISSIONS} REMOVE_ELEMENTS_FROM_LOGISTIC_COLLECTION REMOVE_ELEMENTS_FROM_LABELLED_COLLECTION REMOVE_ELEMENTS_FROM_ROUTING_COLLECTION"
PERMISSIONS="${PERMISSIONS} DELETE_LABELLED_COLLECTION DELETE_LOGISTIC_COLLECTION DELETE_ROUTING_COLLECTION"

for PERMISSION in ${PERMISSIONS}; do
curl -s "https://id4i-develop.herokuapp.com/api/v1/apikeys/${APIKEY_ID}/privileges" \
    -H 'Accept-Encoding: gzip, deflate, br' \
    -H 'Accept-Language: en' \
    -H "${AUTHORIZATION}" \
    -H 'Connection: keep-alive' \
    -H 'Pragma: no-cache' \
    -H 'Content-Type: application/json' \
    -H 'Accept: application/json' \
    -H 'Cache-Control: no-cache' \
    -H 'X-ID4i-Client: ID4i CLI Test' \
    --data-binary '{"privilege":"'${PERMISSION}'"}'
    echo Added permission ${PERMISSION} to API key ${APIKEY_ID}
done
echo Finished adding permissions


# Activate API Key
curl -s "https://id4i-develop.herokuapp.com/api/v1/apikeys/${APIKEY_ID}" -X PUT \
    -H 'Accept-Encoding: gzip, deflate, br' \
    -H 'Accept-Language: en' \
    -H "${AUTHORIZATION}" \
    -H 'Connection: keep-alive' \
    -H 'Pragma: no-cache' \
    -H 'Content-Type: application/json' \
    -H 'Accept: application/json' \
    -H 'Cache-Control: no-cache' \
    -H 'X-ID4i-Client: ID4i CLI Test' \
    --data-binary '{"active":true}'

echo Activated API key ${APIKEY_ID}
