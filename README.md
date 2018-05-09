# RTTD

### Status

This is a toy project built for fun so please don't expect high quality engineering or much attention to issues/requests.

### Screenshot 

![RTTD](https://cloud.githubusercontent.com/assets/46186/16180528/2e06e2d4-36a5-11e6-9ef1-99fd90a23ae5.png)

_Screenshot made of fake data from https://randomuser.me/_

_Originally designed by @skippednote at https://github.com/skippednote/slack-timezone_

## Run locally

### Makefile

Makefile is mostly broken right now but the build work. The workflow is as follows:

```
$ make install
$ make build
$ ./build/timezones --slack-api-token=<token>
```

## Deploy to Heroku in 5 minutes

### Create new Heroku app
```
$ heroku create --buildpack https://github.com/tonyta/heroku-buildpack-custom-binaries.git
Creating app... done, ⬢ <app-name>
Setting buildpack to https://github.com/tonyta/heroku-buildpack-custom-binaries.git... done
https://<app-name>.herokuapp.com/ | https://git.heroku.com/<app-name>.git
```

### Provide access to Slack
Now generate a [Slack API Token](https://get.slack.help/hc/en-us/articles/215770388-Creating-and-regenerating-API-tokens) and add it as an environment variable aka config variable to your heroku app from the heroku dashboard. Variable name should be `SLACK_API_TOKEN`. An alternative to this is to change append `--slack-api-token="<api-token>` to the web command in Procfile below.

### Deploy Heroku app
```
$ git clone https://git.heroku.com/<app-name>.git

$ cd <app-name>

$ echo "rttd_linux: https://github.com/owais/RTTD/releases/download/v0.2.0/RTTD_0.2.0.tar.gz" > .custom_binaries

$ echo "web: rttd_linux" > Procfile

$ git add -A .
$ git commit -m"Deploying to heroku"
$ git push
```

To upgrade to a newer version, change `0.2.0` to the version number in `.customer_binaries` file and push to heroku.