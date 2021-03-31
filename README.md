# CyberChat

## Description

- 采用WebSocket进行通信
- 后端使用Golang编写
- 前端使用React.JS

## Required Tools

1. go or Docker Engine
2. npm

## Use on Docker

- Pull image from the command line:

```bash
docker pull docker.pkg.github.com/fffzlfk/cyberchat/backend:0.1
```

- Use as base image in DockerFile:

```dockerfile
FROM docker.pkg.github.com/fffzlfk/cyberchat/backend:0.1
```

## Frontend

```bash
cd src/frontend/webchat

npm start
```