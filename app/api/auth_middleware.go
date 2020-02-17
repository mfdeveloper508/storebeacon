package api

import (
	"context"
	"net/http"

	"github.com/storebeacon/backend/models/store"
)

// ForceLoginMiddleware returns a 401 response if not logged in
func ForceLoginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		type responseFormat struct {
			Success bool   `json:"success"`
			Message string `json:"message"`
		}
		// Get token from header
		token, err := TokenFromHeader(r)

		if err != nil {
			response := responseFormat{
				Success: false,
				Message: err.Error(),
			}

			ServeJSONWithStatusCode(w, response, http.StatusUnauthorized)
			return
		}

		if token == "" {
			noAuthResponse := responseFormat{
				Success: false,
				Message: "Not Authorised",
			}

			ServeJSONWithStatusCode(w, noAuthResponse, http.StatusUnauthorized)
			return
		}

		if CheckJWT(token) != nil {
			response := responseFormat{
				Success: false,
				Message: "Token is invalid",
			}

			ServeJSONWithStatusCode(w, response, http.StatusUnauthorized)
			return
		}

		// Get subscription user from token UUID
		claims, err := ParseJWTClaims(token)

		if err != nil {
			response := responseFormat{
				Success: false,
				Message: "Token is invalid",
			}

			ServeJSONWithStatusCode(w, response, http.StatusUnauthorized)
			return
		}

		userID := claims["userID"].(float64)

		// Find the user
		user, err := store.GetRecordByIndex(uint(userID))
		if err != nil || user.Email == ""{
			response := responseFormat{
				Success: false,
				Message: "Not Authorised",
			}

			ServeJSONWithStatusCode(w, response, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", user.ID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ForceUILoginMiddleware(w http.ResponseWriter, r *http.Request, token string) int {

	invalidUid := 0

	// Get subscription user from token UUID
	claims, err := ParseJWTClaims(token)

	if err != nil {
		return invalidUid;
	}

	userID := claims["userID"].(float64)

	// Find the user
	user, err := store.GetRecordByIndex(uint(userID))
	if err != nil || user.Email == ""{
		return invalidUid;
	}
	return int(user.ID)
}
