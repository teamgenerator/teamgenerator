import React, { Component } from 'react';
import { Button, Grid, TextField } from '@material-ui/core';
import { withStyles } from '@material-ui/core/styles';
import AppBar from '../components/AppBar';

const styles = {
  container: {
    display: 'flex',
    flexDirection: 'column',
  },
};

class PlayerDetails extends Component {
  constructor(props) {
    super(props);

    this.state = {
      currentAttributes:  {
        name: 'Michael Englo',
        ratings: 7,
      }
    };
  }

  render() {
    const { classes } = this.props;
    return (
      <div className={classes.container}>
        <AppBar title="Add New Player" handleOkay={() => {}} />
        <Grid container spacing={24}>
          <Grid item xs={12}>
            <TextField
              label="Full Name"
              value={this.state.currentAttributes.name}
            />
          </Grid>
          <Grid item xs={12}>
            <TextField
              label="Ratings"
              value={this.state.currentAttributes.ratings}
            />
          </Grid>
        </Grid>
      </div>
    );
  }
}

export default withStyles(styles)(PlayerDetails);
