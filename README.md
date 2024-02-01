# Certinfo - Certificate Checker

## Overview

This is a simple command-line interface (CLI) program written in Go for checking PEM file certificates. It helps you validate and retrieve information about X.509 certificates from PEM-encoded files.

## Features

- Check the validity of X.509 certificates in PEM files.
- Display information about the certificate, such as subject, issuer, validity period, and more.

## Prerequisites

Make sure you have Go installed on your machine.

## Installation

```bash
go get -u github.com/marco-introini/certinfo-go
```

## Usage

```bash
certinfo path/to/certificate.pem
```

Replace `path/to/certificate.pem` with the path to your PEM-encoded certificate file.

### Options

None, for now

## Example

```bash
certinfo /path/to/certificate.pem
```

## Build from Source

```bash
git clone https://github.com/marco-introini/certinfo-go.git
cd certinfo-go
go build
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments


## Contribution

Contributions are welcome! Please open an issue or create a pull request with your suggestions and improvements.

## Contact

For any questions or feedback, feel free to reach out to [marco@mintdev.me](mailto:marco@mintdev.me).

Happy coding!