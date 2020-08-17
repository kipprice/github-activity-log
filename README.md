# Github Activity Log

Prints details about a Codecademy's users activity on Github (as of right now).

## Running the app

### Setup: Getting a Personal Access Token

In order to run this app, you will need to get an Personal Access Token for Github. The easiest way to do so is:

1. Go to [github.com/settings/tokens](https://github.com/settings/tokens)
1. Click "Generate New Token"
1. Select the following permissions:
   - repo
   - read:packages
   - admin:org -> read:org
   - user
1. Copy the token somewhere safe
1. Make sure to select "Enable SSO" to allow pass-through to SSO-protected repos

### Setup: Generate Environment Variables

This application relies on some environment variables to be set, namely for the users you're loading data for, the token you created earlier, and the number of days to lookback.

1. Run `sh setup.sh` to generate `config/export_vars.sh` containing your specific environment variables.
1. Run `source config/export_vars.sh` to set all appropriate environment vars in your terminal.

### Run It

1. If you don't already have `go` installed, [install it](https://golang.org/doc/install).
1. Run `go run main.go > results/index.html` to generate a webpage at results/index.html with all of the results.
