import React, { Component } from 'react';
import { withStyles } from '@material-ui/core/styles';
import PropTypes from 'prop-types';

import { Typography, Paper, Divider, List } from '../../node_modules/@material-ui/core';
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
            <Typography variant="title">{this.props.name}</Typography>
            <Divider />
            <List>
              {this.props.playerListItemProps.map((p, index) =>
                <PlayerListItem key={p.id} name={p.name} index={index + 1} ratings={p.ratings} />
              )}
            </List>
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