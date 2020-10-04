# Nitro.Self V3

[![GitHub release](https://img.shields.io/github/v/release/noto-rious/Nitro.Self-V3)](https://github.com/noto-rious/Nitro.Self-V3/releases)

Multi-Account Discord Nitro sniper and Giveaway joiner written in Go.

Main Snipe functionality credit goes to ![@Vedzaa](https://github.com/Vedzaa).  
Multi-Account threading along with code-caching for dupe protection and a few other nick-nacks are courtesy of me.

It also sends a DM to giveaway host when won.

![Screenshot](screenshot.png)

### Usage

Edit `settings.json`
```
{
  "token": "", // Your main token here
  "nitro_max": 2, // Max Nitro codes redeemed before cooldown
  "cooldown": 24, // How many hours to cooldown for.
  "giveaway_sniper": true // Enable giveaway sniping or not.
  "snipe_on_main": true // Enable sniping on the main account or not.
}
```
Edit `tokens.txt`
```
token1
token2
token3
```

```
 go mod download
 go build
 ./Nitro.Self-V3
 ```
 
### How to obtain your token
https://github.com/Tyrrrz/DiscordChatExporter/wiki/Obtaining-Token-and-Channel-IDs#how-to-get-a-user-token

### Disclaimer
This is a self-bot which is against Discord ToS. Use it at your own risks.
