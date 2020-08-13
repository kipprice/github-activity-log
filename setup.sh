
declare -r FOLDER="server/config"

#! bin/bash
yn() {
    echo "$1 (Y/N)"
    read response
    if [[ "$response" = "Y" ]]; then return 1; fi
    if [[ "$response" = "y" ]]; then return 1; fi
    return 0
}

setup_config() {

    if [ -d "$FOLDER" ]; then return; fi

    echo "\n==> Creating Config Folder"
    mkdir config
}

setup_lookback() {
    
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
    echo $key > $FOLDER/.github_key
}

setup_team() {
    echo "\n==> Setting up GitHub Team"
    if [ -f "$FOLDER/.github_team" ]; then
        yn "Do you want to replace the current team settings?"
        if [[ $? -eq 0 ]]; then return; fi
    fi

    echo "" > $FOLDER/.github_team
    echo "What github IDs do you want to include?"
    read username
    while [[ ! -z "$username" ]]; do
        echo $username >> $FOLDER/.github_team
        read username
    done
}

setup_docker() {
    docker build -t github-activity-log .
    docker run --rm -p 8080:8080 github-acivity-log
}

main() {
    clear
    setup_config
    setup_key
    setup_team
    # setup_docker

    # finalization
    clear
    echo "==> Configured!"
}

main