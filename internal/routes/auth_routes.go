package routes

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"vibed-traveller/internal/config"

	"github.com/gin-gonic/gin"
)

// SetupAuthRoutes configures authenticated routes
func SetupAuthRoutes(router *gin.Engine, cfg *config.Config) {
	// Only setup routes if Auth0 is properly configured
	if !cfg.IsAuth0Configured() {
		panic("Auth configuration is not configured")
	}

	// Public auth routes (no authentication required)
	auth := router.Group("/auth")
	{
		// Login endpoint - redirects to Auth0
		auth.GET("/login", func(c *gin.Context) {
			// Get the return URL from query parameter, default to home page
			returnURL := c.Query("return_url")
			if returnURL == "" {
				returnURL = "/"
			}

			loginURL := config.GenerateAuth0LoginURL(cfg, returnURL)
			c.Redirect(http.StatusTemporaryRedirect, loginURL)
		})

		// Callback endpoint - handles Auth0 response
		auth.GET("/callback", func(c *gin.Context) {
			handleAuth0Callback(c, cfg)
		})

		// Logout endpoint
		auth.GET("/logout", func(c *gin.Context) {
			// Clear the auth token cookie
			config.ClearAuthTokenCookie(c)

			// Redirect to Auth0 logout
			logoutURL := fmt.Sprintf("%s/v2/logout?client_id=%s&returnTo=%s",
				cfg.GetAuth0IssuerURL(),
				cfg.GetAuth0ClientID(),
				url.QueryEscape(cfg.GetBaseURL()),
			)
			c.Redirect(http.StatusTemporaryRedirect, logoutURL)
		})
	}

	// Protected routes group
	protected := router.Group("/api")
	protected.Use(config.AuthMiddleware(cfg))
	{
		// User profile endpoint
		protected.GET("/profile", getUserProfile)

		// Test endpoint to show current user
		protected.GET("/me", func(c *gin.Context) {
			user := config.GetUserFromContext(c)
			c.JSON(http.StatusOK, gin.H{
				"message": "You are authenticated!",
				"user":    user,
			})
		})
	}
}

// handleAuth0Callback handles the Auth0 callback response
func handleAuth0Callback(c *gin.Context, cfg *config.Config) {
	// Check for errors
	if err := c.Query("error"); err != "" {
		errorDescription := c.Query("error_description")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":             err,
			"error_description": errorDescription,
		})
		return
	}

	// Get the authorization code
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization code not provided"})
		return
	}

	// Get the return URL from state parameter
	returnURL := c.Query("state")
	if returnURL == "" {
		returnURL = "/"
	}

	slog.InfoContext(c.Request.Context(), "Callback called", slog.String("code", code))

	// Exchange the authorization code for an access token
	tokenResponse, err := config.ExchangeCodeForToken(cfg, code)
	if err != nil {
		slog.ErrorContext(c.Request.Context(), "Failed to exchange code for token", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange code for token", "details": err.Error()})
		return
	}

	slog.InfoContext(c.Request.Context(), "ExchangeCodeForToken", slog.Any("token", tokenResponse))

	// Extract the access token
	accessToken, ok := tokenResponse["access_token"].(string)
	if !ok {
		slog.ErrorContext(c.Request.Context(), "Access token not found in response")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Access token not found in response"})
		return
	}

	config.SetAuthTokenCookie(c, accessToken)
	slog.InfoContext(c.Request.Context(), "Token ready and cookie set")

	// Redirect to the return URL
	c.Redirect(http.StatusTemporaryRedirect, returnURL)
}

// getUserProfile returns the current user's profile
func getUserProfile(c *gin.Context) {
	user := config.GetUserFromContext(c)
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	c.JSON(http.StatusOK, user)
}
