const messages = require('../../proto/result_pb');
const services = require('../../proto/result_grpc_pb');
const {requestCrawler} = require('./utils');
const log = require('../log');

const grpc = require('grpc');

const TOKEN = 'wzDkh9h2fhfUVuS9jZ8uVbhV3vC5AWX3';

/**
 * startTranslate2
 * @param {string} servAddr - service addr
 * @param {string} srclang - source language
 * @param {string} destlang - destination language
 * @param {string} text - text
 */
function startTranslate2(servAddr, srclang, destlang, text) {
  const client = new services.JarvisCrawlerServiceClient(
      servAddr,
      grpc.credentials.createInsecure(),
  );

  const request = new messages.RequestTranslate2();
  request.setText(text);
  request.setSrclang(srclang);
  request.setDestlang(destlang);

  requestCrawler(
      client,
      TOKEN,
      messages.CrawlerType.CT_TRANSLATE2,
      request,
      (err, reply) => {
        if (err) {
          log.error('err:', err);
        }

        if (reply) {
          log.debug('text:', reply.getText());
        }
      },
  );
}

/**
 * getCrunchBaseCompany
 * @param {string} servAddr - service addr
 * @param {string} company - company code
 */
function getCrunchBaseCompany(servAddr, company) {
  const client = new services.JarvisCrawlerServiceClient(
      servAddr,
      grpc.credentials.createInsecure(),
  );

  const request = new messages.RequestCrunchBaseCompany();
  request.setSearch(company);

  requestCrawler(
      client,
      TOKEN,
      messages.CrawlerType.CT_CB_COMPANY,
      request,
      (err, reply) => {
        if (err) {
          log.error('err:', err);
        }

        if (reply) {
          log.debug('text:', reply);
        }
      },
  );
}

/**
 * analyzePage
 * @param {string} servAddr - servAddr
 * @param {string} url - url
 * @param {number} delay - delay time in seconds
 * @param {number} w - viewport width
 * @param {number} h - viewport height
 */
function analyzePage(servAddr, url, delay, w, h) {
  const client = new services.JarvisCrawlerServiceClient(
      servAddr,
      grpc.credentials.createInsecure(),
  );

  const request = new messages.AnalyzePage();
  request.setUrl(url);
  request.setDelay(delay);
  request.setViewportwidth(w);
  request.setViewportheight(h);

  requestCrawler(
      client,
      TOKEN,
      messages.CrawlerType.CT_ANALYZEPAGE,
      request,
      (err, reply) => {
        if (err) {
          log.error('err:', err);
        }

        if (reply) {
          log.debug('reply:', JSON.stringify(reply.toObject()));
        }
      },
  );
}

/**
 * geoip
 * @param {string} servAddr - servAddr
 * @param {string} ip - ip
 */
function geoip(servAddr, ip) {
  const client = new services.JarvisCrawlerServiceClient(
      servAddr,
      grpc.credentials.createInsecure(),
  );

  const request = new messages.RequestGeoIP();
  request.setIp(ip);

  requestCrawler(
      client,
      TOKEN,
      messages.CrawlerType.CT_GEOIP,
      request,
      (err, reply) => {
        if (err) {
          log.error('err:', err);
        }

        if (reply) {
          log.debug('reply:', JSON.stringify(reply.toObject()));
        }
      },
  );
}

/**
 * techinasiaCompany
 * @param {string} servAddr - servAddr
 * @param {string} companycode - companycode
 */
function techinasiaCompany(servAddr, companycode) {
  const client = new services.JarvisCrawlerServiceClient(
      servAddr,
      grpc.credentials.createInsecure(),
  );

  const request = new messages.RequestTechInAsia();
  request.setMode(messages.TechInAsiaMode.TIAM_COMPANY);
  request.setCompanycode(companycode);

  requestCrawler(
      client,
      TOKEN,
      messages.CrawlerType.CT_TECHINASIA,
      request,
      (err, reply) => {
        if (err) {
          log.error('err:', err);
        }

        if (reply) {
          log.debug('reply:', JSON.stringify(reply.toObject()));
        }
      },
  );
}

/**
 * techinasiaJob
 * @param {string} servAddr - servAddr
 * @param {string} jobcode - jobcode
 */
function techinasiaJob(servAddr, jobcode) {
  const client = new services.JarvisCrawlerServiceClient(
      servAddr,
      grpc.credentials.createInsecure(),
  );

  const request = new messages.RequestTechInAsia();
  request.setMode(messages.TechInAsiaMode.TIAM_JOB);
  request.setJobcode(jobcode);

  requestCrawler(
      client,
      TOKEN,
      messages.CrawlerType.CT_TECHINASIA,
      request,
      (err, reply) => {
        if (err) {
          log.error('err:', err);
        }

        if (reply) {
          log.debug('reply:', JSON.stringify(reply.toObject()));
        }
      },
  );
}

