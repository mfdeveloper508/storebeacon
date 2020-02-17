#!/bin/bash

set -e

TAG=$CI_PIPELINE_ID
PREFIX="$EB_APPLICATION"

mkdir tmp

DOCKERRUN_FILE="Dockerrun.aws.json"
cat "$DOCKERRUN_FILE.template" \
    | sed 's|<BUCKET>|'$EB_BUCKET'|g' \
      | sed 's|<PREFIX>|'$PREFIX'|g' \
        | sed 's|<IMAGE>|'$IMAGE_TAG'|g' \
	  > "tmp/$DOCKERRUN_FILE"

cp -r .ebextensions tmp

cd tmp
zip ../app.zip $DOCKERRUN_FILE .ebextensions/*
cd ..

rm -rf tmp

aws s3 cp app.zip s3://$EB_BUCKET/$PREFIX/app.zip
rm app.zip


echo "Creating new EB version"
echo $CI_PIPELINE_ID
aws elasticbeanstalk create-application-version \
    --application-name $EB_APPLICATION \
      --version-label $CI_PIPELINE_ID \
        --source-bundle S3Bucket=$EB_BUCKET,S3Key=$PREFIX/app.zip

echo "Updating EB environment"
aws elasticbeanstalk update-environment \
    --environment-name $EB_ENV \
      --version-label $CI_PIPELINE_ID
