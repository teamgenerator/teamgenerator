import React, { Component } from "react";
import { Grid, Typography } from "@material-ui/core";
import { withStyles } from "@material-ui/core/styles";
import AppBar from "../components/AppBar";
import { connect } from "react-redux";
import moment from "moment";
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
  pl2: {
    "padding-left": "20px !important"
  },
  pb2: {
    "padding-bottom": "20px !important"
  }
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
    const { classes, match, session } = this.props;
    if (!session) {
      return <div />;
    }
    console.log(this.props.session);
    return (
      <div className={classes.container}>
        <AppBar title={"Sessions"} />
        <Grid container spacing={24} alignItems="center">
          <Grid item xs={12} className={classes.pl2}>
            <Typography
              variant="title"
              className={classes.pb2}
            >{`Session #${match.params.id}`}</Typography>
            <Typography variant="subheading">
              {`Created at: ${moment(session.createdAt).format(
                "MMMM Do YYYY"
              )}`}
            </Typography>
            <Typography variant="subheading">
              {`Is active: ${session.isActive}`}
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
