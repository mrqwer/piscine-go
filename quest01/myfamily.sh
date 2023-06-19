#curl -sS https://01.alem.school/assets/superhero/all.json | jq -r --arg id "$HERO_ID" '.[] | select(.id == ($id | tonumber)) | .connections.relatives' | tr -d '"'

#!/usr/bin/env bash
#curl -s https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json | jq --arg HERO_ID "$HERO_ID" '.[] | select ( .id == ($HERO_ID|tonumber)) | .connections.relatives' | sed 's/"//g'
curl -sS https://01.alem.school/assets/superhero/all.json | jq --arg HERO_ID "$HERO_ID" '.[] | select ( .id == ( $HERO_ID| tonumber )) | .connections.relatives' | tr -d '"'
