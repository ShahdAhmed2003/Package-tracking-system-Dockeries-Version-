import React, { useEffect, useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import "../assets/styles.css";
import { Link } from "react-router-dom";

const OrderDetails = () => {
    const { orderId } = useParams(); //3yza ageeb el order id mn el url
    const [order, setOrder] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const navigate = useNavigate();

    useEffect(() => {
        const fetchOrderDetails = async () => {
            try {
                const response = await fetch(`http://localhost:8080/order/${orderId}`);
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
        return <div>Loading...</div>;
    }

    if (error) {
        return <div>Error: {error}</div>;
    }

    return (
        <div className="order-details-container">
            <h1>Order Details</h1>
            <h2>Order ID: {order.id}</h2>
            <p>Status: {order.status}</p>
            <p>Pick-up Location: {order.pickupLocation}</p>
            <p>Drop-off Location: {order.dropoffLocation}</p>
            <p>Delivery Time: {order.deliveryTime}</p>
            <p>Package Details: {order.packageDetails}</p>
            {order.courier && (
                <div>
                    <h3>Courier Information</h3>
                    <p>Name: {order.courier.name}</p>
                    <p>Contact: {order.courier.contact}</p>
                </div>
            )}
            {order.status === 'Pending' && (
                <button onClick={handleCancelOrder}>Cancel Order</button>
            )}
        </div>
    );
};

export default OrderDetails;
