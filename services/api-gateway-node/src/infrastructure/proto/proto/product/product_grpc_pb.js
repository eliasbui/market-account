// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var proto_product_product_pb = require('../../proto/product/product_pb.js');

function serialize_product_v1_CancelReservationRequest(arg) {
  if (!(arg instanceof proto_product_product_pb.CancelReservationRequest)) {
    throw new Error('Expected argument of type product.v1.CancelReservationRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_product_v1_CancelReservationRequest(buffer_arg) {
  return proto_product_product_pb.CancelReservationRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_product_v1_CancelReservationResponse(arg) {
  if (!(arg instanceof proto_product_product_pb.CancelReservationResponse)) {
    throw new Error('Expected argument of type product.v1.CancelReservationResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_product_v1_CancelReservationResponse(buffer_arg) {
  return proto_product_product_pb.CancelReservationResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_product_v1_CheckProductAvailabilityRequest(arg) {
  if (!(arg instanceof proto_product_product_pb.CheckProductAvailabilityRequest)) {
    throw new Error('Expected argument of type product.v1.CheckProductAvailabilityRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_product_v1_CheckProductAvailabilityRequest(buffer_arg) {
  return proto_product_product_pb.CheckProductAvailabilityRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_product_v1_CheckProductAvailabilityResponse(arg) {
  if (!(arg instanceof proto_product_product_pb.CheckProductAvailabilityResponse)) {
    throw new Error('Expected argument of type product.v1.CheckProductAvailabilityResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_product_v1_CheckProductAvailabilityResponse(buffer_arg) {
  return proto_product_product_pb.CheckProductAvailabilityResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_product_v1_ConfirmReservationRequest(arg) {
  if (!(arg instanceof proto_product_product_pb.ConfirmReservationRequest)) {
    throw new Error('Expected argument of type product.v1.ConfirmReservationRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_product_v1_ConfirmReservationRequest(buffer_arg) {
  return proto_product_product_pb.ConfirmReservationRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_product_v1_ConfirmReservationResponse(arg) {
  if (!(arg instanceof proto_product_product_pb.ConfirmReservationResponse)) {
    throw new Error('Expected argument of type product.v1.ConfirmReservationResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_product_v1_ConfirmReservationResponse(buffer_arg) {
  return proto_product_product_pb.ConfirmReservationResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_product_v1_CreateProductRequest(arg) {
  if (!(arg instanceof proto_product_product_pb.CreateProductRequest)) {
    throw new Error('Expected argument of type product.v1.CreateProductRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_product_v1_CreateProductRequest(buffer_arg) {
  return proto_product_product_pb.CreateProductRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_product_v1_CreateProductResponse(arg) {
  if (!(arg instanceof proto_product_product_pb.CreateProductResponse)) {
    throw new Error('Expected argument of type product.v1.CreateProductResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_product_v1_CreateProductResponse(buffer_arg) {
  return proto_product_product_pb.CreateProductResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_product_v1_DeleteProductRequest(arg) {
  if (!(arg instanceof proto_product_product_pb.DeleteProductRequest)) {
    throw new Error('Expected argument of type product.v1.DeleteProductRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_product_v1_DeleteProductRequest(buffer_arg) {
  return proto_product_product_pb.DeleteProductRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_product_v1_DeleteProductResponse(arg) {
  if (!(arg instanceof proto_product_product_pb.DeleteProductResponse)) {
    throw new Error('Expected argument of type product.v1.DeleteProductResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_product_v1_DeleteProductResponse(buffer_arg) {
  return proto_product_product_pb.DeleteProductResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_product_v1_GetProductRequest(arg) {
  if (!(arg instanceof proto_product_product_pb.GetProductRequest)) {
    throw new Error('Expected argument of type product.v1.GetProductRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_product_v1_GetProductRequest(buffer_arg) {
  return proto_product_product_pb.GetProductRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_product_v1_GetProductResponse(arg) {
  if (!(arg instanceof proto_product_product_pb.GetProductResponse)) {
    throw new Error('Expected argument of type product.v1.GetProductResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_product_v1_GetProductResponse(buffer_arg) {
  return proto_product_product_pb.GetProductResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_product_v1_GetProductsRequest(arg) {
  if (!(arg instanceof proto_product_product_pb.GetProductsRequest)) {
    throw new Error('Expected argument of type product.v1.GetProductsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_product_v1_GetProductsRequest(buffer_arg) {
  return proto_product_product_pb.GetProductsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_product_v1_GetProductsResponse(arg) {
  if (!(arg instanceof proto_product_product_pb.GetProductsResponse)) {
    throw new Error('Expected argument of type product.v1.GetProductsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_product_v1_GetProductsResponse(buffer_arg) {
  return proto_product_product_pb.GetProductsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_product_v1_ReserveProductsRequest(arg) {
  if (!(arg instanceof proto_product_product_pb.ReserveProductsRequest)) {
    throw new Error('Expected argument of type product.v1.ReserveProductsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_product_v1_ReserveProductsRequest(buffer_arg) {
  return proto_product_product_pb.ReserveProductsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_product_v1_ReserveProductsResponse(arg) {
  if (!(arg instanceof proto_product_product_pb.ReserveProductsResponse)) {
    throw new Error('Expected argument of type product.v1.ReserveProductsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_product_v1_ReserveProductsResponse(buffer_arg) {
  return proto_product_product_pb.ReserveProductsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_product_v1_UpdateProductPriceRequest(arg) {
  if (!(arg instanceof proto_product_product_pb.UpdateProductPriceRequest)) {
    throw new Error('Expected argument of type product.v1.UpdateProductPriceRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_product_v1_UpdateProductPriceRequest(buffer_arg) {
  return proto_product_product_pb.UpdateProductPriceRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_product_v1_UpdateProductPriceResponse(arg) {
  if (!(arg instanceof proto_product_product_pb.UpdateProductPriceResponse)) {
    throw new Error('Expected argument of type product.v1.UpdateProductPriceResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_product_v1_UpdateProductPriceResponse(buffer_arg) {
  return proto_product_product_pb.UpdateProductPriceResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_product_v1_UpdateProductRequest(arg) {
  if (!(arg instanceof proto_product_product_pb.UpdateProductRequest)) {
    throw new Error('Expected argument of type product.v1.UpdateProductRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_product_v1_UpdateProductRequest(buffer_arg) {
  return proto_product_product_pb.UpdateProductRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_product_v1_UpdateProductResponse(arg) {
  if (!(arg instanceof proto_product_product_pb.UpdateProductResponse)) {
    throw new Error('Expected argument of type product.v1.UpdateProductResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_product_v1_UpdateProductResponse(buffer_arg) {
  return proto_product_product_pb.UpdateProductResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_product_v1_UpdateProductStockRequest(arg) {
  if (!(arg instanceof proto_product_product_pb.UpdateProductStockRequest)) {
    throw new Error('Expected argument of type product.v1.UpdateProductStockRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_product_v1_UpdateProductStockRequest(buffer_arg) {
  return proto_product_product_pb.UpdateProductStockRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_product_v1_UpdateProductStockResponse(arg) {
  if (!(arg instanceof proto_product_product_pb.UpdateProductStockResponse)) {
    throw new Error('Expected argument of type product.v1.UpdateProductStockResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_product_v1_UpdateProductStockResponse(buffer_arg) {
  return proto_product_product_pb.UpdateProductStockResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// Product service definition for internal microservice communication
var ProductServiceService = exports.ProductServiceService = {
  // Create a new product
createProduct: {
    path: '/product.v1.ProductService/CreateProduct',
    requestStream: false,
    responseStream: false,
    requestType: proto_product_product_pb.CreateProductRequest,
    responseType: proto_product_product_pb.CreateProductResponse,
    requestSerialize: serialize_product_v1_CreateProductRequest,
    requestDeserialize: deserialize_product_v1_CreateProductRequest,
    responseSerialize: serialize_product_v1_CreateProductResponse,
    responseDeserialize: deserialize_product_v1_CreateProductResponse,
  },
  // Get product by ID
getProduct: {
    path: '/product.v1.ProductService/GetProduct',
    requestStream: false,
    responseStream: false,
    requestType: proto_product_product_pb.GetProductRequest,
    responseType: proto_product_product_pb.GetProductResponse,
    requestSerialize: serialize_product_v1_GetProductRequest,
    requestDeserialize: deserialize_product_v1_GetProductRequest,
    responseSerialize: serialize_product_v1_GetProductResponse,
    responseDeserialize: deserialize_product_v1_GetProductResponse,
  },
  // Get products with pagination and filters
getProducts: {
    path: '/product.v1.ProductService/GetProducts',
    requestStream: false,
    responseStream: false,
    requestType: proto_product_product_pb.GetProductsRequest,
    responseType: proto_product_product_pb.GetProductsResponse,
    requestSerialize: serialize_product_v1_GetProductsRequest,
    requestDeserialize: deserialize_product_v1_GetProductsRequest,
    responseSerialize: serialize_product_v1_GetProductsResponse,
    responseDeserialize: deserialize_product_v1_GetProductsResponse,
  },
  // Update product information
updateProduct: {
    path: '/product.v1.ProductService/UpdateProduct',
    requestStream: false,
    responseStream: false,
    requestType: proto_product_product_pb.UpdateProductRequest,
    responseType: proto_product_product_pb.UpdateProductResponse,
    requestSerialize: serialize_product_v1_UpdateProductRequest,
    requestDeserialize: deserialize_product_v1_UpdateProductRequest,
    responseSerialize: serialize_product_v1_UpdateProductResponse,
    responseDeserialize: deserialize_product_v1_UpdateProductResponse,
  },
  // Update product stock
updateProductStock: {
    path: '/product.v1.ProductService/UpdateProductStock',
    requestStream: false,
    responseStream: false,
    requestType: proto_product_product_pb.UpdateProductStockRequest,
    responseType: proto_product_product_pb.UpdateProductStockResponse,
    requestSerialize: serialize_product_v1_UpdateProductStockRequest,
    requestDeserialize: deserialize_product_v1_UpdateProductStockRequest,
    responseSerialize: serialize_product_v1_UpdateProductStockResponse,
    responseDeserialize: deserialize_product_v1_UpdateProductStockResponse,
  },
  // Update product price
updateProductPrice: {
    path: '/product.v1.ProductService/UpdateProductPrice',
    requestStream: false,
    responseStream: false,
    requestType: proto_product_product_pb.UpdateProductPriceRequest,
    responseType: proto_product_product_pb.UpdateProductPriceResponse,
    requestSerialize: serialize_product_v1_UpdateProductPriceRequest,
    requestDeserialize: deserialize_product_v1_UpdateProductPriceRequest,
    responseSerialize: serialize_product_v1_UpdateProductPriceResponse,
    responseDeserialize: deserialize_product_v1_UpdateProductPriceResponse,
  },
  // Delete product
deleteProduct: {
    path: '/product.v1.ProductService/DeleteProduct',
    requestStream: false,
    responseStream: false,
    requestType: proto_product_product_pb.DeleteProductRequest,
    responseType: proto_product_product_pb.DeleteProductResponse,
    requestSerialize: serialize_product_v1_DeleteProductRequest,
    requestDeserialize: deserialize_product_v1_DeleteProductRequest,
    responseSerialize: serialize_product_v1_DeleteProductResponse,
    responseDeserialize: deserialize_product_v1_DeleteProductResponse,
  },
  // Check product availability
checkProductAvailability: {
    path: '/product.v1.ProductService/CheckProductAvailability',
    requestStream: false,
    responseStream: false,
    requestType: proto_product_product_pb.CheckProductAvailabilityRequest,
    responseType: proto_product_product_pb.CheckProductAvailabilityResponse,
    requestSerialize: serialize_product_v1_CheckProductAvailabilityRequest,
    requestDeserialize: deserialize_product_v1_CheckProductAvailabilityRequest,
    responseSerialize: serialize_product_v1_CheckProductAvailabilityResponse,
    responseDeserialize: deserialize_product_v1_CheckProductAvailabilityResponse,
  },
  // Reserve products for order (reduces stock temporarily)
reserveProducts: {
    path: '/product.v1.ProductService/ReserveProducts',
    requestStream: false,
    responseStream: false,
    requestType: proto_product_product_pb.ReserveProductsRequest,
    responseType: proto_product_product_pb.ReserveProductsResponse,
    requestSerialize: serialize_product_v1_ReserveProductsRequest,
    requestDeserialize: deserialize_product_v1_ReserveProductsRequest,
    responseSerialize: serialize_product_v1_ReserveProductsResponse,
    responseDeserialize: deserialize_product_v1_ReserveProductsResponse,
  },
  // Confirm product reservation (finalizes stock reduction)
confirmReservation: {
    path: '/product.v1.ProductService/ConfirmReservation',
    requestStream: false,
    responseStream: false,
    requestType: proto_product_product_pb.ConfirmReservationRequest,
    responseType: proto_product_product_pb.ConfirmReservationResponse,
    requestSerialize: serialize_product_v1_ConfirmReservationRequest,
    requestDeserialize: deserialize_product_v1_ConfirmReservationRequest,
    responseSerialize: serialize_product_v1_ConfirmReservationResponse,
    responseDeserialize: deserialize_product_v1_ConfirmReservationResponse,
  },
  // Cancel product reservation (restores stock)
cancelReservation: {
    path: '/product.v1.ProductService/CancelReservation',
    requestStream: false,
    responseStream: false,
    requestType: proto_product_product_pb.CancelReservationRequest,
    responseType: proto_product_product_pb.CancelReservationResponse,
    requestSerialize: serialize_product_v1_CancelReservationRequest,
    requestDeserialize: deserialize_product_v1_CancelReservationRequest,
    responseSerialize: serialize_product_v1_CancelReservationResponse,
    responseDeserialize: deserialize_product_v1_CancelReservationResponse,
  },
};

exports.ProductServiceClient = grpc.makeGenericClientConstructor(ProductServiceService, 'ProductService');
