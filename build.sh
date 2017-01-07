#!/bin/sh
go-bindata -pkg conf conf.yaml
mv bindata.go src/pig-api-mixier/conf
go build -o pig-api-mixier src/pig-api-mixier/main.go
