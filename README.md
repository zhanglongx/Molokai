# Molokai

Molokai is a trading price alert system.

## Build

```bash
	$ go build .
```

## Usage

1. Prepare the ~/.molokairc. It should be in yaml format

```yaml
  # email from
  from:
    user: "<user@email.com>"
    password: "<password>"
  # smtp info
  smtp:
    smtp: "<smtp.email.com>"
    port: <smtp port>
```

2. Runtime file. Here is an example

```yaml
  # Molokai version
  version: v1
  # tushare token
  token: "<tushare token>"
  holdings:
    # symbol, one of 'ts_code', 'symbol', 'name' from tushare
    # only support CN Market now
  - symbol: "<symbol>"
    runners: 
      # runner name, eg. "MinMax", "MA"
    - name: "MinMax"
      # runner param
      param: {max: 30.00, min: 22.08}
  reminders:
	turnon: true
    to: "<to@email.com>"
```

3. run

```bash
	$ Molokai <runtime.yaml>
```

## About

Moloka’i Coffee refers to coffee grown on the Hawaii, as a variety of Arabica coffee☕.