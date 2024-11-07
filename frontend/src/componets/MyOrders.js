// MyOrders.js
import React, { useEffect, useState } from 'react';
import "../assets/styles.css";
import { Link } from "react-router-dom";

const MyOrders = () => {
    const [orders, setOrders] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const userID = JSON.parse(localStorage.getItem("user")).user_id;

    useEffect(() => {
        const fetchOrders = async () => {
            try {
                const response = await fetch(`http://localhost:8080/api/users/orders?userId=${userID}`);
                if (!response.ok) {
                    throw new Error('Failed to fetch orders');
                }
                const data = await response.json();
                setOrders(data);
            } catch (err) {
                setError(err.message);
            } finally {
                setLoading(false);
            }
        };

        fetchOrders();
    }, []);

    if (loading) {
        return <div className="loading">Loading...</div>;
    }

    if (error) {
        return <div className="error">Error: {error}</div>;
    }

    return (
        <div className="orders-page">
            <h1 className="page-title">My Orders</h1>
            <ul className="orders-list">
                {orders.map(order => (
                    <li key={order.id} className="order-card">
                        <h2>Order ID: {order.id}</h2>
                        <p><strong>Status:</strong> {order.status}</p>
                        <p><strong>Pick-up Location:</strong> {order.pickupLocation}</p>
                        <p><strong>Drop-off Location:</strong> {order.dropoffLocation}</p>
                        <p><strong>Delivery Time:</strong> {order.deliveryTime}</p>
                        <Link to={`/order-details/${order.id}`} className="details-link">
                            View Details
                        </Link>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default MyOrders;
