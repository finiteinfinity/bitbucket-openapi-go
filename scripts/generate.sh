#!/usr/bin/env bash
set -e

# Remove any existing generated files, always starting from a clean state
function cleanExisting() {
  rm -rf ./*.go
}

# Generate the client code. Use docker for this, so we don't have to install
# any external dependencies.
function generateClient() {
  docker run --rm -v "$(pwd):/opt/bitbucket-openapi-go" \
    openapitools/openapi-generator-cli generate \
    --git-repo-id bitbucket-openapi-go \
    --git-user-id finiteinfinity \
    -i https://bitbucket.org/api/swagger.json \
    -c /opt/bitbucket-openapi-go/scripts/generator-config.json \
    -g go \
    -o /opt/bitbucket-openapi-go \
    --package-name bitbucket
}

# Perform any post-generation tidy up as necessary
function postGenerateTidy() {
  # Remove all lines with reference to invalid property 'Type', e.g. `this.Type = type_`
  sed -i '' -e '/this.Type = type_/d' ./*.go
  sed -i '' -e 's/type_ string//g' ./*.go
  sed -i '' -e 's/, type_ string//g' ./*.go
}

cleanExisting
generateClient
postGenerateTidy