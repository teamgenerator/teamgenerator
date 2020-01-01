import React, { Component } from "react";
import {
  Button,
  Grid,
  TextField,
  IconButton,
  Icon,
  Typography
} from "@material-ui/core";
import { withStyles } from "@material-ui/core/styles";
import AppBar from "../components/AppBar";
import { connect } from "react-redux";
import makeRequestApiActionThread from "../actions/apiRequest";

const styles = {
  container: {
    display: "flex",
    flexDirection: "column"
  },
  chevrons: {
    left: "50%",
    transform: "translateX(-50%)"
  },
  ratingsGridItem: {}
};

class SessionDetails extends Component {
  constructor(props) {
    super(props);

    this.state = {};

    this.props.dispatch(
      makeRequestApiActionThread(
        "GET",
        "/sessions",
        undefined,
        "REPLACE",
        "sessions"
      )
    );
    this.props.dispatch(
      makeRequestApiActionThread(
        "GET",
        "/communities",
        undefined,
        "REPLACE",
        "communities"
      )
    );
  }

  setUpdatedAttributes(updatedAttributes) {
    this.setState({
      updatedAttributes: {
        ...this.state.updatedAttributes,
        ...updatedAttributes
      }
    });
  }
  render() {
    const { classes, match } = this.props;
    if (!this.props.session) {
      return <div />;
    }
    console.log(this.props.session);
    return (
      <div className={classes.container}>
        <AppBar title={"Sessions"} />
        <Grid container spacing={24} alignItems="center">
          <Grid item xs={12}>
            <Typography variant="display1">{`Session #${match.params.id}`}</Typography>
            <Typography variant="subheading">
              Rate player's performance for this game. Your ratings will help us
              get some insight of the players performance trend and adjust the
              ratings overtime.
            </Typography>
          </Grid>
        </Grid>
      </div>
    );
  }
}

const mapStateToProps = (state, ownProps) => ({
  session: Object.values(state.sessions).find(
    session => session.id == ownProps.match.params.id
  )
});

const mapDispatchToProps = dispatch => ({ dispatch });

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(withStyles(styles)(SessionDetails));
