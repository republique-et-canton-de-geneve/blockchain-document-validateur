# gethitihteg

Geth Idempotence Toolkit

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

```
golang
```

### Installing

A step by step series of examples that tell you have to get a development env running

Say what the step will be

```
#GOPATH must be set or default to ./
git clone https://github.com/Magicking/gethitihteg.git "${GOPATH:-./}"github.com/Magicking/gethitihteg
```

### Using

 * Deploying contract

```
export PRIVATE_KEY="..."
export RPC_URL="https://ropsten.infura.io"
cd "${GOPATH:-./}"github.com/Magicking/gethitihteg
go run cmd/deployer/deployer.go --rpc-url "${RPC_URL}" --key "${PRIVATE_KEY}" --sol token.sol int:42
```

## License

This project is licensed under the GPLv3 License - see the [LICENSE](LICENSE) file for details
