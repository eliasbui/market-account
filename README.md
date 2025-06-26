# Market Account - Microservices Platform

A comprehensive e-commerce marketplace built with microservices architecture, implementing Clean Architecture principles across multiple programming languages.

## 🏗️ Architecture Overview

This project implements a polyglot microservices architecture with the following services:

- **User Service** (Go) - User authentication and management
- **Product Service** (C#/.NET) - Product catalog and inventory
- **Payment Service** (Python) - Payment processing and transactions
- **API Gateway** (Node.js) - GraphQL gateway with service orchestration

## 🛠️ Technology Stack

### Core Technologies
- **Go** - User Service with Gin framework
- **C#/.NET 8** - Product Service with ASP.NET Core
- **Python** - Payment Service with FastAPI
- **Node.js** - API Gateway with Apollo GraphQL

### Infrastructure
- **PostgreSQL** - Primary database for all services
- **Redis** - Caching and session management
- **Apache Cassandra** - Analytics and time-series data
- **RabbitMQ** - Message queue for event-driven architecture
- **Docker & Docker Compose** - Containerization and orchestration

### Communication
- **GraphQL** - API Gateway and external client communication
- **gRPC** - Internal service-to-service communication
- **REST APIs** - External integrations
- **Message Queues** - Asynchronous event processing

## 📁 Project Structure

```
market-account/
├── services/
│   ├── user-service-go/           # Go microservice
│   │   ├── domain/                # Business entities
│   │   ├── application/           # Use cases & interfaces
│   │   ├── infrastructure/        # Data access & external services
│   │   ├── web/                   # HTTP handlers & routing
│   │   ├── go.mod
│   │   └── main.go
│   │
│   ├── product-service-dotnet/    # .NET microservice
│   │   ├── Domain/                # Domain layer
│   │   ├── Application/           # Application layer
│   │   ├── Infrastructure/        # Infrastructure layer
│   │   ├── WebApi/                # Presentation layer
│   │   └── ProductService.sln
│   │
│   ├── payment-service-python/    # Python microservice
│   │   ├── domain/                # Domain models
│   │   ├── application/           # Application services
│   │   ├── infrastructure/        # Infrastructure layer
│   │   ├── web/                   # Web controllers
│   │   └── requirements.txt
│   │
│   └── api-gateway-node/          # Node.js API Gateway
│       ├── src/
│       │   ├── domain/            # GraphQL schemas
│       │   ├── application/       # Service orchestration
│       │   ├── infrastructure/    # External service clients
│       │   └── web/               # GraphQL resolvers
│       └── package.json
│
├── docker-compose.yml             # Multi-service orchestration
├── .taskmaster/                   # Task management
└── README.md
```

## 🚀 Quick Start

### Prerequisites
- Docker & Docker Compose
- Git

### Development Setup

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd market-account
   ```

2. **Start all services**
   ```bash
   docker-compose up -d
   ```

3. **Verify services are running**
   ```bash
   docker-compose ps
   ```

### Service Endpoints

| Service | Port | Endpoint | Description |
|---------|------|----------|-------------|
| API Gateway | 8000 | http://localhost:8000/graphql | GraphQL API |
| User Service | 8001 | http://localhost:8001/health | Go REST API |
| Product Service | 8002 | http://localhost:8002/health | .NET REST API |
| Payment Service | 8003 | http://localhost:8003/health | Python REST API |
| PostgreSQL | 5432 | localhost:5432 | Database |
| Redis | 6379 | localhost:6379 | Cache |
| RabbitMQ | 15672 | http://localhost:15672 | Management UI |
| Adminer | 8080 | http://localhost:8080 | Database Admin |

### Development Credentials

- **PostgreSQL**: postgres/postgres
- **RabbitMQ**: admin/admin123
- **Database**: marketplace

## 🏛️ Clean Architecture Implementation

Each microservice follows Clean Architecture principles with language-specific adaptations:

### Go Service (Feature-Based)
```
user-service-go/
├── domain/          # Entities, value objects
├── application/     # Use cases, interfaces
├── infrastructure/  # Repository implementations
└── web/            # HTTP handlers, routing
```

### .NET Service (Layer-Based)
```
product-service-dotnet/
├── Domain/          # ProductService.Domain
├── Application/     # ProductService.Application  
├── Infrastructure/  # ProductService.Infrastructure
└── WebApi/         # ProductService.WebApi
```

### Python Service (Module-Based)
```
payment-service-python/
├── domain/          # Models, enums, business rules
├── application/     # Services, use cases
├── infrastructure/  # Repositories, external APIs
└── web/            # FastAPI controllers
```

### Node.js Gateway (Layered)
```
api-gateway-node/src/
├── domain/          # GraphQL schemas, types
├── application/     # Service orchestration
├── infrastructure/  # Service clients, adapters
└── web/            # Resolvers, middleware
```

## 🔧 Development Workflow

This project uses **Taskmaster AI** for task-driven development. Key commands:

```bash
# View all tasks
task-master list

# Get next task to work on
task-master next

# View specific task details
task-master show <task-id>

# Update task progress
task-master update-subtask --id=<subtask-id> --prompt="progress update"

# Mark tasks as complete
task-master set-status --id=<task-id> --status=done
```

## 🧪 Testing Strategy

Each service includes comprehensive testing:

- **Unit Tests** - Domain logic and use cases
- **Integration Tests** - Repository and service interactions
- **Contract Tests** - API endpoint verification
- **End-to-End Tests** - Full service orchestration

## 📈 Current Implementation Status

### ✅ Completed
- [x] Project structure and Clean Architecture foundation
- [x] Docker Compose infrastructure setup
- [x] Basic service scaffolding across all languages
- [x] Database and caching configuration
- [x] GraphQL schema definition

### 🚧 In Progress
- [ ] gRPC service communication
- [ ] REST API implementations
- [ ] Message queue integration
- [ ] Authentication & authorization
- [ ] Database migrations

### 📋 Upcoming
- [ ] Frontend React application
- [ ] AI-powered chatbot
- [ ] Admin dashboard
- [ ] Production deployment pipeline

## 🤝 Contributing

1. Follow Clean Architecture principles
2. Maintain language-specific conventions
3. Update Taskmaster tasks for progress tracking
4. Include comprehensive tests
5. Document significant changes

## 📚 Documentation

- [Clean Architecture Guide](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Microservices Patterns](https://microservices.io/)
- [Docker Compose Reference](https://docs.docker.com/compose/)
- [GraphQL Best Practices](https://graphql.org/learn/best-practices/)

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

---

**Built with ❤️ using Taskmaster AI for task-driven development** 