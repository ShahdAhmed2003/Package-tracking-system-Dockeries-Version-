import React, { useEffect, useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import "../assets/styles.css";

const OrderDetails = () => {
    const { orderId } = useParams();
    const [order, setOrder] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const navigate = useNavigate();

    const user = JSON.parse(localStorage.getItem('user'));
    const userID = user.user_id.toString();

    useEffect(() => {
        const fetchOrderDetails = async () => {
            try {
                const response = await fetch(`http://localhost:8080/api/orders/details/${orderId}?userId=${userID}`);
                if (!response.ok) {
                    throw new Error('Failed to fetch order details');
                }
                const data = await response.json();
                setOrder(data);
            }
            catch (err) {
                setError(err.message);
            }
            finally {
                setLoading(false);
            }
        };

        fetchOrderDetails();
    }, [orderId]);

    const handleCancelOrder = async () => {
        try {
            const response = await fetch(`http://localhost:8080/order/${orderId}/cancel`, {
                method: 'POST',
            });

            if (response.ok) {
                alert("Order cancelled successfully");
                navigate('/my-orders');
            } else {
                const errorData = await response.text();
                alert("Failed to cancel order: " + errorData);
            }
        } catch (error) {
            console.error("Error cancelling the order:", error);
            alert("Something went wrong! Please try again.");
        }
    };

    if (loading) {
        return <div className="loading">Loading...</div>;
    }

    if (error) {
        return <div className="error">Error: {error}</div>;
    }

    return (
        <div className="order-details-container">
            <div className="order-details-header">
                <h1>Order Details</h1>
            </div>
            <div className="order-info">
                <div className="order-section">
                    <h3>Status:</h3>
                    <p>{order.status}</p>
                </div>
                <div className="order-section">
                    <h3>Pick-up Location:</h3>
                    <p>{order.pickup_location.street_address}, {order.pickup_location.city}</p>
                </div>
                <div className="order-section">
                    <h3>Drop-off Location:</h3>
                    <p>{order.drop_off_location.street_address}, {order.drop_off_location.city}</p>
                </div>
                <div className="order-section">
                    <h3>Package Contents:</h3>
                    <p>{order.package_details.contents}</p>
                </div>
                <div className="order-section">
                    <h3>Tracking Number:</h3>
                    <p>{order.tracking_number}</p>
                </div>
                <div className="order-section">
                    <h3>Estimated Delivery:</h3>
                    <p>{order.estimated_delivery}</p>
                </div>
                <div className="order-section">
                    <h3>Recipient Name:</h3>
                    <p>{order.recipient?.name}</p>
                </div>
                <div className="order-section">
                    <h3>Recipient Contact:</h3>
                    <p>{order.recipient?.contact}</p>
                </div>
                <div className="order-section">
                    <h3>Additional Instructions:</h3>
                    <p>{order.package_details.special_requirements}</p>
                </div>
            </div>
            <div className="order-actions">
                <button className="cancel-order-btn" onClick={handleCancelOrder}>Cancel Order</button>
            </div>
        </div>
    );
};

export default OrderDetails;
