import { configureStore } from 'redux-starter-kit';

import rootReducer from '../reducers/root';


const preloadedState = {
  players: {
    '0001': {
      id: '0001',
      name: 'Michael Englo',
      ratings: 2,
    },
    '0002': {
      id: '0002',
      name: 'Nico Alimin',
      ratings: 8,
    },
    '0003': {
      id: '0003',
      name: 'Timothy Situmeang',
      ratings: 3,
    },
    '0004': {
      id: '0004',
      name: 'Kiky Tangerine',
      ratings: 7,
    },
    '0005': {
      id: '0005',
      name: 'Ivan Gunawan',
      ratings: 1,
    },
    '0006': {
      id: '0006',
      name: 'Victor',
      ratings: 9
    },
    '0007': {
      id: '0007',
      name: 'Kynan Bangun',
      ratings: 2,
    },
    '0008': {
      id: '0008',
      name: 'Dito Dakota',
      ratings: 6,
    },
  },
};

export default configureStore({ 
  reducer: rootReducer,
});