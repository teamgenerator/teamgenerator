import React, { Component } from 'react';
import { Button, Grid, TextField, List, ListItem, ListItemText, ListItemSecondaryAction, Divider } from '@material-ui/core';
import { withStyles } from '@material-ui/core/styles';
import AppBar from '../components/AppBar';
import StarRatings from '../components/StarRatings';
import ChevronRight from '@material-ui/icons/ChevronRight';
import { connect } from 'react-redux';
import makeRequestApiActionThread from '../actions/apiRequest';

const styles = {
  container: {
    display: 'flex',
    flexDirection: 'column',
  },
};

class PlayersScreen extends Component {
  constructor(props) {
    super(props);

    this.state = {};

    this.renderPlayerListItem = this.renderPlayerListItem.bind(this);
    this.handleClickPlayerList = this.handleClickPlayerList.bind(this);
  }

  componentDidMount() {
    this.props.dispatch(makeRequestApiActionThread("GET", "/players", undefined, "REPLACE", "players"));
  }

  handleClickPlayerList(playerKey) {
    this.props.history.push('/players/' + playerKey);
  }

  renderPlayerListItem(playerKey, playerName, playerRatings) {
    return (
      <React.Fragment key={playerKey}>
        <ListItem button onClick={() => this.handleClickPlayerList(playerKey)}>
          <ListItemText primary={playerName} />
          <StarRatings ratings={playerRatings}/>
          <ListItemSecondaryAction>
            <ChevronRight />
          </ListItemSecondaryAction>
        </ListItem>
        <Divider />
      </React.Fragment>
    );
  }

  render() {
    const { classes } = this.props;
    return (
      <div className={classes.container}>
        <AppBar title="Players" navpane />
        <List>
          {this.props.playersArray.map(player => this.renderPlayerListItem(player.id, player.name, player.ratings))}
        </List>
      </div>
    );
  }
}

const mapDispatchToProps = dispatch => ({
  dispatch,
});

const mapStateToProps = state => ({
  players: state.players,
  playersArray: Object.values(state.players),
});

export default connect(mapStateToProps, mapDispatchToProps)(withStyles(styles)(PlayersScreen));
