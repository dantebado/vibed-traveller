// Application configuration
export const config = {
  // API configuration
  api: {
    baseUrl: process.env.REACT_APP_API_URL || 'http://localhost:8080',
  },
  
  // Environment information
  env: {
    isDevelopment: process.env.NODE_ENV === 'development',
    isProduction: process.env.NODE_ENV === 'production',
  },
  
  // App metadata
  app: {
    name: 'Vibed Traveller',
    version: '1.0.0',
  },
} as const;

// Helper function to get API URL
export const getApiUrl = (): string => config.api.baseUrl;

// Helper function to get environment info
export const getEnvironmentInfo = () => ({
  isDevelopment: config.env.isDevelopment,
  isProduction: config.env.isProduction,
  mode: config.env.isDevelopment ? 'development' : 'production',
  description: config.env.isDevelopment 
    ? 'Development mode - .env file loaded' 
    : 'Production mode - environment variable baked in at build time',
});
