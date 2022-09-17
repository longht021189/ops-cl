#!/bin/bash
docker build -t tmp --build-arg CACHEBUST=$(date +%s) .