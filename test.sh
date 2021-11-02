#!/bin/bash

set -e

curl -X POST -H "Content-Type: application/json" \
    -d '{"isbn":"978-3-16-148410-0","title":"la caja mágica","description":"un libro de fantasía que cuenta la historia de un joven que encuantra una caja donde todo lo que introduce crea vida",
    "category":"fantasía","author":{"name":"miguel redondo garcía","date_of_birth":"1983/05/12"}}' \
    http://localhost:3000/books

curl -X PUT -H "Content-Type: application/json" \
    -d '{"isbn":"978-3-16-148410-0","title":"la caja mágica 2","description":"un libro de fantasía que cuenta la historia de un joven que encuantra una caja donde todo lo que introduce crea vida",
    "category":"fantasía","author":{"name":"miguel redondo garcía","date_of_birth":"1983/05/12"}}' \
    http://localhost:3000/books/038e8a97-3b51-4cce-b477-5a2cc9ee6e60

curl http://localhost:3000/books/13369f47-e3a7-4f3c-87ca-87e4bb2771b7