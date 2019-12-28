const defaultUI = {
  pendingRequests: 0,
  activeCommunity: "1"
};

const ui = (state = defaultUI, action) => {

  if (action.type.includes("API_REQUEST")) {
    return  {
      ...state,
      pendingRequests: state.pendingRequests + 1,
    };
  } else if (action.type.includes("API_RECEIVE")) {
    return  {
      ...state,
      pendingRequests: state.pendingRequests - 1,
    };
  } else if (action.type === "UI_COMMUNITY_SET") {
    const newState = {
      ...state,
      activeCommunity: action.data
    };
    return newState;
  }

  return state;
};

export default ui;
