import React, { Component } from 'react';
import MenuIcon from '@material-ui/icons/Menu';
import Home from '@material-ui/icons/Home';
import AssignmentTurnedIn from '@material-ui/icons/AssignmentTurnedIn'; 
import DirectionsRun from '@material-ui/icons/DirectionsRun';
import ArrowBack from '@material-ui/icons/ArrowBack';
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
    this.handleGoToPage = this.handleGoToPage.bind(this);
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

  handleGoToPage(pagePath) {
    this.props.history.push(pagePath);
  }
  
  handleBack() {
    this.props.history.goBack();
  }

  renderListItem(icon, label, pagePath) {
    return (
      <ListItem button onClick={() => this.handleGoToPage(pagePath)}>
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

              <Typography className={classes.title} variant="subheading" color="inherit">
                {this.props.title && this.props.title.toUpperCase()}
              </Typography>

              {this.props.rightButton ?
                <IconButton
                  className={classes.rightButton}
                  disabled={this.props.rightButton.disabled}
                  onClick={this.props.rightButton.onClick}
                  color="inherit"
                  aria-label="Menu"
                >
                  {this.props.rightButton.icon}
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
                {this.renderListItem(<Home />, 'Home', '/')}
                {this.renderListItem(<DirectionsRun />, 'Players', '/players')}
                {this.renderListItem(<AssignmentTurnedIn />, 'Sessions', '/sessions')}
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
  rightButton: PropTypes.shape({
    icon: PropTypes.node,
    onClick: PropTypes.func,
    disabled: PropTypes.bool,
  }),
};

CustomAppBar.defaultProps  = {
  navpane: false,
  rightButton: null,
};

export default withRouter(withStyles(styles)(CustomAppBar));