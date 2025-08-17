package config

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// Cookie configuration constants
const (
	// AuthTokenCookieName is the name of the cookie that stores the authentication token
	AuthTokenCookieName = "auth_token"

	// AuthTokenCookieMaxAge is the maximum age of the auth token cookie in seconds (1 hour)
	AuthTokenCookieMaxAge = 3600

	// AuthTokenCookiePath is the path where the auth token cookie is available
	AuthTokenCookiePath = "/"

	// AuthTokenCookieSecure determines if the cookie should only be sent over HTTPS
	// Set to false for development, true for production
	AuthTokenCookieSecure = false
)

// Auth0 path constants
const (
	// Auth0AuthorizePath is the path for the authorization endpoint
	Auth0AuthorizePath = "authorize"

	// Auth0TokenPath is the path for the token endpoint
	Auth0TokenPath = "oauth/token"

	// Auth0LogoutPath is the path for the logout endpoint
	Auth0LogoutPath = "v2/logout"
)

// Auth0 parameter constants
const (
	// Auth0ResponseTypeCode is the response type for authorization code flow
	Auth0ResponseTypeCode = "code"

	// Auth0GrantTypeAuthorizationCode is the grant type for token exchange
	Auth0GrantTypeAuthorizationCode = "authorization_code"

	// Auth0ScopeOpenID is the OpenID scope
	Auth0ScopeOpenID = "openid"

	// Auth0ScopeProfile is the profile scope
	Auth0ScopeProfile = "profile"

	// Auth0ScopeEmail is the email scope
	Auth0ScopeEmail = "email"
)

type Auth0UserInfo struct {
	Sub           string     `json:"sub"`
	GivenName     string     `json:"given_name"`
	FamilyName    string     `json:"family_name"`
	Nickname      string     `json:"nickname"`
	Name          string     `json:"name"`
	Picture       string     `json:"picture"`
	UpdatedAt     *time.Time `json:"updated_at"`
	Email         string     `json:"email"`
	EmailVerified bool       `json:"email_verified"`
}

// SetAuthTokenCookie sets the authentication token cookie with the configured settings
func SetAuthTokenCookie(c *gin.Context, token string) {
	c.SetCookie(
		AuthTokenCookieName,
		token,
		AuthTokenCookieMaxAge,
		AuthTokenCookiePath,
		"",
		AuthTokenCookieSecure,
		false,
	)
}

// ClearAuthTokenCookie clears the authentication token cookie
func ClearAuthTokenCookie(c *gin.Context) {
	c.SetCookie(
		AuthTokenCookieName,
		"",
		-1, // Delete immediately
		AuthTokenCookiePath,
		"",
		AuthTokenCookieSecure,
		true, // httpOnly
	)
}

// GetAuthTokenFromCookie retrieves the authentication token from the cookie
func GetAuthTokenFromCookie(c *gin.Context) (string, error) {
	token, err := c.Cookie(AuthTokenCookieName)
	if err != nil {
		return "", err
	}
	if token == "" {
		return "", fmt.Errorf("auth token cookie is empty")
	}
	return token, nil
}

