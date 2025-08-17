import React from 'react';
import { getApiUrl, getEnvironmentInfo } from '../config/app';

const Home: React.FC = () => {
  const apiUrl = getApiUrl();
  const envInfo = getEnvironmentInfo();

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100">
      <div className="container mx-auto px-4 py-8">
        <div className="text-center">
          <h1 className="text-6xl font-bold text-gray-800 mb-4">
            Welcome to Vibed Traveller! ✈️
          </h1>
          <p className="text-xl text-gray-600 mb-8 max-w-2xl mx-auto">
            Your journey begins here. This is a simple React app with TypeScript and Tailwind CSS, 
            featuring Auth0 authentication and a beautiful profile page.
          </p>
          
          <div className="bg-white rounded-lg shadow-lg p-8 max-w-2xl mx-auto mb-8">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div className="text-center">
                <div className="text-4xl mb-4">🔐</div>
                <h3 className="text-xl font-semibold text-gray-800 mb-2">Authentication</h3>
                <p className="text-gray-600">
                  Secure login with Auth0 integration
                </p>
              </div>
              <div className="text-center">
                <div className="text-4xl mb-4">👤</div>
                <h3 className="text-xl font-semibold text-gray-800 mb-2">Profile Management</h3>
                <p className="text-gray-600">
                  View and manage your user profile
                </p>
              </div>
            </div>
          </div>

          <div className="space-y-4">
            <a
              href="/profile"
              className="inline-block bg-indigo-600 text-white px-8 py-3 rounded-lg hover:bg-indigo-700 transition-colors font-medium text-lg"
            >
              View Your Profile
            </a>
            <p className="text-gray-500 text-sm">
              Click the button above to see your authentication status and profile information
            </p>
          </div>
        </div>
      </div>

      {/* Footer with API URL info */}
      <footer className="bg-white border-t border-gray-200 py-6 mt-auto">
        <div className="container mx-auto px-4">
          <div className="text-center">
            <div className="bg-gray-50 rounded-lg p-4 max-w-md mx-auto">
              <h4 className="text-sm font-semibold text-gray-700 mb-2">🔧 Configuration Info</h4>
              <div className="text-xs text-gray-600 space-y-1">
                <p><span className="font-medium">API URL:</span></p>
                <p className="font-mono bg-gray-100 px-2 py-1 rounded break-all">
                  {apiUrl}
                </p>
                <p className="text-gray-500 mt-2">
                  {envInfo.description}
                </p>
              </div>
            </div>
            <p className="text-gray-500 text-sm mt-4">
              Vibed Traveller Frontend • Built with React, TypeScript & Tailwind CSS
            </p>
          </div>
        </div>
      </footer>
    </div>
  );
};

export default Home;
