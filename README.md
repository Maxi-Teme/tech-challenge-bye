# bye - test app

## init

`go mod init bye && go mod tidy`

## build

`go build -o bye .`

## run

`PORT=8000 ENV=local ./bye`

## deploy locally

### prepare

`./get_envs.sh`

### dev

`act -P ubuntu-latest=local/act_runner --pull=false --secret-file=.env.dev`

### stg

`act -P ubuntu-latest=local/act_runner --pull=false --secret-file=.env.stg`

### prod

`act -P ubuntu-latest=local/act_runner --pull=false --secret-file=.env.prod`
