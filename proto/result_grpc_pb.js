// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var proto_result_pb = require('../proto/result_pb.js');

function serialize_jarviscrawlercore_ReplyArticle(arg) {
  if (!(arg instanceof proto_result_pb.ReplyArticle)) {
    throw new Error('Expected argument of type jarviscrawlercore.ReplyArticle');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_jarviscrawlercore_ReplyArticle(buffer_arg) {
  return proto_result_pb.ReplyArticle.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_jarviscrawlercore_ReplyArticles(arg) {
  if (!(arg instanceof proto_result_pb.ReplyArticles)) {
    throw new Error('Expected argument of type jarviscrawlercore.ReplyArticles');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_jarviscrawlercore_ReplyArticles(buffer_arg) {
  return proto_result_pb.ReplyArticles.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_jarviscrawlercore_ReplyDTData(arg) {
  if (!(arg instanceof proto_result_pb.ReplyDTData)) {
    throw new Error('Expected argument of type jarviscrawlercore.ReplyDTData');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_jarviscrawlercore_ReplyDTData(buffer_arg) {
  return proto_result_pb.ReplyDTData.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_jarviscrawlercore_ReplyTranslate(arg) {
  if (!(arg instanceof proto_result_pb.ReplyTranslate)) {
    throw new Error('Expected argument of type jarviscrawlercore.ReplyTranslate');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_jarviscrawlercore_ReplyTranslate(buffer_arg) {
  return proto_result_pb.ReplyTranslate.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_jarviscrawlercore_RequestArticle(arg) {
  if (!(arg instanceof proto_result_pb.RequestArticle)) {
    throw new Error('Expected argument of type jarviscrawlercore.RequestArticle');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_jarviscrawlercore_RequestArticle(buffer_arg) {
  return proto_result_pb.RequestArticle.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_jarviscrawlercore_RequestArticles(arg) {
  if (!(arg instanceof proto_result_pb.RequestArticles)) {
    throw new Error('Expected argument of type jarviscrawlercore.RequestArticles');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_jarviscrawlercore_RequestArticles(buffer_arg) {
  return proto_result_pb.RequestArticles.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_jarviscrawlercore_RequestDTData(arg) {
  if (!(arg instanceof proto_result_pb.RequestDTData)) {
    throw new Error('Expected argument of type jarviscrawlercore.RequestDTData');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_jarviscrawlercore_RequestDTData(buffer_arg) {
  return proto_result_pb.RequestDTData.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_jarviscrawlercore_RequestTranslate(arg) {
  if (!(arg instanceof proto_result_pb.RequestTranslate)) {
    throw new Error('Expected argument of type jarviscrawlercore.RequestTranslate');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_jarviscrawlercore_RequestTranslate(buffer_arg) {
  return proto_result_pb.RequestTranslate.deserializeBinary(new Uint8Array(buffer_arg));
}


// JarvisCrawlerService - JarvisCrawler service
var JarvisCrawlerServiceService = exports.JarvisCrawlerServiceService = {
  // translate - translate text
  translate: {
    path: '/jarviscrawlercore.JarvisCrawlerService/translate',
    requestStream: false,
    responseStream: false,
    requestType: proto_result_pb.RequestTranslate,
    responseType: proto_result_pb.ReplyTranslate,
    requestSerialize: serialize_jarviscrawlercore_RequestTranslate,
    requestDeserialize: deserialize_jarviscrawlercore_RequestTranslate,
    responseSerialize: serialize_jarviscrawlercore_ReplyTranslate,
    responseDeserialize: deserialize_jarviscrawlercore_ReplyTranslate,
  },
  // exportArticle - export article
  exportArticle: {
    path: '/jarviscrawlercore.JarvisCrawlerService/exportArticle',
    requestStream: false,
    responseStream: true,
    requestType: proto_result_pb.RequestArticle,
    responseType: proto_result_pb.ReplyArticle,
    requestSerialize: serialize_jarviscrawlercore_RequestArticle,
    requestDeserialize: deserialize_jarviscrawlercore_RequestArticle,
    responseSerialize: serialize_jarviscrawlercore_ReplyArticle,
    responseDeserialize: deserialize_jarviscrawlercore_ReplyArticle,
  },
  // getArticles - get articles
  getArticles: {
    path: '/jarviscrawlercore.JarvisCrawlerService/getArticles',
    requestStream: false,
    responseStream: false,
    requestType: proto_result_pb.RequestArticles,
    responseType: proto_result_pb.ReplyArticles,
    requestSerialize: serialize_jarviscrawlercore_RequestArticles,
    requestDeserialize: deserialize_jarviscrawlercore_RequestArticles,
    responseSerialize: serialize_jarviscrawlercore_ReplyArticles,
    responseDeserialize: deserialize_jarviscrawlercore_ReplyArticles,
  },
  // getDTData - get DT data
  getDTData: {
    path: '/jarviscrawlercore.JarvisCrawlerService/getDTData',
    requestStream: false,
    responseStream: false,
    requestType: proto_result_pb.RequestDTData,
    responseType: proto_result_pb.ReplyDTData,
    requestSerialize: serialize_jarviscrawlercore_RequestDTData,
    requestDeserialize: deserialize_jarviscrawlercore_RequestDTData,
    responseSerialize: serialize_jarviscrawlercore_ReplyDTData,
    responseDeserialize: deserialize_jarviscrawlercore_ReplyDTData,
  },
};

exports.JarvisCrawlerServiceClient = grpc.makeGenericClientConstructor(JarvisCrawlerServiceService);
