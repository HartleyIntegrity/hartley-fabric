import axios from "axios";

const BASE_URL = "http://localhost:8080/api";

export const getTenancyAgreements = async () => {
  const response = await axios.get(`${BASE_URL}/tenancy-agreements`);
  return response.data;
};

export const createTenancyAgreement = async (agreement) => {
  const response = await axios.post(`${BASE_URL}/tenancy-agreements`, agreement);
  return response.data;
};

export const getTenancyAgreementById = async (id) => {
  const response = await axios.get(`${BASE_URL}/tenancy-agreements/${id}`);
  return response.data;
};

export const updateTenancyAgreement = async (id, agreement) => {
  const response = await axios.put(`${BASE_URL}/tenancy-agreements/${id}`, agreement);
  return response.data;
};

export const deleteTenancyAgreement = async (id) => {
  const response = await axios.delete(`${BASE_URL}/tenancy-agreements/${id}`);
  return response.data;
};

export const getProperty = async (id) => {
  const response = await axios.get(`${BASE_URL}/properties/${id}`);
  return response.data;
};
