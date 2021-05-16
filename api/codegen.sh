#! /bin/bash

rm -rf generated/go/

java -jar swagger-codegen-cli-3.0.25.jar generate -i swagger.yaml -l go-server -DpackageName=models -o generated

find generated/go/ -type f -name 'model_*' -exec cp '{}' ../pkg/models/ ';'

rm -rf generated/go/
rm -rf generated/.swagger-codegen/
rm -rf generated/api/
rm generated/Dockerfile
rm generated/main.go