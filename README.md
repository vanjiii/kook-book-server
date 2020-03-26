# kook-book-server
Simple recipe book

## Installation

## Quick start

## Documentation

Technologies used
- go
- postgres
- openapi
- implement gql and/or rest
- make
- docker images

Blogposts ideas:
- Structure - what, where, why
- Storing passwords:
  - choosing to encrypt server level(remove server DB connection exploit; remove load of the DB - that encryption process should happen on the server)
  - bcrypt, argon2

TODO:
- [ ] Create user
- [ ] Each recipe should have photo, ingredients, description, time for prep
- [ ] login (use JWT and session cookie)
- [ ] Add avatar
- [ ] login with FB/Twitter, etc (this may be client feature)
- [ ] Add recipe
- [ ] Get recipe
- [ ] Edit recipe
- [ ] Delete recipe
- [ ] Maybe raking of difficulty of preparation
- [ ] List All (search option)
- [ ] List my recipes
- [ ] Advanced GQL
  - [ ] connections
  - [ ] unions
  - [ ] interfaces
  - [ ] enums
  - [ ] APQ
  - [ ] Extract that gql to main (Playground, /query)
- [ ] Extract recipe - by input (usually url) extract the recipe if any
- [ ] Comment on recipe
- [ ] Upvote recipe (Maybe add ranking)
- [ ] Categories (aka tags)
- [ ] Find recipe by ingredients
- [ ] Upload avatar to the recipe
- [ ] cmd command
- [ ] Sending email (for whatever reason)
- [ ] Validate models and rest endpoints
- [ ] Integrate gRPC module
- [ ] Dataloader for gql
- [ ] Repository cache 
- [ ] Stress test
- [ ] Instrumentation test
- [ ] Integration test
- [ ] Regression test
- [ ] Sanitize input fields
- [ ] Make the Graphiql libs versions variables
- [ ] destroy DB on t.Defer()
- [ ] update golang

Nice resource: 
https://github.com/oshalygin/gqlgen-pg-todo-example/
