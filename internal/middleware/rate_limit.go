package middleware

import (
	"net/http"
	"time"

	"github.com/vishal-Choudhary-hi/chotu/internal/repository"
)

func RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ip := r.RemoteAddr
		key := "ratelimit:" + ip

		count, _ := repository.RedisClient.Incr(repository.Ctx, key).Result()

		if count == 1 {
			repository.RedisClient.Expire(repository.Ctx, key, time.Minute)
		}

		if count > 60 {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
