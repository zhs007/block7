/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.jarviscrawlercore.OABTMode', null, global);
goog.exportSymbol('proto.jarviscrawlercore.OABTNode', null, global);
goog.exportSymbol('proto.jarviscrawlercore.OABTPage', null, global);
goog.exportSymbol('proto.jarviscrawlercore.OABTResInfo', null, global);
goog.exportSymbol('proto.jarviscrawlercore.OABTResType', null, global);
goog.exportSymbol('proto.jarviscrawlercore.ReplyOABT', null, global);
goog.exportSymbol('proto.jarviscrawlercore.RequestOABT', null, global);

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.jarviscrawlercore.OABTResInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.jarviscrawlercore.OABTResInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.jarviscrawlercore.OABTResInfo.displayName = 'proto.jarviscrawlercore.OABTResInfo';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.jarviscrawlercore.OABTResInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.jarviscrawlercore.OABTResInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.jarviscrawlercore.OABTResInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.jarviscrawlercore.OABTResInfo.toObject = function(includeInstance, msg) {
  var f, obj = {
    type: jspb.Message.getFieldWithDefault(msg, 1, 0),
    url: jspb.Message.getFieldWithDefault(msg, 2, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.jarviscrawlercore.OABTResInfo}
 */
proto.jarviscrawlercore.OABTResInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.jarviscrawlercore.OABTResInfo;
  return proto.jarviscrawlercore.OABTResInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.jarviscrawlercore.OABTResInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.jarviscrawlercore.OABTResInfo}
 */
proto.jarviscrawlercore.OABTResInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.jarviscrawlercore.OABTResType} */ (reader.readEnum());
      msg.setType(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setUrl(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.jarviscrawlercore.OABTResInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.jarviscrawlercore.OABTResInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.jarviscrawlercore.OABTResInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.jarviscrawlercore.OABTResInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getType();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getUrl();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional OABTResType type = 1;
 * @return {!proto.jarviscrawlercore.OABTResType}
 */
proto.jarviscrawlercore.OABTResInfo.prototype.getType = function() {
  return /** @type {!proto.jarviscrawlercore.OABTResType} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {!proto.jarviscrawlercore.OABTResType} value */
proto.jarviscrawlercore.OABTResInfo.prototype.setType = function(value) {
  jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional string url = 2;
 * @return {string}
 */
proto.jarviscrawlercore.OABTResInfo.prototype.getUrl = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.jarviscrawlercore.OABTResInfo.prototype.setUrl = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.jarviscrawlercore.OABTNode = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.jarviscrawlercore.OABTNode.repeatedFields_, null);
};
goog.inherits(proto.jarviscrawlercore.OABTNode, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.jarviscrawlercore.OABTNode.displayName = 'proto.jarviscrawlercore.OABTNode';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.jarviscrawlercore.OABTNode.repeatedFields_ = [4];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.jarviscrawlercore.OABTNode.prototype.toObject = function(opt_includeInstance) {
  return proto.jarviscrawlercore.OABTNode.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.jarviscrawlercore.OABTNode} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.jarviscrawlercore.OABTNode.toObject = function(includeInstance, msg) {
  var f, obj = {
    fullname: jspb.Message.getFieldWithDefault(msg, 1, ""),
    resid: jspb.Message.getFieldWithDefault(msg, 2, ""),
    cat: jspb.Message.getFieldWithDefault(msg, 3, 0),
    lstList: jspb.Message.toObjectList(msg.getLstList(),
    proto.jarviscrawlercore.OABTResInfo.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.jarviscrawlercore.OABTNode}
 */
proto.jarviscrawlercore.OABTNode.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.jarviscrawlercore.OABTNode;
  return proto.jarviscrawlercore.OABTNode.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.jarviscrawlercore.OABTNode} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.jarviscrawlercore.OABTNode}
 */
proto.jarviscrawlercore.OABTNode.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setFullname(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setResid(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setCat(value);
      break;
    case 4:
      var value = new proto.jarviscrawlercore.OABTResInfo;
      reader.readMessage(value,proto.jarviscrawlercore.OABTResInfo.deserializeBinaryFromReader);
      msg.addLst(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.jarviscrawlercore.OABTNode.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.jarviscrawlercore.OABTNode.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.jarviscrawlercore.OABTNode} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.jarviscrawlercore.OABTNode.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getFullname();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getResid();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getCat();
  if (f !== 0) {
    writer.writeInt32(
      3,
      f
    );
  }
  f = message.getLstList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      4,
      f,
      proto.jarviscrawlercore.OABTResInfo.serializeBinaryToWriter
    );
  }
};


/**
 * optional string fullname = 1;
 * @return {string}
 */
proto.jarviscrawlercore.OABTNode.prototype.getFullname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.jarviscrawlercore.OABTNode.prototype.setFullname = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string resid = 2;
 * @return {string}
 */
proto.jarviscrawlercore.OABTNode.prototype.getResid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.jarviscrawlercore.OABTNode.prototype.setResid = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional int32 cat = 3;
 * @return {number}
 */
proto.jarviscrawlercore.OABTNode.prototype.getCat = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/** @param {number} value */
proto.jarviscrawlercore.OABTNode.prototype.setCat = function(value) {
  jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * repeated OABTResInfo lst = 4;
 * @return {!Array<!proto.jarviscrawlercore.OABTResInfo>}
 */
proto.jarviscrawlercore.OABTNode.prototype.getLstList = function() {
  return /** @type{!Array<!proto.jarviscrawlercore.OABTResInfo>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.jarviscrawlercore.OABTResInfo, 4));
};


/** @param {!Array<!proto.jarviscrawlercore.OABTResInfo>} value */
proto.jarviscrawlercore.OABTNode.prototype.setLstList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 4, value);
};


/**
 * @param {!proto.jarviscrawlercore.OABTResInfo=} opt_value
 * @param {number=} opt_index
 * @return {!proto.jarviscrawlercore.OABTResInfo}
 */
proto.jarviscrawlercore.OABTNode.prototype.addLst = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 4, opt_value, proto.jarviscrawlercore.OABTResInfo, opt_index);
};


