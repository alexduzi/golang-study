# Working with Packaging and Modules in Go

Follow these steps to work with packaging and modules in Go:

## 1. **Initialize a Module**
```sh
go mod init <module-name>
```

## 2. **Organize Code into Packages**
- Place related `.go` files in the same directory.
- Each directory is a package; add `package <name>` at the top of each file.

## 3. **Import Packages**
- Use `import "<module-name>/<package-path>"` in your code.

## 4. **Export Functions and Types**
- Capitalize names to export them (e.g., `FuncName`).

## 5. **Add Dependencies**
```sh
go get <dependency>              # Add or update dependency
go get <dependency>@<version>    # Get specific version (e.g., @v1.2.3, @latest)
go get -u <dependency>           # Update to latest minor/patch version
go get -u=patch <dependency>     # Update to latest patch version only
go mod download                  # Download modules to local cache
go mod vendor                    # Optional: vendor dependencies locally
```

## 6. **Update and Manage Dependencies**
```sh
go mod tidy       # Remove unused and add missing dependencies
go mod verify     # Verify dependencies haven't been modified
```

## 7. **Working with Go Workspaces**

Go workspaces allow you to work with multiple modules simultaneously:

### Initialize a Workspace
```sh
go work init <module1-path> <module2-path>
```

### Add Modules to Workspace
```sh
go work use <module-path>
```

### Remove Modules from Workspace
```sh
go work use -r <module-path>
```

### Edit Workspace File
```sh
go work edit -fmt     # Format go.work file
go work edit -go=1.21 # Set Go version
```

### Sync Workspace Dependencies
```sh
go work sync          # Sync workspace build list back to modules
```

## 8. **Build and Test**
```sh
go mod download       # Download dependencies
go build
go test ./...
```

## 9. **Module Inspection and Troubleshooting**
```sh
go mod graph          # Print module requirement graph
go mod why <package>  # Explain why packages or modules are needed
go list -m all        # List all modules in build
```

## 10. **Distribute Your Module**
- Push your code to a public VCS (e.g., GitHub).
- Tag releases using semantic versioning (e.g., `v1.0.0`).
- Others can import using your module path.

## Example Workspace Structure
```
my-workspace/
├── go.work
├── module1/
│   ├── go.mod
│   └── main.go
└── module2/
    ├── go.mod
    └── lib.go
```

**References:**
- [Go Modules Reference](https://golang.org/doc/go1.11#modules)
- [Go Workspaces Tutorial](https://go.dev/doc/tutorial/workspaces)
- [Effective Go - Packages](https://golang.org/doc/effective_go#packages)
- [Module Management Commands](https://golang.org/ref/mod)