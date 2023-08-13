# spooky_kingdom
Paranormal stories collected from local UK news outlets

Stories are crawled by [spooky_crawler](https://github.com/r-rayns/spooky_crawler) and passed to spooky_kingdom via an API call.

## build

1. Build the UI

```bash
  cd ui
  npm run build
```

2. Compile Go

(From the root directory)

```bash
  go build -o spooky_kingdom main.go
```

## .env 📔
Set in config/config.toml

```toml
[postgresql]
# Data source name
dsn="user= password= database= host= port= sslmode="
username=""
password=""
host=""
port=""

[jwt]
secret=""

[admin]
username=""
password=""

[api]
key=""

```
