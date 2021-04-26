## Quant

This directory contains machine learning models and trading strategies.

## Trading Ideas

- [x] [Daily gap fill](DailyGapFill.ipynb)
- [ ] Elliot Wave
- [ ] Momentum
- [ ] Mean reversion
- [ ] Value investing
- [ ] Earnings
- [ ] IPOs/SPACs
- [ ] News sentiment
- [ ] Scalping
- [ ] Volatility trading
- [ ] Dividend stocks
- [ ] Reinforcement learning

## Development

Tensorflow can be installed locally or be used via Docker.

```shell
# Pull Tensorflow image
docker pull tensorflow/tensorflow:2.4.1-jupyter

# Run Tensorflow container
docker run --detach \
 --name quant \
 --volume ${TRADING_BOT_REPO}/quant:/tf/notebooks \
 --publish 8888:8888 \
 tensorflow/tensorflow:2.4.1-jupyter

# Access logs for Jupyter URL
docker logs quant

# Access shell
docker exec -it quant sh

## Install required dependencies
docker exec -it pip install -r requirements.txt
```
