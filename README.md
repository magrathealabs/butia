# butia

[![Coverage Status](https://coveralls.io/repos/github/magrathealabs/butia/badge.svg?branch=master)](https://coveralls.io/github/magrathealabs/butia?branch=master)
[![Build Status](https://travis-ci.org/magrathealabs/butia.svg?branch=master)](https://travis-ci.org/magrathealabs/butia)

Web microframework to create large web applications.

## Features

- Manage environment variables
  - [x] Autoload .env file
  - [ ] Autoload {{ .environment }}.env file to facilitate the testing, staging and production environment
- HTTP Server
  - [x] Setup gin router in release mode
  - [ ] Create a new pre configured gin router
  - [ ] Allow create multiple servers
- Controllers
  - [ ] Define routes as submodules of other controllers.