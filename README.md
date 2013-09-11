# Tome

A consistent and partition tolerance place to put your data structures.

## API

**Note** This API is not only experimental, but not implemented at all. I plan on using this document for planning purposes as I build out tome so I can see the API before I build it.

#### Creating a new structure

Lets create a hash named `ben`.

```bash
curl -X PUT localhost:7777/v1/structure -d '{"type":"hash", "key":"ben"}'
```

#### Adding a key and value to an existing hash

Now that we have our hash `ben` let's add a key `gender` with a value `male`.

```bash
curl -X PUT localhost:7777/v1/structure/ben -d '{"action":"set", "key":"age", "value":"male"}'
```

#### Getting a key from an existing hash

We can now find out what gender `ben` is.

```bash
curl localhost:7777/v1/structure/ben -d '{"action":"get", "key":"gender"}'
# => {"gender": "male"}
```

## Reference

#### Structure Types

* `hash`

#### hash actions

* `set`
* `get`
