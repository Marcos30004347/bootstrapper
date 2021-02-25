#!/bin/bash -e

CURRENT_DIR=$(echo "$(pwd)/$line")
GEN_DIR=""
REPO_DIR="$CURRENT_DIR"

PROJECT_MODULE="github.com/marcos30004347/kubeapi"
IMAGE_NAME="kubernetes-codegen:latest"

CUSTOM_RESOURCE_NAME="restaurant"
CUSTOM_RESOURCE_VERSION="v1alpha1,v1beta1"

echo "Building codegen Docker image..."
docker build -f "${CURRENT_DIR}/hack/codegen.dockerfile" \
             -t "${IMAGE_NAME}" \
             "${REPO_DIR}"


cmd0="./generate-groups.sh all \
    "$PROJECT_MODULE/pkg/generated" \
    "$PROJECT_MODULE/pkg/apis" \
    $CUSTOM_RESOURCE_NAME:$CUSTOM_RESOURCE_VERSION"

cmd1="./generate-internal-groups.sh "deepcopy,defaulter,conversion" \
    "$PROJECT_MODULE/pkg/generated" \
    "$PROJECT_MODULE/pkg/apis" \
    "$PROJECT_MODULE/pkg/apis" \
    $CUSTOM_RESOURCE_NAME:$CUSTOM_RESOURCE_VERSION"

echo "Generating client codes..."

docker run --rm \
           -v "${REPO_DIR}:/go/src/${PROJECT_MODULE}" \
           "${IMAGE_NAME}" $cmd0

docker run --rm \
           -v "${REPO_DIR}:/go/src/${PROJECT_MODULE}" \
           "${IMAGE_NAME}" $cmd1

sudo chown $USER:$USER -R $REPO_DIR/pkg