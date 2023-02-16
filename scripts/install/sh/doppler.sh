#!/bin/sh

scriptsdir=$( dirname -- "$0"; )
cd $scriptsdir/../../

doppler setup --no-interactive --config doppler.yaml
