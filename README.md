# AWS Cost Notification

## Required

- install go lang
- aws credentials
- slack oauth access token

## Run at local

### Setup

```bash
# download modules
$ go mod tidy

# env variables
$ export SLACK_TOKEN={{ your slack app token }}
```

And else, setup aws credentials.

### Usage

Only `-channel` is required.
Others are optional.

```sh
$ go run app/*.go -channel {{ slack post target }} \
    -prefix {{ message prefix }} \
    -from {{ yyyy/mm }} -to {{ yyyy/mm }}
```

## Run on docker

### Usage

```sh
$ env
AWS_ACCESS_KEY_ID=xxx
AWS_SECRET_ACCESS_KEY=xxx
SLACK_TOKEN=xxx

$ docker-compose up
```

※If you want to assume-role, you must specifiy `AWS_SECURITY_TOKEN`, `AWS_SESSION_TOKEN`.

