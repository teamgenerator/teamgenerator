import { combineReducers } from "redux";
import ui from "./ui";
import createSimpleReducer from './simpleReducer';

const players = createSimpleReducer("PLAYERS");
const sessions = createSimpleReducer("SESSIONS");
const communities = createSimpleReducer("COMMUNITIES");
const ratings = createSimpleReducer("RATINGS");
const sessionPlayers = createSimpleReducer("SESSION_PLAYERS");
const users = createSimpleReducer("USERS");

export default combineReducers({
  players,
  sessions,
  communities,
  ratings,
  sessionPlayers,
  users,
  ui
});
