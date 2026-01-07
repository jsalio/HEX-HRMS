# Agrega los módulos al workspace (edita go.work o usa comandos)
example
```
    go work use ./core
    go work use ./infra/repository
```
# Define dependencias entre módulos

example:
```
    go mod edit -require=hrms/core@v0.0.0   
    go mod edit -replace hrms/core@v0.0.0=../core
```