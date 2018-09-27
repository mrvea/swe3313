# Pizza

Herein lies the code for Pizza.

### Directory Structure
This repository should exist inside of a $GOPATH, and follow golang's conventions for structure inside of
that $GOPATH.  This means that cmd/go has a place to build and install your app and it's dependencies.
For example:

```
$ echo $GOPATH
/some/other/gopath:/path/to/project/pizza
$ tree /path/to/project/pizza
/path/to/project/pizza/
├── bin
├── pkg
└── src
    └── github.com
        └── project
            └── pizza
                ...
```

### Managing Dependencies
External dependencies are now vendored and already available.  If you need/want to update, [use govendor.](https://github.com/kardianos/govendor)

_You **must** vendor any new dependencies before running the app in the container:_

`govendor add +external`

### Dev Environment Setup

#### Docker

* [Install Docker.](https://docs.docker.com/engine/installation/)
* Read the intro documentation.  Run the quickstart terminal, or however you want to do that.
* Get the ip of your docker machine with `docker-machine ip` if using docker-machine.
* [Install docker-compose.](https://docs.docker.com/compose/install/)
* Run `docker-compose up` from the root of the project.
* Find the application at the ip address obtained previously. On linux, this should be localhost.
* After making changes that require the go application to be restarted: `util/restart`

### Migrate the database
#### Run all migrations
`util/run migrate`

This will bring the database up to the current version, from zero.  You should probably do this any time you pull down new code, or it should be your first troubleshooting errors after pulling down new code.

#### Add New Migrations
Migrations should be placed in the directory `schema/migrations`
The **must** be named according to the following format:

`{{id}}_{{multi_word_name}}_{{"up" or "down"}}`

Migrations are simply `.sql` files that contain instructions for performing the migration.
"Up" files move the database forward in time.
"Down" files revert the changes that the up files make.

For example:
The starting point for our schema will contain the entire starting structure as an SQL script.  The script that does this will be called `1_initial_schema_up.sql`.  There will be a counterpart called `1_initial_schema_down.sql`.  This script will drop the database and create it again.

### Running One-Off Commands (pizza cli)
#### `util/dropdb`
Destroy your database.  Careful - it's gone forever!  Consider making an empty database by combo-ing with `util/run migrate`

#### `util/makeseed $filename`
Dump data (not structure) into `schema/seeds/$filename.sql`  This could be useful to create some seed data that can be shared by the team.  Seed data can be loaded using:

#### `util/loadseed $filename`
Execute the SQL script at `schema/seeds/$filename.sql` against the existing database.  This may not succeed if the new data conflicts with existing data!

#### The `pizza` binary exposes a command line interface for certain actions.  These can be run on a local dev machine, or in the dev container using the `util` wrapper.

##### Use `util/run {{options}} {{command}}` to use the pizza cli.
Example (rollback 2 most recent db migrations):

    `util/run -rollback -n=2 migrate`

##### `migrate`
Run any migrations that haven't been run against the database: `util/run migrate`

Options:

* rollback (bool) If this option is present, the most recent migration will be rolled back.
* n (int) This option controls how many recent migrations will be rolled back.  Ignored otherwise. Default is (1).

##### `test`
Print out the configuration of the app, including the app environment and the database connection string being used.  That's all it does right now.


##### `start`
Start serving the web application on the configured port of localhost.

##### `stop`
TODO

##### `restart`
TODO

##### `reload`
TODO

##### `hello`
Print "Hello World." in your terminal.  Exceptionally useful.


### Run tests.
`util/test [args]`
This starts new ephemeral docker containers holding a test database and the application and executes `go test` passing any arguments as well.

Example: `util/test -v ./handler` executes the tests just as though you had run `go test -v ./handler`

### Nuke the Test Database
`util/clear_testdb`
Your tests probably have some effect on the database.  For the sake of development speed, the database is not nuked every time you run util/test, allowing you to quickly run tests while developing them.  If you need a reset, run this command.  You should definitely do this for a clean test before merging to a major shared branch.


### Load the testing database into your development environment.
`util/load_testdb`
When you update your code, you sometimes need to update the data that will be used for testing.  You can load the test database into your dev environment with this command.  This will allow you (a) migrate your test database up to the current version and (b) insert any new records that your tests might rely upon.


### Update the testing environment.   *** CAUTION: KNOW WHAT YOU ARE DOING ***
`util/overwrite_testdb`
This dumps the contents of your development environment to be used **in all tests ever**. Be careful what ends up in the testing data.  **NEVER** do this unless you have first run `util/load_testdb` to get the existing test database and go from there.


### Save your data to a seed file.
`util/makeseed`

`util/makeseed ${filename}`

This command dumps the data from your database into `schema/seed.sql` or the specified filename.
This file is ignored by source control.  Use this if you want to keep some data in
your local development database only, eg. a user(s) and associated data.
It *will* overwrite the existing file.


### Build and run the production application in your dev environment
`util/prodtest`
This command builds the application for production, shuts down your dev environment, and starts
the app in production mode.

## Testing
Tests for all API endpoints should be located in handler/handler_test.go.

### Writing Route Tests
To test a route, start by setting up a RouteTest (TODO link to type def) and defining the required fields.  Adding optional fields affects the types of checks that performed automatically on the response:

* `ExpectedCode` - Populating this fields asserts that the handler responds with the correct status.
* `ExpectedResponse` - Populating this field with **a path to a file containing expected response data** asserts that the response matches the contents of the file **exactly**.
* `ExpectedResponseFragment` - Populating this field with a string asserts that the response contains the given string.
* `ExpectedResponseFragments` - Populating this field with a slice of strings asserts that the response contains each string in the slice.

Once the route test has run, you may perform any additonal tests of the expected conditions, such as confirming that a record was inserted or deleted successfully.  The response is available at `RouteTest.RR` if needed.

### Updating the Test Database
If you need to add data to the test database, you should follow these guidelines - if the order is not followed, and proper defaults given to new fields, test data from previous versions may not be inserted properly and will have to be fixed manually.

1. New migrations should be built and saved in your development environment.  Fully tested migrations should be applied to the test database in step 5 of this process.
2. When adding `NOT NULL` columns to the schema, always assign a default value.  This ensures that old data can be migrated.
3. The testing database should only contain very specific records that meet the needs of a test case.
4. Before inserting new test data into the database **always** flush any records from your development database by running `util/load_testdb`.  This keeps records created in the development database out of the testing database. Be aware that data in your development environment will be gone forever, unless you `util/makeseed` first.
5. Immediately after `util/load_testdb` you should `util/run migrate`.  This brings the test database up to the latest version, and hopefully sets intelligent defaults to records that already exist.
6. At this point you can insert/delete/update any records required for your test case(s).
7. Once you have achieved the desired state of the test database, cast `util/overwrite_testdb` to write it in stone for future generations of testers.
8. If desired, use `util/dropdb`, `util/run migrate`, and `util/loadseed` to get your development database to where you want it to be.


## Build Script Usage

* `util/build`  -- defaults to linux
* `util/build linux`
* `util/build windows`
* `util/build darwin`

Architecture is specified as amd64 for now, and cannot be altered by command line args.
This script compiles the go application, and copies files required to run the app in production
to the dist folder.


## Deployment (mac/linux only for now)

* `util/deploy {environment}` eg. `util/deploy staging` or `util/deploy production`


### Notes on Deployment
Deployment also executes the production build - there is no requirement to build before deploy.

Deployment does not clean previous deployments currently.  Existing files are overwritten and new files are added.  If files move or are renamed, they will not be deleted from the production server.
# swe3313
