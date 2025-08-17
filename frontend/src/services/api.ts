import { AuthStatus } from '../types/user';
import { getApiUrl } from '../config/app';

const API_BASE_URL = getApiUrl();

export const api = {
  // Check authentication status by calling /api/profile
  async getAuthStatus(): Promise<AuthStatus> {
    try {
      const response = await fetch(`${API_BASE_URL}/api/profile`, {
        credentials: 'include',
        redirect: 'manual',
      });
      
      if (response.status === 200) {
        const user = await response.json();
        return { authenticated: true, user };
      } else if (response.status === 302 || response.status === 307) {
        return { authenticated: false, message: 'Not authenticated' };
      } else {
        return { authenticated: false, message: `HTTP error! status: ${response.status}` };
      }
    } catch (error) {
      console.error('Failed to get auth status:', error);
      return { authenticated: false, message: 'Failed to check authentication' };
    }
  },

  // Login redirect
  login(): void {
    window.location.href = `${API_BASE_URL}/auth/login?return_url=${window.location.href}`;
  },

  // Logout
  logout(): void {
    window.location.href = `${API_BASE_URL}/auth/logout`;
  }
};
