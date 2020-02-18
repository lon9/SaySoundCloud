# SaySoundCloud

SaySoundCloud is a framework to build a cloud native SaySound service for many games and applications.

<div align="center">
<img src="./art/saysoundcloud.png" alt="desc" title="SaySoundCloud">
</div>

## Features
* **Scalable**, it provides sounds for not only several game servers or chat services but something can treat http protocol.
* **Client side plugin free**, user just need a browser to browse it, so no need specific client side plugins.
* **No downloading sounds**, because the sounds is stored on cloud storage, users no need to download them.

## Usage

1. Prepare sounds and execute tools for make database of sounds. (see tools directory)
1. Make `.env` file.
1. Run `docker-compose up -d` to start service.


## Contribution

The framework supports multi languages, but I can use just English and Japanese. So please add more languages.
Also, please contribute to develop plugins for other applications and games.
