// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var proto_payment_payment_pb = require('../../proto/payment/payment_pb.js');

function serialize_payment_v1_CancelPaymentRequest(arg) {
  if (!(arg instanceof proto_payment_payment_pb.CancelPaymentRequest)) {
    throw new Error('Expected argument of type payment.v1.CancelPaymentRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_payment_v1_CancelPaymentRequest(buffer_arg) {
  return proto_payment_payment_pb.CancelPaymentRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_payment_v1_CancelPaymentResponse(arg) {
  if (!(arg instanceof proto_payment_payment_pb.CancelPaymentResponse)) {
    throw new Error('Expected argument of type payment.v1.CancelPaymentResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_payment_v1_CancelPaymentResponse(buffer_arg) {
  return proto_payment_payment_pb.CancelPaymentResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_payment_v1_ConfirmPaymentRequest(arg) {
  if (!(arg instanceof proto_payment_payment_pb.ConfirmPaymentRequest)) {
    throw new Error('Expected argument of type payment.v1.ConfirmPaymentRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_payment_v1_ConfirmPaymentRequest(buffer_arg) {
  return proto_payment_payment_pb.ConfirmPaymentRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_payment_v1_ConfirmPaymentResponse(arg) {
  if (!(arg instanceof proto_payment_payment_pb.ConfirmPaymentResponse)) {
    throw new Error('Expected argument of type payment.v1.ConfirmPaymentResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_payment_v1_ConfirmPaymentResponse(buffer_arg) {
  return proto_payment_payment_pb.ConfirmPaymentResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_payment_v1_CreatePaymentRequest(arg) {
  if (!(arg instanceof proto_payment_payment_pb.CreatePaymentRequest)) {
    throw new Error('Expected argument of type payment.v1.CreatePaymentRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_payment_v1_CreatePaymentRequest(buffer_arg) {
  return proto_payment_payment_pb.CreatePaymentRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_payment_v1_CreatePaymentResponse(arg) {
  if (!(arg instanceof proto_payment_payment_pb.CreatePaymentResponse)) {
    throw new Error('Expected argument of type payment.v1.CreatePaymentResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_payment_v1_CreatePaymentResponse(buffer_arg) {
  return proto_payment_payment_pb.CreatePaymentResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_payment_v1_GetOrderPaymentsRequest(arg) {
  if (!(arg instanceof proto_payment_payment_pb.GetOrderPaymentsRequest)) {
    throw new Error('Expected argument of type payment.v1.GetOrderPaymentsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_payment_v1_GetOrderPaymentsRequest(buffer_arg) {
  return proto_payment_payment_pb.GetOrderPaymentsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_payment_v1_GetOrderPaymentsResponse(arg) {
  if (!(arg instanceof proto_payment_payment_pb.GetOrderPaymentsResponse)) {
    throw new Error('Expected argument of type payment.v1.GetOrderPaymentsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_payment_v1_GetOrderPaymentsResponse(buffer_arg) {
  return proto_payment_payment_pb.GetOrderPaymentsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_payment_v1_GetPaymentRequest(arg) {
  if (!(arg instanceof proto_payment_payment_pb.GetPaymentRequest)) {
    throw new Error('Expected argument of type payment.v1.GetPaymentRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_payment_v1_GetPaymentRequest(buffer_arg) {
  return proto_payment_payment_pb.GetPaymentRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_payment_v1_GetPaymentResponse(arg) {
  if (!(arg instanceof proto_payment_payment_pb.GetPaymentResponse)) {
    throw new Error('Expected argument of type payment.v1.GetPaymentResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_payment_v1_GetPaymentResponse(buffer_arg) {
  return proto_payment_payment_pb.GetPaymentResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_payment_v1_GetUserPaymentsRequest(arg) {
  if (!(arg instanceof proto_payment_payment_pb.GetUserPaymentsRequest)) {
    throw new Error('Expected argument of type payment.v1.GetUserPaymentsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_payment_v1_GetUserPaymentsRequest(buffer_arg) {
  return proto_payment_payment_pb.GetUserPaymentsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_payment_v1_GetUserPaymentsResponse(arg) {
  if (!(arg instanceof proto_payment_payment_pb.GetUserPaymentsResponse)) {
    throw new Error('Expected argument of type payment.v1.GetUserPaymentsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_payment_v1_GetUserPaymentsResponse(buffer_arg) {
  return proto_payment_payment_pb.GetUserPaymentsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_payment_v1_ProcessPaymentRequest(arg) {
  if (!(arg instanceof proto_payment_payment_pb.ProcessPaymentRequest)) {
    throw new Error('Expected argument of type payment.v1.ProcessPaymentRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_payment_v1_ProcessPaymentRequest(buffer_arg) {
  return proto_payment_payment_pb.ProcessPaymentRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_payment_v1_ProcessPaymentResponse(arg) {
  if (!(arg instanceof proto_payment_payment_pb.ProcessPaymentResponse)) {
    throw new Error('Expected argument of type payment.v1.ProcessPaymentResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_payment_v1_ProcessPaymentResponse(buffer_arg) {
  return proto_payment_payment_pb.ProcessPaymentResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_payment_v1_RefundPaymentRequest(arg) {
  if (!(arg instanceof proto_payment_payment_pb.RefundPaymentRequest)) {
    throw new Error('Expected argument of type payment.v1.RefundPaymentRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_payment_v1_RefundPaymentRequest(buffer_arg) {
  return proto_payment_payment_pb.RefundPaymentRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_payment_v1_RefundPaymentResponse(arg) {
  if (!(arg instanceof proto_payment_payment_pb.RefundPaymentResponse)) {
    throw new Error('Expected argument of type payment.v1.RefundPaymentResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_payment_v1_RefundPaymentResponse(buffer_arg) {
  return proto_payment_payment_pb.RefundPaymentResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_payment_v1_UpdatePaymentStatusRequest(arg) {
  if (!(arg instanceof proto_payment_payment_pb.UpdatePaymentStatusRequest)) {
    throw new Error('Expected argument of type payment.v1.UpdatePaymentStatusRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_payment_v1_UpdatePaymentStatusRequest(buffer_arg) {
  return proto_payment_payment_pb.UpdatePaymentStatusRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_payment_v1_UpdatePaymentStatusResponse(arg) {
  if (!(arg instanceof proto_payment_payment_pb.UpdatePaymentStatusResponse)) {
    throw new Error('Expected argument of type payment.v1.UpdatePaymentStatusResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_payment_v1_UpdatePaymentStatusResponse(buffer_arg) {
  return proto_payment_payment_pb.UpdatePaymentStatusResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_payment_v1_ValidatePaymentProviderRequest(arg) {
  if (!(arg instanceof proto_payment_payment_pb.ValidatePaymentProviderRequest)) {
    throw new Error('Expected argument of type payment.v1.ValidatePaymentProviderRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_payment_v1_ValidatePaymentProviderRequest(buffer_arg) {
  return proto_payment_payment_pb.ValidatePaymentProviderRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_payment_v1_ValidatePaymentProviderResponse(arg) {
  if (!(arg instanceof proto_payment_payment_pb.ValidatePaymentProviderResponse)) {
    throw new Error('Expected argument of type payment.v1.ValidatePaymentProviderResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_payment_v1_ValidatePaymentProviderResponse(buffer_arg) {
  return proto_payment_payment_pb.ValidatePaymentProviderResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// Payment service definition for internal microservice communication
var PaymentServiceService = exports.PaymentServiceService = {
  // Create a new payment
createPayment: {
    path: '/payment.v1.PaymentService/CreatePayment',
    requestStream: false,
    responseStream: false,
    requestType: proto_payment_payment_pb.CreatePaymentRequest,
    responseType: proto_payment_payment_pb.CreatePaymentResponse,
    requestSerialize: serialize_payment_v1_CreatePaymentRequest,
    requestDeserialize: deserialize_payment_v1_CreatePaymentRequest,
    responseSerialize: serialize_payment_v1_CreatePaymentResponse,
    responseDeserialize: deserialize_payment_v1_CreatePaymentResponse,
  },
  // Get payment by ID
getPayment: {
    path: '/payment.v1.PaymentService/GetPayment',
    requestStream: false,
    responseStream: false,
    requestType: proto_payment_payment_pb.GetPaymentRequest,
    responseType: proto_payment_payment_pb.GetPaymentResponse,
    requestSerialize: serialize_payment_v1_GetPaymentRequest,
    requestDeserialize: deserialize_payment_v1_GetPaymentRequest,
    responseSerialize: serialize_payment_v1_GetPaymentResponse,
    responseDeserialize: deserialize_payment_v1_GetPaymentResponse,
  },
  // Get payments by user ID
getUserPayments: {
    path: '/payment.v1.PaymentService/GetUserPayments',
    requestStream: false,
    responseStream: false,
    requestType: proto_payment_payment_pb.GetUserPaymentsRequest,
    responseType: proto_payment_payment_pb.GetUserPaymentsResponse,
    requestSerialize: serialize_payment_v1_GetUserPaymentsRequest,
    requestDeserialize: deserialize_payment_v1_GetUserPaymentsRequest,
    responseSerialize: serialize_payment_v1_GetUserPaymentsResponse,
    responseDeserialize: deserialize_payment_v1_GetUserPaymentsResponse,
  },
  // Get payments by order ID
getOrderPayments: {
    path: '/payment.v1.PaymentService/GetOrderPayments',
    requestStream: false,
    responseStream: false,
    requestType: proto_payment_payment_pb.GetOrderPaymentsRequest,
    responseType: proto_payment_payment_pb.GetOrderPaymentsResponse,
    requestSerialize: serialize_payment_v1_GetOrderPaymentsRequest,
    requestDeserialize: deserialize_payment_v1_GetOrderPaymentsRequest,
    responseSerialize: serialize_payment_v1_GetOrderPaymentsResponse,
    responseDeserialize: deserialize_payment_v1_GetOrderPaymentsResponse,
  },
  // Process payment (initiate payment with provider)
processPayment: {
    path: '/payment.v1.PaymentService/ProcessPayment',
    requestStream: false,
    responseStream: false,
    requestType: proto_payment_payment_pb.ProcessPaymentRequest,
    responseType: proto_payment_payment_pb.ProcessPaymentResponse,
    requestSerialize: serialize_payment_v1_ProcessPaymentRequest,
    requestDeserialize: deserialize_payment_v1_ProcessPaymentRequest,
    responseSerialize: serialize_payment_v1_ProcessPaymentResponse,
    responseDeserialize: deserialize_payment_v1_ProcessPaymentResponse,
  },
  // Confirm payment (mark as completed)
confirmPayment: {
    path: '/payment.v1.PaymentService/ConfirmPayment',
    requestStream: false,
    responseStream: false,
    requestType: proto_payment_payment_pb.ConfirmPaymentRequest,
    responseType: proto_payment_payment_pb.ConfirmPaymentResponse,
    requestSerialize: serialize_payment_v1_ConfirmPaymentRequest,
    requestDeserialize: deserialize_payment_v1_ConfirmPaymentRequest,
    responseSerialize: serialize_payment_v1_ConfirmPaymentResponse,
    responseDeserialize: deserialize_payment_v1_ConfirmPaymentResponse,
  },
  // Cancel payment
cancelPayment: {
    path: '/payment.v1.PaymentService/CancelPayment',
    requestStream: false,
    responseStream: false,
    requestType: proto_payment_payment_pb.CancelPaymentRequest,
    responseType: proto_payment_payment_pb.CancelPaymentResponse,
    requestSerialize: serialize_payment_v1_CancelPaymentRequest,
    requestDeserialize: deserialize_payment_v1_CancelPaymentRequest,
    responseSerialize: serialize_payment_v1_CancelPaymentResponse,
    responseDeserialize: deserialize_payment_v1_CancelPaymentResponse,
  },
  // Refund payment
refundPayment: {
    path: '/payment.v1.PaymentService/RefundPayment',
    requestStream: false,
    responseStream: false,
    requestType: proto_payment_payment_pb.RefundPaymentRequest,
    responseType: proto_payment_payment_pb.RefundPaymentResponse,
    requestSerialize: serialize_payment_v1_RefundPaymentRequest,
    requestDeserialize: deserialize_payment_v1_RefundPaymentRequest,
    responseSerialize: serialize_payment_v1_RefundPaymentResponse,
    responseDeserialize: deserialize_payment_v1_RefundPaymentResponse,
  },
  // Update payment status (webhook handling)
updatePaymentStatus: {
    path: '/payment.v1.PaymentService/UpdatePaymentStatus',
    requestStream: false,
    responseStream: false,
    requestType: proto_payment_payment_pb.UpdatePaymentStatusRequest,
    responseType: proto_payment_payment_pb.UpdatePaymentStatusResponse,
    requestSerialize: serialize_payment_v1_UpdatePaymentStatusRequest,
    requestDeserialize: deserialize_payment_v1_UpdatePaymentStatusRequest,
    responseSerialize: serialize_payment_v1_UpdatePaymentStatusResponse,
    responseDeserialize: deserialize_payment_v1_UpdatePaymentStatusResponse,
  },
  // Validate payment provider
validatePaymentProvider: {
    path: '/payment.v1.PaymentService/ValidatePaymentProvider',
    requestStream: false,
    responseStream: false,
    requestType: proto_payment_payment_pb.ValidatePaymentProviderRequest,
    responseType: proto_payment_payment_pb.ValidatePaymentProviderResponse,
    requestSerialize: serialize_payment_v1_ValidatePaymentProviderRequest,
    requestDeserialize: deserialize_payment_v1_ValidatePaymentProviderRequest,
    responseSerialize: serialize_payment_v1_ValidatePaymentProviderResponse,
    responseDeserialize: deserialize_payment_v1_ValidatePaymentProviderResponse,
  },
};

exports.PaymentServiceClient = grpc.makeGenericClientConstructor(PaymentServiceService, 'PaymentService');
