This script enables developers to import recipes to their local MongoDB instance.

### How to run the utility

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