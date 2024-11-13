# nokku

## Prepare python environment

```
python3 -m venv .venv

source .venv/bin/activate

python3 -m pip install scapy
```

## Run packet.py

```bash
sudo .venv/bin/python3 packet.py
```

### Alternative way?:

```bash
sudo setcap cap_net_raw+eip $(which python3)
```

## Run service

```bash
# Run in place or build
./nokku
```

## Send package

### Using nc to send payload(s)

```bash
printf '\x00%.0s' {1..99} > payload.bin && printf '\x01' >> payload.bin

nc 127.0.0.1 4404 < payload.bin
```

## Build static binary

```bash
CGO_ENABLED=0 GOOS=linux go build -a -o nokku
```

