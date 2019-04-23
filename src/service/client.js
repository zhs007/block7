const messages = require('../../proto/result_pb');
const services = require('../../proto/result_grpc_pb');

const grpc = require('grpc');

const TOKEN = 'wzDkh9h2fhfUVuS9jZ8uVbhV3vC5AWX3';

/**
 * startTranslate
 * @param {string} servAddr - service addr
 * @param {string} srclang - source language
 * @param {string} destlang - destination language
 * @param {string} text - text
 */
function startTranslate(servAddr, srclang, destlang, text) {
  const client = new services.JarvisCrawlerServiceClient(servAddr,
      grpc.credentials.createInsecure());

  const request = new messages.RequestTranslate();
  request.setText(text);
  request.setSrclang(srclang);
  request.setDestlang(destlang);
  request.setToken(TOKEN);

  client.translate(request, function(err, response) {
    if (err) {
      console.log('err:', err);
    }

    if (response) {
      console.log('text:', response.getText());
    }
  });
}

/**
 * startArticle
 * @param {string} servAddr - service addr
 * @param {string} url - url
 * @param {bool} attachJQuery - is attach jquery
 */
function startArticle(servAddr, url, attachJQuery) {
  const client = new services.JarvisCrawlerServiceClient(servAddr,
      grpc.credentials.createInsecure());

  const request = new messages.RequestArticle();
  request.setUrl(url);
  request.setAttachjquery(attachJQuery);
  request.setToken(TOKEN);

  const call = client.exportArticle(request);
  call.on('data', (msg) =>{
    const result = msg.getResult();
    if (result) {
      console.log(result.getTitle());
    } else {
      console.log(msg.getTotallength(), msg.getCurlength());
    }
  });
  call.on('end', ()=>{
    console.log('end.');
  });
  call.on('error', (err)=>{
    console.log('err', err);
  });
}

/**
 * startGetArticles
 * @param {string} servAddr - service addr
 * @param {string} website - website
 */
function startGetArticles(servAddr, website) {
  const client = new services.JarvisCrawlerServiceClient(servAddr,
      grpc.credentials.createInsecure());

  const request = new messages.RequestArticles();
  request.setWebsite(website);
  request.setToken(TOKEN);
  // request.setUrl(url);
  // request.setAttachjquery(jquery);

  client.getArticles(request, function(err, response) {
    if (err) {
      console.log('err:', err);
    }

    if (response) {
      console.log('text:', JSON.stringify(response.getArticles().toObject()));
    }
  });
}

/**
 * startGetDTData
 * @param {string} servAddr - service addr
 * @param {string} mode - mode
 * @param {string} startTime - startTime
 * @param {string} endTime - endTime
 */
function startGetDTData(servAddr, mode, startTime, endTime) {
  const client = new services.JarvisCrawlerServiceClient(servAddr,
      grpc.credentials.createInsecure());

  const request = new messages.RequestDTData();

  request.setMode(mode);
  request.setStarttime(startTime);
  request.setEndtime(endTime);
  request.setToken(TOKEN);

  client.getDTData(request, function(err, response) {
    if (err) {
      console.log('err:', err);
    }

    if (response) {
      console.log('result:', JSON.stringify(response.toObject()));
    }
  });
}

startTranslate('127.0.0.1:7051', 'en', 'zh-CN',
    '@Peter Walker I am sure there is a problem with excel file, I need more time to check it.');

startGetArticles('127.0.0.1:7051', 'baijingapp');
startGetArticles('127.0.0.1:7051', '36kr');
startGetArticles('127.0.0.1:7051', 'geekpark');
startGetArticles('127.0.0.1:7051', 'huxiu');
startGetArticles('127.0.0.1:7051', 'lieyunwang');
startGetArticles('127.0.0.1:7051', 'tmtpost');
startGetArticles('127.0.0.1:7051', 'techcrunch');

// startGetDTData('127.0.0.1:7051', 'gametodaydata', '', '');
// startGetDTData('127.0.0.1:7051', 'gamedatareport', '2019-04-17', '2019-04-17');

startArticle('127.0.0.1:7051', 'https://post.smzdm.com/p/alpzl63o/', true);
