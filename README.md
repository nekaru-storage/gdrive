gdrive
======


## Important
~~This tool is no longer maintained.~~ This fork intends to collate important fixes that are necessary to keep this thing working. This fork also does **not** include binaries, on purpose.

To build your own binaries, follow the steps [here](https://github.com/prasmussen/gdrive/issues/426) to get your own client ID and secret.

Then, supply the obtained Client ID and secret to `go build`:

```shell
go build -ldflags "-X main.ClientId=${CLIENT_ID} -X main.ClientSecret=${CLIENT_SECRET}"
```

## Overview
gdrive is a command line utility for interacting with Google Drive.

## Important
~~This tool is no longer maintained.~~ This fork intends to collate important fixes that are necessary to keep this thing working. This fork also does **not** include binaries, on purpose.
