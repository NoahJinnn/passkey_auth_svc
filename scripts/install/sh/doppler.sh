#!/bin/sh

scriptsdir=$( dirname -- "$0"; )
cd $scriptsdir/../../../

if !command -v "doppler" &> /dev/null; then
    brew install dopplerhq/cli/doppler
fi

doppler setup --no-interactive