/**
 * techinasiaJobTags
 * @param {string} servAddr - servAddr
 * @param {string} maintag - main tag
 */
function techinasiaJobTags(servAddr, maintag) {
  const client = new services.JarvisCrawlerServiceClient(
      servAddr,
      grpc.credentials.createInsecure(),
  );

  const request = new messages.RequestTechInAsia();
  request.setMode(messages.TechInAsiaMode.TIAM_JOBTAG);
  request.setJobtag(maintag);

  requestCrawler(
      client,
      TOKEN,
      messages.CrawlerType.CT_TECHINASIA,
      request,
      (err, reply) => {
        if (err) {
          log.error('err:', err);
        }

        if (reply) {
          log.debug('reply:', JSON.stringify(reply.toObject()));
        }
      },
  );
}

/**
 * steepandcheapProducts
 * @param {string} servAddr - servAddr
 * @param {string} url - url
 * @param {string} page - page
 */
function steepandcheapProducts(servAddr, url, page) {
  const client = new services.JarvisCrawlerServiceClient(
      servAddr,
      grpc.credentials.createInsecure(),
  );

  const request = new messages.RequestSteepAndCheap();
  request.setMode(messages.SteepAndCheapMode.SACM_PRODUCTS);
  request.setUrl(url);
  request.setPage(page);

  requestCrawler(
      client,
      TOKEN,
      messages.CrawlerType.CT_STEEPANDCHEAP,
      request,
      (err, reply) => {
        if (err) {
          log.error('err:', err);
        }

        if (reply) {
          log.debug('reply:', JSON.stringify(reply.toObject()));
        }
      },
  );
}

/**
 * steepandcheapProduct
 * @param {string} servAddr - servAddr
 * @param {string} url - url
 */
function steepandcheapProduct(servAddr, url) {
  const client = new services.JarvisCrawlerServiceClient(
      servAddr,
      grpc.credentials.createInsecure(),
  );

  const request = new messages.RequestSteepAndCheap();
  request.setMode(messages.SteepAndCheapMode.SACM_PRODUCT);
  request.setUrl(url);

  requestCrawler(
      client,
      TOKEN,
      messages.CrawlerType.CT_STEEPANDCHEAP,
      request,
      (err, reply) => {
        if (err) {
          log.error('err:', err);
        }

        if (reply) {
          log.debug('reply:', JSON.stringify(reply.toObject()));
        }
      },
  );
}

/**
 * jrjFunds
 * @param {string} servAddr - servAddr
 */
function jrjFunds(servAddr) {
  const client = new services.JarvisCrawlerServiceClient(
      servAddr,
      grpc.credentials.createInsecure(),
  );

  const request = new messages.RequestJRJ();
  request.setMode(messages.JRJMode.JRJM_FUNDS);

  requestCrawler(
      client,
      TOKEN,
      messages.CrawlerType.CT_JRJ,
      request,
      (err, reply) => {
        if (err) {
          log.error('err:', err);
        }

        if (reply) {
          log.debug('reply:', JSON.stringify(reply.toObject()));
        }
      },
  );
}

/**
 * jrjFund
 * @param {string} servAddr - servAddr
 * @param {string} code - code
 */
function jrjFund(servAddr, code) {
  const client = new services.JarvisCrawlerServiceClient(
      servAddr,
      grpc.credentials.createInsecure(),
  );

  const request = new messages.RequestJRJ();
  request.setMode(messages.JRJMode.JRJM_FUND);
  request.setFundcode(code);

  requestCrawler(
      client,
      TOKEN,
      messages.CrawlerType.CT_JRJ,
      request,
      (err, reply) => {
        if (err) {
          log.error('err:', err);
        }

        if (reply) {
          log.debug('reply:', JSON.stringify(reply.toObject()));
        }
      },
  );
}

/**
 * jrjFundManager
 * @param {string} servAddr - servAddr
 * @param {string} code - code
 */
function jrjFundManager(servAddr, code) {
  const client = new services.JarvisCrawlerServiceClient(
      servAddr,
      grpc.credentials.createInsecure(),
  );

  const request = new messages.RequestJRJ();
  request.setMode(messages.JRJMode.JRJM_FUNDMANAGER);
  request.setFundcode(code);

  requestCrawler(
      client,
      TOKEN,
      messages.CrawlerType.CT_JRJ,
      request,
      (err, reply) => {
        if (err) {
          log.error('err:', err);
        }

        if (reply) {
          log.debug('reply:', JSON.stringify(reply.toObject()));
        }
      },
  );
}

