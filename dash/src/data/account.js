import { getJSON } from "../utils/http";

export const updateWatchlist = () => {};

export const fetchWatchlists = () => getJSON("/api/account/watchlists");

export const deleteWatchlist = () => {};
