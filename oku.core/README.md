# oku.core

This is a small snippet engine that allows you to insert some snippets into files.

## Usage

Download script, place on your path and call:

```bash
oku.core <N> <data.json> <language> <snippet_name> <file>

oku.core 12 data.json elixir read_file my_elixir_file.ex
```

```
N            - line number
data.json    - your path to data.json
language     - language | key under which snippet resides
snippet_name - name of the snippet you want to insert
file         - name of the file you are inserting to
```

### data.json

This script now relies on project [bentoi](https://github.com/regularpoe/bentoi) to build data.json file.

Example of 'data.json'

```json
{
    "misc": {
        "foo": "Li9zbmlwcGV0LW1hbmFnZXIgcnVieSByZWFkX2ZpbGUgJ2RlZiByZWFkX2ZpbGUocGF0aCk7IGVuZCc="
    },
    "shell": {
        "sed_copy_lines": "c2VkIC1uICdYOllwJyA8ZmlsZT4gfCB4Y2xpcCAtc2VsIGNsaXA=",
        "socat": "c29jYXQgVENQLUxJU1RFTjo4NDQzLHJldXNlYWRkcixmb3JrIFNZU1RFTTonZWNobyAiSGVsbG8gZnJvbSBzb2NhdCI7IGNhdCc="
    },
    "vim": {
        "add_text_to_each_line": "OiVzLyQvcGF0dGVybi8=",
        "change_until": "Y3Rf",
        "delete_with_pattern": "OmcvcGF0dGVybi9k"
    }
}
```

Running the following would insert value of foo into file called 'tmp.txt' on the third (3.) line.

```bash
oku.core 3 data.json misc foo tmp.txt
```

