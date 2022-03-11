# GitLab CI Microsoft Teams Bot

This CLI tool simplifies integrating Microsoft Teams bot functionality into GitLab CI piplines.

## Usage

1. Clone repository
2. `make`
3. `export WEBHOOK_URL=https://teamschannelwebhookurl/`
4. `gitlab-teams-bot --title "Card Title" --description "HTML for the card body"`

### Notes:

* By default, the following environment variables will be used:
  * `WEBHOOK_URL` the URL Microsoft Teams webhook URL
  * `CI_JOB_STATUS` the GitLab CI job status
  * `CI_JOB_URL` the URL to the failed GitLab CI job
* If you wish you can override various CLI options to further customise
