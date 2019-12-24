const defaultCommunities = {
  "1": {
    ID: 1,
    CreatedAt: "2019-06-24T03:12:00.987197Z",
    UpdatedAt: "2019-06-24T05:24:22.166288Z",
    Name: "Update",
    Location: "Richmond"
  }
};

const communities = (state = defaultCommunities, action) => {
  switch (action.type) {
    case "COMMUNITIES_GET": {
      return {
        ...state
      };
    }
    case "COMMUNITIES_REPLACE": {
      return {
        ...action.data
      };
    }
    default: {
      return state;
    }
  }
};

export default communities;
