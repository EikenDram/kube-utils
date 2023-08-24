#!/bin/sh
echo "Publishing release: $1"

echo "Cleaning directory..."
rm -r publish
mkdir publish

echo "Making binaries for all platforms..."
GOOS=linux GOARCH=amd64 go build -o publish/ keycloak
GOOS=windows GOARCH=amd64 go build -o publish/ keycloak
echo "Making archives..."

tar -czvf ./publish/keycloak-linux-amd64-$1.tar.gz ./publish/keycloak
tar -czvf ./publish/keycloak-windows-x64-$1.tar.gz ./publish/keycloak.exe

rm publish/keycloak
rm publish/keycloak.exe

echo "Done"