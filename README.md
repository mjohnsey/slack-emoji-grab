# Slack Emoji Grab

This will take the output of grabbing all your slack Emoji urls, saving it to json, and save the resulting images to file.

## Grabbing from Chrome Dev Tools
Adapted from here: https://gist.github.com/lmarkus/8722f56baf8c47045621

1. Login to your team through the browser.
1. Go to: `https://<team name>.slack.com/customize/emoji`
1. Open browser's dev tools javascript console, look for `emoji.adminList` in the Network tab
1. Either copy the response JSON object or copy the request as curl to save to file named `emojis.json`

## emojiGrab
Pass the json file name into the go program, along with where you want to save the new emoji files.

```shell
go run emojiGrab.go ~/Downloads/name_of_emoji_file.json ./emojis
```