proto.jarviscrawlercore.OABTNode.prototype.clearLstList = function() {
  this.setLstList([]);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.jarviscrawlercore.OABTPage = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.jarviscrawlercore.OABTPage.repeatedFields_, null);
};
goog.inherits(proto.jarviscrawlercore.OABTPage, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.jarviscrawlercore.OABTPage.displayName = 'proto.jarviscrawlercore.OABTPage';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.jarviscrawlercore.OABTPage.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.jarviscrawlercore.OABTPage.prototype.toObject = function(opt_includeInstance) {
  return proto.jarviscrawlercore.OABTPage.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.jarviscrawlercore.OABTPage} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.jarviscrawlercore.OABTPage.toObject = function(includeInstance, msg) {
  var f, obj = {
    lstList: jspb.Message.toObjectList(msg.getLstList(),
    proto.jarviscrawlercore.OABTNode.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.jarviscrawlercore.OABTPage}
 */
proto.jarviscrawlercore.OABTPage.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.jarviscrawlercore.OABTPage;
  return proto.jarviscrawlercore.OABTPage.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.jarviscrawlercore.OABTPage} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.jarviscrawlercore.OABTPage}
 */
proto.jarviscrawlercore.OABTPage.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.jarviscrawlercore.OABTNode;
      reader.readMessage(value,proto.jarviscrawlercore.OABTNode.deserializeBinaryFromReader);
      msg.addLst(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.jarviscrawlercore.OABTPage.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.jarviscrawlercore.OABTPage.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.jarviscrawlercore.OABTPage} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.jarviscrawlercore.OABTPage.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getLstList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.jarviscrawlercore.OABTNode.serializeBinaryToWriter
    );
  }
};


/**
 * repeated OABTNode lst = 1;
 * @return {!Array<!proto.jarviscrawlercore.OABTNode>}
 */
proto.jarviscrawlercore.OABTPage.prototype.getLstList = function() {
  return /** @type{!Array<!proto.jarviscrawlercore.OABTNode>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.jarviscrawlercore.OABTNode, 1));
};


/** @param {!Array<!proto.jarviscrawlercore.OABTNode>} value */
proto.jarviscrawlercore.OABTPage.prototype.setLstList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.jarviscrawlercore.OABTNode=} opt_value
 * @param {number=} opt_index
 * @return {!proto.jarviscrawlercore.OABTNode}
 */
proto.jarviscrawlercore.OABTPage.prototype.addLst = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.jarviscrawlercore.OABTNode, opt_index);
};


proto.jarviscrawlercore.OABTPage.prototype.clearLstList = function() {
  this.setLstList([]);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.jarviscrawlercore.RequestOABT = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.jarviscrawlercore.RequestOABT, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.jarviscrawlercore.RequestOABT.displayName = 'proto.jarviscrawlercore.RequestOABT';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.jarviscrawlercore.RequestOABT.prototype.toObject = function(opt_includeInstance) {
  return proto.jarviscrawlercore.RequestOABT.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.jarviscrawlercore.RequestOABT} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.jarviscrawlercore.RequestOABT.toObject = function(includeInstance, msg) {
  var f, obj = {
    mode: jspb.Message.getFieldWithDefault(msg, 1, 0),
    pageindex: jspb.Message.getFieldWithDefault(msg, 2, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.jarviscrawlercore.RequestOABT}
 */
proto.jarviscrawlercore.RequestOABT.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.jarviscrawlercore.RequestOABT;
  return proto.jarviscrawlercore.RequestOABT.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.jarviscrawlercore.RequestOABT} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.jarviscrawlercore.RequestOABT}
 */
proto.jarviscrawlercore.RequestOABT.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.jarviscrawlercore.OABTMode} */ (reader.readEnum());
      msg.setMode(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setPageindex(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.jarviscrawlercore.RequestOABT.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.jarviscrawlercore.RequestOABT.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.jarviscrawlercore.RequestOABT} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.jarviscrawlercore.RequestOABT.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getMode();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getPageindex();
  if (f !== 0) {
    writer.writeInt32(
      2,
      f
    );
  }
};


