# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: proto/product/product.proto
# Protobuf Python Version: 6.31.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    6,
    31,
    0,
    '',
    'proto/product/product.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1bproto/product/product.proto\x12\nproduct.v1\"\xad\x02\n\x07Product\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12\r\n\x05price\x18\x04 \x01(\x01\x12\x16\n\x0estock_quantity\x18\x05 \x01(\x05\x12\x10\n\x08\x63\x61tegory\x18\x06 \x01(\t\x12\x0b\n\x03sku\x18\x07 \x01(\t\x12\x11\n\tis_active\x18\x08 \x01(\x08\x12\x12\n\ncreated_at\x18\t \x01(\x03\x12\x12\n\nupdated_at\x18\n \x01(\x03\x12\x0c\n\x04tags\x18\x0b \x03(\t\x12\x33\n\x08metadata\x18\x0c \x03(\x0b\x32!.product.v1.Product.MetadataEntry\x1a/\n\rMetadataEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\xd8\x01\n\x0eProductFilters\x12\x15\n\x08\x63\x61tegory\x18\x01 \x01(\tH\x00\x88\x01\x01\x12\x16\n\tmin_price\x18\x02 \x01(\x01H\x01\x88\x01\x01\x12\x16\n\tmax_price\x18\x03 \x01(\x01H\x02\x88\x01\x01\x12\x15\n\x08in_stock\x18\x04 \x01(\x08H\x03\x88\x01\x01\x12\x0c\n\x04tags\x18\x05 \x03(\t\x12\x16\n\tis_active\x18\x06 \x01(\x08H\x04\x88\x01\x01\x42\x0b\n\t_categoryB\x0c\n\n_min_priceB\x0c\n\n_max_priceB\x0b\n\t_in_stockB\x0c\n\n_is_active\"M\n\nPagination\x12\x0c\n\x04page\x18\x01 \x01(\x05\x12\r\n\x05limit\x18\x02 \x01(\x05\x12\r\n\x05total\x18\x03 \x01(\x05\x12\x13\n\x0btotal_pages\x18\x04 \x01(\x05\">\n\x16ProductReservationItem\x12\x12\n\nproduct_id\x18\x01 \x01(\t\x12\x10\n\x08quantity\x18\x02 \x01(\x05\"\x80\x02\n\x14\x43reateProductRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x02 \x01(\t\x12\r\n\x05price\x18\x03 \x01(\x01\x12\x16\n\x0estock_quantity\x18\x04 \x01(\x05\x12\x10\n\x08\x63\x61tegory\x18\x05 \x01(\t\x12\x0b\n\x03sku\x18\x06 \x01(\t\x12\x0c\n\x04tags\x18\x07 \x03(\t\x12@\n\x08metadata\x18\x08 \x03(\x0b\x32..product.v1.CreateProductRequest.MetadataEntry\x1a/\n\rMetadataEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\x1f\n\x11GetProductRequest\x12\n\n\x02id\x18\x01 \x01(\t\"\x92\x01\n\x12GetProductsRequest\x12\x30\n\x07\x66ilters\x18\x01 \x01(\x0b\x32\x1a.product.v1.ProductFiltersH\x00\x88\x01\x01\x12/\n\npagination\x18\x02 \x01(\x0b\x32\x16.product.v1.PaginationH\x01\x88\x01\x01\x42\n\n\x08_filtersB\r\n\x0b_pagination\"\x9b\x03\n\x14UpdateProductRequest\x12\n\n\x02id\x18\x01 \x01(\t\x12\x11\n\x04name\x18\x02 \x01(\tH\x00\x88\x01\x01\x12\x18\n\x0b\x64\x65scription\x18\x03 \x01(\tH\x01\x88\x01\x01\x12\x12\n\x05price\x18\x04 \x01(\x01H\x02\x88\x01\x01\x12\x1b\n\x0estock_quantity\x18\x05 \x01(\x05H\x03\x88\x01\x01\x12\x15\n\x08\x63\x61tegory\x18\x06 \x01(\tH\x04\x88\x01\x01\x12\x10\n\x03sku\x18\x07 \x01(\tH\x05\x88\x01\x01\x12\x16\n\tis_active\x18\x08 \x01(\x08H\x06\x88\x01\x01\x12\x0c\n\x04tags\x18\t \x03(\t\x12@\n\x08metadata\x18\n \x03(\x0b\x32..product.v1.UpdateProductRequest.MetadataEntry\x1a/\n\rMetadataEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x42\x07\n\x05_nameB\x0e\n\x0c_descriptionB\x08\n\x06_priceB\x11\n\x0f_stock_quantityB\x0b\n\t_categoryB\x06\n\x04_skuB\x0c\n\n_is_active\"9\n\x19UpdateProductStockRequest\x12\n\n\x02id\x18\x01 \x01(\t\x12\x10\n\x08quantity\x18\x02 \x01(\x05\"6\n\x19UpdateProductPriceRequest\x12\n\n\x02id\x18\x01 \x01(\t\x12\r\n\x05price\x18\x02 \x01(\x01\"\"\n\x14\x44\x65leteProductRequest\x12\n\n\x02id\x18\x01 \x01(\t\"T\n\x1f\x43heckProductAvailabilityRequest\x12\x31\n\x05items\x18\x01 \x03(\x0b\x32\".product.v1.ProductReservationItem\"x\n\x16ReserveProductsRequest\x12\x16\n\x0ereservation_id\x18\x01 \x01(\t\x12\x31\n\x05items\x18\x02 \x03(\x0b\x32\".product.v1.ProductReservationItem\x12\x13\n\x0b\x65xpiry_time\x18\x03 \x01(\x03\"3\n\x19\x43onfirmReservationRequest\x12\x16\n\x0ereservation_id\x18\x01 \x01(\t\"2\n\x18\x43\x61ncelReservationRequest\x12\x16\n\x0ereservation_id\x18\x01 \x01(\t\"m\n\x15\x43reateProductResponse\x12&\n\x07product\x18\x01 \x01(\x0b\x32\x13.product.v1.ProductH\x00\x12\"\n\x05\x65rror\x18\x02 \x01(\x0b\x32\x11.product.v1.ErrorH\x00\x42\x08\n\x06result\"j\n\x12GetProductResponse\x12&\n\x07product\x18\x01 \x01(\x0b\x32\x13.product.v1.ProductH\x00\x12\"\n\x05\x65rror\x18\x02 \x01(\x0b\x32\x11.product.v1.ErrorH\x00\x42\x08\n\x06result\"\xad\x01\n\x13GetProductsResponse\x12%\n\x08products\x18\x01 \x03(\x0b\x32\x13.product.v1.Product\x12/\n\npagination\x18\x02 \x01(\x0b\x32\x16.product.v1.PaginationH\x00\x88\x01\x01\x12%\n\x05\x65rror\x18\x03 \x01(\x0b\x32\x11.product.v1.ErrorH\x01\x88\x01\x01\x42\r\n\x0b_paginationB\x08\n\x06_error\"m\n\x15UpdateProductResponse\x12&\n\x07product\x18\x01 \x01(\x0b\x32\x13.product.v1.ProductH\x00\x12\"\n\x05\x65rror\x18\x02 \x01(\x0b\x32\x11.product.v1.ErrorH\x00\x42\x08\n\x06result\"r\n\x1aUpdateProductStockResponse\x12&\n\x07product\x18\x01 \x01(\x0b\x32\x13.product.v1.ProductH\x00\x12\"\n\x05\x65rror\x18\x02 \x01(\x0b\x32\x11.product.v1.ErrorH\x00\x42\x08\n\x06result\"r\n\x1aUpdateProductPriceResponse\x12&\n\x07product\x18\x01 \x01(\x0b\x32\x13.product.v1.ProductH\x00\x12\"\n\x05\x65rror\x18\x02 \x01(\x0b\x32\x11.product.v1.ErrorH\x00\x42\x08\n\x06result\"Y\n\x15\x44\x65leteProductResponse\x12\x0f\n\x07success\x18\x01 \x01(\x08\x12%\n\x05\x65rror\x18\x02 \x01(\x0b\x32\x11.product.v1.ErrorH\x00\x88\x01\x01\x42\x08\n\x06_error\"[\n\x13ProductAvailability\x12\x12\n\nproduct_id\x18\x01 \x01(\t\x12\x14\n\x0cis_available\x18\x02 \x01(\x08\x12\x1a\n\x12\x61vailable_quantity\x18\x03 \x01(\x05\"\x8a\x01\n CheckProductAvailabilityResponse\x12\x35\n\x0c\x61vailability\x18\x01 \x03(\x0b\x32\x1f.product.v1.ProductAvailability\x12%\n\x05\x65rror\x18\x02 \x01(\x0b\x32\x11.product.v1.ErrorH\x00\x88\x01\x01\x42\x08\n\x06_error\"\xac\x01\n\x17ReserveProductsResponse\x12\x0f\n\x07success\x18\x01 \x01(\x08\x12\x16\n\x0ereservation_id\x18\x02 \x01(\t\x12\x37\n\x0ereserved_items\x18\x03 \x03(\x0b\x32\x1f.product.v1.ProductAvailability\x12%\n\x05\x65rror\x18\x04 \x01(\x0b\x32\x11.product.v1.ErrorH\x00\x88\x01\x01\x42\x08\n\x06_error\"^\n\x1a\x43onfirmReservationResponse\x12\x0f\n\x07success\x18\x01 \x01(\x08\x12%\n\x05\x65rror\x18\x02 \x01(\x0b\x32\x11.product.v1.ErrorH\x00\x88\x01\x01\x42\x08\n\x06_error\"]\n\x19\x43\x61ncelReservationResponse\x12\x0f\n\x07success\x18\x01 \x01(\x08\x12%\n\x05\x65rror\x18\x02 \x01(\x0b\x32\x11.product.v1.ErrorH\x00\x88\x01\x01\x42\x08\n\x06_error\"\x87\x01\n\x05\x45rror\x12\x0c\n\x04\x63ode\x18\x01 \x01(\t\x12\x0f\n\x07message\x18\x02 \x01(\t\x12/\n\x07\x64\x65tails\x18\x03 \x03(\x0b\x32\x1e.product.v1.Error.DetailsEntry\x1a.\n\x0c\x44\x65tailsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x32\x93\x08\n\x0eProductService\x12T\n\rCreateProduct\x12 .product.v1.CreateProductRequest\x1a!.product.v1.CreateProductResponse\x12K\n\nGetProduct\x12\x1d.product.v1.GetProductRequest\x1a\x1e.product.v1.GetProductResponse\x12N\n\x0bGetProducts\x12\x1e.product.v1.GetProductsRequest\x1a\x1f.product.v1.GetProductsResponse\x12T\n\rUpdateProduct\x12 .product.v1.UpdateProductRequest\x1a!.product.v1.UpdateProductResponse\x12\x63\n\x12UpdateProductStock\x12%.product.v1.UpdateProductStockRequest\x1a&.product.v1.UpdateProductStockResponse\x12\x63\n\x12UpdateProductPrice\x12%.product.v1.UpdateProductPriceRequest\x1a&.product.v1.UpdateProductPriceResponse\x12T\n\rDeleteProduct\x12 .product.v1.DeleteProductRequest\x1a!.product.v1.DeleteProductResponse\x12u\n\x18\x43heckProductAvailability\x12+.product.v1.CheckProductAvailabilityRequest\x1a,.product.v1.CheckProductAvailabilityResponse\x12Z\n\x0fReserveProducts\x12\".product.v1.ReserveProductsRequest\x1a#.product.v1.ReserveProductsResponse\x12\x63\n\x12\x43onfirmReservation\x12%.product.v1.ConfirmReservationRequest\x1a&.product.v1.ConfirmReservationResponse\x12`\n\x11\x43\x61ncelReservation\x12$.product.v1.CancelReservationRequest\x1a%.product.v1.CancelReservationResponseBa\n\x1a\x63om.marketplace.product.v1Z$user-service/proto/product;productpb\xaa\x02\x1cProductService.Proto.Productb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'proto.product.product_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\032com.marketplace.product.v1Z$user-service/proto/product;productpb\252\002\034ProductService.Proto.Product'
  _globals['_PRODUCT_METADATAENTRY']._loaded_options = None
  _globals['_PRODUCT_METADATAENTRY']._serialized_options = b'8\001'
  _globals['_CREATEPRODUCTREQUEST_METADATAENTRY']._loaded_options = None
  _globals['_CREATEPRODUCTREQUEST_METADATAENTRY']._serialized_options = b'8\001'
  _globals['_UPDATEPRODUCTREQUEST_METADATAENTRY']._loaded_options = None
  _globals['_UPDATEPRODUCTREQUEST_METADATAENTRY']._serialized_options = b'8\001'
  _globals['_ERROR_DETAILSENTRY']._loaded_options = None
  _globals['_ERROR_DETAILSENTRY']._serialized_options = b'8\001'
  _globals['_PRODUCT']._serialized_start=44
  _globals['_PRODUCT']._serialized_end=345
  _globals['_PRODUCT_METADATAENTRY']._serialized_start=298
  _globals['_PRODUCT_METADATAENTRY']._serialized_end=345
  _globals['_PRODUCTFILTERS']._serialized_start=348
  _globals['_PRODUCTFILTERS']._serialized_end=564
  _globals['_PAGINATION']._serialized_start=566
  _globals['_PAGINATION']._serialized_end=643
  _globals['_PRODUCTRESERVATIONITEM']._serialized_start=645
  _globals['_PRODUCTRESERVATIONITEM']._serialized_end=707
  _globals['_CREATEPRODUCTREQUEST']._serialized_start=710
  _globals['_CREATEPRODUCTREQUEST']._serialized_end=966
  _globals['_CREATEPRODUCTREQUEST_METADATAENTRY']._serialized_start=298
  _globals['_CREATEPRODUCTREQUEST_METADATAENTRY']._serialized_end=345
  _globals['_GETPRODUCTREQUEST']._serialized_start=968
  _globals['_GETPRODUCTREQUEST']._serialized_end=999
  _globals['_GETPRODUCTSREQUEST']._serialized_start=1002
  _globals['_GETPRODUCTSREQUEST']._serialized_end=1148
  _globals['_UPDATEPRODUCTREQUEST']._serialized_start=1151
  _globals['_UPDATEPRODUCTREQUEST']._serialized_end=1562
  _globals['_UPDATEPRODUCTREQUEST_METADATAENTRY']._serialized_start=298
  _globals['_UPDATEPRODUCTREQUEST_METADATAENTRY']._serialized_end=345
  _globals['_UPDATEPRODUCTSTOCKREQUEST']._serialized_start=1564
  _globals['_UPDATEPRODUCTSTOCKREQUEST']._serialized_end=1621
  _globals['_UPDATEPRODUCTPRICEREQUEST']._serialized_start=1623
  _globals['_UPDATEPRODUCTPRICEREQUEST']._serialized_end=1677
  _globals['_DELETEPRODUCTREQUEST']._serialized_start=1679
  _globals['_DELETEPRODUCTREQUEST']._serialized_end=1713
  _globals['_CHECKPRODUCTAVAILABILITYREQUEST']._serialized_start=1715
  _globals['_CHECKPRODUCTAVAILABILITYREQUEST']._serialized_end=1799
  _globals['_RESERVEPRODUCTSREQUEST']._serialized_start=1801
  _globals['_RESERVEPRODUCTSREQUEST']._serialized_end=1921
  _globals['_CONFIRMRESERVATIONREQUEST']._serialized_start=1923
  _globals['_CONFIRMRESERVATIONREQUEST']._serialized_end=1974
  _globals['_CANCELRESERVATIONREQUEST']._serialized_start=1976
  _globals['_CANCELRESERVATIONREQUEST']._serialized_end=2026
  _globals['_CREATEPRODUCTRESPONSE']._serialized_start=2028
  _globals['_CREATEPRODUCTRESPONSE']._serialized_end=2137
  _globals['_GETPRODUCTRESPONSE']._serialized_start=2139
  _globals['_GETPRODUCTRESPONSE']._serialized_end=2245
  _globals['_GETPRODUCTSRESPONSE']._serialized_start=2248
  _globals['_GETPRODUCTSRESPONSE']._serialized_end=2421
  _globals['_UPDATEPRODUCTRESPONSE']._serialized_start=2423
  _globals['_UPDATEPRODUCTRESPONSE']._serialized_end=2532
  _globals['_UPDATEPRODUCTSTOCKRESPONSE']._serialized_start=2534
  _globals['_UPDATEPRODUCTSTOCKRESPONSE']._serialized_end=2648
  _globals['_UPDATEPRODUCTPRICERESPONSE']._serialized_start=2650
  _globals['_UPDATEPRODUCTPRICERESPONSE']._serialized_end=2764
  _globals['_DELETEPRODUCTRESPONSE']._serialized_start=2766
  _globals['_DELETEPRODUCTRESPONSE']._serialized_end=2855
  _globals['_PRODUCTAVAILABILITY']._serialized_start=2857
  _globals['_PRODUCTAVAILABILITY']._serialized_end=2948
  _globals['_CHECKPRODUCTAVAILABILITYRESPONSE']._serialized_start=2951
  _globals['_CHECKPRODUCTAVAILABILITYRESPONSE']._serialized_end=3089
  _globals['_RESERVEPRODUCTSRESPONSE']._serialized_start=3092
  _globals['_RESERVEPRODUCTSRESPONSE']._serialized_end=3264
  _globals['_CONFIRMRESERVATIONRESPONSE']._serialized_start=3266
  _globals['_CONFIRMRESERVATIONRESPONSE']._serialized_end=3360
  _globals['_CANCELRESERVATIONRESPONSE']._serialized_start=3362
  _globals['_CANCELRESERVATIONRESPONSE']._serialized_end=3455
  _globals['_ERROR']._serialized_start=3458
  _globals['_ERROR']._serialized_end=3593
  _globals['_ERROR_DETAILSENTRY']._serialized_start=3547
  _globals['_ERROR_DETAILSENTRY']._serialized_end=3593
  _globals['_PRODUCTSERVICE']._serialized_start=3596
  _globals['_PRODUCTSERVICE']._serialized_end=4639
# @@protoc_insertion_point(module_scope)
