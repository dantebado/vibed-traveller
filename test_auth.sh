#!/bin/bash

echo "Testing Vibed Traveller Authentication Flow"
echo "=========================================="
echo ""

# Read base URL from .env file, default to localhost:8080
if [ -f .env ]; then
    BASE_URL=$(grep BASE_URL .env | cut -d'=' -f2)
    if [ -z "$BASE_URL" ]; then
        BASE_URL="http://localhost:8080"
    fi
else
    BASE_URL="http://localhost:8080"
fi

echo "1. Testing protected endpoint without authentication (should redirect to login):"
echo "   GET $BASE_URL/api/me"
echo "   Expected: Redirect to Auth0 login"
echo ""

echo "2. Testing login page:"
echo "   GET $BASE_URL/auth/login-page"
echo "   Expected: Simple HTML login page"
echo ""

echo "3. Testing login redirect:"
echo "   GET $BASE_URL/auth/login"
echo "   Expected: Redirect to Auth0 with proper parameters"
echo ""

echo "4. Testing logout:"
echo "   GET $BASE_URL/auth/logout"
echo "   Expected: Redirect to Auth0 logout"
echo ""

echo "5. Testing callback endpoint:"
echo "   GET $BASE_URL/auth/callback?state=/api/me"
echo "   Expected: Redirect back to /api/me"
echo ""

echo ""
echo "To test the redirect flow:"
echo "1. Open your browser and go to: $BASE_URL/api/me"
echo "2. You should be redirected to Auth0 login"
echo "3. After login, you'll be redirected back to /api/me"
echo "4. The /api/me endpoint should show your user information"
echo ""

echo "Note: Make sure you have configured your .env file with proper Auth0 credentials!"
echo "Current .env values:"
if [ -f .env ]; then
    echo "BASE_URL: $BASE_URL"
    echo "AUTH0_DOMAIN: $(grep AUTH0_DOMAIN .env | cut -d'=' -f2)"
    echo "AUTH0_CLIENT_ID: $(grep AUTH0_CLIENT_ID .env | cut -d'=' -f2)"
    echo "AUTH0_AUDIENCE: $(grep AUTH0_AUDIENCE .env | cut -d'=' -f2)"
else
    echo ".env file not found!"
fi
