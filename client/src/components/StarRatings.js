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
    width: '200px',
  },
  center: {
    margin: 'auto',
  },
  starJumbo:  {
    fontSize: '20px',
  }
};

class StarRatings extends Component {

  constructor(props) {
    super(props);

    this.state = {};
  }

  render() {
    const { classes, justStars } = this.props;
    const stars = [];

    for (let i = 1; i <= this.props.ratings; i++) {
      stars.push(<Star key={i} onClick={() => this.props.onStarChange(i)} className={[classes.star, justStars ? classes.starJumbo : ''].join(' ')} />)
    }

    for (let i = this.props.ratings + 1; i <= 10; i++) {
      stars.push(<Star key={i} onClick={() => this.props.onStarChange(i)} className={[classes.starDisabled, justStars ? classes.starJumbo : '']. join(' ')} />);
    }

    return (
      <div className={[classes.container, justStars ? classes.center : ''].join(' ')}>
        {stars}
        {
          !this.props.justStars
          ? <Typography variant="caption" className={classes.ratingsFraction}>({this.props.ratings} / 10)</Typography>
          : null
        }
      </div>
    );
  }
}

StarRatings.propTypes = {
  ratings: PropTypes.number.isRequired,
  justStars: PropTypes.bool,
  onStarChange: PropTypes.func,
};

StarRatings.defaultProps = {
  justStars: false,
  onStarChange: null,
};

export default withStyles(styles)(StarRatings);