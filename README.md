Redis seeder
============

[![Build Status](https://travis-ci.com/obukhov/redis-seeder.svg?branch=master)](https://travis-ci.com/obukhov/redis-seeder)
[![Coverage Status](https://coveralls.io/repos/github/obukhov/redis-seeder/badge.svg?branch=main)](https://coveralls.io/github/obukhov/redis-seeder?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/obukhov/redis-seeder)](https://goreportcard.com/report/github.com/obukhov/redis-seeder)

Cli tool to seed random data to a redis instance

## Usage

```bash
redis-seeder seed host:port
```

For now to change strategy of seeding you have to edit cmd/seed.go. In the future seeding plan can be laoded from yaml.
