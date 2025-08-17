import React, { useState, useEffect } from 'react';
import {AuthStatus, User} from '../types/user';
import { api } from '../services/api';

const Profile: React.FC = () => {
  const [authStatus, setAuthStatus] = useState<AuthStatus | null>(null);
  const [userProfile, setUserProfile] = useState<User | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    checkAuthStatus();
  }, []);

  const checkAuthStatus = async () => {
    try {
      setLoading(true);
      const status = await api.getAuthStatus();
      setAuthStatus(status);
      setUserProfile(status.user);
    } catch (err) {
      setError('Failed to check authentication status');
      console.error('Auth check failed:', err);
    } finally {
      setLoading(false);
    }
  };

  const handleLogin = () => {
    api.login();
  };

  const handleLogout = () => {
    api.logout();
  };

  const handleRefresh = () => {
    checkAuthStatus();
  };

  if (loading) {
    return (
      <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex items-center justify-center">
        <div className="text-center">
          <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-indigo-600 mx-auto mb-4"></div>
          <p className="text-gray-600">Checking authentication...</p>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex items-center justify-center">
        <div className="text-center">
          <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
            <p className="font-bold">Error</p>
            <p>{error}</p>
          </div>
          <button
            onClick={handleRefresh}
            className="bg-indigo-600 text-white px-4 py-2 rounded hover:bg-indigo-700 transition-colors"
          >
            Try Again
          </button>
        </div>
      </div>
    );
  }

  if (!authStatus?.authenticated) {
    return (
      <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex items-center justify-center">
        <div className="text-center">
          <div className="bg-white rounded-lg shadow-lg p-8 max-w-md mx-auto">
            <div className="text-4xl mb-4">🔒</div>
            <h2 className="text-2xl font-bold text-gray-800 mb-4">Not Authenticated</h2>
            <p className="text-gray-600 mb-6">
              Please log in to view your profile.
            </p>
            <button
              onClick={handleLogin}
              className="bg-indigo-600 text-white px-6 py-3 rounded-lg hover:bg-indigo-700 transition-colors font-medium"
            >
              Log In
            </button>
          </div>
        </div>
      </div>
    );
  }

  const user = authStatus.user!;

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100">
      <div className="container mx-auto px-4 py-8">
        {/* Header */}
        <div className="text-center mb-8">
          <h1 className="text-4xl font-bold text-gray-800 mb-2">Profile</h1>
          <p className="text-gray-600">Welcome back, {user.username || user.email}!</p>
        </div>

        {/* Profile Card */}
        <div className="max-w-2xl mx-auto">
          <div className="bg-white rounded-lg shadow-lg overflow-hidden">
            {/* Profile Header */}
            <div className="bg-gradient-to-r from-indigo-500 to-purple-600 px-6 py-8 text-white">
              <div className="flex items-center space-x-4">
                <div className="w-20 h-20 bg-white bg-opacity-20 rounded-full flex items-center justify-center text-3xl">
                  {user.metadata?.picture ? (
                    <img 
                      src={user.metadata.picture} 
                      alt="Profile" 
                      className="w-20 h-20 rounded-full object-cover"
                    />
                  ) : (
                    '👤'
                  )}
                </div>
                <div>
                  <h2 className="text-2xl font-bold">
                    {user.metadata?.name || user.username || 'User'}
                  </h2>
                  <p className="text-indigo-100">{user.email}</p>
                </div>
              </div>
            </div>

            {/* Profile Details */}
            <div className="p-6">
              <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                {/* Basic Info */}
                <div>
                  <h3 className="text-lg font-semibold text-gray-800 mb-4">Basic Information</h3>
                  <div className="space-y-3">
                    <div>
                      <label className="text-sm font-medium text-gray-500">User ID</label>
                      <p className="text-gray-800 font-mono text-sm">{user.id}</p>
                    </div>
                    <div>
                      <label className="text-sm font-medium text-gray-500">Username</label>
                      <p className="text-gray-800">{user.username || 'Not set'}</p>
                    </div>
                    <div>
                      <label className="text-sm font-medium text-gray-500">Email</label>
                      <p className="text-gray-800">{user.email}</p>
                    </div>
                  </div>
                </div>

                {/* Additional Info */}
                <div>
                  <h3 className="text-lg font-semibold text-gray-800 mb-4">Additional Information</h3>
                  <div className="space-y-3">
                    {user.metadata?.given_name && (
                      <div>
                        <label className="text-sm font-medium text-gray-500">First Name</label>
                        <p className="text-gray-800">{user.metadata.given_name}</p>
                      </div>
                    )}
                    {user.metadata?.family_name && (
                      <div>
                        <label className="text-sm font-medium text-gray-500">Last Name</label>
                        <p className="text-gray-800">{user.metadata.family_name}</p>
                      </div>
                    )}
                    {user.metadata?.updated_at && (
                      <div>
                        <label className="text-sm font-medium text-gray-500">Last Updated</label>
                        <p className="text-gray-800">{new Date(user.metadata.updated_at).toLocaleDateString()}</p>
                      </div>
                    )}
                  </div>
                </div>
              </div>

              {/* Actions */}
              <div className="mt-8 pt-6 border-t border-gray-200">
                <div className="flex flex-col sm:flex-row gap-3 justify-center">
                  <button
                    onClick={handleRefresh}
                    className="bg-gray-600 text-white px-6 py-2 rounded-lg hover:bg-gray-700 transition-colors"
                  >
                    Refresh Profile
                  </button>
                  <button
                    onClick={handleLogout}
                    className="bg-red-600 text-white px-6 py-2 rounded-lg hover:bg-red-700 transition-colors"
                  >
                    Log Out
                  </button>
                </div>
              </div>
            </div>
          </div>

          {/* Debug Info (only in development) */}
          {process.env.NODE_ENV === 'development' && (
            <div className="mt-6 bg-gray-100 rounded-lg p-4">
              <h4 className="text-sm font-semibold text-gray-700 mb-2">Debug Information</h4>
              <pre className="text-xs text-gray-600 overflow-auto">
                {JSON.stringify({ authStatus, userProfile }, null, 2)}
              </pre>
            </div>
          )}
        </div>
      </div>
    </div>
  );
};

export default Profile;
