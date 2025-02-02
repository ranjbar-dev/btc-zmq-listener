# btczmq 

you can run a bitcoin node on prune mode ( use less Desk and it is configable ) and run this service beside bitcoin node, then you can connect to ws server port 8000 ( confiable on /config folder ) and receive new connections that gets into Bitcoin mempool

add you client ip address to give access to connect to ws server( /config/config.yaml file ).

when client connects to ws server, server pushes latest 1000 pending transactions to client.

screenshot
![image](https://github.com/user-attachments/assets/838f6d2f-26e1-45e6-b332-2d4555e3bcb7)

new transction data example: 
```
{
    "c": 1,
    "d": {
        "txid": "596f3ae6bde68fa114289830c5f3204e29b586b29614c6b00ef4c0ebccf9de90",
        "inputs": [
            {
                "txid": "99030fe530a08080b0802c5540b2ab1e5e5caf20786383d0081cfa6915f39076",
                "index": 56,
                "address": "1GNfXR51TvWE8TenQRPVWmKHwUauhyyRF4"
            }
        ],
        "outputs": [
            {
                "address": "1M6gbBVHBxkMWNiTQ8kXgZpHz7nevwH7fR",
                "value": 737686
            }
        ],
        "confirmations": 0
    }
}
```

### Build binary 

`go build -o ./build/main ./cmd/main.go`


### Run binary 

`./build/main /path-to-src/config/config.yaml`

