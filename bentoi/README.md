# bentoi

## Build

```bash
go build -o bentoi
```

## Usage

```bash
Usage:
  Add snippet:           bentoi <language> <snippet_name> <snippet_content>
  Add encoded snippet:   bentoi -encoded <language> <snippet_name> <base64_content>
  Export to JSON:        bentoi -export <filename.json>
  Export encoded:        bentoi -export <filename.json> -export-encoded
  List all snippets:     bentoi -list

Example:
  bentoi ruby read_file 'def read_file(path); end'
  bentoi -encoded python write_file 'ZGVmIHdyaXRlX2ZpbGUocGF0aCwgY29udGVudCk6Cg=='
```

