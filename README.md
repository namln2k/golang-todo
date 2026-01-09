# Project structure

# Rate limit middleware
## Tutorials
https://youtu.be/REX4cuSOkAQ?si=h7QOgDZVTgqva9v6
## Notes
- Simply limit request rates by IP address
- Using package `rate`: <a href="https://pkg.go.dev/golang.org/x/time/rate">golang.org/x/time/rate</a>
- Middleware declaration
```
func (engine *gin.Engine) Use(middleware ...gin.HandlerFunc) gin.IRoutes
```
- Extract IP from `ctx *gin.Context`
```
ctx.ClientIP()

ctx.Request.RemoteAddr
```
- Rate limiter implements <a href="https://viblo.asia/p/cac-thuat-toan-thuong-dung-cho-bo-gioi-han-truy-cap-V3m5WRPblO7#_token-bucket-0">**token bucket algorithm**</a>
- Func `NewLimiter`
  - `r`: Rate to fill *tokens* to the *bucket*
  - `b`: Allowed number of burst tokens at once
```
func NewLimiter(r Limit, b int) *Limiter
```
- Load testing with <a href="https://www.baeldung.com/linux/ab-apachebench-load-testing">Apache Benchmark</a>

# Database connection
# User CRUD
# Sign up
# Log in
# Access token
# Refresh token
# Todo CRUD
# Logging
## Tutorials
- https://youtu.be/UE_4XagcU0A?si=hUSkeIGp4UI-Ij71
- https://youtu.be/kGXeeHWJw24?si=R3Qj7GvKn_QHxmKS
- https://youtu.be/Q5hJByOrTYE?si=3x-YAwH2_Yrgu-1f
- https://youtu.be/SLY87qRDF-Y?si=g6-A_7HvMCJ7KcQR
