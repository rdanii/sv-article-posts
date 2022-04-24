# sv-article-posts

## Quick Start

- Clone the repo: git clone https://github.com/rdanii/sv-article-posts.git

## Instalation

```bash
$ npm install
```

or

```bash
$ yarn install
```

for Server Side

```
$ go get -u ./...
```

## Basic usage

```bash
# dev server with hot reload at http://localhost:3000
$ npm start

# if you use Node 17+ use this command instead of `npm start`
$ npm run start:n17
```

or

```bash
# dev server with hot reload at http://localhost:3000
$ yarn start

# if you use Node 17+ use this command instead of `yarn start`
$ yarn start:n17
```

for Server side

```
$ go run .
```

Navigate to [http://localhost:3000](http://localhost:3000). The app will automatically reload if you change any of the source files.

Server running on http://localhost:8000

Set your database equal to database in yours localhost

### Build

Run `build` to build the project. The build artifacts will be stored in the `build/` directory.

```bash
# build for production with minification
$ npm run build

# if you use Node 17+ use this command instead of `build run build`
$ npm run build:n17
```

or

```bash
# build for production with minification
$ yarn build

# if you use Node 17+ use this command instead of `yarn build`
$ yarn build:n17
```

## What's included

Within the download you'll find the following directories and files, logically grouping common assets and providing both compiled and minified variations. You'll see something like this:

```
coreui-free-react-admin-template
├── public/          # static files
│   └── index.html   # html template
│
├── src/             # project root
│   ├── assets/      # images, icons, etc.
│   ├── components/  # common components - header, footer, sidebar, etc.
│   ├── layouts/     # layout containers
│   ├── scss/        # scss styles
│   ├── views/       # application views
│   ├── _nav.js      # sidebar navigation config
│   ├── App.js
│   ├── ...
│   ├── index.js
│   ├── routes.js    # routes config
│   └── store.js     # template state example
│
└── package.json
```

## Documentation

The documentation for the CoreUI Admin Template is hosted at our website [CoreUI for React](https://coreui.io/react/)
