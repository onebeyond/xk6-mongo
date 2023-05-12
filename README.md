# xk6-mongo

This is an extension for performing <a href="https://www.mongodb.com/">MongoDB</a> load testing using <a href="https://k6.io/">K6</a>

## Build

In my case this repository is cloned at `/projects/my/xk6-mongo`

```bash
xk6 build --with github.com/grafana/xk6-redis --replace module=/projects/my/xk6-mongo
```