import React, { Component } from 'react';
import {
  Button,
  Grid,
  Divider,
  ListItem,
  ListItemSecondaryAction,
  MobileStepper,
  Typography,
  Checkbox,
  ListItemText,
  Paper,
  List,
  Radio,
  RadioGroup,
  FormControlLabel,
  IconButton
} from '@material-ui/core';
import { withStyles } from '@material-ui/core/styles';
import { connect } from 'react-redux';
import StarRatings from '../components/StarRatings';
import { withRouter } from 'react-router-dom';
import ChevronLeft from '@material-ui/icons/ChevronLeft';
import ChevronRight from '@material-ui/icons/ChevronRight';


const styles = {
  content: {
    marginTop: '2%',
  },
  paper: {
    marginBottom: '10px',
  },
  playerList: {
    backgroundColor: '#fafafa',
    position: 'fixed',
    height: '78%',
    width: '100%',
    bottom: 0,
    overflow: 'scroll',
  },
  chevrons: {
    left: '50%',
    transform: 'translateX(-50%)',
  },
};

class GenerateScreen extends Component {
  constructor(props) {
    super(props);

    this.state = {
      step: 1,
      selectedPlayersKeys: [],
      extraPlayersSelectedOption: null,
      numberOfTeams: 1,
    };
    this.handleToggleCheckbox = this.handleToggleCheckbox.bind(this);
    this.handleMoveToAnotherStep = this.handleMoveToAnotherStep.bind(this);
    this.handleRadioChange = this.handleRadioChange.bind(this);
    this.handleChangeNumberOfTeams = this.handleChangeNumberOfTeams.bind(this);
    this.getStepAttributes = this.getStepAttributes.bind(this);
    this.getStepOneAttributes = this.getStepOneAttributes.bind(this);
    this.getStepTwoAttributes = this.getStepTwoAttributes.bind(this);
  }

  componentDidUpdate(prevProps, prevState) {
    if (
      this.state.step === 2 &&
      this.state.numberOfTeams > this.state.selectedPlayersKeys.length
    ) {
      this.setState({ numberOfTeams: 1 });
    }
  }

  getStepAttributes(step) {
    switch (step) {
      case 1:
        return this.getStepOneAttributes();
      case 2:
        return this.getStepTwoAttributes();
      default:
        break;
    }
  }

  getStepOneAttributes() {
    const { classes } = this.props;

    const content = (
      <Grid container spacing={24} className={classes.content}>
        <Grid item xs={12}>
          <Typography variant="body1">
            Pick at least one player to go to the next step
          </Typography>
          <Typography variant="caption">
            {this.state.selectedPlayersKeys.length} of {this.props.playersArray.length} players selected
          </Typography>
        </Grid>
        <Grid item xs={12} className={classes.playerList}>
          <List>
            {this.props.playersArray.map(player =>
              this.renderPlayerListItem(player.id, player.name, player.ratings, this.state.selectedPlayersKeys.includes(player.id)))
            }
          </List>
        </Grid>
      </Grid>
    );

    return {
      content,
      title: 'Select Players',
      leftButtonLabel: 'Back',
      rightButtonDisabled: this.state.selectedPlayersKeys <= 0,
      rightButtonLabel: 'Next',
    }
  }

  getStepTwoAttributes() {
    const { classes } = this.props;

    const content = (
      <Grid container spacing={24} className={classes.content} alignContent="center">
        <Grid item xs={12}>
          <Typography variant="subheading">
            Number of Players per team:
          </Typography>
        </Grid>
        <Grid item xs={4}>
          <IconButton className={classes.chevrons} onClick={() => this.handleChangeNumberOfTeams(-1)} disabled={this.state.numberOfTeams <= 1}>
            <ChevronLeft />
          </IconButton>
        </Grid>
        <Grid item xs={4}>
          <Typography variant="display2" align="center">
            {this.state.numberOfTeams}
          </Typography>
        </Grid>
        <Grid item xs={4}>
          <IconButton className={classes.chevrons} onClick={() => this.handleChangeNumberOfTeams(1)} disabled={this.state.numberOfTeams >= this.state.selectedPlayersKeys.length}>
            <ChevronRight />
          </IconButton>
        </Grid>
        <Grid item xs={12}>
          <Typography variant="subheading">
            Extra Players:
          </Typography>
        </Grid>
        <Grid item xs={12}>
          <RadioGroup
            name="extra-players-options"
            value={this.state.extraPlayersSelectedOption}
            onChange={this.handleRadioChange}
          >
            <FormControlLabel value="another-team" control={<Radio />} label="Make Another Team" />
            <FormControlLabel value="substitutes" control={<Radio />} label="Add As Substitutes" />
          </RadioGroup>
        </Grid>
      </Grid>
    );

    return {
      content,
      title: 'Additional Settings',
      leftButtonLabel: 'Back',
      rightButtonLabel: 'Roll',
      rightButtonDisabled: !this.state.extraPlayersSelectedOption,
    }
  }

  renderPlayerListItem(playerKey, playerName, playerRatings, playerSelected) {
    return (
      <Paper key={playerKey} className={this.props.classes.paper}>
        <ListItem>
          <ListItemText primary={playerName} />
          <StarRatings ratings={playerRatings}/>
          <ListItemSecondaryAction>
            <Checkbox
              onChange={() => this.handleToggleCheckbox(playerKey)}
              checked={playerSelected}
            />
          </ListItemSecondaryAction>
        </ListItem>
        <Divider />
      </Paper>
    );
  }

  handleChangeNumberOfTeams(change) {
    this.setState({
      numberOfTeams: this.state.numberOfTeams + change,
    });
  }

  handleRadioChange(e) {
    this.setState({ extraPlayersSelectedOption: e.currentTarget.value});
  }

  handleToggleCheckbox(playerKey) {
    if (this.state.selectedPlayersKeys.includes(playerKey)) {
      this.setState({
        selectedPlayersKeys: this.state.selectedPlayersKeys.filter(key => key !== playerKey)
      });
    } else {
      this.setState({
        selectedPlayersKeys: [...this.state.selectedPlayersKeys, playerKey],
      });
    }
  }

  handleMoveToAnotherStep(stepChanges) {
    const currentStep = this.state.step + stepChanges;

    switch (currentStep) {
      case 0:
        this.props.history.goBack();
        break;
      default:
        this.setState({
          step: currentStep,
        });
        break;
    }
  }

  render() {
    const { classes } = this.props;
    const stepAttributes = this.getStepAttributes(this.state.step);
    return (
      <div className={classes.container}>
        <MobileStepper
          variant="progress"
          steps={3}
          position="static"
          activeStep={this.state.step - 1}
          nextButton={
              <Button size="small" onClick={() => this.handleMoveToAnotherStep(1)} disabled={stepAttributes.rightButtonDisabled}>
                {stepAttributes.rightButtonLabel}
              </Button>
          }
          backButton={
              <Button size="small" onClick={() => this.handleMoveToAnotherStep(-1)}>
                {stepAttributes.leftButtonLabel}
              </Button>
          }
        />
        <Typography variant="headline">Step {this.state.step}: {stepAttributes.title}</Typography>
        {stepAttributes.content}  
      </div>
    );
  }
}

const mapStateToProps = state => ({
  players: state.players,
  playersArray: Object.values(state.players),
});

const mapDispatchToProps = dispatch => ({
  
});



export default connect(mapStateToProps, mapDispatchToProps)(withRouter(withStyles(styles)(GenerateScreen)));
