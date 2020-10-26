import React, { useCallback, useState } from "react";
import clsx from "clsx";
import { ChevronLeft, TrendingUp, Menu, PieChart } from "@material-ui/icons";
import {
  AppBar,
  Divider,
  Drawer,
  IconButton,
  TextField,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
  Toolbar,
  Typography,
  makeStyles,
} from "@material-ui/core";
import { Autocomplete, createFilterOptions } from "@material-ui/lab";
import { Helmet } from "react-helmet-async";
import { Link, useHistory } from "react-router-dom";
import PropTypes from "prop-types";
import { createSelector } from "@reduxjs/toolkit";
import { useSelector } from "react-redux";
import { useTitleContext } from "../state/title-context";

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
  title: {
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
  const [drawerVisible, setDrawerVisible] = useState(false);

  const handleDrawerOpen = () => setDrawerVisible(true);
  const handleDrawerClose = () => setDrawerVisible(false);

  const handleStockClick = useCallback(
    (e, option) => history.push(`/stocks/${option.symbol}`),
    [history]
  );

  return (
    <div className={classes.root}>
      <AppBar
        position="fixed"
        className={clsx(classes.appBar, {
          [classes.appBarShift]: drawerVisible,
        })}
      >
        <Toolbar className={classes.toolbar}>
          <div className={classes.title}>
            <IconButton
              color="inherit"
              aria-label="open drawer"
              onClick={handleDrawerOpen}
              edge="start"
              className={clsx(
                classes.menuButton,
                drawerVisible && classes.hide
              )}
            >
              <Menu />
            </IconButton>
            <Title />
          </div>
          <div>
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
          </div>
        </Toolbar>
      </AppBar>
      <Drawer
        className={classes.drawer}
        variant="persistent"
        anchor="left"
        open={drawerVisible}
        classes={{
          paper: classes.drawerPaper,
        }}
      >
        <div className={classes.drawerHeader}>
          <IconButton onClick={handleDrawerClose}>
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
              <ListItemText primary="Overview" />
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
      </Drawer>
      <main
        className={clsx(classes.content, {
          [classes.contentShift]: drawerVisible,
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
