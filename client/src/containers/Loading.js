import React, { Component } from 'react';
import { Button, Grid, CircularProgress } from '@material-ui/core';
import { withStyles } from '@material-ui/core/styles';

import AppBar from '../components/AppBar';
import { connect } from 'react-redux';

const styles = {
  container: {
    zIndex: 20000,
    display: 'flex',
    height: '100vh',
    width: '100vw',
    alignItems: 'center',
    justifyContent: 'center',
    backgroundColor: '#3535358a',
    position: 'fixed',
  },
};

class Loading extends Component {
  constructor(props) {
    super(props);

    this.state = {

    };
  }

  render() {
    const pendingRequests = this.props.ui.pendingRequests;

    const { classes } = this.props;
    return (
      <React.Fragment>
        {pendingRequests > 0 && 
        <div className={classes.container}>
          <CircularProgress size={50} />
        </div>}
        {this.props.children}
      </React.Fragment>
    );
  }
}

const mapStateToProps = state => ({
  ui: state.ui,
});

const mapDispatchToProps = dispatch => ({});

export default connect(mapStateToProps, mapDispatchToProps)(withStyles(styles)(Loading));
