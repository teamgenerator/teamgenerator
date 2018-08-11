import React, { Component } from 'react';
import { withStyles } from '@material-ui/core/styles';
import PropTypes from 'prop-types';

import Star from '@material-ui/icons/Star';
import { Typography } from '../../node_modules/@material-ui/core';

const styles = {
  star: {
    fontSize: '12px',
    color: '#ffd700',
  },
  starDisabled: {
    fontSize: '12px',
    color: '#dddddd',
  },
  ratingsFraction: {
    marginLeft: '2px',
    fontSize: '10px',
    width: '50px',
  },
  container: {
    width: '120px',
  },
};

class StarRatings extends Component {

  constructor(props) {
    super(props);

    this.state = {};
  }

  render() {
    const { classes } = this.props;
    const stars = [];

    for (let i = 0; i < this.props.ratings; i++) {
      stars.push(<Star key={i} className={classes.star}/>)
    }

    for (let i = this.props.ratings; i < 10; i++) {
      stars.push(<Star key={i} className={classes.starDisabled} />);
    }

    return (
      <div className={classes.container}>
        {stars}
        <Typography variant="caption" className={classes.ratingsFraction}>({this.props.ratings} / 10)</Typography>
      </div>
    );
  }
}

StarRatings.propTypes = {
  ratings: PropTypes.number.isRequired,
};

export default withStyles(styles)(StarRatings);