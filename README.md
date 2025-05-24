# MongoDB Server

A Go server that exposes API to enable multi-tenancy on single MongoDB deployment.

## 🐰 Quick Start

### 📄 Pre-requisites

1. A MongoDB server Deployment
2. Environment variable `MONGODB_URI` with the value of of MongoDB server URL
3. Environment variable `JWT_SECRET` for generating JWT tokens

### 🏗️ Build

```bash
make
```

### 🏁 Run Binary

```bash
./bin/Mongo-Server
```

## API
| Routes      | Description                           | Method |
| ----------- | ------------------------------------- | ------ |
| /register   | Registers the tenant                  |POST    |
| /auth/token | Creates JWT Token                     |POST    |
| /schema     | Registers schema to create collection |POST    |

For more information, refer [API Docs](./docs/API.md)