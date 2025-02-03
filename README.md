# VLC-Slave
Combination of a tiny server (running on your Multimedia/Home Theater PC) and a small script for your clients to initiate streams.
This is achieved by first sending a quick authentication request to the server, which then uses the IP from that request to open the stream provided by the client.

The goal is to make it easy to stream shows and movies from your primary device to your TV.

## Usage
Either run the script locally from the same folder or move it to your PATH.
```
❯ ./vlcstream
vlcstream: Stream from this device to a computer running vlcslave.
Available Arguments:

    screen [slave-address]          Stream your current screen. Useful for streaming from a browser.

    file [slave-address] [file]     Stream a file. Use for media that is saved locally.

    stream [slave-address] [url]    Just tell the slave to listen for an already existing stream on the network.

```
## Dependencies
- [VLC](https://www.videolan.org/vlc) for streaming and receiving the stream
- [bash](https://www.gnu.org/software/bash/) for running the client script
- [curl](https://curl.se/) for the request from client to server

## Installation
1. Place the server executable on your server and create a [systemd service](https://www.freedesktop.org/software/systemd/man/latest/systemd.service.html) for it.
2. Copy the `vlcstream` script to your client(s)
3. Ensure your client and server are on the same network
4. Run script with the address of your server