/**
 * jrjFundValue
 * @param {string} servAddr - servAddr
 * @param {string} code - code
 * @param {string} year - year
 */
function jrjFundValue(servAddr, code, year) {
  const client = new services.JarvisCrawlerServiceClient(
      servAddr,
      grpc.credentials.createInsecure(),
  );

  const request = new messages.RequestJRJ();
  request.setMode(messages.JRJMode.JRJM_FUNDVALUE);
  request.setFundcode(code);
  request.setYear(year);

  requestCrawler(
      client,
      TOKEN,
      messages.CrawlerType.CT_JRJ,
      request,
      (err, reply) => {
        if (err) {
          log.error('err:', err);
        }

        if (reply) {
          log.debug('reply:', JSON.stringify(reply.toObject()));
        }
      },
  );
}

/**
 * jdActive
 * @param {string} servAddr - servAddr
 * @param {string} url - url
 */
function jdActive(servAddr, url) {
  const client = new services.JarvisCrawlerServiceClient(
      servAddr,
      grpc.credentials.createInsecure(),
  );

  const request = new messages.RequestJD();
  request.setMode(messages.JDMode.JDM_ACTIVE);
  request.setUrl(url);

  requestCrawler(
      client,
      TOKEN,
      messages.CrawlerType.CT_JD,
      request,
      (err, reply) => {
        if (err) {
          log.error('err:', err);
        }

        if (reply) {
          log.debug('reply:', JSON.stringify(reply.toObject()));
        }
      },
  );
}

/**
 * jdProduct
 * @param {string} servAddr - servAddr
 * @param {string} url - url
 */
function jdProduct(servAddr, url) {
  const client = new services.JarvisCrawlerServiceClient(
      servAddr,
      grpc.credentials.createInsecure(),
  );

  const request = new messages.RequestJD();
  request.setMode(messages.JDMode.JDM_PRODUCT);
  request.setUrl(url);

  requestCrawler(
      client,
      TOKEN,
      messages.CrawlerType.CT_JD,
      request,
      (err, reply) => {
        if (err) {
          log.error('err:', err);
        }

        if (reply) {
          log.debug('reply:', JSON.stringify(reply.toObject()));
        }
      },
  );
}

/**
 * jdActivePage
 * @param {string} servAddr - servAddr
 * @param {string} url - url
 */
function jdActivePage(servAddr, url) {
  const client = new services.JarvisCrawlerServiceClient(
      servAddr,
      grpc.credentials.createInsecure(),
  );

  const request = new messages.RequestJD();
  request.setMode(messages.JDMode.JDM_ACTIVEPAGE);
  request.setUrl(url);

  requestCrawler(
      client,
      TOKEN,
      messages.CrawlerType.CT_JD,
      request,
      (err, reply) => {
        if (err) {
          log.error('err:', err);
        }

        if (reply) {
          log.debug('reply:', JSON.stringify(reply.toObject()));
        }
      },
  );
}

/**
 * alimamaKeepalive
 * @param {string} servAddr - servAddr
 * @param {string} url - url
 */
function alimamaKeepalive(servAddr) {
  const client = new services.JarvisCrawlerServiceClient(
      servAddr,
      grpc.credentials.createInsecure(),
  );

  const request = new messages.RequestAlimama();
  request.setMode(messages.AlimamaMode.ALIMMM_KEEPALIVE);

  requestCrawler(
      client,
      TOKEN,
      messages.CrawlerType.CT_ALIMAMA,
      request,
      (err, reply) => {
        if (err) {
          log.error('err:', err);
        }

        if (reply) {
          log.debug('reply:', JSON.stringify(reply.toObject()));
        }

        alimamaSearch('127.0.0.1:7051', 'montbell 羽绒服 女');
      // alimamaGetTop('127.0.0.1:7051');
      },
  );
}

/**
 * alimamaGetTop
 * @param {string} servAddr - servAddr
 */
function alimamaGetTop(servAddr) {
  const client = new services.JarvisCrawlerServiceClient(
      servAddr,
      grpc.credentials.createInsecure(),
  );

  const request = new messages.RequestAlimama();
  request.setMode(messages.AlimamaMode.ALIMMM_GETTOP);

  requestCrawler(
      client,
      TOKEN,
      messages.CrawlerType.CT_ALIMAMA,
      request,
      (err, reply) => {
        if (err) {
          log.error('err:', err);
        }

        if (reply) {
          log.debug('reply:', JSON.stringify(reply.toObject()));
        }
      },
  );
}

