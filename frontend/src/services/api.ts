import { AuthStatus, User } from '../types/user';
import { getApiUrl } from '../config/app';

const API_BASE_URL = getApiUrl();

export const api = {
  // Check authentication status
  async getAuthStatus(): Promise<AuthStatus> {
    try {
      const response = await fetch(`${API_BASE_URL}/auth/status`, {
        credentials: 'include', // Include cookies
      });
      
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      
      return await response.json();
    } catch (error) {
      console.error('Failed to get auth status:', error);
      return { authenticated: false, message: 'Failed to check authentication' };
    }
  },

  // Get user profile
  async getUserProfile(): Promise<User> {
    try {
      const response = await fetch(`${API_BASE_URL}/api/me`, {
        credentials: 'include', // Include cookies
      });
      
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      
      return await response.json();
    } catch (error) {
      console.error('Failed to get user profile:', error);
      throw error;
    }
  },

  // Login redirect
  login(): void {
    window.location.href = `${API_BASE_URL}/auth/login`;
  },

  // Logout
  logout(): void {
    window.location.href = `${API_BASE_URL}/auth/logout`;
  }
};
