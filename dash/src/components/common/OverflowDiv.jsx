import PropTypes from "prop-types";
import React from "react";

const OverflowDiv = ({ children }) => (
  <div style={{ maxHeight: "100px", overflow: "auto" }}>{children}</div>
);

OverflowDiv.propTypes = {
  children: PropTypes.node,
};

OverflowDiv.defaultProps = {
  children: undefined,
};

export default OverflowDiv;
