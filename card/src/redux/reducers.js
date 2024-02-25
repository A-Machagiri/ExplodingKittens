import { combineReducers } from "redux";
import * as actionTypes from "./actions";

const initialState = {
  elements: ["cat", "cat", "shuffle", "bomb", "defuse"],
  shuffledElements: [],
  defuseActivated: false,
  gameWon: false,
  lostGame: false,
};

const gameReducer = (state = initialState, action) => {
  switch (action.type) {
    case actionTypes.INITIALIZE_ELEMENTS:
      return {
        ...state,
        shuffledElements: action.elements,
      };
    case actionTypes.GAME_WON:
      return {
        ...state,
        gameWon: true,
      };
    case actionTypes.GAME_LOST:
      return {
        ...state,
        lostGame: true,
      };
      case actionTypes.CLICK_ELEMENT:
  if (action.element === "cat") {
    const catIndex = state.shuffledElements.indexOf("cat");
    if (catIndex !== -1) {
      const newElements = [
        ...state.shuffledElements.slice(0, catIndex),
        ...state.shuffledElements.slice(catIndex + 1),
      ];
      return {
        ...state,
        shuffledElements: newElements,
      };
    }
  } else if (action.element === "shuffle") {
    const shuffledArray = [...state.elements];
    for (let i = shuffledArray.length - 1; i > 0; i--) {
      const j = Math.floor(Math.random() * (i + 1));
      [shuffledArray[i], shuffledArray[j]] = [
        shuffledArray[j],
        shuffledArray[i],
      ];
    }
    return {
      ...state,
      shuffledElements: shuffledArray,
    };
  } else if (action.element === "defuse") {
    const indexOfDefuse = state.shuffledElements.indexOf("defuse");
    if (indexOfDefuse !== -1) {
      const newElements = [
        ...state.shuffledElements.slice(0, indexOfDefuse),
        ...state.shuffledElements.slice(indexOfDefuse + 1),
      ];
      return {
        ...state,
        defuseActivated: true,
        shuffledElements: newElements,
      };
    }
  } else if (action.element === "bomb" && !state.defuseActivated) {
    return {
      ...state,
      lostGame: true,
      shuffledElements: [],
    };
  } else if (action.element === "bomb" && state.defuseActivated) {
    const bombIndex = state.shuffledElements.indexOf("bomb");
    if (bombIndex !== -1) {
      const newElements = [
        ...state.shuffledElements.slice(0, bombIndex),
        ...state.shuffledElements.slice(bombIndex + 1),
      ];
      return {
        ...state,
        shuffledElements: newElements,
      };
    }
  }
  return state;
    default:
      return state;
  }
};

const rootReducer = combineReducers({
  game: gameReducer,
});

export default rootReducer;
