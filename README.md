# Slack Emoji Grab

This will take the output of grabbing all your slack Emoji urls, saving it to json, and save the resulting images to file.

## Grabbing from Chrome Dev Tools
Adapted from here: https://gist.github.com/dogeared/f8af0c03d96f75c8215731a29faf172c

1. Login to your team through the browser.
1. Go to: https://<team name>.slack.com/customize/emoji
1. Run the snippet on the browser's dev tools javascript console
1. Save resulting JSON string to file

## emojiGrab
Pass the json file name into the go program, along with where you want to save the new emoji files.

```shell
go run emojiGrab.go ~/Downloads/name_of_emoji_file.json ./emojis
```
