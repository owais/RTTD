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