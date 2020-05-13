package request_handler

import (
	"common"
	"github.com/dgrijalva/jwt-go"
	"model/service_error"
	"net/http"
	"service/security_service"
)

type config struct {
	Auth authConfig `yaml:"auth"`
}

type authConfig struct {
	Secret string `yaml:"secret"`
}

var secret string

func init() {
	cfg := config{}
	common.ReadConfig(&cfg)

	secret = cfg.Auth.Secret
}

func TokenValidator(inner http.Handler, grants []string) http.Handler {
	if len(grants) == 0 {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			inner.ServeHTTP(w, r)
		})
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := security_service.VerifyToken(r)
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			actualGrants, ok := claims["grants"].([]interface{})
			if !ok {
				panic(service_error.ForbiddenError{})
			}
			grantsMap := map[interface{}]bool{}
			for _, grant := range actualGrants {
				grantsMap[grant] = true
			}
			for _, grant := range grants {
				if _, hasKey := grantsMap[grant]; !hasKey {
					panic(service_error.ForbiddenError{})
				}
			}
		}

		inner.ServeHTTP(w, r)
	})
}



