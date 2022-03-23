### signup

POST http://localhost:8080/v1/signup

### signin

POST http://localhost:8080/v1/signin

### profile

POST http://localhost:8080/v1/profile

### profile/update

POST http://localhost:8080/v1/profile/update

```
.
│  .env
│  config.toml //port config
│  go.mod
│  go.sum
│  main.go
│  README.md
│  
│            
├─controller
│      auth.go
│      
├─db
│  └─migrations //database
│          000001_create_users_table.down.sql
│          000001_create_users_table.up.sql
│          
├─domain
│      user.go
│      
├─infra
│      postgres.go //database config
│      router.go
│      
├─middleware
│  ├─auth
│  │      auth.go
│  │      
│  ├─cache
│  │      redis.go
│  │      
│  ├─config
│  │      config.go
│  │      
│  ├─constants
│  │      const.go
│  │      
│  ├─jwt
│  │      jwt.go
│  │      
│  └─response
│          response.go
│          
├─repository
│  │  user.go
│  │  
│  └─impl
│          user.go
│          
├─testdata
│      auth.http //Examples of requests
│      http-client.private.env.json //set token
│      
└─usecase
    │  auth.go
    │  
    └─impl
            auth.go

```