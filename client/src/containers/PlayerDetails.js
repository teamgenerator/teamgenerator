import React, { Component } from 'react';
import { Button, Grid, TextField, IconButton, Icon } from '@material-ui/core';
import { withStyles } from '@material-ui/core/styles';
import AppBar from '../components/AppBar';
import { connect } from 'react-redux';
import Save from '@material-ui/icons/Save';
import { diff } from 'deep-object-diff';
import StarRatings from '../components/StarRatings';
import ChevronLeft from '@material-ui/icons/ChevronLeft';
import ChevronRight from '@material-ui/icons/ChevronRight';
import makeRequestApiActionThread from '../actions/apiRequest';

const styles = {
  container: {
    display: 'flex',
    flexDirection: 'column',
  },
  chevrons: {
    left: '50%',
    transform: 'translateX(-50%)',
  },
  ratingsGridItem: {
  },
};

class PlayerDetails extends Component {
  constructor(props) {
    super(props);

    this.state = {
      updatedAttributes: {
        id:  this.props.player.id || null,
        name: this.props.player.name || '',
        ratings: this.props.player.ratings || 5,
        isDirty: false,
      },
    };

    this.handleNameTextfieldChange = this.handleNameTextfieldChange.bind(this);
    this.handleEditRatings = this.handleEditRatings.bind(this);
    this.setUpdatedAttributes = this.setUpdatedAttributes.bind(this);
  }

  componentDidMount() {
    this.props.dispatch(makeRequestApiActionThread("GET", "/players", undefined, "REPLACE", "player"));
  }

  componentDidUpdate(prevProps, prevState) {
    const playerDiffBetweenProps

    const playerDiffBetweenPrevStateAndCurrState = diff(prevState.updatedAttributes, this.state.updatedAttributes);
    if (Object.keys(playerDiffBetweenPrevStateAndCurrState).length > 0) {
      const playerDiff = diff(this.props.player, this.state.updatedAttributes);
      this.setState({ isDirty: Object.keys(playerDiff).length > 0 });
    }
  }

  setUpdatedAttributes(updatedAttributes) {
    this.setState({
      updatedAttributes: {
        ...this.state.updatedAttributes,
        ...updatedAttributes,
      }
    });
  }

  handleNameTextfieldChange(e) {
    this.setUpdatedAttributes({
        name: e.currentTarget.value,
      });
  }

  handleEditRatings(ratingsChange) {
    this.setUpdatedAttributes({
      ratings: this.state.updatedAttributes.ratings + ratingsChange,
    });
  }

  render() {
    const { classes } = this.props;
    const title = this.props.match.params.id !== 'new' ? this.props.player.name : 'New Player';
    console.warn(title);
    return (
      <div className={classes.container}>
        <AppBar title={title} rightButton={{
          icon: <Save />,
          onClick: () => alert('clicked'),
          disabled: !this.state.isDirty,
        }}/>
        <Grid container spacing={24} alignItems="center">
          <Grid item xs={12}>
            <TextField
              label="Full Name"
              fullWidth
              value={this.state.updatedAttributes.name}
              onChange={this.handleNameTextfieldChange}
            />
          </Grid>
          <Grid item xs={2}>
            <IconButton
              className={classes.chevrons}
              onClick={() => this.handleEditRatings(-1)}
              disabled={this.state.updatedAttributes.ratings <= 0}
            >
              <ChevronLeft />
            </IconButton>
          </Grid>
          <Grid item xs={8} className={classes.ratingsGridItem}>
            <StarRatings onStarChange={newRating => this.handleEditRatings(newRating - this.state.updatedAttributes.ratings)} justStars ratings={this.state.updatedAttributes.ratings} />
          </Grid>
          <Grid item xs={2}>
            <IconButton className={classes.chevrons}
              onClick={() => this.handleEditRatings(1)}
              disabled={this.state.updatedAttributes.ratings >= 10}
            >
              <ChevronRight />
            </IconButton>
          </Grid>
        </Grid>
      </div>
    );
  }
}

const mapStateToProps = (state, ownProps) => ({
  players: state.players,
  player: Object.values(state.players).find(player => !Number.isNaN(ownProps.match.params.id) && player.id === Number(ownProps.match.params.id)) || {},
});

const mapDispatchToProps = dispatch => ({
  dispatch,
});



export default connect(mapStateToProps, mapDispatchToProps)(withStyles(styles)(PlayerDetails));
