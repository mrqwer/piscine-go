#!/bin/bash

name=$(curl -sS https://01.alem.school/assets/superhero/all.json | jq -r --arg id 70 '.[] | select(.id == ($id | tonumber)) |  .name')

echo '"'$name'"'
