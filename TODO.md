# TODO

- [ ] client do not allow to perform anything other then registration or login without token
- [ ] store different credentials functionality
- [ ] load different credentials functionality
- [ ] grpc encryption
- [ ] makefile unix, osx, windows
- [ ] login functionality
- [ ] client commandline global flags
- [ ] prevent user from able to store or load before registering
- [ ] force store functionality (overwrite) client (flag) and server (update)
- [ ] client error printer for user
- [ ] log format for client should be plain text and not json
- [ ] password salt must be unique for each user

## In progress

- [ ] server store usecase

## Done

- [X] server auth interceptor
- [X] client store passes auth
- [X] login functionality
- [X] test client login
- [X] server login returns token
- [X] server storage password by id
- [X] server log in use case test
- [X] refactor token builder
- [X] refactor password hashing code
- [X] server log in grpcservice impl
- [X] client log in gateway
- [X] log in proto
- [X] client usecase log in
- [X] client controller login
- [X] client store token after registration
- [X] client restore register storage failure test
- [X] client read token on registration
- [X] client remove credentials storage
- [X] server controller register returns auth token
- [X] server default token lifetime from config
- [X] server token signing key from config
- [X] use same environment variables for client and server
- [X] server password salt from config
- [X] server register usecase returns token
- [X] credentials storage load impl
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
