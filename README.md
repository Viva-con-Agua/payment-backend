# payment-backend
payment service implementation with go and echo

## use
install docker-compose

## redis
run `docker-compose up -d redis`
## database
run `docker-compose up -d payment-database`

## development

### 1. Install Go language 
Like here: https://itrig.de/index.php?/archives/2377-Installation-einer-aktuellen-Go-Version-auf-Ubuntu.html

### 2. Install dependecies
```
go get github.com/stripe/stripe-go
go get github.com/stripe/stripe-go/customer
go get github.com/stripe/stripe-go/paymentintent
```

### 3. Checkout payment-backend-go
git clone https://github.com/Viva-con-Agua/payment-backend-go.git

### 4. Run server
Start server wiht `go run server.go`

### 5. update nginx
Update IP to you local IP in develop-pool branch at `routes/nginx-pool/pool.upstream` and restart nginx-pool docker with `docker restart pool-nginx`