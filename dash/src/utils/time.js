import dayjs from "dayjs";

// eslint-disable-next-line import/prefer-default-export
export const formatDateTime = (dateString) => {
  return dayjs(dateString).format("YYYY-MM-DD HH:mm:ss");
};
