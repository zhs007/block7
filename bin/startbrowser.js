const chromeLauncher = require('chrome-launcher');
const axios = require('axios');
const fs = require('fs');

(async () => {
  const chrome = await chromeLauncher.launch({
    chromeFlags: [],
  });

  console.log(`http://localhost:${chrome.port}/json/version`);

  const response = await axios.get(
      `http://localhost:${chrome.port}/json/version`
  );
  const {webSocketDebuggerUrl} = response.data;

  console.log(webSocketDebuggerUrl);
  fs.writeFileSync('./startbrowser.txt', webSocketDebuggerUrl);
})();
