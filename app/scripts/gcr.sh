#!bin/bash

pull_push() {
    REGISTRY=us-docker.pkg.dev
    REPO_NAME=main
    REPO_FORMAT=DOCKER
    LOCATION=us

    if [[ "$3" == "" ]]; then
        export SOURCE_REG=docker.io
    else
        export SOURCE_REG=$3
    fi

    DESTINATION_REG=$REGISTRY
    IMAGE=$1
    TAG=$2
    docker pull $SOURCE_REG/$IMAGE:$TAG
    docker tag $SOURCE_REG/$IMAGE:$TAG  $DESTINATION_REG/$PROJECT_ID/$REPO_NAME/$IMAGE:$TAG
    docker push $DESTINATION_REG/$PROJECT_ID/$REPO_NAME/$IMAGE:$TAG
}

if [ -z "$PROJECT_ID"]
then
    echo "Please setup project var: export PROJECT_ID=<your project id>"
else
    echo "Working on project: $PROJECT_ID"
    pull_push $1 $2 $3
fi