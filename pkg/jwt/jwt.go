package jwt

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// type Service interface {
// 	GenerateToken(userID int) (string, error)
// 	ValidateToken(encodedToken string) (*jwt.Token, error)
// }

// type JwtService struct {
// }

// var Jwt *JwtService

var secret = []byte("secret")

// var exp = 120

// func NewService() *JwtService {
// 	return &JwtService{}
// }

// TokenMetadata struct to describe metadata in JWT.
type TokenMetadata struct {
	Email   string
	Expires int64
}

func GenerateToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   jwt.NewNumericDate(time.Now().Add(24 * time.Hour)).Unix(),
	}
	// claim["user_id"] = email

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(secret)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

// // GenerateNewAccessToken func for generate a new Access token.
// func GenerateNewAccessToken(email string) (string, error) {
// 	// Set secret key from .env file.
// 	// secret := os.Getenv("JWT_SECRET_KEY")

// 	// Set expires minutes count for secret key from .env file.
// 	minutesCount := exp

// 	// Create a new claims.
// 	claims := jwt.MapClaims{}

// 	// Set public claims:
// 	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()
// 	claims["email"] = email

// 	// Create a new JWT access token with claims.
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	// Generate token.
// 	t, err := token.SignedString(secret)
// 	if err != nil {
// 		// Return error, it JWT token generation failed.
// 		return "", err
// 	}

// 	return t, nil
// }

// func ValidateToken(encodedToken string) (jwt.MapClaims, error) {
// 	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
// 		_, ok := token.Method.(*jwt.SigningMethodHMAC)
// 		if !ok {
// 			return nil, errors.New("invalid token")
// 		}

// 		return secret, nil
// 	})

// 	if err != nil {
// 		return jwt.MapClaims{}, err
// 	}

// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if ok && token.Valid {
// 		return claims, nil
// 	}

// 	return jwt.MapClaims{}, fmt.Errorf("unable to extract claims")

// }

// ExtractTokenMetadata func to extract metadata from JWT.
func ExtractTokenMetadata(c *fiber.Ctx) (*TokenMetadata, error) {
	token, err := verifyToken(c)
	if err != nil {
		return nil, err
	}

	// Setting and checking token and credentials.
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		// Expires time.
		expires := int64(claims["exp"].(float64))
		email := claims["email"].(string)

		return &TokenMetadata{
			Expires: expires,
			Email:   email,
		}, nil
	}

	return nil, err
}

func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")
	// log.Println(bearToken)

	// Normally Authorization HTTP header.
	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(c)

	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return secret, nil
}
