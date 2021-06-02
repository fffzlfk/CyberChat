<h1 align="center">Welcome to CyberChat 👋</h1>
<p>
  <a href="https://twitter.com/fffzlfk" target="_blank">
    <img alt="Twitter: fffzlfk" src="https://img.shields.io/twitter/follow/fffzlfk.svg?style=social" />
  </a>
</p>

> Go实现的基于WebSocket的实时聊天室

## Description

- 采用WebSocket进行通信
- 后端使用Golang编写
- 前端使用React.JS

## Required Tools

1. go or Docker Engine
2. npm

## Install Backend With Docker

### Pull image from the command line

```sh
docker pull docker.pkg.github.com/fffzlfk/cyberchat/backend:0.1
```

### Use as base image in DockerFile

```sh
FROM docker.pkg.github.com/fffzlfk/cyberchat/backend:0.1
```

## Usage

### Backend

```sh
cd src/backend

go run main.go [-p $port]   # -p为可选参数 用来指定端口号，默认为8080
```

### Frontend

```sh
cd src/frontend/webchat

npm start
```

## Author

👤 **fffzlfk**

* Website: https://fffzlfk.netlify.app/
* Twitter: [@fffzlfk](https://twitter.com/fffzlfk)
* Github: [@fffzlfk](https://github.com/fffzlfk)

## 🤝 Contributing

Contributions, issues and feature requests are welcome!<br />Feel free to check [issues page](https://github.com/fffzlfk/CyberChat/issues). 

## Show your support

Give a ⭐️ if this project helped you!

***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_