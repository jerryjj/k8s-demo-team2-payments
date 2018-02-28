#!/bin/bash

CWD=`pwd`
CONTAINER_REGISTRY_HOST=eu.gcr.io
NS=$1
VERSION=$2
CONTAINER_NAME=payments
BUILD_DIR="$CWD/build"

C_RED='\033[0;31m'
C_GREEN='\033[0;32m'
NC='\033[0;0m'

printUsage() {
  echo "Required environment variables:"
  echo "  PROJECT_ID:     Google Project ID"
  echo ""
  echo "Usage ./cloud-build.sh [local|dev|stg|prod] [version]"
}
printUsageAndExit() {
  printUsage
  exit
}

exitWithError() {
    echo "$1" 1>&2
    exit 1
}

if [ "$PROJECT_ID" = "" ]; then
  echo -e "${C_RED}No Google Project ID defined!${NC}"
  printUsageAndExit
fi

if [ "$NS" = "" ]; then
  echo -e "${C_RED}No environment defined!${NC}"
  printUsageAndExit
fi

if [ "$NS" != "dev" ]; then
  if [ "$VERSION" = "" ]; then
    echo -e "${C_RED}No version defined!${NC}"
    printUsageAndExit
  fi
fi

prepareBuild() {
  rm "$BUILD_DIR/Dockerfile" 2> /dev/null
  rm -fR "$BUILD_DIR/src" 2> /dev/null
  mkdir -p "$BUILD_DIR/src"
  cp -r "$CWD/Dockerfile" "$BUILD_DIR/"
  cp -r "$CWD/src/qvik.fi" "$BUILD_DIR/src/"
}

# Arguments:
buildInCloud() {
  local CONTAINER_VERSION="$NS"
  if [ "$VERSION" != "" ]; then
    CONTAINER_VERSION="$NS-$VERSION"
  fi

  gcloud container builds submit \
    --project $PROJECT_ID \
    --substitutions _VERSION=$CONTAINER_VERSION,_REGISTRY=$CONTAINER_REGISTRY_HOST \
    --config build-config.yml $BUILD_DIR
}

echo "Building Docker container for environment '$NS'"

prepareBuild

buildInCloud

echo -e "${C_GREEN}Finished!${NC}"
