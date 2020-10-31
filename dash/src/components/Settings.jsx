import React, { useCallback } from "react";
import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  FormControl,
  FormControlLabel,
  FormHelperText,
  FormGroup,
  MenuItem,
  Select,
  Switch,
} from "@material-ui/core";
import PropTypes from "prop-types";
import { useDispatch, useSelector } from "react-redux";
import {
  getCandleSize,
  getShowExtendedHours,
  setCandleSize,
  toggleShowExtendedHours,
} from "../state/account";

const Settings = ({ visible, handleCloseSettings }) => {
  const dispatch = useDispatch();
  const candleSize = useSelector(getCandleSize);
  const handleCandleSizeChange = useCallback(
    (e) => {
      dispatch(setCandleSize(e.target.value));
    },
    [dispatch]
  );
  const showExtendedHours = useSelector(getShowExtendedHours);
  const handleExtendedHoursToggle = useCallback(() => {
    dispatch(toggleShowExtendedHours());
  }, [dispatch]);
  return (
    <Dialog open={visible} onClose={handleCloseSettings}>
      <DialogTitle>Settings</DialogTitle>
      <DialogContent>
        <div style={{ width: 300 }}>
          <FormControl style={{ width: "100%" }}>
            <FormGroup>
              <Select
                id="input-with-icon-adornment"
                value={candleSize}
                onChange={handleCandleSizeChange}
              >
                <MenuItem value="1min">1 minute</MenuItem>
                <MenuItem value="3min">3 minute</MenuItem>
                <MenuItem value="5min">5 minute</MenuItem>
                <MenuItem value="10min">10 minute</MenuItem>
                <MenuItem value="30min">30 minute</MenuItem>
                <MenuItem value="1hour">1 hour</MenuItem>
              </Select>
              <FormHelperText>Candle Size</FormHelperText>
              <FormControlLabel
                control={
                  <Switch
                    name="extended"
                    checked={showExtendedHours}
                    onChange={handleExtendedHoursToggle}
                  />
                }
                label="Show Extended Hours"
              />
              <FormControlLabel
                control={<Switch name="volume" checked />}
                label="Volume"
              />
              <hr style={{ width: "100%" }} />
              <FormControlLabel
                control={<Switch name="ema" checked />}
                label="EMA"
              />
              <FormControlLabel
                control={<Switch name="macd" disabled />}
                label="MACD"
              />
              <FormControlLabel
                control={<Switch name="rsi" disabled />}
                label="RSI"
              />
              <FormControlLabel
                control={<Switch name="vwap" checked />}
                label="VWAP"
              />
            </FormGroup>
          </FormControl>
        </div>
      </DialogContent>
      <DialogActions>
        <Button onClick={handleCloseSettings} color="primary">
          OK
        </Button>
      </DialogActions>
    </Dialog>
  );
};

Settings.propTypes = {
  visible: PropTypes.bool.isRequired,
  handleCloseSettings: PropTypes.func.isRequired,
};

export default Settings;