/**
 * alimamaGetShop
 * @param {string} servAddr - servAddr
 * @param {string} url - url
 */
function alimamaGetShop(servAddr, url) {
  const client = new services.JarvisCrawlerServiceClient(
      servAddr,
      grpc.credentials.createInsecure(),
  );

  const request = new messages.RequestAlimama();
  request.setMode(messages.AlimamaMode.ALIMMM_GETSHOP);
  request.setUrl(url);

  requestCrawler(
      client,
      TOKEN,
      messages.CrawlerType.CT_ALIMAMA,
      request,
      (err, reply) => {
        if (err) {
          log.error('err:', err);
        }

        if (reply) {
          log.debug('reply:', JSON.stringify(reply.toObject()));
        }
      },
  );
}

/**
 * alimamaSearch
 * @param {string} servAddr - servAddr
 * @param {string} text - text
 */
function alimamaSearch(servAddr, text) {
  const client = new services.JarvisCrawlerServiceClient(
      servAddr,
      grpc.credentials.createInsecure(),
  );

  const request = new messages.RequestAlimama();
  request.setMode(messages.AlimamaMode.ALIMMM_SEARCH);
  request.setText(text);

  requestCrawler(
      client,
      TOKEN,
      messages.CrawlerType.CT_ALIMAMA,
      request,
      (err, reply) => {
        if (err) {
          log.error('err:', err);
        }

        if (reply) {
          log.debug('reply:', JSON.stringify(reply.toObject()));
        }
      },
  );
}

/**
 * tmallProduct
 * @param {string} servAddr - servAddr
 * @param {string} url - url
 */
function tmallProduct(servAddr, url) {
  const client = new services.JarvisCrawlerServiceClient(
      servAddr,
      grpc.credentials.createInsecure(),
  );

  const request = new messages.RequestTmall();
  request.setMode(messages.TmallMode.TMM_PRODUCT);
  request.setUrl(url);

  requestCrawler(
      client,
      TOKEN,
      messages.CrawlerType.CT_TMALL,
      request,
      (err, reply) => {
        if (err) {
          log.error('err:', err);
        }

        if (reply) {
          log.debug('reply:', JSON.stringify(reply.toObject()));
        }

      // alimamaSearch('127.0.0.1:7051', 'montbell 羽绒服 女');
      // alimamaGetTop('127.0.0.1:7051');
      },
  );
}

/**
 * taobaoProduct
 * @param {string} servAddr - servAddr
 * @param {string} itemid - itemid
 */
function taobaoProduct(servAddr, itemid) {
  const client = new services.JarvisCrawlerServiceClient(
      servAddr,
      grpc.credentials.createInsecure(),
  );

  const request = new messages.RequestTaobao();
  request.setMode(messages.TaobaoMode.TBM_PRODUCT);
  request.setItemid(itemid);

  requestCrawler(
      client,
      TOKEN,
      messages.CrawlerType.CT_TAOBAO,
      request,
      (err, reply) => {
        if (err) {
          log.error('err:', err);
        }

        if (reply) {
          log.debug('reply:', JSON.stringify(reply.toObject()));
        }

      // alimamaSearch('127.0.0.1:7051', 'montbell 羽绒服 女');
      // alimamaGetTop('127.0.0.1:7051');
      },
  );
}

/**
 * mountainstealsProduct
 * @param {string} servAddr - servAddr
 * @param {string} url - url
 */
function mountainstealsProduct(servAddr, url) {
  const client = new services.JarvisCrawlerServiceClient(
      servAddr,
      grpc.credentials.createInsecure(),
  );

  const request = new messages.RequestMountainSteals();
  request.setMode(messages.MountainStealsMode.MSM_PRODUCT);
  request.setUrl(url);

  requestCrawler(
      client,
      TOKEN,
      messages.CrawlerType.CT_MOUNTAINSTEALS,
      request,
      (err, reply) => {
        if (err) {
          log.error('err:', err);
        }

        if (reply) {
          log.debug('reply:', JSON.stringify(reply.toObject()));
        }

      // alimamaSearch('127.0.0.1:7051', 'montbell 羽绒服 女');
      // alimamaGetTop('127.0.0.1:7051');
      },
  );
}

/**
 * mountainstealsSale
 * @param {string} servAddr - servAddr
 * @param {string} url - url
 */
