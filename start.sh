#!/usr/bin/env bash

mount -t tmpfs -o size=1024m tmpfs ./smart-cache
docker-compose up 
umount ./smart-cache
