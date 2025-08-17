# Auth0 Setup Guide for Vibed Traveller

This guide will help you set up Auth0 authentication for your Go backend application.

## Prerequisites

- An Auth0 account (free tier available)
- Go 1.21+ installed
- Your Go backend running

## Step 1: Create an Auth0 Application

1. **Sign up/Login to Auth0**
   - Go to [auth0.com](https://auth0.com)
   - Create a free account or sign in

2. **Create a New Application**
   - In your Auth0 dashboard, go to "Applications" → "Applications"
   - Click "Create Application"
   - Choose "Machine to Machine Applications" (for API authentication)
   - Give it a name like "Vibed Traveller API"

3. **Configure the Application**
   - Set the **Allowed Callback URLs** to: `http://localhost:8080/callback`
   - Set the **Allowed Logout URLs** to: `http://localhost:8080`
   - Set the **Allowed Web Origins** to: `http://localhost:8080`
   - Save the changes

## Step 2: Create an API

1. **Go to APIs Section**
   - In your Auth0 dashboard, go to "Applications" → "APIs"
   - Click "Create API"

2. **Configure the API**
   - **Name**: `Vibed Traveller API`
   - **Identifier**: `https://vibed-traveller-api`
   - **Signing Algorithm**: `RS256`
   - Click "Create"

## Step 3: Get Your Configuration Values

1. **Domain**: Found in your Auth0 dashboard under "Settings" → "Domain"
   - Format: `your-tenant.auth0.com`

2. **Client ID**: Found in your application settings
   - Copy the "Client ID" value

3. **Client Secret**: Found in your application settings
   - Copy the "Client Secret" value

4. **Audience**: Use the API identifier you created
   - Format: `https://vibed-traveller-api`

5. **Issuer URL**: Your Auth0 domain with `https://` prefix
   - Format: `https://your-tenant.auth0.com/`

## Step 4: Configure Environment Variables

1. **Copy the example file**:
   ```bash
   cp env.example .env
   ```

2. **Edit the `.env` file** with your actual values:
   ```bash
   # Server Configuration
   PORT=8080
   LOG_LEVEL=info
   
   # Auth0 Configuration
   AUTH0_DOMAIN=your-tenant.auth0.com
   AUTH0_AUDIENCE=https://vibed-traveller-api
   AUTH0_ISSUER_URL=https://your-tenant.auth0.com
   AUTH0_CLIENT_ID=your-client-id
   AUTH0_CLIENT_SECRET=your-client-secret
   ```

## Step 5: Test the Setup

1. **Start your Go backend**:
   ```bash
   go run cmd/main.go
   ```

2. **Check the logs** - you should see:
   ```
   Auth0 routes configured successfully
   ```

3. **Test the protected endpoint**:
   ```bash
   # This should return 401 Unauthorized
   curl http://localhost:8080/api/profile
   ```

## Step 6: Frontend Integration (Optional)

To integrate with your React frontend, you'll need to:

1. **Install Auth0 React SDK**:
   ```bash
   cd frontend
   npm install @auth0/auth0-react
   ```

2. **Configure Auth0 Provider** in your React app
3. **Use the `useAuth0` hook** to get user information
4. **Include the JWT token** in API requests

## API Endpoints

### Protected Endpoints (Require Authentication)
- `GET /api/profile` - Get user profile

## Testing with Postman/Insomnia

1. **Get a JWT token** from Auth0 (you can use the Auth0 playground)
2. **Add Authorization header**:
   ```
   Authorization: Bearer YOUR_JWT_TOKEN
   ```
3. **Test the protected endpoint**

## Troubleshooting

### Common Issues

1. **"Auth0 not configured" warning**
   - Check your `.env` file has all required variables
   - Ensure no extra spaces or quotes around values

2. **"Invalid token" errors**
   - Verify your Auth0 domain and audience are correct
   - Check that the token hasn't expired
   - Ensure you're using the correct API identifier

3. **"Failed to set up the jwt validator"**
   - Check your Auth0 domain format
   - Verify your issuer URL is correct

### Debug Mode

To see more detailed logs, set:
```bash
LOG_LEVEL=debug
```

## Security Notes

- Never commit your `.env` file to version control
- Keep your Auth0 client secret secure
- Use HTTPS in production
- Regularly rotate your Auth0 client secrets
- Monitor your Auth0 dashboard for suspicious activity

## Next Steps

1. **Add database integration** for user data persistence
2. **Implement user registration flows**
3. **Add role-based access control**
4. **Set up social login providers** (Google, Facebook, etc.)
5. **Add multi-factor authentication**
6. **Implement password reset flows**

## Support

- [Auth0 Documentation](https://auth0.com/docs)
- [Auth0 Go SDK](https://github.com/auth0/go-jwt-middleware)
- [Auth0 Community](https://community.auth0.com/)
