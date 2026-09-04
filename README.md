## This project will be rewrite
See https://github.com/lantongxue/rustdesk-api-server-pro/issues/30

Rustdesk Api Server Pro
============

[English](https://github.com/rustdesk/rustdesk) | [简体中文](https://github.com/lantongxue/rustdesk-api-server-pro/blob/master/README_CN.md)

This is an open source Api server based on the open source [RustDesk](https://github.com/rustdesk/rustdesk) client, the implementation of the client all Api interfaces, and provides a Web-UI for the management of data.

![Dashboard](./img/1.jpeg "Dashboard")

> We strive to achieve functionality with the simplest possible code and structure!

## Special Sponsor

CDN acceleration and security protection for his project are sponsored by Tencent EdgeOne.

<a href="https://edgeone.ai/?from=github" target="_blank">Best Asian CDN, Edge, and Secure Solutions - Tencent EdgeOne</a>

<a href="https://edgeone.ai/?from=github" target="_blank">![edgeone](https://edgeone.ai/media/34fe3a45-492d-4ea4-ae5d-ea1087ca7b4b.png)</a>

## Features

- Synchronized RuskDesk version (Currently adapted client: 1.4.6)
- Pure Go implementation of all interfaces
- Visual management interface
  - Internationalization support
  - Statistics panel
  - User Management
  - 2FA & Email Verify Code
  - Session Management
  - Log Audit
- Lightweight & Cross Platform
  - Minimal sqlite
  - Support for major operating systems and architectures

## Compatibility Statement (RustDesk 1.4.6)

- Target client baseline: `1.4.6`
- Covered in this adaptation:
  - Heartbeat/sysinfo payload compatibility
  - Version capability gate (`translate_mode` enabled at `>=1.4.6`)
  - Auth payload compatibility (strict required fields, tolerant unknown fields)
  - `rustdesk install --version` supports both `1.4.6` and `Branch_1.4.6`
- Verification commands:
  - `cd backend && go test ./...`
  - `cd soybean-admin && pnpm typecheck && pnpm lint && pnpm build`

## Playwright E2E (Full-stack)

- Covered cases: `login`, `devices`, `users`, `audit`
- E2E test files are under `soybean-admin/tests/e2e`

### Prerequisites

1. Start backend API and create admin user:

```shell
cd backend
go run . sync
go run . user add admin admin123456 --admin
E2E_SKIP_CAPTCHA=true go run . start
```

2. Install frontend dependencies and Playwright browser:

```shell
cd soybean-admin
pnpm i
npx playwright install chromium
```

### Run tests

```shell
cd soybean-admin
E2E_ADMIN_USER=admin E2E_ADMIN_PASS=admin123456 pnpm test:e2e
```

### CI

- `build-release.yml` supports optional full-stack Playwright E2E.
- Trigger `workflow_dispatch` with `run_playwright_e2e=true`.

## Deploying with Docker

The included Compose file builds this fork, runs it as an unprivileged user, stores state under `./data`, reads the initial administrator password from a file, and binds the API to loopback by default.

```shell
mkdir -p data secrets
read -rsp "Initial administrator password: " ADMIN_PASS
printf '%s' "$ADMIN_PASS" > secrets/admin_password
unset ADMIN_PASS
chmod 600 secrets/admin_password
sudo chown -R 10001:10001 data secrets/admin_password
ADMIN_USER=admin docker compose up -d --build
```

The password file is read only during first-time initialization. The generated `data/server.yaml` contains a random signing key and is created with mode `0600`.

The default listener is `127.0.0.1:8080`; put an HTTPS reverse proxy in front of it. To use another loopback port, stop the container, edit `httpConfig.port` in `data/server.yaml`, and start it again.

### Environment variables

| Variable | Default | Description |
|:--|:--|:--|
| `ADMIN_USER` | required on first boot | Initial administrator username |
| `ADMIN_PASS_FILE` | `/run/secrets/admin_password` in Compose | File containing the initial administrator password |
| `TZ` | `UTC` | Container timezone; keep it aligned with `db.timeZone` |

## Build from source

### Required

- Golang >= 1.21.4
- NodeJs ~= latest(recommend LTS)version
- pnpm ~= latest

### Build

1. Get source code

```shell
git clone https://github.com/lantongxue/rustdesk-api-server-pro.git
```

2. Build the api-server

```shell
cd backend && go build
```

3. Build the frontend
   
```shell
cd soybean-admin && pnpm i && pnpm build
```

### Run

#### api-server

Assuming the compiled binary file is called `rustdesk-api-server-pro.exe`.

1. Synchronize the database table structure
   
```shell
rustdesk-api-server-pro.exe sync
```

2. Add your first user
   
```shell
rustdesk-api-server-pro.exe user add admin yourpassword --admin
```

> --admin is optional, when enabled the added user is an administrator user, otherwise it is a regular user

3. Start the server
   
```shell
rustdesk-api-server-pro.exe start
```

> Listening on port `8080` by default

#### Web Management Interface

For this step you need a web server software (e.g. nginx, apache, etc.), by copying the packaged product to the web root directory.

Typically, the packaged product is in the `soybean-admin/dist` directory.

Reverse Proxy Configuration, you need to configure reverse proxy in `nginx` or other WEB servers, through the reverse proxy server can access the interface address correctly.

Here's my backend reverse proxy configuration for you to refer to:

```nginx
#PROXY-START /api for rustdesk client
location ^~ /api
{
    proxy_pass http://127.0.0.1:8080;
    proxy_set_header Host 127.0.0.1;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header REMOTE-HOST $remote_addr;
    proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection $connection_upgrade;
    proxy_http_version 1.1;
    # proxy_hide_header Upgrade;

    add_header X-Cache $upstream_cache_status;
}
#PROXY-END/

#PROXY-START /admin for web-ui
location ^~ /admin
{
    proxy_pass http://127.0.0.1:8080/admin;
    proxy_set_header Host 127.0.0.1;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header REMOTE-HOST $remote_addr;
    proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection $connection_upgrade;
    proxy_http_version 1.1;
    # proxy_hide_header Upgrade;

    add_header X-Cache $upstream_cache_status;
}
#PROXY-END/
```

## CLI help

```shell
Usage:
  rustdesk-api-server-pro [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  rustdesk    About rustdesk-server command
  start       Start the api-server
  sync        The api-server database synchronization
  user        User management

Flags:
  -h, --help   help for rustdesk-api-server-pro

Use "rustdesk-api-server-pro [command] --help" for more information about a command.
```

## Follow-up plan

We will continue to follow up the RustDesk client and implement the corresponding interfaces, which will be a long-term plan.

## Sponsorship

If you found this project helpful, why not buy the developers a cup of coffee :)

![Sponsorship](./soybean-admin/src/assets/imgs/sponsorships.png "Sponsorship")

**Thank you for your sponsorship**

## License

This fork is distributed under the [GNU Affero General Public License v3.0](https://github.com/wiredmind/rustdesk-api-server-pro/blob/master/LICENSE).
