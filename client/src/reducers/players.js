const players = (state = {}, action) => {
  switch (action.type) {
    case 'PLAYERS_MERGE': {
      return Object.assign({}, state, action.data);
    }
    case 'PLAYERS_REPLACE': {
      return Object.assign({}, action.data);
    }
    default: {
      return state;
    }
  }
};

export default players;
