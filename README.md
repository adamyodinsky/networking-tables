# Networking Tables
Networking Tables is a smart event app that enhances networking. It efficiently shuffles seating arrangements, ensuring everyone meets with minimal table switches, fostering a dynamic and engaging experience.

Copy code

# Networking Tables

Networking Tables is a smart event app designed to optimize social interactions at gatherings. It efficiently shuffles seating arrangements, ensuring everyone meets with minimal table switches, fostering a dynamic and engaging experience.

## Installation

### Python Version

Navigate to the `nt-python` directory and install the dependencies using Poetry:

```bash
make setup-python
```

### Go Version

```bash
make setup-go
```


## Usage

```bash
make run-python
```

You will be prompted to enter the names of the people (separated by spaces), the capacity of each table (as a series of numbers separated by spaces), and the number of combinations to compute.

## Tests

There is a sanity E2E test under the tests folder. To run it, use the command:
```bash
make test-python
```


Contribute
Contributions are welcome! Please fork the repository and create a pull request with your changes.

License
This project is licensed under the terms of the MIT license.

Copy code


This README provides a comprehensive overview of the project. However, the specific commands and directories mentioned (such as `requirements.txt` and `tests`) should be adjusted to match your actual project structure and setup.

