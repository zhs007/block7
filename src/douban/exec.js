const {startBrowser} = require('../browser');
const {search} = require('./search');
const log = require('../log');

/**
 * doubanexec
 * @param {object} program - program
 * @param {string} version - version
 */
async function doubanexec(program, version) {
  program
      .command('douban [mode]')
      .description('bt')
      .option('-h, --headless [isheadless]', 'headless mode')
      .option('-d, --debug [isdebug]', 'debug mode')
      .option('-s, --search [search]', 'search string')
      .option('-t, --type [type]', 'type')
      .action(function(mode, options) {
        log.console('version is ', version);

        if (!mode) {
          log.debug(
              'command wrong, please type ' + 'jarviscrawler douban --help'
          );

          return;
        }

        log.console('mode - ', mode);

        const headless = options.headless === 'true';
        log.console('headless - ', headless);

        const debugmode = options.debug === 'true';
        log.console('debug - ', debugmode);

        (async () => {
          const browser = await startBrowser(headless);

          await search(browser, options.type, options.search, debugmode);

          if (!debugmode) {
            await browser.close();
          }
        })().catch((err) => {
          log.console('catch a err ', err);

          if (headless) {
            process.exit(-1);
          }
        });
      });
}

exports.doubanexec = doubanexec;
