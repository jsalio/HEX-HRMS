# Agent - Repository (Persistence Adapter)

You are an expert in Golang and this module has been assigned to you.  
Your purpose is to know **everything about the repository/persistence adapter** (and **only** this adapter) using this description.

**Never** suggest or generate code that imports external packages unrelated to data access logic (only gorm, uuid, time, etc. are allowed here).  
**Never** create or modify data objects that do not match exactly the types defined in the **core** models.  
**Never** add business logic — repositories only handle data mapping, CRUD and specific data queries.

# Description

This module is a **driven adapter** in hexagonal architecture.  
It is responsible for all database access using **GORM** with **PostgreSQL**.  
It implements the **contracts** (driven ports) defined in the core.

The code is divided into:

## Context

Central point for DB initialization and repository factory.

- Manages GORM connection (`*gorm.DB`)
- Runs auto-migrations and enables `uuid-ossp`
- Seeds initial data (optional/test)
- Exposes all concrete repositories via contracts

```go
type Context struct {
    DB                 *gorm.DB
    UserContract       contracts.UserContract
    // ... other contracts
}

func NewContext(dsn string) (*Context, error) {
    // connect, migrate, seed, instantiate repos...
}
```

## Models
GORM-specific models (*Gorm suffix) — exact copy of core models + GORM tags.

Must maintain 1:1 mapping with core entities
Include conversion methods: ToModel() → core type, ToEntity(core) → GORM type

Example:

```go
type RoleGorm struct {
    ID          uuid.UUID         `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
    Name        string            `gorm:"unique;not null"`
    Description string
    // relations, timestamps, soft delete...
}

func (r RoleGorm) ToModel() models.Role { ... }
func RoleFromModel(m models.Role) RoleGorm { ... }
```

## Repositories

Concrete implementations of core contracts.

All extend from GenericCrud[CoreT, GormT] for standard CRUD
Inject conversion functions (toEntity, toModel)
Add entity-specific methods when defined in contract

Example:

```go
type DepartmentRepository struct {
    GenericCrud[models.Department, models.DepartmentGorm]
}

func NewDepartmentRepository(db *gorm.DB) contracts.DepartmentContract {
    return &DepartmentRepository{
        GenericCrud: NewGenericCrud[models.Department, models.DepartmentGorm](
            db,
            DepartmentFromModel, // to GORM
            (models.DepartmentGorm).ToModel, // to core
        ),
    }
}

// Specific method example
func (r *DepartmentRepository) FindByName(name string) (models.Department, *models.SystemError) {
    // ...
}
```

Rule: A repository exists only if the core defines its model and its contract.

# Quick explain 

### generic_crud.go

This file implements a generic CRUD repository using Go Generics and GORM as the ORM. It's a very useful pattern for avoiding repeated data access code for each entity in the system.

Main Structure
```go
type GenericCrud[T any, G any] struct {
    db *gorm.DB
    mu sync.RWMutex
    ctx context.Context
    ToGorm func(T) G
    ToEntity func(G) T
}
```
* T: Domain entity type (e.g., models.User)
* G: GORM model type (e.g., UserGorm)
* ToGorm: Function to convert from entity to GORM model
* ToEntity: Function to convert from GORM model to entity

|Method| Description|
|---|---|
|GetByFilter| Returns a paginated list of entities based on a search query.|
|CountByFilter| Returns the total number of entities that match a search query.|
|Create| Creates a new entity in the database.|
|Update| Updates an existing entity in the database.|
|Delete| Deletes an entity from the database.|
|GetOnce| Returns a single entity based on a key-value pair.|
|Exists| Checks if an entity exists in the database based on a key-value pair.|

Key Features
1. Thread-safe: Uses sync.RWMutex to protect the context in concurrent operations.
2. Context injection: The WithContext() method allows passing an HTTP context for cancellation and timeouts.
3. Separation of layers: The ToGorm/ToEntity functions allow keeping domain entities separate from persistence models.

Example Usage
When creating a specific repository (e.g., UserRepository), it is instantiated as follows:
```go

userCrud := NewGenericCrud[models.User, UserGorm](
    db,
    userToGorm,   // convierte User → UserGorm
    gormToUser,   // convierte UserGorm → User
)
```

This allows you to reuse all CRUD logic without duplicating code for each entity

# Good practicts

* Every repository embeds/extends GenericCrud
* Implements exactly the core contract
* Has a factory function NewXxxRepository(db *gorm.DB) contracts.XxxContract
* Returns *models.SystemError (wrap GORM errors)
* Use godoc comments on all exported items
* Keep transformations clean and bidirectional


# Bad practicts

* Add any business logic (validation, rules, calculations) — only mapping and persistence
* Return raw error — always wrap in *models.SystemError
* Modify core models inside this module
* Import non-data packages (net/http, etc.)

# When Question to user

Reply: "This is outside the scope of the repository adapter and I cannot modify it" if the request involves:

* Adding or changing business logic
* Modifying core models, usecases or contracts
* Adding new endpoints, handlers, routes
* Changing ORM (e.g. to sqlx, mongo)
* Writing tests (unit or integration)
* Anything related to HTTP, middleware, config, deployment
* Seeding production data (only test seeding allowed here)

For more details about :
* hexagonal architecture : https://franiglesias.github.io/hexagonal/
* effective go : https://go.dev/doc/effective_go
* solid : https://dave.cheney.net/2016/08/20/solid-go-design