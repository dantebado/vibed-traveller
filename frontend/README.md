# Vibed Traveller Frontend

A React application with TypeScript and Tailwind CSS that provides a beautiful user interface for the Vibed Traveller backend.

## Features

- **Authentication**: Integrated with Auth0 for secure user login/logout
- **Profile Page**: Beautiful profile display with user information from Auth0
- **Responsive Design**: Mobile-friendly interface using Tailwind CSS
- **TypeScript**: Full type safety for better development experience

## Getting Started

### Prerequisites

- Node.js (v14 or higher)
- npm or yarn
- Backend server running (see main README)

### Installation

1. Install dependencies:
   ```bash
   npm install
   ```

2. Configure environment variables:
   ```bash
   cp env.example .env
   ```
   
   Update the `.env` file with your backend API URL:
   ```
   REACT_APP_API_URL=http://localhost:8080
   ```

3. Start the development server:
   ```bash
   npm start
   ```

   The app will open at [http://localhost:3000](http://localhost:3000)

## Available Scripts

- `npm start` - Start development server
- `npm run build` - Build for production
- `npm test` - Run tests
- `npm run eject` - Eject from Create React App (not recommended)

## Project Structure

```
src/
├── components/          # React components
│   ├── Home.tsx        # Home page component
│   ├── Profile.tsx     # User profile component
│   └── Navigation.tsx  # Navigation header
├── services/           # API services
│   └── api.ts         # Backend API communication
├── types/              # TypeScript type definitions
│   └── user.ts        # User and authentication types
├── App.tsx            # Main application component
└── index.tsx          # Application entry point
```

## Authentication Flow

1. **Home Page**: Users see a welcome message and can navigate to their profile
2. **Profile Page**: 
   - If not authenticated: Shows login button
   - If authenticated: Displays user information from Auth0
3. **Navigation**: Header with logout functionality

## API Integration

The frontend communicates with the backend through the following endpoints:

- `GET /auth/status` - Check authentication status
- `GET /api/me` - Get user profile information
- `GET /auth/login` - Redirect to Auth0 login
- `GET /auth/logout` - Logout and redirect to Auth0

## Styling

Built with Tailwind CSS for:
- Responsive design
- Beautiful gradients and shadows
- Consistent spacing and typography
- Hover effects and transitions

## Development

The app uses simple client-side routing based on the current pathname. For production, consider using React Router for more robust routing.

## Troubleshooting

- **Build errors**: Make sure all dependencies are installed with `npm install`
- **API connection**: Verify the backend is running and `REACT_APP_API_URL` is correct
- **Authentication issues**: Check that Auth0 is properly configured in the backend
