import React from 'react';
import Navigation from './components/Navigation';
import Home from './components/Home';
import Profile from './components/Profile';

function App() {
  // Simple routing based on current pathname
  const getCurrentPage = () => {
    const path = window.location.pathname;
    
    switch (path) {
      case '/profile':
        return <Profile />;
      case '/':
      default:
        return <Home />;
    }
  };

  return (
    <div className="App">
      <Navigation />
      {getCurrentPage()}
    </div>
  );
}

export default App;
