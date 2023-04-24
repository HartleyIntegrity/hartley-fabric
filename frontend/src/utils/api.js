import axios from "axios";

const BASE_URL = "http://localhost:8080/api";
const apiUrl = 'http://localhost:8080/api';


export const getTenancyAgreements = async () => {
  const response = await axios.get(`${BASE_URL}/tenancy-agreements`);
  return response.data;
};

export async function createTenancyAgreement(tenancyAgreement) {
  const response = await fetch(`${apiUrl}/tenancy-agreements`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(tenancyAgreement),
  });
  const data = await response.json();
  
  if (response.status !== 201) {
    const error = new Error(data.message || 'Failed to create tenancy agreement');
    error.status = response.status;
    throw error;
  }

  const latestHashResponse = await fetch(`${apiUrl}/latest-hash`);
  const latestHashData = await latestHashResponse.json();
  const latestHash = latestHashData.hash;

  return { transaction: data, latestHash };
}






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
