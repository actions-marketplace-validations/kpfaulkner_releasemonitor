# Release Monitor
Github action to react to release events

Update slack when "in test" or "in stage" releases are modified.
Update slack when "release" is created/published.

Simply configure the action with the yaml:

on: gollum
name: ReleaseMonitor
jobs:
  slackNotification:
    name: Slack Notification
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v2
    - name: Slack Notification
      uses: kpfaulkner/releasemonitor@0.2.0
      env:
        SLACK_WEBHOOK: ${{ secrets.SLACK_WEBHOOK }}

secrets.SLACK_WEBHOOK is a regular Slack webhook ( https://slack.com/help/articles/115005265063-Incoming-Webhooks-for-Slack )

WIKI_TITLES_TO_ALERT is a comma separated list of wiki titles that you want to alert on. If you happen to have a comma in your title, then 
I might need to change things.
