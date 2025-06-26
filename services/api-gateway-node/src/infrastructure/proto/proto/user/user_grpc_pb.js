// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var proto_user_user_pb = require('../../proto/user/user_pb.js');

function serialize_user_v1_CreateUserRequest(arg) {
  if (!(arg instanceof proto_user_user_pb.CreateUserRequest)) {
    throw new Error('Expected argument of type user.v1.CreateUserRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_user_v1_CreateUserRequest(buffer_arg) {
  return proto_user_user_pb.CreateUserRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_user_v1_CreateUserResponse(arg) {
  if (!(arg instanceof proto_user_user_pb.CreateUserResponse)) {
    throw new Error('Expected argument of type user.v1.CreateUserResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_user_v1_CreateUserResponse(buffer_arg) {
  return proto_user_user_pb.CreateUserResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_user_v1_DeleteUserRequest(arg) {
  if (!(arg instanceof proto_user_user_pb.DeleteUserRequest)) {
    throw new Error('Expected argument of type user.v1.DeleteUserRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_user_v1_DeleteUserRequest(buffer_arg) {
  return proto_user_user_pb.DeleteUserRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_user_v1_DeleteUserResponse(arg) {
  if (!(arg instanceof proto_user_user_pb.DeleteUserResponse)) {
    throw new Error('Expected argument of type user.v1.DeleteUserResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_user_v1_DeleteUserResponse(buffer_arg) {
  return proto_user_user_pb.DeleteUserResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_user_v1_GetUserByEmailRequest(arg) {
  if (!(arg instanceof proto_user_user_pb.GetUserByEmailRequest)) {
    throw new Error('Expected argument of type user.v1.GetUserByEmailRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_user_v1_GetUserByEmailRequest(buffer_arg) {
  return proto_user_user_pb.GetUserByEmailRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_user_v1_GetUserRequest(arg) {
  if (!(arg instanceof proto_user_user_pb.GetUserRequest)) {
    throw new Error('Expected argument of type user.v1.GetUserRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_user_v1_GetUserRequest(buffer_arg) {
  return proto_user_user_pb.GetUserRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_user_v1_GetUserResponse(arg) {
  if (!(arg instanceof proto_user_user_pb.GetUserResponse)) {
    throw new Error('Expected argument of type user.v1.GetUserResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_user_v1_GetUserResponse(buffer_arg) {
  return proto_user_user_pb.GetUserResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_user_v1_GetUsersRequest(arg) {
  if (!(arg instanceof proto_user_user_pb.GetUsersRequest)) {
    throw new Error('Expected argument of type user.v1.GetUsersRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_user_v1_GetUsersRequest(buffer_arg) {
  return proto_user_user_pb.GetUsersRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_user_v1_GetUsersResponse(arg) {
  if (!(arg instanceof proto_user_user_pb.GetUsersResponse)) {
    throw new Error('Expected argument of type user.v1.GetUsersResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_user_v1_GetUsersResponse(buffer_arg) {
  return proto_user_user_pb.GetUsersResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_user_v1_UpdateUserRequest(arg) {
  if (!(arg instanceof proto_user_user_pb.UpdateUserRequest)) {
    throw new Error('Expected argument of type user.v1.UpdateUserRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_user_v1_UpdateUserRequest(buffer_arg) {
  return proto_user_user_pb.UpdateUserRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_user_v1_UpdateUserResponse(arg) {
  if (!(arg instanceof proto_user_user_pb.UpdateUserResponse)) {
    throw new Error('Expected argument of type user.v1.UpdateUserResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_user_v1_UpdateUserResponse(buffer_arg) {
  return proto_user_user_pb.UpdateUserResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_user_v1_ValidateUserRequest(arg) {
  if (!(arg instanceof proto_user_user_pb.ValidateUserRequest)) {
    throw new Error('Expected argument of type user.v1.ValidateUserRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_user_v1_ValidateUserRequest(buffer_arg) {
  return proto_user_user_pb.ValidateUserRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_user_v1_ValidateUserResponse(arg) {
  if (!(arg instanceof proto_user_user_pb.ValidateUserResponse)) {
    throw new Error('Expected argument of type user.v1.ValidateUserResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_user_v1_ValidateUserResponse(buffer_arg) {
  return proto_user_user_pb.ValidateUserResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// User service definition for internal microservice communication
var UserServiceService = exports.UserServiceService = {
  // Create a new user
createUser: {
    path: '/user.v1.UserService/CreateUser',
    requestStream: false,
    responseStream: false,
    requestType: proto_user_user_pb.CreateUserRequest,
    responseType: proto_user_user_pb.CreateUserResponse,
    requestSerialize: serialize_user_v1_CreateUserRequest,
    requestDeserialize: deserialize_user_v1_CreateUserRequest,
    responseSerialize: serialize_user_v1_CreateUserResponse,
    responseDeserialize: deserialize_user_v1_CreateUserResponse,
  },
  // Get user by ID
getUser: {
    path: '/user.v1.UserService/GetUser',
    requestStream: false,
    responseStream: false,
    requestType: proto_user_user_pb.GetUserRequest,
    responseType: proto_user_user_pb.GetUserResponse,
    requestSerialize: serialize_user_v1_GetUserRequest,
    requestDeserialize: deserialize_user_v1_GetUserRequest,
    responseSerialize: serialize_user_v1_GetUserResponse,
    responseDeserialize: deserialize_user_v1_GetUserResponse,
  },
  // Get user by email
getUserByEmail: {
    path: '/user.v1.UserService/GetUserByEmail',
    requestStream: false,
    responseStream: false,
    requestType: proto_user_user_pb.GetUserByEmailRequest,
    responseType: proto_user_user_pb.GetUserResponse,
    requestSerialize: serialize_user_v1_GetUserByEmailRequest,
    requestDeserialize: deserialize_user_v1_GetUserByEmailRequest,
    responseSerialize: serialize_user_v1_GetUserResponse,
    responseDeserialize: deserialize_user_v1_GetUserResponse,
  },
  // Update user information
updateUser: {
    path: '/user.v1.UserService/UpdateUser',
    requestStream: false,
    responseStream: false,
    requestType: proto_user_user_pb.UpdateUserRequest,
    responseType: proto_user_user_pb.UpdateUserResponse,
    requestSerialize: serialize_user_v1_UpdateUserRequest,
    requestDeserialize: deserialize_user_v1_UpdateUserRequest,
    responseSerialize: serialize_user_v1_UpdateUserResponse,
    responseDeserialize: deserialize_user_v1_UpdateUserResponse,
  },
  // Delete user
deleteUser: {
    path: '/user.v1.UserService/DeleteUser',
    requestStream: false,
    responseStream: false,
    requestType: proto_user_user_pb.DeleteUserRequest,
    responseType: proto_user_user_pb.DeleteUserResponse,
    requestSerialize: serialize_user_v1_DeleteUserRequest,
    requestDeserialize: deserialize_user_v1_DeleteUserRequest,
    responseSerialize: serialize_user_v1_DeleteUserResponse,
    responseDeserialize: deserialize_user_v1_DeleteUserResponse,
  },
  // Validate user credentials (for authentication)
validateUser: {
    path: '/user.v1.UserService/ValidateUser',
    requestStream: false,
    responseStream: false,
    requestType: proto_user_user_pb.ValidateUserRequest,
    responseType: proto_user_user_pb.ValidateUserResponse,
    requestSerialize: serialize_user_v1_ValidateUserRequest,
    requestDeserialize: deserialize_user_v1_ValidateUserRequest,
    responseSerialize: serialize_user_v1_ValidateUserResponse,
    responseDeserialize: deserialize_user_v1_ValidateUserResponse,
  },
  // Get multiple users by IDs (batch operation)
getUsers: {
    path: '/user.v1.UserService/GetUsers',
    requestStream: false,
    responseStream: false,
    requestType: proto_user_user_pb.GetUsersRequest,
    responseType: proto_user_user_pb.GetUsersResponse,
    requestSerialize: serialize_user_v1_GetUsersRequest,
    requestDeserialize: deserialize_user_v1_GetUsersRequest,
    responseSerialize: serialize_user_v1_GetUsersResponse,
    responseDeserialize: deserialize_user_v1_GetUsersResponse,
  },
};

exports.UserServiceClient = grpc.makeGenericClientConstructor(UserServiceService, 'UserService');
