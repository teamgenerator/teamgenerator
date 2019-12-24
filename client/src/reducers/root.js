import { combineReducers } from "redux";
import players from "./players";
import sessions from "./sessions";
import communities from "./communities";
import ui from "./ui";

export default combineReducers({
  players,
  sessions,
  communities,
  ui
});
