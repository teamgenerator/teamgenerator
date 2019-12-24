const defaultSessions = {
  "1": {
    ID: 1,
    CreatedAt: "2019-06-24T05:27:56.577219Z",
    UpdatedAt: "2019-06-24T05:29:50.233697Z",
    is_active: false,
    community_id: 1
  },
  "2": {
    ID: 4,
    CreatedAt: "2019-06-24T05:32:47.873579Z",
    UpdatedAt: "2019-06-24T05:32:47.873579Z",
    is_active: false,
    community_id: 1
  }
};

const sessions = (state = defaultSessions, action) => {
  switch (action.type) {
    case "SESSIONS_MERGE": {
      return {
        ...state,
        ...action.data
      };
    }
    case "SESSIONS_REPLACE": {
      return {
        ...action.data
      };
    }
    default: {
      return state;
    }
  }
};

export default sessions;
