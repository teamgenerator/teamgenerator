import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import MenuIcon from '@material-ui/icons/Menu';
import {
  AppBar,
  Toolbar,
  IconButton,
} from '@material-ui/core';
import { withStyles } from '@material-ui/core/styles';
import { BrowserRouter, Route } from 'react-router-dom';
import Home from './containers/Home';

const styles = {
  root: {
    flexGrow: 1,
  },
  flex: {
    flexGrow: 1,
  },
  menuButton: {
    marginLeft: -12,
    marginRight: 20,
  },
};
class App extends Component {
  render() {
    const { classes } = this.props;
    return (
      <div>
        <AppBar position="static">
          <Toolbar>
            <IconButton className={classes.menuButton} color="inherit" aria-label="Menu">
              <MenuIcon />
            </IconButton>
          </Toolbar>
        </AppBar>
        
        <BrowserRouter>
          <Route exact path="/" component={Home} />
        </BrowserRouter>
      </div>
    );
  }
}

export default withStyles(styles)(App);