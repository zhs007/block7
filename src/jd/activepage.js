const log = require('../log');
const {WaitAllResponse} = require('../waitallresponse');

/**
 * jdActivePage - jd active page
 * @param {object} browser - browser
 * @param {string} url - url
 * @param {number} timeout - timeout in microseconds
 * @return {object} ret - {error, ret}
 */
async function jdActivePage(browser, url, timeout) {
  let awaiterr = undefined;
  const page = await browser.newPage();

  const waitAllResponse = new WaitAllResponse(page);

  await page
      .setViewport({
        width: 1280,
        height: 600,
        deviceScaleFactor: 1,
      })
      .catch((err) => {
        awaiterr = err;
      });

  if (awaiterr) {
    log.error('jdActivePage.setViewport', awaiterr);

    await page.close();

    return {error: awaiterr.toString()};
  }

  await page
      .goto(url, {
        timeout: timeout,
      })
      .catch((err) => {
        awaiterr = err;
      });

  if (awaiterr) {
    log.error('jdActivePage.goto', awaiterr);

    await page.close();

    return {error: awaiterr.toString()};
  }

  const isok = await waitAllResponse.waitDone(timeout);
  if (!isok) {
    const err = new Error('jdActivePage.waitDone timeout.');

    log.error('jdActivePage.waitDone ', err);

    await page.close();

    return {erroor: err};
  }

  waitAllResponse.reset();

  const ret = await page
      .$$eval('a', (eles) => {
        if (eles.length > 0) {
          activeret = {
            urlActive: [],
            urlProduct: [],
            title: document.title,
          };

          for (let i = 0; i < eles.length; ++i) {
            if (eles[i].href.indexOf('https://pro.jd.com/mall/active/') == 0) {
              activeret.urlActive.push(
                  eles[i].href.split('https://pro.jd.com/mall/active/')[1]
              );
            } else if (eles[i].href.indexOf('https://item.jd.com/') == 0) {
              activeret.urlProduct.push(
                  eles[i].href.split('https://item.jd.com/')[1]
              );
            }
          }

          return activeret;
        }

        return undefined;
      })
      .catch((err) => {
        awaiterr = err;
      });

  if (awaiterr) {
    log.error('jdActivePage.a', awaiterr);

    await page.close();

    return {error: awaiterr.toString()};
  }

  ret.url = url;

  await page.close();

  return {ret: ret};
}

exports.jdActivePage = jdActivePage;
