import React, { Component } from 'react';
import MenuIcon from '@material-ui/icons/Menu';
import Home from '@material-ui/icons/Home';
import AssignmentTurnedIn from '@material-ui/icons/AssignmentTurnedIn'; 
import DirectionsRun from '@material-ui/icons/DirectionsRun';
import {
  AppBar,
  Toolbar,
  IconButton,
  Drawer,
  List,
  Typography,
  ListItemIcon,
  ListItemText,
  ListItem,
  Divider,
} from '@material-ui/core';
import { withStyles } from '@material-ui/core/styles';
import { BrowserRouter, Route } from 'react-router-dom';
import HomeScreen from './containers/HomeScreen';

const styles = {
  menuButton: {
    marginLeft: -12,
    marginRight: 20,
  },
  listItem: {
    width: 250,
  },
  projectName: {
    padding: 10,
    marginBottom: 40,
  },
};
class App extends Component {

  constructor(props) {
    super(props);

    this.state = {
      navPaneOpened: false,
    }

    this.handleCloseNavPane = this.handleCloseNavPane.bind(this);
    this.handleOpenNavPane = this.handleOpenNavPane.bind(this);
  }

  handleCloseNavPane() {
    this.setState({
      navPaneOpened: false,
    });
  }

  
  handleOpenNavPane() {
    this.setState({
      navPaneOpened: true,
    });
  }

  renderListItem(icon, label) {
    return (
      <ListItem button>
        <ListItemIcon>
          {icon}
        </ListItemIcon>
        <ListItemText primary={label} />
      </ListItem>
    )
  }

  render() {
    const { classes } = this.props;
    return (
      <div>
        <AppBar position="static">
          <Toolbar>
            <IconButton className={classes.menuButton} onClick={this.handleOpenNavPane} color="inherit" aria-label="Menu">
              <MenuIcon />
            </IconButton>
          </Toolbar>
        </AppBar>

        <Drawer open={this.state.navPaneOpened} className={classes.drawer} onClose={this.handleCloseNavPane}>
          <List className={classes.projectName}>
            <Typography variant="title">Community 1</Typography>
          </List>
          <Divider />

          <List
            className={classes.listItem}
            tabIndex={0}
            role="button"
          >
            {this.renderListItem(<Home />, 'Home')}
            {this.renderListItem(<DirectionsRun />, 'Players')}
            {this.renderListItem(<AssignmentTurnedIn />, 'Sessions')}
          </List>
        </Drawer>

        <BrowserRouter>
          <Route exact path="/" component={HomeScreen} />
        </BrowserRouter>
      </div>
    );
  }
}

export default withStyles(styles)(App);