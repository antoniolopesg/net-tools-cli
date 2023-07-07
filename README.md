# net-tools-cli

Net-Tools-CLI is a Go-based command-line application that serves as my first foray into CLI development. It offers two simple yet handy network tool functionalities. The first allows users to retrieve the IP address associated with a given host, providing quick and easy access to this information. The second functionality enables users to query the nameserver associated with a specific host, facilitating DNS-related investigations. Net-Tools-CLI's minimalist design and user-friendly interface make it a convenient choice for performing these network-related tasks effortlessly.

## Usage

This is not a production-ready software so you need the go binary to run the application

### Features

- Get Addresses given a host:<br/>
`go run main.go ip --host google.com`

- Get name server given a host:<br/>
`go run main.go servers --host google.com`