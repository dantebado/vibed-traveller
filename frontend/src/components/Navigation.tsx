import React from 'react';
import { api } from '../services/api';

const Navigation: React.FC = () => {
  const handleLogout = () => {
    api.logout();
  };

  return (
    <nav className="bg-white shadow-lg">
      <div className="container mx-auto px-4">
        <div className="flex justify-between items-center h-16">
          {/* Logo/Brand */}
          <div className="flex items-center space-x-2">
            <div className="text-2xl">✈️</div>
            <span className="text-xl font-bold text-gray-800">Vibed Traveller</span>
          </div>

          {/* Navigation Links */}
          <div className="hidden md:flex items-center space-x-6">
            <a 
              href="/" 
              className="text-gray-600 hover:text-indigo-600 transition-colors font-medium"
            >
              Home
            </a>
            <a 
              href="/profile" 
              className="text-gray-600 hover:text-indigo-600 transition-colors font-medium"
            >
              Profile
            </a>
          </div>

          {/* Actions */}
          <div className="flex items-center space-x-4">
            <button
              onClick={handleLogout}
              className="bg-red-600 text-white px-4 py-2 rounded-lg hover:bg-red-700 transition-colors text-sm font-medium"
            >
              Log Out
            </button>
          </div>
        </div>
      </div>
    </nav>
  );
};

export default Navigation;
