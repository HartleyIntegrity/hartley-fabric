const API_BASE_URL = 'http://localhost:8080';

export async function getContracts() {
  const response = await fetch(`${API_BASE_URL}/contracts`);
  const data = await response.json();
  return data;
}

export async function createContract(contract) {
  const response = await fetch(`${API_BASE_URL}/contracts`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(contract),
  });
  const data = await response.json();
  return data;
}

export async function deleteContract(id) {
  await fetch(`${API_BASE_URL}/contracts/${id}`, {
    method: 'DELETE',
  });
}
