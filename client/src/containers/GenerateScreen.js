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
import Casino from '@material-ui/icons/Casino';
import TeamPanel from '../components/TeamPanel';
import generateRandomTeam from '../utils/generateRandomTeam';


const styles = {
  content: {
    marginTop: '1%',
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
  rerollButton: {
    color: 'grey',
    float: 'right',
    marginRight: '10px',
  },
  casino: {
    marginRight: '5px',
  },
};

class GenerateScreen extends Component {
  constructor(props) {
    super(props);
    this.state = {
      step: 1,
      selectedPlayersKeys: [],
      extraPlayersSelectedOption: null,
      numberOfPlayersPerTeam: 1,
      teams: null,
    };
    this.handleToggleCheckbox = this.handleToggleCheckbox.bind(this);
    this.handleMoveToAnotherStep = this.handleMoveToAnotherStep.bind(this);
    this.handleRadioChange = this.handleRadioChange.bind(this);
    this.handleChangeNumberOfPlayersPerTeam = this.handleChangeNumberOfPlayersPerTeam.bind(this);
    this.getStepAttributes = this.getStepAttributes.bind(this);
    this.getStepOneAttributes = this.getStepOneAttributes.bind(this);
    this.getStepTwoAttributes = this.getStepTwoAttributes.bind(this);
    this.getStepThreeAttributes = this.getStepThreeAttributes.bind(this);
    this.handleGenerateNewTeams = this.handleGenerateNewTeams.bind(this);
  }

  componentDidUpdate(prevProps, prevState) {
    if (
      this.state.step === 2 &&
      this.state.numberOfPlayersPerTeam > this.state.selectedPlayersKeys.length
    ) {
      this.setState({ numberOfPlayersPerTeam: 1 });
    }

    if (this.state.step === 3 && !this.state.teams) {
      this.handleGenerateNewTeams();
    }

    if (this.state.step === 3 && prevState.step !== 3 ) {
      this.handleGenerateNewTeams();
    }
  }

  handleGenerateNewTeams() {
    const selectedPlayers = this.props.playersArray.filter(p => this.state.selectedPlayersKeys.includes(p.id));
    this.setState({ teams: generateRandomTeam(selectedPlayers, this.state.numberOfPlayersPerTeam)})
  }

  getStepAttributes(step) {
    switch (step) {
      case 1:
        return this.getStepOneAttributes();
      case 2:
        return this.getStepTwoAttributes();
      case 3:
        return this.getStepThreeAttributes();
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
          <IconButton className={classes.chevrons} onClick={() => this.handleChangeNumberOfPlayersPerTeam(-1)} disabled={this.state.numberOfPlayersPerTeam <= 1}>
            <ChevronLeft />
          </IconButton>
        </Grid>
        <Grid item xs={4}>
          <Typography variant="display2" align="center">
            {this.state.numberOfPlayersPerTeam}
          </Typography>
        </Grid>
        <Grid item xs={4}>
          <IconButton className={classes.chevrons} onClick={() => this.handleChangeNumberOfPlayersPerTeam(1)} disabled={this.state.numberOfPlayersPerTeam >= this.state.selectedPlayersKeys.length}>
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

  getStepThreeAttributes() {
    const { classes } = this.props;

    const content = (
      <Grid container spacing={12} className={classes.content} alignContent="center">
        <Grid item xs={12} >
          <Button variant="raised" size="small" className={classes.rerollButton} onClick={this.handleGenerateNewTeams}>
            <Casino className={classes.casino} /> Reroll
          </Button>
        </Grid>
        <Grid item xs={12}>
          {(this.state.teams || []).map((team, i) => {
            return <TeamPanel name={`Team ${i + 1}`} playerListItemProps={team}/>
          })}
        </Grid>
      </Grid>
    );

    return {
      content,
      title: 'Confirm Team',
      leftButtonLabel: 'Back',
      rightButtonLabel: 'Play',
    };
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

  handleChangeNumberOfPlayersPerTeam(change) {
    this.setState({
      numberOfPlayersPerTeam: this.state.numberOfPlayersPerTeam + change,
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
