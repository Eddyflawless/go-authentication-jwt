#!/bin/bash

echo -e "\nList available commands"

cat <<EOF

Usage:

make [COMMAND] [args]

1. migrate m-status=up Run mongo migrations

2. migrate m-status=down Teardown mongo migrations

2. run-api Start Go application service

3. platform-db Start MongoDB database service

4. platform-down Stop MongoDB database service

EOF