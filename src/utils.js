
const fs = require('fs');
const jarviscrawlercore = require('../proto/result_pb.js');
const AdmZip = require('adm-zip');
const crypto = require('crypto');

/**
 * save protobuf message
 * @param {string} filename - output file
 * @param {string} msg - message
 */
function saveMessage(filename, msg) {
  fs.writeFileSync(filename,
      msg.serializeBinary());
}


/**
 * save protobuf message with zip
 * @param {string} filename - output file
 * @param {string} msg - message
 */
function saveZipMessage(filename, msg) {
  const zip = new AdmZip();
  zip.addFile('msg.pb',
      msg.serializeBinary());
  zip.writeZip(filename);
}

/**
 * hash md5
 * @param {buffer} buf - buffer
 * @return {string} md5 - md5 string
 */
function hashMD5(buf) {
  return crypto.createHash('md5').update(buf).digest('hex');
}

/**
 * set ImageInfo with img
 * @param {object} img - img object
 * @param {object} mapResponse - map response
 * @param {bool} isoutpurimages - is output images
 * @return {ImageInfo} imginfo - imginfo
 */
function setImageInfo(img, mapResponse, isoutpurimages) {
  if (mapResponse[img.url]) {
    img.data = mapResponse[img.url];

    img.hashName = hashMD5(img.data);

    if (isoutpurimages) {
      fs.writeFileSync('./output/' + img.hashName + '.jpg', img.data);
    }
  }

  return img;
}

/**
 * set ImageInfo with img
 * @param {string} url - url
 * @param {object} mapResponse - map response
 * @return {string} hashname - maybe undefined
 */
function getImageHashName(url, mapResponse) {
  if (mapResponse[url]) {
    return hashMD5(mapResponse[url]);
  }

  return undefined;
}

/**
 * new Paragraph with object
 * @param {object} obj - Paragraph object
 * @return {jarviscrawlercore.Paragraph} paragraph - Paragraph
 */
function newParagraph(obj) {
  const result = new jarviscrawlercore.Paragraph();

  if (obj.pt) {
    result.setPt(obj.pt);
  }

  if (obj.imgHashName) {
    result.setImghashname(obj.imgHashName);
  }

  if (obj.text) {
    result.setText(obj.text);
  }

  if (obj.imgURL) {
    result.setImgURL(obj.imgURL);
  }

  return result;
}

/**
 * new ImageInfo with object
 * @param {object} obj - ImageInfo object
 * @return {jarviscrawlercore.ImageInfo} imginfo - ImageInfo
 */
function newImageInfo(obj) {
  const result = new jarviscrawlercore.ImageInfo();

  if (obj.hashName) {
    result.setHashname(obj.hashName);
  }

  if (obj.url) {
    result.setUrl(obj.url);
  }

  if (obj.width) {
    result.setWidth(obj.width);
  }

  if (obj.height) {
    result.setHeight(obj.height);
  }

  if (obj.data) {
    result.setData(obj.data);
  }

  return result;
}

/**
 * new ExportArticleResult with object
 * @param {object} obj - ExportArticleResult object
 * @return {jarviscrawlercore.ExportArticleResult} ear - ExportArticleResult
 */
function newExportArticleResult(obj) {
  const result = new jarviscrawlercore.ExportArticleResult();

  if (obj.title) {
    result.setTitle(obj.title);
  }

  if (obj.author) {
    result.setAuthor(obj.author);
  }

  if (obj.writeTime) {
    result.setWritetime(obj.writeTime);
  }

  if (obj.article) {
    result.setArticle(obj.article);
  }

  if (obj.url) {
    result.setUrl(obj.url);
  }

  if (obj.imgs) {
    for (let i = 0; i < obj.imgs.length; ++i) {
      result.addImgs(newImageInfo(obj.imgs[i]), i);
    }
  }

  if (obj.titleImage) {
    result.setTitleimage(newImageInfo(obj.titleImage));
  }

  if (obj.tags) {
    for (let i = 0; i < obj.tags.length; ++i) {
      result.addTags(obj.tags[i], i);
    }
  }

  if (obj.paragraphs) {
    for (let i = 0; i < obj.paragraphs.length; ++i) {
      result.addParagraphs(newParagraph(obj.paragraphs[i]), i);
    }
  }

  if (obj.summary) {
    result.setSummary(obj.summary);
  }

  return result;
}

exports.saveMessage = saveMessage;
exports.saveZipMessage = saveZipMessage;
exports.hashMD5 = hashMD5;
exports.setImageInfo = setImageInfo;
exports.getImageHashName = getImageHashName;
exports.newParagraph = newParagraph;
exports.newImageInfo = newImageInfo;
exports.newExportArticleResult = newExportArticleResult;