/**
 * optional OABTMode mode = 1;
 * @return {!proto.jarviscrawlercore.OABTMode}
 */
proto.jarviscrawlercore.RequestOABT.prototype.getMode = function() {
  return /** @type {!proto.jarviscrawlercore.OABTMode} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {!proto.jarviscrawlercore.OABTMode} value */
proto.jarviscrawlercore.RequestOABT.prototype.setMode = function(value) {
  jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional int32 pageIndex = 2;
 * @return {number}
 */
proto.jarviscrawlercore.RequestOABT.prototype.getPageindex = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.jarviscrawlercore.RequestOABT.prototype.setPageindex = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.jarviscrawlercore.ReplyOABT = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, proto.jarviscrawlercore.ReplyOABT.oneofGroups_);
};
goog.inherits(proto.jarviscrawlercore.ReplyOABT, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.jarviscrawlercore.ReplyOABT.displayName = 'proto.jarviscrawlercore.ReplyOABT';
}
/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.jarviscrawlercore.ReplyOABT.oneofGroups_ = [[100]];

/**
 * @enum {number}
 */
proto.jarviscrawlercore.ReplyOABT.ReplyCase = {
  REPLY_NOT_SET: 0,
  PAGE: 100
};

/**
 * @return {proto.jarviscrawlercore.ReplyOABT.ReplyCase}
 */
proto.jarviscrawlercore.ReplyOABT.prototype.getReplyCase = function() {
  return /** @type {proto.jarviscrawlercore.ReplyOABT.ReplyCase} */(jspb.Message.computeOneofCase(this, proto.jarviscrawlercore.ReplyOABT.oneofGroups_[0]));
};



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.jarviscrawlercore.ReplyOABT.prototype.toObject = function(opt_includeInstance) {
  return proto.jarviscrawlercore.ReplyOABT.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.jarviscrawlercore.ReplyOABT} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.jarviscrawlercore.ReplyOABT.toObject = function(includeInstance, msg) {
  var f, obj = {
    mode: jspb.Message.getFieldWithDefault(msg, 1, 0),
    page: (f = msg.getPage()) && proto.jarviscrawlercore.OABTPage.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.jarviscrawlercore.ReplyOABT}
 */
proto.jarviscrawlercore.ReplyOABT.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.jarviscrawlercore.ReplyOABT;
  return proto.jarviscrawlercore.ReplyOABT.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.jarviscrawlercore.ReplyOABT} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.jarviscrawlercore.ReplyOABT}
 */
proto.jarviscrawlercore.ReplyOABT.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.jarviscrawlercore.OABTMode} */ (reader.readEnum());
      msg.setMode(value);
      break;
    case 100:
      var value = new proto.jarviscrawlercore.OABTPage;
      reader.readMessage(value,proto.jarviscrawlercore.OABTPage.deserializeBinaryFromReader);
      msg.setPage(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.jarviscrawlercore.ReplyOABT.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.jarviscrawlercore.ReplyOABT.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.jarviscrawlercore.ReplyOABT} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.jarviscrawlercore.ReplyOABT.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getMode();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getPage();
  if (f != null) {
    writer.writeMessage(
      100,
      f,
      proto.jarviscrawlercore.OABTPage.serializeBinaryToWriter
    );
  }
};


/**
 * optional OABTMode mode = 1;
 * @return {!proto.jarviscrawlercore.OABTMode}
 */
proto.jarviscrawlercore.ReplyOABT.prototype.getMode = function() {
  return /** @type {!proto.jarviscrawlercore.OABTMode} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {!proto.jarviscrawlercore.OABTMode} value */
proto.jarviscrawlercore.ReplyOABT.prototype.setMode = function(value) {
  jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional OABTPage page = 100;
 * @return {?proto.jarviscrawlercore.OABTPage}
 */
proto.jarviscrawlercore.ReplyOABT.prototype.getPage = function() {
  return /** @type{?proto.jarviscrawlercore.OABTPage} */ (
    jspb.Message.getWrapperField(this, proto.jarviscrawlercore.OABTPage, 100));
};


/** @param {?proto.jarviscrawlercore.OABTPage|undefined} value */
proto.jarviscrawlercore.ReplyOABT.prototype.setPage = function(value) {
  jspb.Message.setOneofWrapperField(this, 100, proto.jarviscrawlercore.ReplyOABT.oneofGroups_[0], value);
};


proto.jarviscrawlercore.ReplyOABT.prototype.clearPage = function() {
  this.setPage(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.jarviscrawlercore.ReplyOABT.prototype.hasPage = function() {
  return jspb.Message.getField(this, 100) != null;
};


/**
 * @enum {number}
 */
proto.jarviscrawlercore.OABTResType = {
  OABTRT_ED2K: 0,
  OABTRT_MAGNET: 1
};

/**
 * @enum {number}
 */
proto.jarviscrawlercore.OABTMode = {
  OABTM_PAGE: 0
};

goog.object.extend(exports, proto.jarviscrawlercore);
