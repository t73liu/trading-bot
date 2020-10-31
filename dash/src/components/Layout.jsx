import React, { useCallback, useState } from "react";
import clsx from "clsx";
import {
  ChevronLeft,
  Menu,
  PieChart,
  Settings as SettingsIcon,
  TrendingUp,
} from "@material-ui/icons";
import {
  AppBar,
  Divider,
  IconButton,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
  makeStyles,
  SwipeableDrawer,
  TextField,
  Toolbar,
  Typography,
} from "@material-ui/core";
import { Autocomplete, createFilterOptions } from "@material-ui/lab";
import { Helmet } from "react-helmet-async";
import { Link, useHistory } from "react-router-dom";
import PropTypes from "prop-types";
import { createSelector } from "@reduxjs/toolkit";
import { useSelector } from "react-redux";
import { useTitleContext } from "../state/title-context";
import Settings from "./Settings";

const drawerWidth = 240;

// TODO Simplify layout styles
const useStyles = makeStyles((theme) => ({
  root: {
    display: "flex",
  },
  appBar: {
    transition: theme.transitions.create(["margin", "width"], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
  },
  appBarShift: {
    width: `calc(100% - ${drawerWidth}px)`,
    marginLeft: drawerWidth,
    transition: theme.transitions.create(["margin", "width"], {
      easing: theme.transitions.easing.easeOut,
      duration: theme.transitions.duration.enteringScreen,
    }),
  },
  toolbar: {
    justifyContent: "space-between",
  },
  toolbarContent: {
    display: "flex",
    alignItems: "center",
  },
  menuButton: {
    marginRight: theme.spacing(2),
  },
  hide: {
    display: "none",
  },
  drawer: {
    width: drawerWidth,
    flexShrink: 0,
  },
  drawerPaper: {
    width: drawerWidth,
  },
  drawerHeader: {
    display: "flex",
    alignItems: "center",
    padding: theme.spacing(0, 1),
    ...theme.mixins.toolbar,
    justifyContent: "flex-end",
  },
  content: {
    flexGrow: 1,
    padding: theme.spacing(3),
    transition: theme.transitions.create("margin", {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
    marginLeft: -drawerWidth,
  },
  contentShift: {
    transition: theme.transitions.create("margin", {
      easing: theme.transitions.easing.easeOut,
      duration: theme.transitions.duration.enteringScreen,
    }),
    marginLeft: 0,
  },
}));

const Title = () => {
  const { title } = useTitleContext();
  return (
    <>
      <Helmet>
        <title>{title}</title>
      </Helmet>
      <Typography variant="h6" noWrap>
        {title}
      </Typography>
    </>
  );
};

const getStocks = createSelector(
  (state) => state.stocks,
  (stocks) => stocks.allStocks
);

const filterOptions = createFilterOptions({
  limit: 50,
});

const getStockLabel = (stock) => `${stock.symbol} - ${stock.company}`;

const Layout = ({ children }) => {
  const classes = useStyles();
  const history = useHistory();
  const stocks = useSelector(getStocks);
  const [showMenu, setShowMenu] = useState(false);
  const handleOpenMenu = () => setShowMenu(true);
  const handleCloseMenu = () => setShowMenu(false);
  const [showSettings, setShowSettings] = useState(false);
  const handleOpenSettings = () => setShowSettings(true);
  const handleCloseSettings = () => setShowSettings(false);

  const handleStockClick = useCallback(
    (e, option) => {
      if (option) {
        history.push(`/stocks/${option.symbol}`);
      }
    },
    [history]
  );

  return (
    <div className={classes.root}>
      <AppBar
        position="fixed"
        className={clsx(classes.appBar, {
          [classes.appBarShift]: showMenu,
        })}
      >
        <Toolbar className={classes.toolbar}>
          <div className={classes.toolbarContent}>
            <IconButton
              color="inherit"
              aria-label="open drawer"
              onClick={handleOpenMenu}
              edge="start"
              className={clsx(classes.menuButton, showMenu && classes.hide)}
            >
              <Menu />
            </IconButton>
            <Title />
          </div>
          <div className={classes.toolbarContent}>
            <Autocomplete
              onChange={handleStockClick}
              filterOptions={filterOptions}
              options={stocks}
              getOptionLabel={getStockLabel}
              renderInput={(params) => (
                <TextField
                  {...params}
                  margin="normal"
                  variant="outlined"
                  placeholder="Symbol"
                  style={{
                    backgroundColor: "white",
                    width: 300,
                  }}
                />
              )}
            />
            <div>
              <IconButton
                onClick={handleOpenSettings}
                style={{ color: "white" }}
              >
                <SettingsIcon />
              </IconButton>
            </div>
          </div>
        </Toolbar>
      </AppBar>
      <SwipeableDrawer
        className={classes.drawer}
        variant="persistent"
        anchor="left"
        open={showMenu}
        classes={{
          paper: classes.drawerPaper,
        }}
      >
        <div className={classes.drawerHeader}>
          <IconButton onClick={handleCloseMenu}>
            <ChevronLeft />
          </IconButton>
        </div>
        <Divider />
        <List>
          <Link to="/">
            <ListItem button>
              <ListItemIcon>
                <PieChart color="secondary" />
              </ListItemIcon>
              <ListItemText primary="Market" />
            </ListItem>
          </Link>
          <Link to="/stocks">
            <ListItem button>
              <ListItemIcon>
                <TrendingUp color="secondary" />
              </ListItemIcon>
              <ListItemText primary="Stocks" />
            </ListItem>
          </Link>
        </List>
      </SwipeableDrawer>
      <Settings
        visible={showSettings}
        handleCloseSettings={handleCloseSettings}
      />
      <main
        className={clsx(classes.content, {
          [classes.contentShift]: showMenu,
        })}
      >
        <div className={classes.drawerHeader} />
        {children}
      </main>
    </div>
  );
};

Layout.propTypes = {
  children: PropTypes.node.isRequired,
};

export default Layout;
