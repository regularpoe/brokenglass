# Nocturno

nocturno is a file watcher in Elixir, using Jason for JSON, and file_system as a monitor.

## Usage

Clone the repo, build it:

```bash
mix escript.build
```

Create config.json file

```json
{
  "paths": ["/path/to/watch1", "/path/to/watch/2"],
  "notifications": false
}
```

Set notifications to `true` if you want to receive desktop notifications.

Run program:

```bash
./nocturno -c config.json
```