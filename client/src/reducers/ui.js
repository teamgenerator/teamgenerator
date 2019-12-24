const defaultUI = {
  activeCommunity: "1"
};

const ui = (state = defaultUI, action) => {
  switch (action.type) {
    case "UI_COMMUNITY_GET": {
      return {
        ...state
      };
    }
    case "UI_COMMUNITY_SET": {
      const newState = {
        ...state,
        activeCommunity: action.data
      };
      return newState;
    }
    default: {
      return state;
    }
  }
};

export default ui;
