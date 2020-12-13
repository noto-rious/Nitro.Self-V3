# Nitro.Self V3
[![GitHub release](https://img.shields.io/github/v/release/noto-rious/Nitro.Self-V3?style=plastic)](https://github.com/noto-rious/Nitro.Self-V3/releases) [![GitHub All Releases](https://img.shields.io/github/downloads/noto-rious/Nitro.Self-V3/total?style=plastic)](https://github.com/noto-rious/Nitro.Self-V3/releases)

Multi-Account Discord Nitro sniper and Giveaway joiner written in Go 1.15.2.

I have to give <a href="https://github.com/Vedza">@Vedza</a> some credit as I used his nitro sniper as a base.  
Multi-Account threading along with code-caching for dupe protection and a few other nick-nacks are courtesy of me.

If you need any help or have any suggestions you can chekout my main profile for contact details.

### Features 
* Supports multiple accounts (optional).
* Redeems all codes on the main account.
* Uses Fast HTTP package for Go (faster than Python requests).
* Cooldown for # hour(s) after redeeming # nitro code(s).
* Optional Nitro Giveway joiner.
* DM host with custom DM message if giveaway won.
* Removes some gift-code obfuscation.
* Fake/duplicate code detection to avoid being banned.
* Webhook support with extended features like 'ping user' and 'report failed codes/giveaway entries'.
* Webhook returns color based on success or failure (green/red)
* Cross-platform binaries for Linux and Windows.


![Screenshot](screenshot.png)

### Download & Build with Go
```
 go get github.com/noto-rious/Nitro.Self-V3
 go mod download
 go build
 ./Nitro.Self-V3
 ```
### Configure
Edit `settings.json`
```
{
  "token": "X",                            // Replace X with your main token.
  "nitro_max": 2,                          // Maximum Nitro code redeems allowed before cooldown
  "cooldown": 24,                          // How many hours to cooldown for.
  "giveaway_sniper": true,                 // Enable(true) or Disable(false) giveaway sniping or not.

  "giveaway_delay": {
    "minimum": 60,                         // Configure minimum seconds that the account will wait before entering a giveaway.
    "maximum": 120                         // Configure maximum seconds that the account will wait before entering a giveaway.
  },

  "snipe_on_main": true,                   // Enable(true) or Disable(false) sniping on the main account or not.
  "dm_host": true,                         // Enable(true) or Disable(false) the option to DM the giveaway host if you win.
  "dm_message": "hi, i won your giveaway!" // Custom DM Message.
  "webhook_url": "",                       // this is optional, if you're not sure what goes here then you don't need it.
  "webhook_ping_id": "",                   // this is also optional, this value would be your numerical user id(obtained by enabling developer mode in settings).
  "report_fails_to_webhook": false,        // you can set this to true if you want to webhook log failed events.
  "save_cache": true                       // Allows for permanent code caching.
}
```
Edit `tokens.txt`
```
NDI4Mjc31DExNzZyNTQ1NTQ2.X // Token #1
NzYxORQyMDkwNtU1NjA1PDEz.X // Token #2
NzYxODF1OikyNDEyOTE1NzKz.X // Token #2 - Add as many or as little as you want
```
***
### How to obtain your token
**1.** Press **Ctrl+Shift+I** (⌘⌥I on Mac) on Discord to show developer tools<br/>
**2.** Navigate to the **Application** tab<br/>
**3.** Select **Local Storage** > **https://discordapp.com** on the left<br/>
**4.** Press **Ctrl+R** (⌘R) to reload<br/>
**5.** Find **token** at the bottom and copy the value<br/>
***
### Disclaimer
This is a self-bot which goes against the Discord ToS agreement. Use it at your own risk.
