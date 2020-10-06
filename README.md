# Nitro.Self V3

[![GitHub release](https://img.shields.io/github/v/release/noto-rious/Nitro.Self-V3)](https://github.com/noto-rious/Nitro.Self-V3/releases)

Multi-Account Discord Nitro sniper and Giveaway joiner written in Go.

Main Snipe functionality credit goes to ![@Vedzaa](https://github.com/Vedzaa).  
Multi-Account threading along with code-caching for dupe protection and a few other nick-nacks are courtesy of me.

### Features 
* Supports multiple accounts.
* Cooldown settings.
* Optional giveway joiner.
* DM host with custom DM message if giveaway won.
* Removes some code obfuscation.
* Fake/duplicate code detection to avoid being banned.


Might look into adding webhook support later.

![Screenshot](screenshot.png)

### Usage
Edit `settings.json`
```
{
  "token": "", // Your main token here
  "nitro_max": 2, // Maximum Nitro code redeems allowed before cooldown
  "cooldown": 24, // How many hours to cooldown for.
  "giveaway_sniper": true, // Enable giveaway sniping or not.
  "snipe_on_main": true, // Enable sniping on the main account or not.
  "dm_host": true, // Enable the option to DM the giveaway host if you win.
  "dm_message": "hi, i won your giveaway!" // Custom DM Message.
}
```
Edit `tokens.txt`
```
NDI4Mjc31DExNzZyNTQ1NTQ2.X // Token #1
NzYxORQyMDkwNtU1NjA1PDEz.X // Token #2
NzYxODF1OikyNDEyOTE1NzKz.X // Token #2 - Add as many as you want
```

```
 go get https://gopkg.in/noto-rious/Nitro.Self-V3.v3
 go mod download
 go build
 ./Nitro.Self-V3
 ```
***
### How to obtain your token
https://github.com/Tyrrrz/DiscordChatExporter/wiki/Obtaining-Token-and-Channel-IDs#how-to-get-a-user-token
***
### Disclaimer
This is a self-bot which is against Discord ToS. Use it at your own risk.
