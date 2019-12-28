// This function returns a simpleReducer that contains 
// name: must be uppercase snake case and plural (e.g. PLAYER_SESSIONS).
const createSimpleReducer  = (name, defaultState) => (state = defaultState || {}, action) => {
    switch (action.type) {
      case `${name}_MERGE`: {
        return Object.assign({}, state, action.data);
      }
      case `${name}_REPLACE`: {
        return Object.assign({}, action.data);
      }
      default: {
        return state;
      }
    }
  };

  export default createSimpleReducer;