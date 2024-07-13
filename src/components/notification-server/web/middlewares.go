package web

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/authentication/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func AuthMiddleware(c *gin.Context) {
	// Auth middleware
	tokenList := strings.Split(c.Request.Header.Get("Authorization"), "Bearer ")
	if len(tokenList) != 2 {
		c.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}
	token := tokenList[1]

	k8sClient := getK8sClient()
	tokenReview, err := k8sClient.AuthenticationV1().TokenReviews().Create(
		c,
		&v1.TokenReview{
			Spec: v1.TokenReviewSpec{
				Token: token,
			},
		},
		metav1.CreateOptions{},
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
			"msg":   err.Error(),
		})
		c.Abort()
		return
	}

	if tokenReview.Status.Authenticated {
		c.Set("account", tokenReview.Status.User.UID)
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}
}
