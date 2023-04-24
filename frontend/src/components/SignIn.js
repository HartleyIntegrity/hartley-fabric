import React, { useState } from 'react';
import { signIn } from '../api';
import styled from '@emotion/styled';
import { TextField, Button } from '@mui/material';

const SignInWrapper = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 24px;
`;

const SignInForm = styled.form`
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 16px;
`;

const SignIn = ({ onSignIn }) => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const { token } = await signIn(username, password);
      onSignIn(token);
    } catch (err) {
      setError('Invalid username or password');
    }
  };

  return (
    <SignInWrapper>
    <h1>Sign In</h1>
    {error && <p>{error}</p>}
    <SignInForm onSubmit={handleSubmit}>
      <TextField
        label="Username"
        value={username}
        onChange={(e) => setUsername(e.target.value)}
      />
      <TextField
        label="Password"
        type="password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
      />
      <Button type="submit" variant="contained" color="primary">
        Sign In
      </Button>
    </SignInForm>
  </SignInWrapper>
);
};

export default SignIn;

