import { fetchJSON } from "../utils/http";

export const updateWatchlist = () => {};

export const fetchWatchlists = () => fetchJSON("/api/account/watchlists");

export const deleteWatchlist = (id) =>
  fetchJSON(`/api/account/watchlists/${id}`, { method: "DELETE" });
