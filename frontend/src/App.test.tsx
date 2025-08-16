import React from 'react';
import { render, screen } from '@testing-library/react';
import App from './App';

test('renders hello world message', () => {
  render(<App />);
  const helloElement = screen.getByText(/Hello World!/i);
  expect(helloElement).toBeInTheDocument();
});

test('renders welcome message', () => {
  render(<App />);
  const welcomeElement = screen.getByText(/Welcome to Vibed Traveller/i);
  expect(welcomeElement).toBeInTheDocument();
});
