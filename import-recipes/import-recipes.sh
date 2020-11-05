#!/usr/bin/env bash

if [[ -z "$1" ]]
  then
    echo "Please supply the mongo connection string as the first parameter, e.g mongodb://localhost:27017"
    exit 1
fi

mongo $1 <<EOF

 var file = cat('./recipes.json');
 use recipes
 var recipes = JSON.parse(file);
 db.recipes.insert(recipes)

EOF

