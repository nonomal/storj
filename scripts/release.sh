#!/usr/bin/env bash

set -eu
set -o pipefail

echo -n "Build timestamp: "
TIMESTAMP=$(date +%s)
echo $TIMESTAMP

echo -n "Git commit: "
if [[ "$(git diff --stat)" != '' ]] || [[ -n "$(git status -s)" ]]; then
  COMMIT=$(git rev-parse HEAD)-dirty
  RELEASE=false
else
  COMMIT=$(git rev-parse HEAD)
  RELEASE=true
fi
echo $COMMIT

echo -n "Tagged version: "
if git describe --tags --exact-match --match "v[0-9]*.[0-9]*.[0-9]*"; then
  VERSION=$(git describe --tags --exact-match --match "v[0-9]*.[0-9]*.[0-9]*")
  echo $VERSION
else
  VERSION=v0.0.0
  RELEASE=false
fi

echo Running "go $@"
exec go "$1" -ldflags \
	"-s -w -X storj.io/storj/internal/version.buildTimestamp=$TIMESTAMP
         -X storj.io/storj/internal/version.buildCommitHash=$COMMIT
         -X storj.io/storj/internal/version.buildVersion=$VERSION
         -X storj.io/storj/internal/version.buildRelease=$RELEASE" "${@:2}"
