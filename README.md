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

    screen [slave-address]      Stream your current screen. Useful for streaming from a browser.


    file [slave-address]        Stream a file. Use for media that is saved locally.

```
