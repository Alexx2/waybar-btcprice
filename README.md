Very simple Bitcoin price indicator for waybar

Build the go file with "go build pricefetch.go"
You must have Go installed in your computer. 

Example config to add to the waybar config file:

"modules-center": ["custom/bitcoinprice"]
"custom/bitcoinprice": {"exec": ~/pricefetch", "interval: 30", "format": "{}"}
