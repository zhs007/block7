/**
 * cbcompany
 * @param {object} browser - browser
 * @param {string} company - company name
 */
async function cbcompany(browser, company) {
  const page = await browser.newPage();
  await page
      .setViewport({
        width: 1280,
        height: 600,
        deviceScaleFactor: 1,
      })
      .catch((err) => {
        console.log('cbcompany.setViewport', err);
      });

  await page
      .goto('https://www.crunchbase.com/organization/' + company, {
        waitUntil: 'domcontentloaded',
        timeout: 0,
      })
      .catch((err) => {
        console.log('cbcompany.goto', err);
      });

  console.log('haha');

  const cbobj = await page
      .$$eval('.layout-row.section-header.ng-star-inserted', (objs) => {
        console.log(objs);

        const nameobj = document.getElementsByClassName('cb-overflow-ellipsis');

        const cbobj = {};

        if (nameobj.length > 0) {
          cbobj.name = nameobj[0].innerText;
        }

        for (let i = 0; i < objs.length; ++i) {
          const curobj = objs[i];

          //   console.log(curobj.innerText);

          if (curobj.innerText == 'Overview') {
            const lsteles = curobj.parentElement.getElementsByClassName(
                'cb-text-color-medium field-label flex-100 flex-gt-sm-25 ng-star-inserted'
            );

            // console.log(lsteles);

            for (let j = 0; j < lsteles.length; ++j) {
              const cursubobj = lsteles[j];

              console.log(cursubobj.innerText);
              console.log(cursubobj.nextElementSibling.innerText);

              if (cursubobj.innerText == 'Categories ') {
                cbobj.categories = cursubobj.nextElementSibling.innerText.split(
                    ', '
                );
              } else if (cursubobj.innerText == 'Headquarters Regions ') {
                cbobj.headquartersregions = cursubobj.nextElementSibling.innerText.split(
                    ', '
                );
              } else if (cursubobj.innerText == 'Founded Date ') {
                cbobj.foundeddate = cursubobj.nextElementSibling.innerText;
              } else if (cursubobj.innerText == 'Founders ') {
                cbobj.founders = cursubobj.nextElementSibling.innerText.split(
                    ', '
                );
              } else if (cursubobj.innerText == 'Operating Status ') {
                cbobj.operatingstatus = cursubobj.nextElementSibling.innerText;
              } else if (cursubobj.innerText == 'Funding Status ') {
                cbobj.fundingstatus = cursubobj.nextElementSibling.innerText;
              } else if (cursubobj.innerText == 'Last Funding Type ') {
                cbobj.lastfundingtype = cursubobj.nextElementSibling.innerText;
              } else if (cursubobj.innerText == 'Legal Name ') {
                cbobj.legalname = cursubobj.nextElementSibling.innerText;
              }
            }
          } else if (curobj.innerText == 'IPO & Stock Price') {
            const lsteles = curobj.parentElement.getElementsByClassName(
                'cb-text-color-medium field-label flex-100 flex-gt-sm-25 ng-star-inserted'
            );

            // console.log(lsteles);

            for (let j = 0; j < lsteles.length; ++j) {
              const cursubobj = lsteles[j];

              console.log(cursubobj.innerText);
              console.log(cursubobj.nextElementSibling.innerText);

              if (cursubobj.innerText == 'Stock Symbol ') {
                cbobj.stocksymbol = cursubobj.nextElementSibling.innerText;
              } else if (cursubobj.innerText == 'Valuation at IPO ') {
                cbobj.valuationipo = cursubobj.nextElementSibling.innerText;
              } else if (cursubobj.innerText == 'Money Raised at IPO ') {
                cbobj.moneyraisedipo = cursubobj.nextElementSibling.innerText;
              } else if (cursubobj.innerText == 'IPO Share Price ') {
                cbobj.priceipo = cursubobj.nextElementSibling.innerText;
              } else if (cursubobj.innerText == 'IPO Date ') {
                cbobj.dateipo = cursubobj.nextElementSibling.innerText;
              }
            }
          }
        }

        return cbobj;
      })
      .catch((err) => {
        console.log(
            'cbcompany.$$eval:.layout-row.section-header.ng-star-inserted',
            err
        );
      });

  console.log(cbobj);
}

exports.cbcompany = cbcompany;
