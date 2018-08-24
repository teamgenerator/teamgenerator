import React, { Component } from 'react';
import { Button, Grid, TextField, IconButton, Icon, Typography } from '@material-ui/core';
import { withStyles } from '@material-ui/core/styles';
import AppBar from '../components/AppBar';
import { connect } from 'react-redux';
import Save from '@material-ui/icons/Save';
import { diff } from 'deep-object-diff';
import StarRatings from '../components/StarRatings';
import ChevronLeft from '@material-ui/icons/ChevronLeft';
import ChevronRight from '@material-ui/icons/ChevronRight';

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

class SessionDetails extends Component {
  constructor(props) {
    super(props);

    this.state = {
      updatedAttributes: {
        id: '001',
        isActive: true,

      },
    };
  }

  componentDidUpdate(prevProps, prevState) {
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
    const title = this.props.match.params.id !== 'new' ? this.props.player.name : 'Current Session';
    return (
      <div className={classes.container}>
        <AppBar title={title} rightButton={{
          icon: <Save />,
          onClick: () => alert('clicked'),
          disabled: !this.state.isDirty,
        }}/>
        <Grid container spacing={24} alignItems="center">
          <Grid item xs={12}>
            <Typography variant="display1">New session has started</Typography>
            <Typography variant="subheading">Rate player's performance for this game. 
            Your ratings will help us get some insight of the players performance trend and adjust the ratings overtime.</Typography>
          </Grid>
        </Grid>
      </div>
    );
  }
}

const mapStateToProps = (state, ownProps) => ({
  player: Object.values(state.players).find(player => player.id === ownProps.match.params.id) || {},
});

const mapDispatchToProps = dispatch => ({
  
});



export default connect(mapStateToProps, mapDispatchToProps)(withStyles(styles)(SessionDetails));
