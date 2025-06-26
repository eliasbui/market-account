from dataclasses import dataclass
from decimal import Decimal
from datetime import datetime
from enum import Enum
from typing import Optional
import uuid


class PaymentStatus(Enum):
    PENDING = "pending"
    PROCESSING = "processing"
    COMPLETED = "completed"
    FAILED = "failed"
    CANCELLED = "cancelled"
    REFUNDED = "refunded"


class PaymentMethod(Enum):
    CREDIT_CARD = "credit_card"
    DEBIT_CARD = "debit_card"
    BANK_TRANSFER = "bank_transfer"
    DIGITAL_WALLET = "digital_wallet"
    CRYPTO = "crypto"


@dataclass
class Payment:
    """Core payment entity representing a payment transaction"""
    
    id: str
    order_id: str
    user_id: str
    amount: Decimal
    currency: str
    payment_method: PaymentMethod
    status: PaymentStatus
    provider_transaction_id: Optional[str]
    created_at: datetime
    updated_at: datetime
    completed_at: Optional[datetime] = None
    failure_reason: Optional[str] = None
    refund_amount: Optional[Decimal] = None
    
    @classmethod
    def create(cls, order_id: str, user_id: str, amount: Decimal, 
               currency: str, payment_method: PaymentMethod) -> 'Payment':
        """Factory method to create a new payment"""
        now = datetime.utcnow()
        return cls(
            id=str(uuid.uuid4()),
            order_id=order_id,
            user_id=user_id,
            amount=amount,
            currency=currency,
            payment_method=payment_method,
            status=PaymentStatus.PENDING,
            provider_transaction_id=None,
            created_at=now,
            updated_at=now
        )
    
    def mark_as_processing(self, provider_transaction_id: str) -> None:
        """Mark payment as processing with provider transaction ID"""
        self.status = PaymentStatus.PROCESSING
        self.provider_transaction_id = provider_transaction_id
        self.updated_at = datetime.utcnow()
    
    def mark_as_completed(self) -> None:
        """Mark payment as completed"""
        self.status = PaymentStatus.COMPLETED
        self.completed_at = datetime.utcnow()
        self.updated_at = datetime.utcnow()
    
    def mark_as_failed(self, reason: str) -> None:
        """Mark payment as failed with reason"""
        self.status = PaymentStatus.FAILED
        self.failure_reason = reason
        self.updated_at = datetime.utcnow()
    
    def refund(self, amount: Optional[Decimal] = None) -> None:
        """Process refund for the payment"""
        if self.status != PaymentStatus.COMPLETED:
            raise ValueError("Can only refund completed payments")
        
        refund_amount = amount or self.amount
        if refund_amount > self.amount:
            raise ValueError("Refund amount cannot exceed payment amount")
        
        self.status = PaymentStatus.REFUNDED
        self.refund_amount = refund_amount
        self.updated_at = datetime.utcnow()
    
    def is_successful(self) -> bool:
        """Check if payment was successful"""
        return self.status == PaymentStatus.COMPLETED
    
    def can_be_refunded(self) -> bool:
        """Check if payment can be refunded"""
        return self.status == PaymentStatus.COMPLETED


@dataclass
class PaymentProvider:
    """Represents a payment provider configuration"""
    
    id: str
    name: str
    supported_methods: list[PaymentMethod]
    supported_currencies: list[str]
    is_active: bool
    api_endpoint: str
    api_key: str
    fee_percentage: Decimal
    
    def supports_method(self, method: PaymentMethod) -> bool:
        """Check if provider supports payment method"""
        return method in self.supported_methods
    
    def supports_currency(self, currency: str) -> bool:
        """Check if provider supports currency"""
        return currency in self.supported_currencies 