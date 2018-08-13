
import React, { Component } from 'react';
import StarRatings from './StarRatings';
import { withStyles } from '@material-ui/core/styles';
import {
    Paper,
    ListItem,
    ListItemText,
    ListItemSecondaryAction,
    Checkbox,
    Divider,
} from '@material-ui/core';

const styles = () => ({
  paper: {
    margin: '0 10px',
  },
});

const PlayerListItem = ({key, name, ratings, secondaryAction, inPaper, index, classes}) => {
    const content = (
      <React.Fragment>
        <ListItem>
          <ListItemText primary={`${index ? `${index}. ` : ''}${name}`} />
          <StarRatings ratings={ratings}/>
          <ListItemSecondaryAction>
            {secondaryAction}
          </ListItemSecondaryAction>
        </ListItem>
        <Divider />
      </React.Fragment>
    );

    return (
        <React.Fragment key={key}>
          {inPaper ? <Paper className={classes.paper}>{content}</Paper> : content}
        </React.Fragment>
    );
}

export default withStyles(styles)(PlayerListItem);