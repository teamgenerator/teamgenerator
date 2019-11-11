import { combineReducers } from "redux";
import players from "./players";
import sessions from "./sessions";
import activeCommunity from "./activeCommunity";

export default combineReducers({
  players,
  sessions,
  activeCommunity
});
