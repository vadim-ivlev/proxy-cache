#!/usr/bin/env bash

echo "
*****************************************
10 запросов напрямую к серверу. 
8 одновременных запроса.
Тест длится 10 секунд.
*****************************************
"

ab -t 10 -n 10 -c 8  http://localhost:8085/
