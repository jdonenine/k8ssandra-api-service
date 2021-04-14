package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	ggauth "github.com/shaj13/go-guardian/auth"
	"github.com/shaj13/go-guardian/auth/strategies/basic"
	"github.com/shaj13/go-guardian/auth/strategies/bearer"
	"github.com/shaj13/go-guardian/store"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type AuthManager struct {
	Kubeclient     *kubernetes.Clientset
	KubeNamespace  string
	UserSecretName string
	authenticator  ggauth.Authenticator
}

func (manager *AuthManager) Init() {
	authenticator := ggauth.New()
	cache := store.NewFIFO(context.Background(), time.Minute*10)
	basicStrategy := basic.New(manager.validateUser, cache)
	tokenStrategy := bearer.New(manager.verifyToken, cache)
	authenticator.EnableStrategy(basic.StrategyKey, basicStrategy)
	authenticator.EnableStrategy(bearer.CachedStrategyKey, tokenStrategy)
	manager.authenticator = authenticator
}

func (manager *AuthManager) WithAuthentication(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := manager.authenticator.Authenticate(r)
		if err != nil {
			code := http.StatusUnauthorized
			http.Error(w, http.StatusText(code), code)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (manager *AuthManager) validateUser(ctx context.Context, r *http.Request, userName, password string) (ggauth.Info, error) {
	if manager.Kubeclient == nil {
		log.Println("Unable to perform authentication, no valid client configured.")
		return nil, fmt.Errorf("unable to perform authentication, no valid context provided")
	}

	secret, err := manager.Kubeclient.CoreV1().Secrets(manager.KubeNamespace).Get(context.TODO(), manager.UserSecretName, metav1.GetOptions{})
	if secret == nil || err != nil {
		log.Printf("Unable to retrieve cluster secret '%s' in namespace '%s', authentication cannot be performed.", manager.UserSecretName, manager.KubeNamespace)
		if err != nil {
			log.Print(err)
		}
		return nil, fmt.Errorf("unable to perform authentication")
	}

	if userName == string(secret.Data["username"]) && password == string(secret.Data["password"]) {
		return ggauth.NewDefaultUser("userName", "userName", nil, nil), nil
	}

	return nil, fmt.Errorf("invalid credentials")
}

func (manager *AuthManager) verifyToken(ctx context.Context, r *http.Request, tokenString string) (ggauth.Info, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user := ggauth.NewDefaultUser(claims["sub"].(string), claims["sub"].(string), nil, nil)
		return user, nil
	}

	return nil, fmt.Errorf("invaled token")
}
