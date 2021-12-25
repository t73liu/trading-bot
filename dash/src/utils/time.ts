import dayjs from "dayjs";

// eslint-disable-next-line import/prefer-default-export
export const formatDateTime = (dt: string): string => {
  return dayjs(dt).format("YYYY-MM-DD HH:mm:ss");
};
