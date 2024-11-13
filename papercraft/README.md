# papercraft

Pull JSON out of thin air.

## Usage

Call script with schema, like so:

```bash
chmod +x papercraft

./papercraft id:number,name:string,roles:array,active:boolean,description:string

{
  "id": 0,
  "name": "",
  "roles": [],
  "active": false,
  "description": ""
}
```

Pipe it to the `jq` and smoke it.
