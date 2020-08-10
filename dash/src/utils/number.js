const USD_CURRENCY_FORMAT = new Intl.NumberFormat("en-US", {
  style: "currency",
  currency: "USD",
});

// eslint-disable-next-line import/prefer-default-export
export const formatAsCurrency = (num) => {
  return USD_CURRENCY_FORMAT.format(num);
};

const US_FORMAT = new Intl.NumberFormat("en-US");

export const formatWithCommas = (num) => {
  return US_FORMAT.format(num);
};
