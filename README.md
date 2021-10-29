# Molokai

Molokai is a trading price alert system.

Molokai fetches the lasted price data(only CN market symbol supported now) from tushare. Then A runner is used to see if it meet certain conditions. The common condition is MinMax, MA. 

If the runner outputs a result, then a reminder is automatically started. For example, the email reminder will send an email with result to the user.

In general, you will want Molokai to be executed daily by the cron service on Linux.

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

Moloka’i Coffee refers to coffee grown on the Hawaii. It's a variety of Arabica coffee☕.