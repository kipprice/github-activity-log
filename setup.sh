#!/bin/bash

# constants for what file gets generated
declare -r FOLDER="config"
declare -r FILENAME=".env"
declare -r FILEPATH="$FOLDER/$FILENAME"

# =========
#  HELPERS
# =========

yn() {
    echo "$1" >$(tty)
    printf "[Y/N] > " >$(tty)
    read -r response
    echo "" >$(tty)
    if [[ "$response" = "Y" ]]; then return 1; fi
    if [[ "$response" = "y" ]]; then return 1; fi
    return 0
}

collect_array() {
    local out=""

    read ipt

    while [[ ! -z "$ipt" ]]; do
        if [[ -z "$out" ]]; then
            out="$ipt"
        else
            out="$out,$ipt"
        fi
        read ipt
    done

    echo $out
}

prompt_for_change() {
    label=$1
    val=$2

    if [[ ! -z "$val" ]]; then
        yn "$label: $val\nDo you want to change this?"
        if [[ $? -eq 0 ]]; then
            echo $val
        fi
    fi

    echo ""
}

line_count() {
    echo $(wc -l < $FILEPATH | awk '{print $1}')
}

# =======
#  SETUP
# =======

setup_key() {
    echo "\n==> Setting up GitHub Key"

    local key="$(prompt_for_change "GitHub Key" $GITHUB_TOKEN)"
    if [[ -z "$key" ]]; then
        echo "Enter the key for your GitHub personal access token: "
        read key
    fi

    echo "GITHUB_TOKEN=$key" >> $FILEPATH
}

setup_team() {
    echo "\n==> Setting up GitHub Team"

    local usernames="$(prompt_for_change "GitHub Users" $GITHUB_TEAM)"
    if [[ -z "$usernames" ]]; then

        echo "What GitHub usernames do you want to include? (Enter a blank line to end)"
        usernames=$(collect_array)
    fi

    echo "GITHUB_TEAM=$usernames" >> $FILEPATH
}

setup_orgs() {
    echo "\n==> Setting up organizations"
    local orgs="$(prompt_for_change "GitHub Orgs" $GITHUB_ORGS)"
    if [[ -z "$orgs" ]]; then
        echo "What GitHub organizations do you want to include? (Enter a blank line to end)"
        orgs=$(collect_array)
    fi

    echo "GITHUB_ORGS=$orgs" >> $FILEPATH
}

setup_branches() {
    echo "\n==> Setting up branches"
    local branches="$(prompt_for_change "GitHub Branches" $GITHUB_BRANCHES)"

    if [[ -z "$branches" ]]; then
        echo "What GitHub branches do you want to include? (Enter a blank line to end)"
        branches=$(collect_array)
    fi

    echo "GITHUB_BRANCHES=$branches" >> $FILEPATH
}

setup_lookback() {
    echo "\n==> Setting up lookback"
    local lookback_days="$(prompt_for_change "Lookback Days" $LOOKBACK_DAYS)"

    if [[ -z "$lookback_days" ]]; then
        echo "How many days do you want to look back for data? \r"
        read lookback_days
    fi

    echo "LOOKBACK_DAYS=$lookback_days" >> $FILEPATH
}

setup_port() {
    echo "\n==> Setting up port"

    # Set a default value
    if [[ -z $PORT ]]; then
        PORT=8080
    fi

    local port="$(prompt_for_change "Port" $PORT)"

    if [[ -z "$port" ]]; then
        echo "What port would you like to host the application on?"
        read port
    fi

    echo "PORT=$port" >> $FILEPATH
}

setup_mode() {
    echo "GIN_MODE=release" >> $FILEPATH
}

setup_docker() {
    docker build . -t github-activity-log
}

setup_config() {

    # check if the folder already exists
    if [ -d "$FOLDER" ]; then
        yn "Would you like to rerun the Docker build? (Only required if there are code changes)"
        if [[ $? -eq 1 ]]; then
            setup_docker
        fi

        lc=$(line_count)
        if [[ lc -gt 6 ]]; then
            yn "You already have the appropriate setup; do you want to update?"
            if [[ $? -eq 0 ]]; then return; fi

            # get the token out of the file
            source $FILEPATH
            rm $FILEPATH

            
        fi
        
    else
        # make the appropriate folder
        mkdir $FOLDER
        setup_docker
    fi

    # generate the config file
    touch $FILEPATH

    # perform the additional setup
    setup_key
    setup_team
    setup_orgs
    setup_branches
    setup_lookback
    setup_port
    setup_mode
}

main() {
    clear
    setup_config
    clear
    echo "==> Configured!"
}

main