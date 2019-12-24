import React, { Component } from "react";
import {
  Grid,
  Typography,
  List,
  ListItem,
  ListItemText,
  ListItemSecondaryAction,
  Divider
} from "@material-ui/core";
import { withStyles } from "@material-ui/core/styles";
import AppBar from "../components/AppBar";
import { connect } from "react-redux";
import Save from "@material-ui/icons/Save";
import ChevronRight from "@material-ui/icons/ChevronRight";
import { diff } from "deep-object-diff";

const styles = {
  container: {
    display: "flex",
    flexDirection: "column"
  },
  fontRed: {
    color: "red"
  },
  chevrons: {
    left: "50%",
    transform: "translateX(-50%)"
  }
};

class SessionIndex extends Component {
  constructor(props) {
    super(props);
    this.state = {};
  }

  componentDidUpdate(prevProps, prevState) {
    const playerDiffBetweenPrevStateAndCurrState = diff(
      prevState.updatedAttributes,
      this.state.updatedAttributes
    );
    if (Object.keys(playerDiffBetweenPrevStateAndCurrState).length > 0) {
      const playerDiff = diff(this.props.player, this.state.updatedAttributes);
      this.setState({ isDirty: Object.keys(playerDiff).length > 0 });
    }
  }

  setUpdatedAttributes(updatedAttributes) {
    this.setState({
      updatedAttributes: {
        ...this.state.updatedAttributes,
        ...updatedAttributes
      }
    });
  }

  renderSessionListItem(sessionKey, sessionDate, sessionNumPlayers) {
    return (
      <React.Fragment key={sessionKey}>
        <ListItem button>
          <ListItemText primary={sessionDate} />
          <ListItemText secondary={`Num TPlayers: ${sessionNumPlayers}`} />
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
        <AppBar
          title="Sessions"
          rightButton={{
            icon: <Save />,
            onClick: () => alert("clicked"),
            disabled: !this.state.isDirty
          }}
        />
        <Grid container spacing={24} alignItems="center">
          <Grid item xs={12}>
            <Typography variant="title">
              <span className={classes.fontRed}>
                {this.props.activeCommunity.Name}
              </span>
              {" sessions"}
            </Typography>
          </Grid>
        </Grid>
        <List>
          {this.props.sessionsArray.map(session =>
            this.renderSessionListItem(session.ID, session.CreatedAt, 1)
          )}
        </List>
      </div>
    );
  }
}

const mapStateToProps = state => ({
  sessions: state.sessions,
  sessionsArray: Object.values(state.sessions),
  activeCommunity: state.activeCommunity
});

const mapDispatchToProps = dispatch => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(withStyles(styles)(SessionIndex));
