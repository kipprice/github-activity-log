# Github Activity Log

Generates a live webpage about a set of users' activity on Github. It currently is pretty slow in loading, but can get a lot of information.

## Running the app

### Setup: Getting a Personal Access Token

In order to run this app, you will need to get an Personal Access Token for Github. The easiest way to do so is:

1. Go to [github.com/settings/tokens](https://github.com/settings/tokens)
1. Click "Generate New Token"
1. Select the following permissions:
   - repo
   - admin:org -> read:org
   - user -> read:user
1. Copy the token key somewhere safe
1. Make sure to select "Enable SSO" and authorize with any relevant SSO orgs to include SSO-protected repos

### Running the server

1. Run `sh setup.sh` to setup the application to run within a docker container. All environment variables will be saved to `config/.env`. Through this script, you will be prompted for the following variables:
    - **GitHub Token**: the key to the token you generated above
    - **GitHub Usernames**: a new-line delimited list of github usernames to search on
    - **GitHub Organizations**: the GitHub organizations to look within. Leave blank for all.
    - **GitHub Branches**: within the specified orgs, what branches should have PRs included. Leave blank for all.
    - **Lookback Days**: the number of days this search should run on
    - **Port**: the port to run the app on; defaults to 8080
1. Run the docker container via `sh run.sh`
1. Go to [http://localhost:8080](http://localhost:8080) (or to whatever port you configured) to see the GitHub activity data