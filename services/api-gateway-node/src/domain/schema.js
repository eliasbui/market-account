import { gql } from 'apollo-server-express';

export const typeDefs = gql`
  # User Types
  type User {
    id: ID!
    email: String!
    username: String!
    isActive: Boolean!
    createdAt: String!
    updatedAt: String!
  }

  type UserCreateResponse {
    user: User
    error: String
  }

  # Product Types
  type Product {
    id: ID!
    name: String!
    description: String
    price: Float!
    stockQuantity: Int!
    category: String!
    sku: String
    isActive: Boolean!
    createdAt: String!
    updatedAt: String!
  }

  type ProductListResponse {
    products: [Product!]!
    total: Int!
    page: Int!
    limit: Int!
  }

  # Payment Types
  enum PaymentStatus {
    PENDING
    PROCESSING
    COMPLETED
    FAILED
    CANCELLED
    REFUNDED
  }

  enum PaymentMethod {
    CREDIT_CARD
    DEBIT_CARD
    BANK_TRANSFER
    DIGITAL_WALLET
    CRYPTO
  }

  type Payment {
    id: ID!
    orderId: String!
    userId: String!
    amount: Float!
    currency: String!
    paymentMethod: PaymentMethod!
    status: PaymentStatus!
    providerTransactionId: String
    createdAt: String!
    updatedAt: String!
    completedAt: String
    failureReason: String
    refundAmount: Float
  }

  # Input Types
  input CreateUserInput {
    email: String!
    username: String!
    password: String!
  }

  input CreateProductInput {
    name: String!
    description: String
    price: Float!
    stockQuantity: Int!
    category: String!
    sku: String
  }

  input CreatePaymentInput {
    orderId: String!
    userId: String!
    amount: Float!
    currency: String!
    paymentMethod: PaymentMethod!
  }

  input ProductFilters {
    category: String
    minPrice: Float
    maxPrice: Float
    inStock: Boolean
  }

  # Root Types
  type Query {
    # Health Check
    health: String!
    
    # User Queries
    user(id: ID!): User
    userByEmail(email: String!): User
    
    # Product Queries
    product(id: ID!): Product
    products(page: Int = 1, limit: Int = 10, filters: ProductFilters): ProductListResponse!
    
    # Payment Queries
    payment(id: ID!): Payment
    userPayments(userId: ID!, page: Int = 1, limit: Int = 10): [Payment!]!
  }

  type Mutation {
    # User Mutations
    createUser(input: CreateUserInput!): UserCreateResponse!
    deleteUser(id: ID!): Boolean!
    
    # Product Mutations
    createProduct(input: CreateProductInput!): Product!
    updateProductStock(id: ID!, quantity: Int!): Product!
    updateProductPrice(id: ID!, price: Float!): Product!
    
    # Payment Mutations
    createPayment(input: CreatePaymentInput!): Payment!
    processPayment(id: ID!): Payment!
    refundPayment(id: ID!, amount: Float): Payment!
  }

  type Subscription {
    # Real-time updates
    paymentStatusUpdated(userId: ID!): Payment!
    productStockUpdated(productId: ID!): Product!
  }
`;

export default typeDefs; 