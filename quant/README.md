## Quant

This directory contains machine learning models and trading strategies.

## Trading Ideas

- [ ] Momentum
- [ ] Mean reversion
- [ ] Value investing
- [ ] Earnings
- [ ] IPOs
- [ ] News sentiment
- [ ] Scalping
- [ ] Volatility trading
- [ ] Dividend stocks
- [ ] Reinforcement learning

## Development

Tensorflow can be installed locally or be used via Docker.

```shell
# Pull Tensorflow image
docker pull tensorflow/tensorflow:2.4.0-jupyter

# Run Tensorflow container
docker run --detach \
 --name quant \
 --volume ${TRADING_BOT_REPO}/quant:/tf \
 --publish 8888:8888 \
 tensorflow/tensorflow:2.4.0-jupyter

# Access logs for Jupyter URL
docker logs quant

# Access shell
docker exec -it quant sh
```
