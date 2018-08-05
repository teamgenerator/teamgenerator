import React, { Component } from 'react';
import MenuIcon from '@material-ui/icons/Menu';
import Home from '@material-ui/icons/Home';
import AssignmentTurnedIn from '@material-ui/icons/AssignmentTurnedIn'; 
import DirectionsRun from '@material-ui/icons/DirectionsRun';
import ArrowBack from '@material-ui/icons/ArrowBack';
import Check from '@material-ui/icons/Check';
import { withStyles } from '@material-ui/core/styles';
import { withRouter } from 'react-router-dom';
import PropTypes from 'prop-types';

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

const styles = {
  container: {
    position: 'fixed',
    width: '100%',
    left: '0',
    top: '0',
  },
  space: {
    height: '60px',
  },
  leftButton: {
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
  title: {
    position: 'absolute',
    left: '50%',
    transform: 'translate(-50%)',
  },
  rightButton: {
    position: 'absolute',
    right: 15,
  },
};

class CustomAppBar extends Component {

  constructor(props) {
    super(props);

    this.state = {
      navPaneOpened: false,
    }

    this.handleCloseNavPane = this.handleCloseNavPane.bind(this);
    this.handleOpenNavPane = this.handleOpenNavPane.bind(this);
    this.handleBack = this.handleBack.bind(this);
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
  
  handleBack() {
    this.props.history.goBack();
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
      <React.Fragment>
        <div className={classes.container}>
          <AppBar position="static">
            <Toolbar>
              {this.props.navpane ?
                <IconButton className={classes.leftButton} onClick={this.handleOpenNavPane} color="inherit" aria-label="Menu">
                  <MenuIcon />
                </IconButton>
              : 
                <IconButton className={classes.leftButton} onClick={this.handleBack} color="inherit" aria-label="Menu">
                  <ArrowBack />
                </IconButton>
              }

              <Typography className={classes.title} variant="title" color="inherit">
                {this.props.title.toUpperCase()}
              </Typography>

              {this.props.handleOkay ?
                <IconButton className={classes.rightButton} onClick={this.props.handleOkay} color="inherit" aria-label="Menu">
                  <Check />
                </IconButton>
              : null}

            </Toolbar>
          </AppBar>

          {this.props.navpane ?
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
          : null}
          </div>
          <div className={classes.space} />
        </React.Fragment>
    );
  }
}

CustomAppBar.propTypes = {
  title: PropTypes.string.isRequired,
  navpane: PropTypes.bool,
  handleOkay: PropTypes.func,
};

CustomAppBar.defaultProps  = {
  navpane: false,
  handleOkay: null,
};

export default withRouter(withStyles(styles)(CustomAppBar));