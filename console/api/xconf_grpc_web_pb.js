/**
 * @fileoverview gRPC-Web generated client stub for api
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


goog.provide('proto.api.XconfClient');
goog.provide('proto.api.XconfPromiseClient');

goog.require('grpc.web.GrpcWebClientBase');
goog.require('grpc.web.AbstractClientBase');
goog.require('grpc.web.ClientReadableStream');
goog.require('grpc.web.Error');
goog.require('grpc.web.MethodDescriptor');
goog.require('grpc.web.MethodType');
goog.require('proto.api.Config');
goog.require('proto.api.CreateConfigRequest');
goog.require('proto.api.CreateGroupRequest');
goog.require('proto.api.CreateNamespaceRequest');
goog.require('proto.api.DeleteConfigRequest');
goog.require('proto.api.GetConfigRequest');
goog.require('proto.api.Group');
goog.require('proto.api.ListConfigsRequest');
goog.require('proto.api.ListConfigsResponse');
goog.require('proto.api.ListGroupsRequest');
goog.require('proto.api.ListGroupsResponse');
goog.require('proto.api.ListNamespacesResponse');
goog.require('proto.api.Namespace');
goog.require('proto.api.UpdateConfigRequest');
goog.require('proto.google.protobuf.Empty');



goog.scope(function() {

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.api.XconfClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.api.XconfPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.api.CreateNamespaceRequest,
 *   !proto.api.Namespace>}
 */
const methodDescriptor_Xconf_CreateNamespace = new grpc.web.MethodDescriptor(
  '/api.Xconf/CreateNamespace',
  grpc.web.MethodType.UNARY,
  proto.api.CreateNamespaceRequest,
  proto.api.Namespace,
  /**
   * @param {!proto.api.CreateNamespaceRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.api.Namespace.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.api.CreateNamespaceRequest,
 *   !proto.api.Namespace>}
 */
const methodInfo_Xconf_CreateNamespace = new grpc.web.AbstractClientBase.MethodInfo(
  proto.api.Namespace,
  /**
   * @param {!proto.api.CreateNamespaceRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.api.Namespace.deserializeBinary
);


/**
 * @param {!proto.api.CreateNamespaceRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.api.Namespace)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.api.Namespace>|undefined}
 *     The XHR Node Readable Stream
 */
proto.api.XconfClient.prototype.createNamespace =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/api.Xconf/CreateNamespace',
      request,
      metadata || {},
      methodDescriptor_Xconf_CreateNamespace,
      callback);
};


/**
 * @param {!proto.api.CreateNamespaceRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.api.Namespace>}
 *     A native promise that resolves to the response
 */
proto.api.XconfPromiseClient.prototype.createNamespace =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/api.Xconf/CreateNamespace',
      request,
      metadata || {},
      methodDescriptor_Xconf_CreateNamespace);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.google.protobuf.Empty,
 *   !proto.api.ListNamespacesResponse>}
 */
const methodDescriptor_Xconf_ListNamespaces = new grpc.web.MethodDescriptor(
  '/api.Xconf/ListNamespaces',
  grpc.web.MethodType.UNARY,
  proto.google.protobuf.Empty,
  proto.api.ListNamespacesResponse,
  /**
   * @param {!proto.google.protobuf.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.api.ListNamespacesResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.google.protobuf.Empty,
 *   !proto.api.ListNamespacesResponse>}
 */
const methodInfo_Xconf_ListNamespaces = new grpc.web.AbstractClientBase.MethodInfo(
  proto.api.ListNamespacesResponse,
  /**
   * @param {!proto.google.protobuf.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.api.ListNamespacesResponse.deserializeBinary
);


/**
 * @param {!proto.google.protobuf.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.api.ListNamespacesResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.api.ListNamespacesResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.api.XconfClient.prototype.listNamespaces =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/api.Xconf/ListNamespaces',
      request,
      metadata || {},
      methodDescriptor_Xconf_ListNamespaces,
      callback);
};


/**
 * @param {!proto.google.protobuf.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.api.ListNamespacesResponse>}
 *     A native promise that resolves to the response
 */
proto.api.XconfPromiseClient.prototype.listNamespaces =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/api.Xconf/ListNamespaces',
      request,
      metadata || {},
      methodDescriptor_Xconf_ListNamespaces);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.api.CreateGroupRequest,
 *   !proto.api.Group>}
 */
