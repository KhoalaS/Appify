# Appify

Ionic ordered on Temu.

- go 1.24.2+

## Building

```bash
go build -o=build/appify main.go
```

## Usage

Generate a example config and userscripts:

```bash
appify scaffold [--typescript]
```

Generate a project:

```bash
appify generate -c ./config.json
```

Build and install the app:

```bash
appify build -c ./config.json [--type "release"|"debug"]
```

Example config:

```json
{
  "appName": "MyApp",
  "website": "https://example.com",
  "blockedHosts": ["blocked.site"],
  "packageName": "com.example.app",
  "userAgentString": "Mozilla/5.0 (Linux; Android 8.0.0; SM-G955U Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/140.0.0.0 Mobile Safari/537.36",
  "sslBypass": ["localhost", "lan-host"],
  "globals": "./userscripts/globals/globals.js",
  "onloadScripts": "./userscripts/onload",
  "projectDirectory": "./myapp-appify"
}
```

## TODO

- Watch mode that copies userscripts into project
