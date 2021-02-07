import { fetchJSON } from "../utils/http";

export const fetchWatchlists = () => fetchJSON("/api/account/watchlists");

export const deleteWatchlist = (id) =>
  fetchJSON(`/api/account/watchlists/${id}`, { method: "DELETE" });

export const updateWatchlist = (watchlist) =>
  fetchJSON(`/api/account/watchlists/${watchlist.id}`, {
    method: "PUT",
    body: JSON.stringify(watchlist),
  });

export const createWatchlist = () =>
  fetchJSON("/api/account/watchlists", {
    method: "POST",
    body: JSON.stringify({ name: "New Watchlist", stockIDs: [] }),
  });
