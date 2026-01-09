package middleware

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type Client struct {
	limiter      *rate.Limiter
	lastAccessAt time.Time
}

var (
	clients = make(map[string]*Client)
	mu      sync.Mutex
)

func getClientIp(ctx *gin.Context) string {
	ip := ctx.ClientIP()

	if ip == "" {
		ip = ctx.Request.RemoteAddr
	}

	return ip
}

func CleanupRateLimiters() {
	for {
		mu.Lock()

		for ip, client := range clients {
			if time.Since(client.lastAccessAt) > 5*time.Minute {
				delete(clients, ip)
			}
		}

		mu.Unlock()
	}
}

func getRateLimiter(ip string) *rate.Limiter {
	fmt.Printf("IP: %s\n", ip)

	mu.Lock()
	defer mu.Unlock()

	_, exists := clients[ip]

	if !exists {
		limiter := rate.NewLimiter(20, 50)
		newClient := &Client{limiter, time.Now()}
		clients[ip] = newClient
	}

	return clients[ip].limiter
}

func RateLimitByIp() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := getClientIp(ctx)

		limiter := getRateLimiter(ip)

		if !limiter.Allow() {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests. Please try again in some minutes."})
			return
		}
	}
}
