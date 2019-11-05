import { combineReducers } from "redux";
import players from "./players";
import sessions from "./sessions";

export default combineReducers({
  players,
  sessions
});
