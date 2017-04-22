[![Go Report Card](https://goreportcard.com/badge/github.com/DavyJ0nes/qotd)](https://goreportcard.com/report/github.com/DavyJ0nes/qotd)
# Quote of the Day

## Description
- Toy App that displays a quote of the day in a simple web page.
- Gets Quotes from [theysaidso](https://theysaidso.com/api/)
- Tries to adhere to the [twelve-factor-app methodology](https://12factor.net/).
- The deployment artifact is a docker container.
- Caches API request for one day (it is a quote of the day after all ;) ).
- Having the simple cache means that only one request a day is made.

## Deployment
1. Copy env.example to .env and update with relevant values
2. You can deploy using the [deploy](./deploy) script, which automates compiling, build container and deploying to docker host

## Ideas for future
- Could abstract out to allow use with any API
- If set it for any API then would need to allow for authorisation

# License
This package is distributed under the BSD-style license found in the [LICENSE](./LICENSE) file.
