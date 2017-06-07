//  Adapted from here: https://gist.github.com/dogeared/f8af0c03d96f75c8215731a29faf172c

// Login to your team through the browser.
// Go to: https://<team name>.slack.com/customize/emoji
// Run this on the browser's dev tools javascript console
var emojis = $('.emoji_row');
var numEmojis = emojis.size();
var emojiResults = [];

emojis.each(function (index) {
  var url = $(this).find('td:nth-child(1) span').attr('data-original');
  var extension = url.substring(url.lastIndexOf('.'));
  var name = $(this).find('td:nth-child(2)').html().replace(/:|\s/g, '');
  var obj = {url: url, extension: extension, name: name};
  emojiResults.push(obj);
});

console.log(JSON.stringify({emojis: emojiResults}));
