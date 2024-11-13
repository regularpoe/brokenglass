# from scapy.all import *

# payload = bytes([0] * 99 + [1])

# packet = IP(dst="localhost") / TCP(dport=4404) / payload
# send(packet)

import socket

sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

server_address = ('localhost', 4404)
sock.connect(server_address)

try:
    # payload = bytes([0] * 99 + [1])
    payload = bytes([0] * 100)
    sock.sendall(payload)
finally:
    sock.close()
