# xk6-mongo

This is an extension for performing <a href="https://www.mongodb.com/">MongoDB</a> load testing using <a href="https://k6.io/">k6</a>

## Build

```bash
xk6 build --with github.com/Bounteous17/xk6-mongo
```

## Development

In my case this repository is cloned at `/projects/my/xk6-mongo`

This would allow you to import the local development version on your k6 Javascript tests by using:

```bash
xk6 build --with github.com/Bounteous17/xk6-mongo=/projects/my/xk6-mongo
```

You can now test your local changes by importing the new module build:
```js
import xk6_mongo from 'k6/x/mongo';
```