# Dash

This is a dashboard for visualizing trades and account balances.

## Prerequisites

`trader` and `traderdb` must be setup prior to running `dash`. Instructions can
be found in the respective READMEs.

## Development

The UI was bootstrapped with [create-react-app] and can be shared via ngrok:

```bash
# Open http://localhost:4040 for ngrok UI and inspect traffic
ngrok http --host-header=rewrite 3000
```

[create-react-app]: https://create-react-app.dev/
