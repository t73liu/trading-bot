import React from "react";
import PropTypes from "prop-types";

const Watchlists = ({ watchlists }) => {
  return (
    <div>
      <h2>Watchlists</h2>
      <ul>
        {watchlists.map((watchlist) => (
          <li key={watchlist.id}>{watchlist.name}</li>
        ))}
      </ul>
    </div>
  );
};

Watchlists.propTypes = {
  watchlists: PropTypes.arrayOf(
    PropTypes.shape({
      id: PropTypes.number.isRequired,
      name: PropTypes.string.isRequired,
    })
  ),
};

Watchlists.defaultProps = {
  watchlists: [],
};

export default Watchlists;
