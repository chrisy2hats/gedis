import json
import logging
import socket

logger = logging.getLogger()
logging.basicConfig(level=logging.INFO, format='%(message)s')

HOST = "localhost"
PORT = 8080

with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as sock:
    sock.connect((HOST, PORT))
    logger.info(f"Connected to {HOST}:{PORT}")
    while line := input(">"):
        if line.strip() == "exit":
            logger.info(f"Goodbye!")
            break

        encoded = (line + "\n").encode()
        sock.sendall(encoded)
        result = sock.recv(4096)
        jsn = json.loads(result)
        if jsn.get("successful", False) is False:
            logging.error(f"Error in previous command: {jsn.get('error')}")
            continue

        if 'result' in jsn:
            logger.info(f"{jsn['result']}")
