import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import PropertiesPage from './pages/PropertiesPage';
import ContractsPage from './pages/ContractsPage';
import "./index.css"

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<h1>Welcome to Hartley-Fabric</h1>} />
        <Route path="/properties" element={<PropertiesPage />} />
        <Route path="/contracts" element={<ContractsPage />} />
      </Routes>
    </Router>
  );
}

export default App;
