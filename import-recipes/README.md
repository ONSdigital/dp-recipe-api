This script enables developers to import recipes to their local MongoDB instance. It assumes that no recipes already exist. If you already have recipes you may get an insert error and need to remove the existing recipes.

### Prerequisites

This script requires the Mongo shell to be installed locally. If you attempt to run the script without MongoDB shell, you will see the error message: `command not found: mongo`

To install the MongoDB shell on a Mac:
```
brew tap mongodb/brew
brew install mongodb-community-shell
```

### How to run the utility

Ensure you are in the `import-recipes` directory:
```
cd import-recipes
```

Run
```
./import-recipes.sh <mongo_url> 
```

The `<mongo_url>` part should look like:
- `mongodb://localhost:27017`
  - if authentication is needed, use:
    `mongodb://<username>:<password>@<host>:<port>?authSource=admin`
    (use single-quotes for protection from your shell)

Full example 

```
./import-recipes.sh mongodb://localhost:27017
```