import React, { Component } from 'react';
import { Provider } from 'react-redux';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import store from './store';

import HomeScreen from "./containers/HomeScreen";
import PlayersScreen from "./containers/PlayersScreen";
import PlayerDetails from "./containers/PlayerDetails";
import GenerateScreen from "./containers/GenerateScreen";
import SessionDetails from "./containers/SessionDetails";
import SessionIndex from "./containers/SessionIndex";
import Loading from './containers/Loading';

class App extends Component {
  constructor(props) {
    super(props);

    this.state = {};
  }

  render() {
    return (
      <Provider store={store}>
        <Loading>
          <BrowserRouter>
            <Switch>
              <Route exact path="/" component={HomeScreen} />
              // Router for players
              <Route exact path="/players" component={PlayersScreen} />
              <Route path="/players/:id" component={PlayerDetails} />
              <Route exact path="/generate" component={GenerateScreen} />
              <Route exact path="/sessions" component={SessionIndex} />
              <Route path="/sessions/:id" component={SessionDetails} />
            </Switch>
          </BrowserRouter>
        </Loading>
      </Provider>
    );
  }
}

export default App;
