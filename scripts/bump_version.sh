#!/bin/bash

set -e

PART=$1
VERSION=$(cat VERSION)
echo "Current version $VERSION"

# strip leading 'v' if present
VERSION=${VERSION#v}

# splits version into components
IFS='.' read -r MAJOR MINOR PATCH <<< "$VERSION"

# bump the version
case "$PART" in
    major)
        MAJOR=$((MAJOR + 1))
        MINOR=0
        PATCH=0
        ;;
    minor)
        MINOR=$((MINOR + 1))
        PATCH=0
        ;;
    patch)
        PATCH=$((PATCH + 1))
        ;;
       *)
    echo "Usage: bump_version.sh patch|minor|major"
    exit 1
    ;;
esac

# Reassemble and add 'v' prefix
NEW_VERSION="v${MAJOR}.${MINOR}.${PATCH}"
echo "$NEW_VERSION" > VERSION
echo "Version bumped to $NEW_VERSION"