const methodDescriptor_Xconf_CreateGroup = new grpc.web.MethodDescriptor(
  '/api.Xconf/CreateGroup',
  grpc.web.MethodType.UNARY,
  proto.api.CreateGroupRequest,
  proto.api.Group,
  /**
   * @param {!proto.api.CreateGroupRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.api.Group.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.api.CreateGroupRequest,
 *   !proto.api.Group>}
 */
const methodInfo_Xconf_CreateGroup = new grpc.web.AbstractClientBase.MethodInfo(
  proto.api.Group,
  /**
   * @param {!proto.api.CreateGroupRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.api.Group.deserializeBinary
);


/**
 * @param {!proto.api.CreateGroupRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.api.Group)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.api.Group>|undefined}
 *     The XHR Node Readable Stream
 */
proto.api.XconfClient.prototype.createGroup =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/api.Xconf/CreateGroup',
      request,
      metadata || {},
      methodDescriptor_Xconf_CreateGroup,
      callback);
};


/**
 * @param {!proto.api.CreateGroupRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.api.Group>}
 *     A native promise that resolves to the response
 */
proto.api.XconfPromiseClient.prototype.createGroup =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/api.Xconf/CreateGroup',
      request,
      metadata || {},
      methodDescriptor_Xconf_CreateGroup);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.api.ListGroupsRequest,
 *   !proto.api.ListGroupsResponse>}
 */
const methodDescriptor_Xconf_ListGroups = new grpc.web.MethodDescriptor(
  '/api.Xconf/ListGroups',
  grpc.web.MethodType.UNARY,
  proto.api.ListGroupsRequest,
  proto.api.ListGroupsResponse,
  /**
   * @param {!proto.api.ListGroupsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.api.ListGroupsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.api.ListGroupsRequest,
 *   !proto.api.ListGroupsResponse>}
 */
const methodInfo_Xconf_ListGroups = new grpc.web.AbstractClientBase.MethodInfo(
  proto.api.ListGroupsResponse,
  /**
   * @param {!proto.api.ListGroupsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.api.ListGroupsResponse.deserializeBinary
);


/**
 * @param {!proto.api.ListGroupsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.api.ListGroupsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.api.ListGroupsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.api.XconfClient.prototype.listGroups =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/api.Xconf/ListGroups',
      request,
      metadata || {},
      methodDescriptor_Xconf_ListGroups,
      callback);
};


/**
 * @param {!proto.api.ListGroupsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.api.ListGroupsResponse>}
 *     A native promise that resolves to the response
 */
proto.api.XconfPromiseClient.prototype.listGroups =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/api.Xconf/ListGroups',
      request,
      metadata || {},
      methodDescriptor_Xconf_ListGroups);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.api.CreateConfigRequest,
 *   !proto.api.Config>}
 */
const methodDescriptor_Xconf_CreateConfig = new grpc.web.MethodDescriptor(
  '/api.Xconf/CreateConfig',
  grpc.web.MethodType.UNARY,
  proto.api.CreateConfigRequest,
  proto.api.Config,
  /**
   * @param {!proto.api.CreateConfigRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.api.Config.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.api.CreateConfigRequest,
 *   !proto.api.Config>}
 */
const methodInfo_Xconf_CreateConfig = new grpc.web.AbstractClientBase.MethodInfo(
  proto.api.Config,
  /**
   * @param {!proto.api.CreateConfigRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.api.Config.deserializeBinary
);


/**
 * @param {!proto.api.CreateConfigRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.api.Config)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.api.Config>|undefined}
 *     The XHR Node Readable Stream
 */
proto.api.XconfClient.prototype.createConfig =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/api.Xconf/CreateConfig',
      request,
      metadata || {},
      methodDescriptor_Xconf_CreateConfig,
      callback);
};


/**
 * @param {!proto.api.CreateConfigRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.api.Config>}
 *     A native promise that resolves to the response
 */
proto.api.XconfPromiseClient.prototype.createConfig =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/api.Xconf/CreateConfig',
      request,
      metadata || {},
      methodDescriptor_Xconf_CreateConfig);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.api.UpdateConfigRequest,
 *   !proto.api.Config>}
 */
const methodDescriptor_Xconf_UpdateConfig = new grpc.web.MethodDescriptor(
  '/api.Xconf/UpdateConfig',
  grpc.web.MethodType.UNARY,
  proto.api.UpdateConfigRequest,
  proto.api.Config,
  /**
   * @param {!proto.api.UpdateConfigRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.api.Config.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.api.UpdateConfigRequest,
 *   !proto.api.Config>}
 */
const methodInfo_Xconf_UpdateConfig = new grpc.web.AbstractClientBase.MethodInfo(
  proto.api.Config,
  /**
   * @param {!proto.api.UpdateConfigRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.api.Config.deserializeBinary
);


/**
 * @param {!proto.api.UpdateConfigRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.api.Config)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.api.Config>|undefined}
 *     The XHR Node Readable Stream
 */
proto.api.XconfClient.prototype.updateConfig =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/api.Xconf/UpdateConfig',
      request,
      metadata || {},
      methodDescriptor_Xconf_UpdateConfig,
      callback);
};