function mountainstealsSale(servAddr, url) {
  const client = new services.JarvisCrawlerServiceClient(
      servAddr,
      grpc.credentials.createInsecure(),
  );

  const request = new messages.RequestMountainSteals();
  request.setMode(messages.MountainStealsMode.MSM_SALE);
  request.setUrl(url);

  requestCrawler(
      client,
      TOKEN,
      messages.CrawlerType.CT_MOUNTAINSTEALS,
      request,
      (err, reply) => {
        if (err) {
          log.error('err:', err);
        }

        if (reply) {
          log.debug('reply:', JSON.stringify(reply.toObject()));
        }

      // alimamaSearch('127.0.0.1:7051', 'montbell 羽绒服 女');
      // alimamaGetTop('127.0.0.1:7051');
      },
  );
}

// startTranslate2(
//     '127.0.0.1:7051',
//     'en',
//     'zh-CN',
//     '@Peter Walker I am sure there is a problem with excel file, I need more time to check it.'
// );

// getCrunchBaseCompany('127.0.0.1:7051', 'slack');

// analyzePage(
//     '127.0.0.1:7051',
//     'http://47.90.46.159:8090/game.html?gameCode=nightclub&language=zh_CN&isCheat=true&slotKey=',
//     10, 1280, 800
// );

// geoip('127.0.0.1:7051', '60.250.112.36');

// techinasiaCompany('127.0.0.1:7051', 'niki-dot-ai');
// techinasiaJob('127.0.0.1:7051', 'b6f2b504-e3b5-4f33-9732-5d0d59af828f');
// techinasiaJobTags('127.0.0.1:7051', 'TYPE');
// steepandcheapProducts('127.0.0.1:7051', 'rc/arcteryx-on-sale', 0);
// steepandcheapProduct(
//     '127.0.0.1:7051',
//     'arc-teryx-rho-lt-zip-neck-top-womens?skid=ARC3698-HARCOR-XL&ti=UExQIFJ1bGUgQmFzZWQ6QXJjJ3Rlcnl4IE9uIFNhbGU6MzoxOg==',
// );

// jrjFunds('127.0.0.1:7051');
// jrjFund('127.0.0.1:7051', '110011');
// jrjFundManager('127.0.0.1:7051', '110011');
// jrjFundValue('127.0.0.1:7051', '110011', '2019');

// jdActive('127.0.0.1:7051', '3nTQQZ66AGtiwwtRcikGFnT1DVjX/index.html');
// jdProduct('127.0.0.1:7051', '100006585530.html');
// jdActivePage('127.0.0.1:7051', 'https://h5.m.jd.com/pc/dev/391BqWHzwykzEcW9DR3zTek4PC8h/index.html');

// alimamaKeepalive('127.0.0.1:7051');
// alimamaKeepalive('10.211.55.4:7052');
// alimamaGetTop('127.0.0.1:7051');
// alimamaSearch('127.0.0.1:7051', '土拨鼠 羽绒服 女');
// alimamaGetShop(
//     '127.0.0.1:7051',
//     'https://pub.alimama.com/myunion.htm?#!/promo/self/shop_detail?userNumberId=2783286164',
// );

// tmallProduct('127.0.0.1:7051', '525967713966');
// tmallProduct('127.0.0.1:7051', '595765750524');
// taobaoProduct('127.0.0.1:7051', '607627559703');

taobaoProduct('127.0.0.1:7051', '23986840005');

// mountainstealsProduct(
//     '127.0.0.1:7051',
//     'smartwool-women-s-dasher-crew-sock_10384064',
// );
// mountainstealsSale('127.0.0.1:7051', 'promo/msbf19');

exports.startTranslate2 = startTranslate2;
exports.getCrunchBaseCompany = getCrunchBaseCompany;
exports.analyzePage = analyzePage;
exports.geoip = geoip;
exports.techinasiaCompany = techinasiaCompany;
exports.techinasiaJob = techinasiaJob;
exports.techinasiaJobTags = techinasiaJobTags;
exports.steepandcheapProducts = steepandcheapProducts;
exports.steepandcheapProduct = steepandcheapProduct;
exports.jrjFunds = jrjFunds;
exports.jrjFund = jrjFund;
exports.jrjFundManager = jrjFundManager;
exports.jrjFundValue = jrjFundValue;
exports.jdActive = jdActive;
exports.jdProduct = jdProduct;
exports.jdActivePage = jdActivePage;
exports.alimamaKeepalive = alimamaKeepalive;
exports.alimamaGetTop = alimamaGetTop;
exports.alimamaSearch = alimamaSearch;
exports.alimamaGetShop = alimamaGetShop;
exports.tmallProduct = tmallProduct;
exports.taobaoProduct = taobaoProduct;
exports.mountainstealsSale = mountainstealsSale;
exports.mountainstealsProduct = mountainstealsProduct;
