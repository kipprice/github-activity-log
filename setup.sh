#! bin/bash

# constants for what file gets generated
declare -r FOLDER="config"
declare -r FILENAME="export_vars.sh"
declare -r FILEPATH="$FOLDER/$FILENAME"

yn() {
    echo "$1 (Y/N)\r"
    read response
    if [[ "$response" = "Y" ]]; then return 1; fi
    if [[ "$response" = "y" ]]; then return 1; fi
    return 0
}

setup_key() {
    echo "\n==> Setting up GitHub Key"
    if [ -f "$FOLDER/.github_key" ]; then
        yn "Do you want to replace the current key?"
        if [[ $? -eq 0 ]]; then return; fi
    fi

    echo "Enter your GitHub API key: "
    read key
    
    if [ -z "$key" ]; then return; fi
    echo "export GITHUB_TOKEN=\"$key\"" >> $FILEPATH
}

setup_team() {
    echo "\n==> Setting up GitHub Team"

    local usernames=""

    echo "What github IDs do you want to include? (Enter a blank line to end)"
    read username

    while [[ ! -z "$username" ]]; do
        usernames="$usernames,$username"
        read username
    done

    echo "export GITHUB_TEAM=\"$usernames\"" >> $FILEPATH
}

setup_lookback() {
    echo "How many days do you want to lookback? \r"
    read lookback_days

    echo "export LOOKBACK_DAYS=$lookback_days" >> $FILEPATH
}

setup_docker() {
    # TODO: finish this all up
    docker build -t github-activity-log .
    docker run --rm github-acivity-log
}

setup_config() {

    # check if the folder already exists
    if [ -d "$FOLDER" ]; then
        yn "You already have the appropriate setup; do you want to update?"
        if [[ $? -eq 0 ]]; then return; fi

        # clear out the existing file
        rm $FILEPATH
        
    else
        # make the appropriate folder
        mkdir $FOLDER
    fi

    # generate the config file
    touch $FILEPATH

    # perform the additional setup
    setup_key
    setup_team
    setup_lookback
}

run_config() {
    set -o allexport
    source $FILEPATH
    set +o allexport
}

run_app() {
    go run main.go > results/test.html
}

main() {
    clear

    # set up the appropriate config
    setup_config
    # run_config
    # run_app

    # finalization
    clear
    echo "==> Configured!"
}

main