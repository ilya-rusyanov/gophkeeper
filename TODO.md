# TODO

- [ ] server store usecase
- [ ] load different credentials functionality
- [ ] grpc encryption
- [ ] makefile unix, osx, windows
- [ ] server password salt from config
- [ ] login functionality
- [ ] client commandline global flags
- [ ] prevent user from able to store or load before registering
- [ ] use same environment variables for client and server
- [ ] force store functionality (overwrite) client (flag) and server (update)
- [ ] client error printer for user
- [ ] log format for client should be plain text and not json

## In progress

- [ ] store different credentials functionality
- [ ] client gw store impl

## Done

- [X] client store usecase
- [X] refactor credentials with payload and type
- [X] rename credentials to record
- [X] test client controller store auth impl
- [x] test client controller
- [X] client must output help even if env variable is not set
- [X] test server store controller
- [X] store proto method
- [X] rename register argument in proto
- [X] proto go generate
- [X] server users repo store impl
- [X] server database
- [X] server first migration
- [X] store hashed user password
- [X] client MyCredentials instead of passing login/password
- [X] makefile
- [X] server controller handling user already exists test
- [X] server register user already exists test
- [X] server register usecase
- [X] test server register usecase
- [X] retest client register usecase
- [X] make mocks makefile
- [X] move mock generation to test files
- [X] client server gateway register
- [X] client usercred store impl
- [X] client controller register impl
- [X] server grpc service
- [X] server stub
- [X] design server
- [X] client register usecase
- [X] client register usecase test
- [X] client stub is buildable
