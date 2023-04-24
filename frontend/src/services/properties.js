const API_BASE_URL = 'http://localhost:8080';

export async function getProperties() {
  const response = await fetch(`${API_BASE_URL}/properties`);
  const data = await response.json();
  return data;
}
