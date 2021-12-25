const TRILLION = 10 ** 12;

const BILLION = 10 ** 9;

const MILLION = 10 ** 6;

const THOUSAND = 10 ** 3;

const USD_CURRENCY_FORMAT = new Intl.NumberFormat("en-US", {
  style: "currency",
  currency: "USD",
});

export const formatAsCurrency = (num: number): string =>
  USD_CURRENCY_FORMAT.format(num);

const US_FORMAT = new Intl.NumberFormat("en-US");

export const formatWithCommas = (num: number): string => US_FORMAT.format(num);

const US_PERCENT_FORMAT = new Intl.NumberFormat("en-US", { style: "percent" });

export const formatAsPercent = (num: number): string =>
  US_PERCENT_FORMAT.format(num);

export const numberOfDigits = (num: number): number => {
  return Math.max(Math.floor(Math.log10(Math.abs(num))), 0) + 1;
};

export const humanizeNumber = (num: number): string => {
  if (num >= TRILLION) {
    return `${(num / TRILLION).toFixed(1)}T`;
  }
  if (num >= BILLION) {
    return `${(num / BILLION).toFixed(1)}B`;
  }
  if (num >= MILLION) {
    return `${(num / MILLION).toFixed(1)}M`;
  }
  if (num >= THOUSAND) {
    return `${(num / THOUSAND).toFixed(1)}K`;
  }
  return String(num);
};