/**
 * @param {!proto.api.UpdateConfigRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.api.Config>}
 *     A native promise that resolves to the response
 */
proto.api.XconfPromiseClient.prototype.updateConfig =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/api.Xconf/UpdateConfig',
      request,
      metadata || {},
      methodDescriptor_Xconf_UpdateConfig);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.api.GetConfigRequest,
 *   !proto.api.Config>}
 */
const methodDescriptor_Xconf_GetConfig = new grpc.web.MethodDescriptor(
  '/api.Xconf/GetConfig',
  grpc.web.MethodType.UNARY,
  proto.api.GetConfigRequest,
  proto.api.Config,
  /**
   * @param {!proto.api.GetConfigRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.api.Config.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.api.GetConfigRequest,
 *   !proto.api.Config>}
 */
const methodInfo_Xconf_GetConfig = new grpc.web.AbstractClientBase.MethodInfo(
  proto.api.Config,
  /**
   * @param {!proto.api.GetConfigRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.api.Config.deserializeBinary
);


/**
 * @param {!proto.api.GetConfigRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.api.Config)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.api.Config>|undefined}
 *     The XHR Node Readable Stream
 */
proto.api.XconfClient.prototype.getConfig =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/api.Xconf/GetConfig',
      request,
      metadata || {},
      methodDescriptor_Xconf_GetConfig,
      callback);
};


/**
 * @param {!proto.api.GetConfigRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.api.Config>}
 *     A native promise that resolves to the response
 */
proto.api.XconfPromiseClient.prototype.getConfig =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/api.Xconf/GetConfig',
      request,
      metadata || {},
      methodDescriptor_Xconf_GetConfig);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.api.DeleteConfigRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_Xconf_DeleteConfig = new grpc.web.MethodDescriptor(
  '/api.Xconf/DeleteConfig',
  grpc.web.MethodType.UNARY,
  proto.api.DeleteConfigRequest,
  proto.google.protobuf.Empty,
  /**
   * @param {!proto.api.DeleteConfigRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.google.protobuf.Empty.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.api.DeleteConfigRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_Xconf_DeleteConfig = new grpc.web.AbstractClientBase.MethodInfo(
  proto.google.protobuf.Empty,
  /**
   * @param {!proto.api.DeleteConfigRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.google.protobuf.Empty.deserializeBinary
);


/**
 * @param {!proto.api.DeleteConfigRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.api.XconfClient.prototype.deleteConfig =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/api.Xconf/DeleteConfig',
      request,
      metadata || {},
      methodDescriptor_Xconf_DeleteConfig,
      callback);
};


/**
 * @param {!proto.api.DeleteConfigRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     A native promise that resolves to the response
 */
proto.api.XconfPromiseClient.prototype.deleteConfig =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/api.Xconf/DeleteConfig',
      request,
      metadata || {},
      methodDescriptor_Xconf_DeleteConfig);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.api.ListConfigsRequest,
 *   !proto.api.ListConfigsResponse>}
 */
const methodDescriptor_Xconf_ListConfigs = new grpc.web.MethodDescriptor(
  '/api.Xconf/ListConfigs',
  grpc.web.MethodType.UNARY,
  proto.api.ListConfigsRequest,
  proto.api.ListConfigsResponse,
  /**
   * @param {!proto.api.ListConfigsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.api.ListConfigsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.api.ListConfigsRequest,
 *   !proto.api.ListConfigsResponse>}
 */
const methodInfo_Xconf_ListConfigs = new grpc.web.AbstractClientBase.MethodInfo(
  proto.api.ListConfigsResponse,
  /**
   * @param {!proto.api.ListConfigsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.api.ListConfigsResponse.deserializeBinary
);


/**
 * @param {!proto.api.ListConfigsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.api.ListConfigsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.api.ListConfigsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.api.XconfClient.prototype.listConfigs =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/api.Xconf/ListConfigs',
      request,
      metadata || {},
      methodDescriptor_Xconf_ListConfigs,
      callback);
};


/**
 * @param {!proto.api.ListConfigsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.api.ListConfigsResponse>}
 *     A native promise that resolves to the response
 */
proto.api.XconfPromiseClient.prototype.listConfigs =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/api.Xconf/ListConfigs',
      request,
      metadata || {},
      methodDescriptor_Xconf_ListConfigs);
};


}); // goog.scope

