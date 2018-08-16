import React, { Component } from 'react';
import { withStyles } from '@material-ui/core/styles';
import PropTypes from 'prop-types';

import { Typography, Paper, Divider, List, Grid } from '../../node_modules/@material-ui/core';
import PlayerListItem from './PlayerListItem';

const styles = {
  container: {
    padding: '10px',
    margin: '10px',
  },
};

class TeamPanel extends Component {

  constructor(props) {
    super(props);

    this.state = {};
  }

  render() {
    const { classes } = this.props;
    return (
        <Paper className={classes.container}>
            <Grid container>
              <Grid item xs={4}>
                <Typography variant="title">{this.props.name}</Typography>
              </Grid>
              <Grid item xs={4}>
                <Typography variant="caption">Players: {this.props.playerListItemProps.length}</Typography>
                <Typography variant="caption">Total Stars: {this.props.playerListItemProps.reduce((accumulator, player) => accumulator + player.ratings, 0)}</Typography>
              </Grid>
              <Grid item xs={4} />
              <Grid item xs={12}>
                <Divider />
                <List>
                  {this.props.playerListItemProps.map((p, index) =>
                    <PlayerListItem key={p.id} name={p.name} index={index + 1} ratings={p.ratings} />
                  )}
                </List>
              </Grid>
            </Grid>
        </Paper>
    );
  }
}

TeamPanel.propTypes = {
  name: PropTypes.string.isRequired,
  playerListItemProps: PropTypes.arrayOf(PropTypes.shape({
        id: PropTypes.string,
        name: PropTypes.string.isRequired,
        ratings: PropTypes.number.isRequired,
    })).isRequired,
};

export default withStyles(styles)(TeamPanel);