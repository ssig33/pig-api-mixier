#!/bin/sh
go-bindata -pkg conf conf.yaml
mv bindata.go src/pig-api-mixier/conf
go build src/pig-api-mixier/main.go
