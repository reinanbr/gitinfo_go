# Contributing to GitInfo Go

Thank you for your interest in contributing to GitInfo Go! We welcome contributions from the community.

## How to Contribute

### Reporting Issues

If you find a bug or have a feature request, please open an issue on GitHub with:
- A clear description of the problem or feature
- Steps to reproduce (for bugs)
- Expected vs actual behavior
- Your Go version and operating system

### Submitting Pull Requests

1. Fork the repository
2. Create a new branch for your feature or bugfix
3. Write tests for your changes
4. Ensure all tests pass: `go test`
5. Run `go fmt` to format your code
6. Submit a pull request with a clear description

### Development Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/reinanbr/gitinfo_go.git
   cd gitinfo_go
   ```

2. Run tests:
   ```bash
   go test
   ```

3. Run the example:
   ```bash
   cd example
   go run main.go
   ```

### Code Style

- Follow standard Go formatting (`go fmt`)
- Write clear, descriptive function and variable names
- Add comments for public functions
- Write tests for new functionality

### Testing

- All new features should include tests
- Tests should pass consistently
- Aim for good test coverage

## Questions?

Feel free to open an issue for any questions about contributing!
