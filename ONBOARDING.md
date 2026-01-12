# 🚀 Onboarding Guide - HEX-HRMS

Welcome to the team! This document serves as a guide to understanding the architecture, design patterns, and coding standards used in this project.

## 🏗 Architectural Overview

This project follows **Hexagonal Architecture (Clean Architecture)** principles strictly. The goal is to decouple the business logic (Core) from external concerns (Database, UI, Frameworks).

### 📂 Folder Structure

```
/
├── core/             # 🧠 Domain Logic (No external dependencies)
│   ├── models/       # Structs/Entities
│   ├── contracts/    # Interfaces (Ports) for UseCases and Repositories
│   └── usecases/     # Application Logic (Interactors)
│
├── infra/            # 🔌 Adapters & External Implementations
│   ├── api/          # HTTP Layer (Controllers, Middleware, Gin)
│   ├── repository/   # Database Implementations (GORM, Postgres)
│   └── hrms-web/     # Frontend Application (Angular)
│
└── cmd/              # 🏁 Entry point
```

---

## 🟢 Backend Guidelines (Go)

### 1. The Flow of Data
Every request follows this path:
`Controller (Infra) -> UseCase (Core) -> Contract/Interface (Core) -> Repository Implementation (Infra)`

### 2. Coding Rules
- **Dependency Rule**: `core` must NEVER import `infra`. 
- **Contracts First**: Always define an interface in `core/contracts` before implementing the repository.
- **Dependency Injection**: We use **Manual Dependency Injection** (no magic frameworks). Dependencies are wired in `server.go`.
- **Generics**: Use `GenericCrud[T, G]` for standard CRUD operations to avoid boilerplate.
    - `T`: Domain Model (from `core/models`)
    - `G`: GORM Model (from `infra/repository`)
- **Error Handling**: Use `models.SystemError` for passing errors across layers.

### 3. How to Add a New Feature (Backend)
1. **Model**: Define the struct in `core/models`.
2. **Contract**: Define the interface in `core/contracts`.
3. **Repository**: Implement the interface in `infra/repository/postgress/repo`.
    - Use `repo.NewGenericCrud` if possible.
    - Create mapping functions `ToModel` and `ToEntity`.
4. **UseCase**: Create a logic handler in `core/usecases`.
    - It must depend on the **Contract**, not the Repository struct.
5. **Controller**: Create the HTTP handler in `infra/api/controller`.
6. **Wiring**: Register everything in `infra/api/server.go`.

---

## 🔵 Frontend Guidelines (Angular)

We do **NOT** write standard Angular code. We apply Clean Architecture to the frontend as well.

### 1. Architecture
```
src/app/
├── core/             # Frontend Business Logic
│   ├── domain/       # Models & Ports (Repository Interfaces)
│   └── usecases/     # Classes that execute specific actions
│
├── infrastructure/   # API Clients (Http implementations of Ports)
└── ui/               # Components & Pages
```

### 2. Coding Rules
- **Standalone Components**: All new components must be `standalone: true`.
- **Signals**: Use **Angular Signals** for state management instead of `BehaviorSubject` where possible.
- **No Logic in UI**: Components should only call **UseCases**. They should not make HTTP calls directly.
- **Tailwind CSS**: Use utility classes. Avoid custom CSS files unless necessary.

### 3. How to Add a New Feature (Frontend)
1. **Model**: Define interface in `core/domain/models`.
2. **Port**: Define the Repository `interface` and `InjectionToken` in `core/domain/ports`.
3. **Infrastructure**: Implement the service in `infrastructure/` implementing the Port.
4. **UseCase**: Create a simple class for the action (e.g., `CreateUserUseCase`) in `core/usecases`.
5. **UI**: Inject the `UseCase` into your component.

---

## 📝 General Best Practices
- **Naming**: Use clear, descriptive names.
- **Testing**: Write unit tests for UseCases (Core is easy to test!).
- **Commits**: Follow conventional commits (feat, fix, refactor, etc.).

---
*Happy Coding!*
