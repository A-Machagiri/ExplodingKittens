import axios from 'axios';
const API_URL = 'http://localhost:8000';

export const startGameApi = (username) => {
  return axios.post(`${API_URL}/start`, { username });
};

export const getLeaderboardAPI = () => {
  return axios.get(`${API_URL}/leaderboard`);
};
