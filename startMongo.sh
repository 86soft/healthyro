#!/bin/bash

nerdctl compose up -d

sleep 10

nerdctl exec mongodb /scripts/rs-init.sh