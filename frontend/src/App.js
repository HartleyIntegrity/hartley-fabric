import React, { useState } from 'react';
import { BrowserRouter as Router, Route, Navigate, Routes } from 'react-router-dom';
import SignIn from './components/SignIn';
import Properties from './components/Properties';

const App = () => {
  const [token, setToken] = useState('');

  const handleSignIn = (newToken) => {
    setToken(newToken);
  };

  return (
    <Router>
      <Routes>
        <Route path="/" element={token ? <Navigate to="/properties" replace /> : <SignIn onSignIn={handleSignIn} />} />
        <Route path="/properties" element={!token ? <Navigate to="/" replace /> : <Properties />} />
      </Routes>
    </Router>
  );
};

export default App;
