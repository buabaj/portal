# Portal

Portal is a basic network scanner and port mapper written in Go. It allows you to scan IP addresses or entire subnets for open TCP ports.

## Usage

1. Clone the repository or copy the `scanner.go` file to your local machine.
2. Open a terminal and navigate to the project directory.
3. Run the following command, with the IP address or subnet you want to scan. Example:

```shell
go run scanner.go 127.0.0.1
```

The program will scan the specified IP addresses or subnet and print the list of open TCP ports for each IP address.

```shell
127.0.0.1: Open ports: [5000 5432 7000 44950 44960 59869 63280 63816]
```

## Why?

Mostly because i need something simple for a project i am working on. its worth knowing that port scanning can be illegal or unethical on networks that you do not own/manage. 