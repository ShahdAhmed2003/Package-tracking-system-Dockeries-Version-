import React, { useEffect, useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import "../assets/styles.css";
import { decodeToken } from "../utils/auth";

const OrderDetails = () => {
    const { orderId } = useParams();
    const [order, setOrder] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const navigate = useNavigate();

    const token = localStorage.getItem("token");
    let userID = null;
    if (token) {
        const decodedToken = decodeToken(token);
        userID = decodedToken.userID;
    }

    useEffect(() => {
        const fetchOrderDetails = async () => {
            try {
                const response = await fetch(`https://backend-shahdahmed851-dev.apps.rm2.thpm.p1.openshiftapps.com/api/orders/details/${orderId}?userId=${userID}`, {
                    method: 'GET',
                    headers: {
                        'Authorization': `Bearer ${token}`
                    }
                });

                if (!response.ok) {
                    throw new Error('Failed to fetch order details');
                }
                const data = await response.json();
                setOrder(data);
            } catch (err) {
                setError(err.message);
            } finally {
                setLoading(false);
            }
        };

        fetchOrderDetails();
    }, [orderId, userID, token]);

    const handleCancelOrder = async () => {
        try {
            const response = await fetch(`https://backend-shahdahmed851-dev.apps.rm2.thpm.p1.openshiftapps.com/api/orders/cancel/${orderId}`, {
                method: 'DELETE',
                headers: {
                    'Authorization': `Bearer ${token}`
                }
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

    const handleVerifyOrder = async () => {
        try {
            const response = await fetch(`https://backend-shahdahmed851-dev.apps.rm2.thpm.p1.openshiftapps.com/api/orders/verify?orderId=${orderId}`, {
                method: 'POST',
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            });

            if (response.ok) {
                alert("Order verified successfully");
                setOrder({ ...order, status: "Verified" });
            } else {
                const errorData = await response.text();
                alert("Failed to verify order: " + errorData);
            }
        } catch (error) {
            console.error("Error verifying the order:", error);
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
                    <p>{order.estimated_delivery_time}</p>
                </div>
                <div className="order-section">
                    <h3>Additional Instructions:</h3>
                    <p>{order.package_details.special_requirements ? order.package_details.special_requirements : "None"}</p>
                </div>
            </div>
            <div className="order-actions">
                {(order.status === "pending" || order.status === "declined") ? (
                    <button className="cancel-order-btn" onClick={handleCancelOrder}>Cancel Order</button>
                ) : (
                    <span className="cancel-disabled-message">Order cannot be canceled because it is already accepted</span>
                )}
            </div>
        </div>
    );
};

export default OrderDetails;
