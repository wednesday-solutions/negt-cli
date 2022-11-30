#!/bin/sh

if ! [[ $GITHUB_ACTION ]]; then 
  echo "formatting.."
  golines cmd gqlgenUtils hbs -w --shorten-comments --reformat-tags -m 300
fi