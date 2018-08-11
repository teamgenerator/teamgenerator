import React, { Component } from 'react';
import { Button, Grid } from '@material-ui/core';
import { withStyles } from '@material-ui/core/styles';
import { withRouter } from 'react-router-dom';

import AppBar from '../components/AppBar';

const styles = {
  container: {
    display: 'flex',
    flexDirection: 'column',
  },
  actions: {
    marginTop: 30,
    flex: 1,
  },
};

class HomeScreen extends Component {
  constructor(props) {
    super(props);

    this.state = {
      
    };
  
    this.handleGoToCreatePlayer = this.handleGoToCreatePlayer.bind(this);
  }

  handleGoToCreatePlayer() {
    this.props.history.push('/players/new');
  }

  render() {
    const { classes } = this.props;
    return (
      <div className={classes.container}>
        <AppBar title="home" navpane/>
        <Grid container>
          <Grid item xs>
            <Button variant="extendedFab" color="secondary" fullWidth className={classes.actions}>
              Generate Teams
            </Button>

            <Button variant="extendedFab"  color="secondary" onClick={this.handleGoToCreatePlayer} fullWidth className={classes.actions}>
                Add Players
            </Button>

            <Button variant="extendedFab"  color="secondary" fullWidth className={classes.actions}>
              Edit Players
            </Button>
          </Grid>
        </Grid>
      </div>
    );
  }
}

export default withRouter(withStyles(styles)(HomeScreen));
