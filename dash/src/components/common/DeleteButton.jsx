import React from "react";
import { IconButton } from "@material-ui/core";
import { Delete } from "@material-ui/icons";
import PropTypes from "prop-types";

const DeleteButton = ({ onClick, edge }) => (
  <IconButton edge={edge} onClick={onClick}>
    <Delete />
  </IconButton>
);

DeleteButton.defaultProps = {
  edge: false,
};

DeleteButton.propTypes = {
  edge: PropTypes.oneOf(["end", "start", false]),
  onClick: PropTypes.func.isRequired,
};

export default DeleteButton;
