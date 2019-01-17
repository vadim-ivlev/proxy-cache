#!/usr/bin/env bash

mount -t tmpfs -o size=1024m tmpfs ./cache
docker-compose up 
umount ./cache
