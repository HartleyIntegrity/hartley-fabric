import axios from 'axios';

const API_BASE_URL = 'http://localhost:8000';

export const getTransactions = async () => {
  const response = await axios.get(`${API_BASE_URL}/transactions`);
  return response.data;
};

export const createTransaction = async (transaction) => {
  const response = await axios.post(`${API_BASE_URL}/transactions`, transaction);
  return response.data;
};

// ...

export const signIn = async (username, password) => {
    const response = await axios.post(`${API_BASE_URL}/signin`, { username, password });
    return response.data
};  
