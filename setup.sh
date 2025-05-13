#!/usr/bin/env bash
set -e

###### README ######
# This script is meant to be executed once after a fresh clone of the repo.
# It will setup some prerequisites.

# Change to directory of this script
cd "$(dirname "$0")"

# >>> Git submodule
git submodule update --init --recursive
git config submodule.recurse true
