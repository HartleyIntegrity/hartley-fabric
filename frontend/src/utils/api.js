import axios from "axios";
import jwt_decode from "jwt-decode";

const BASE_URL = "http://localhost:3001/api";

const getDecodedToken = () => {
  const token = localStorage.getItem("token");
  if (token) {
    return jwt_decode(token);
  }
  return null;
};

const api = {
  get: async (url) => {
    const response = await axios.get(`${BASE_URL}${url}`);
    return response;
  },

  post: async (url, data) => {
    const response = await axios.post(`${BASE_URL}${url}`, data, {
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
      },
    });
    return response;
  },

  delete: async (url) => {
    const decodedToken = getDecodedToken();
    if (decodedToken && decodedToken.role === "admin") {
      const response = await axios.delete(`${BASE_URL}${url}`, {
        headers: {
          Authorization: `Bearer ${localStorage.getItem("token")}`,
        },
      });
      return response;
    }
    throw new Error("Unauthorized");
  },

  put: async (url, data) => {
    const decodedToken = getDecodedToken();
    if (decodedToken && decodedToken.role === "admin") {
      const response = await axios.put(`${BASE_URL}${url}`, data, {
        headers: {
          Authorization: `Bearer ${localStorage.getItem("token")}`,
        },
      });
      return response;
    }
    throw new Error("Unauthorized");
  },
};

export default api;