// User represents an authenticated user
type User struct {
	ID       string            `json:"id"`
	Email    string            `json:"email"`
	Username string            `json:"username"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

// AuthMiddleware creates a Gin middleware for JWT authentication
func AuthMiddleware(config *Config) gin.HandlerFunc {
	// Validate Auth0 configuration before creating middleware
	if err := validateAuth0Config(config); err != nil {
		panic(err.Error())
	}

	return func(c *gin.Context) {
		// Extract token from Authorization header or cookie
		var token string

		loginURL := GenerateAuth0LoginURL(config, c.Request.URL.String())

		// First try to get token from Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
			slog.InfoContext(c.Request.Context(), "Extracted token from Authorization header")
			token = strings.TrimPrefix(authHeader, "Bearer ")
		}
		if cookieToken, err := GetAuthTokenFromCookie(c); err == nil {
			slog.InfoContext(c.Request.Context(), "Extracted token from cookie")
			token = cookieToken
		}

		if token == "" {
			slog.InfoContext(c.Request.Context(), "No token found in header or cookie, redirecting to login")
			c.Redirect(http.StatusTemporaryRedirect, loginURL)
			c.Abort()
			return
		}

		// Validate JWT expiration
		err := ValidateJWT(token, config)
		if err != nil {
			slog.ErrorContext(c.Request.Context(), "Invalid token", slog.Any("error", err))
			// If token is invalid or expired, redirect to login
			c.Redirect(http.StatusTemporaryRedirect, loginURL)
			c.Abort()
			return
		}

		// Extract user information from the token
		user, err := ExtractUserFromToken(token, config)
		if err != nil {
			slog.ErrorContext(c.Request.Context(), "Failed to extract user info from token", slog.Any("error", err))
			c.Redirect(http.StatusTemporaryRedirect, loginURL)
			c.Abort()
			return
		}

		slog.InfoContext(c.Request.Context(), "User authenticated successfully", slog.String("user_id", user.ID))

		// Store user in context
		c.Set("user", user)
		c.Next()
	}
}

func ValidateJWT(token string, config *Config) error {
	// Parse JWT and validate expiration
	// Split the token into parts (header.payload.signature)
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return fmt.Errorf("invalid JWT format: expected 3 parts, got %d", len(parts))
	}

	// Decode the payload (second part)
	payload := parts[1]
	// Add padding if needed for base64 decoding
	if len(payload)%4 != 0 {
		payload += strings.Repeat("=", 4-len(payload)%4)
	}

	// Decode base64
	decodedPayload, err := base64.URLEncoding.DecodeString(payload)
	if err != nil {
		return fmt.Errorf("failed to decode JWT payload: %v", err)
	}

	// Parse the JSON payload
	var claims struct {
		Exp int64 `json:"exp"`
		Iat int64 `json:"iat"`
	}
	if err := json.Unmarshal(decodedPayload, &claims); err != nil {
		return fmt.Errorf("failed to parse JWT claims: %v", err)
	}

	// Check if token is expired
	now := time.Now().Unix()
	if claims.Exp > 0 && now > claims.Exp {
		return fmt.Errorf("JWT token is expired: exp=%d, now=%d", claims.Exp, now)
	}

	// Check if token is issued in the future (clock skew tolerance)
	if claims.Iat > 0 && now < claims.Iat-300 { // 5 minutes tolerance
		return fmt.Errorf("JWT token issued in the future: iat=%d, now=%d", claims.Iat, now)
	}

	return nil
}

// GenerateAuth0LoginURL generates the Auth0 login URL with the current page as the return URL
func GenerateAuth0LoginURL(config *Config, returnURL string) string {
	// Encode the return URL
	encodedReturnURL := url.QueryEscape(returnURL)

	// Build the Auth0 authorize URL
	authorizeURL := buildAuth0URL(config.GetAuth0IssuerURL(), Auth0AuthorizePath)

	// Build the Auth0 login URL using authorization code flow
	loginURL := fmt.Sprintf("%s?response_type=%s&client_id=%s&redirect_uri=%s&scope=%s%%20%s%%20%s&audience=%s",
		authorizeURL,
		Auth0ResponseTypeCode,
		config.GetAuth0ClientID(),
		url.QueryEscape(buildCallbackURL(config)),
		Auth0ScopeOpenID,
		Auth0ScopeProfile,
		Auth0ScopeEmail,
		url.QueryEscape(config.GetAuth0Audience()),
	)

	// Add the return URL as a state parameter
	loginURL += fmt.Sprintf("&state=%s", encodedReturnURL)

	return loginURL
}

// ExchangeCodeForToken exchanges an authorization code for an access token
func ExchangeCodeForToken(config *Config, code string) (map[string]interface{}, error) {
	// Prepare the token exchange request
	data := url.Values{}
	data.Set("grant_type", Auth0GrantTypeAuthorizationCode)
	data.Set("client_id", config.GetAuth0ClientID())
	data.Set("client_secret", config.GetAuth0ClientSecret())
	data.Set("code", code)
	data.Set("redirect_uri", buildCallbackURL(config))

	// Build the Auth0 token URL
	tokenURL := buildAuth0URL(config.GetAuth0IssuerURL(), Auth0TokenPath)

	req, err := http.NewRequest("POST", tokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create token request: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make token request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			slog.Error("failed to close response body: %v", slog.Any("err", err))
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("token exchange failed with status %d: %s", resp.StatusCode, string(body))
	}

	// Parse the response
	var tokenResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return nil, fmt.Errorf("failed to decode token response: %v", err)
	}

	return tokenResponse, nil
}

// ExtractUserFromToken extracts user information from a validated JWT token
func ExtractUserFromToken(accessToken string, config *Config) (*User, error) {
	if accessToken == "" {
		return nil, fmt.Errorf("no access token available")
	}

	// Call Auth0 userinfo endpoint to get user profile
	userinfoURL := fmt.Sprintf("%s/userinfo", config.GetAuth0IssuerURL())

	req, err := http.NewRequest("GET", userinfoURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create userinfo request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call userinfo endpoint: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			slog.Error("failed to close response body: %v", slog.Any("err", err))
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("userinfo endpoint returned status %d: %s", resp.StatusCode, string(body))
	}

	var auth0UserInfo Auth0UserInfo
	if err := json.NewDecoder(resp.Body).Decode(&auth0UserInfo); err != nil {
		return nil, fmt.Errorf("failed to unmarshal userinfo response: %v", err)
	}

	user := &User{
		ID:       auth0UserInfo.Sub,
		Email:    auth0UserInfo.Email,
		Username: auth0UserInfo.Nickname,
		Metadata: map[string]string{
			"name":        auth0UserInfo.Name,
			"given_name":  auth0UserInfo.GivenName,
			"family_name": auth0UserInfo.FamilyName,
			"picture":     auth0UserInfo.Picture,
			"updated_at":  auth0UserInfo.UpdatedAt.Format(time.RFC3339),
		},
	}

	return user, nil
}

// validateAuth0Config validates that all required Auth0 configuration is present
func validateAuth0Config(config *Config) error {
	if !config.IsAuth0Configured() {
		return fmt.Errorf("Auth0 configuration is missing. Please check your environment variables")
	}

	issuerURL := config.GetAuth0IssuerURL()
	if issuerURL == "" {
		return fmt.Errorf("AUTH0_ISSUER_URL environment variable is not set or is empty")
	}

	return nil
}

// buildAuth0URL builds a properly formatted Auth0 URL by ensuring no double slashes
func buildAuth0URL(baseURL, path string) string {
	// Ensure the base URL doesn't end with a slash
	cleanBase := strings.TrimRight(baseURL, "/")
	// Ensure the path starts with a slash
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return cleanBase + path
}

// buildCallbackURL builds the callback URL for Auth0
func buildCallbackURL(config *Config) string {
	return fmt.Sprintf("%s/auth/callback", config.APIURL)
}

// GetUserFromContext extracts the authenticated user from Gin context
func GetUserFromContext(c *gin.Context) *User {
	if user, exists := c.Get("user"); exists {
		if u, ok := user.(*User); ok {
			return u
		}
	}
	return nil
}
