## Dash

This is a dashboard for visualizing trades and account balances.

## Prerequisites

`trader` and `traderdb` must be setup prior to running `dash`. Instructions can
be found on the respective READMEs.

## Development

The UI is built with React and CLI instructions can be found in [CRA.md](CRA.md).

The UI can be shared via ngrok:

```sh
# Open http://localhost:4040 for ngrok UI and inspect traffic
ngrok http --host-header=rewrite 3000
```
