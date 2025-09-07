# Working with Packaging in Go

Follow these steps to work with packaging in Go:

1. **Create a Module**
    ```sh
    go mod init <module-name>
    ```

2. **Organize Code into Packages**
    - Place related `.go` files in the same directory.
    - Each directory is a package; add `package <name>` at the top of each file.

3. **Import Packages**
    - Use `import "<module-name>/<package-path>"` in your code.

4. **Export Functions and Types**
    - Capitalize names to export them (e.g., `FuncName`).

5. **Build and Test**
    ```sh
    go build
    go test ./...
    ```

6. **Add Dependencies**
    ```sh
    go get <dependency>
    ```

7. **Tidy Up**
    ```sh
    go mod tidy
    ```

8. **Distribute Your Module**
    - Push your code to a public VCS (e.g., GitHub).
    - Others can import using your module path.

**References:**
- [Go Modules Reference](https://golang.org/doc/go1.11#modules)
- [Effective Go - Packages](https://golang.org/doc/effective_go#packages)