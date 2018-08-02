import React, { Component } from 'react';
import { Button, Grid } from '@material-ui/core';
import { withStyles } from '@material-ui/core/styles';


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
  render() {
    const { classes } = this.props;
    return (
      <div className={classes.container}>
        <Grid container>
          <Grid item xs>
            <Button variant="extendedFab" color="secondary" fullWidth className={classes.actions}>
              Generate Teams
            </Button>

            <Button variant="extendedFab"  color="secondary" fullWidth className={classes.actions}>
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

export default withStyles(styles)(HomeScreen);
