# Networking Tables

Networking Tables is a smart event app that enhances networking. It efficiently shuffles seating arrangements, ensuring everyone meets with minimal table switches, fostering a dynamic and engaging experience.

## Dependencies

- Python 3.7+
- Poetry (python package manager)
- Go 1.13+

## Installation

#### Python Version

```bash
make setup-python
```

#### Go Version

```bash
make setup-go
```

## Usage

You will be prompted to enter the names of the people (separated by spaces), the capacity of each table (as a series of numbers separated by spaces), and the number of combinations to compute.

**Note:** *The Go version runs much faster than the Python version.*

**Example:**

```bash
Enter the names of the people (separated by spaces): Alice Bob Charlie David Eve

Enter the capacity of each table (as a series of numbers separated by spaces): 2 3 4 5

Enter the number of combinations to compute: 10000

```

#### Python Version

```bash
make run-python
```

#### Go Version
  
```bash
make run-go
```

## Tests

There is a sanity E2E test under the tests folder. To run it, use the command:
```bash
make test-python
```

## Contribute

Contributions are welcome! Please fork the repository and create a pull request with your changes.
