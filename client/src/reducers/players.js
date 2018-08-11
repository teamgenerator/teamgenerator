const defaultState = {
  '0001': {
    id: '0001',
    name: 'Michael Englo',
    ratings: 5,
  },
  '0002': {
    id: '0002',
    name: 'Nico Alimin',
    ratings: 8,
  },
};

const players = (state = defaultState, action) => {
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
