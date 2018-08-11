import React, { Component } from 'react';
import { Provider } from 'react-redux';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import configureStore from './store';

import HomeScreen from './containers/HomeScreen';
import CreatePlayerScreen from './containers/CreatePlayerScreen';

const store = configureStore();
class App extends Component {

  constructor(props) {
    super(props);

    this.state = {};
  }

  render() {
    return (
      <Provider store={store}>
        <BrowserRouter>
          <Switch>
            <Route exact path="/" component={HomeScreen} />

            // Router for players
            <Route path="/players/new" component={CreatePlayerScreen} />
          </Switch>
        </BrowserRouter>
      </Provider>
    );
  }
}

export default App;