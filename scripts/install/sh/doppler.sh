#!/bin/sh

scriptsdir=$( dirname -- "$0"; )
cd $scriptsdir/../../../

brew install dopplerhq/cli/doppler
doppler setup --no-interactive
