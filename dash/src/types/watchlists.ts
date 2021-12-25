import { Stock } from "./stocks";

export interface Watchlist {
  id: string;
  name: string;
  stocks: Stock[];
}
