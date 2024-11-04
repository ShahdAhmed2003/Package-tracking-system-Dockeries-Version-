import React, { useEffect, useState } from 'react';
import "../assets/styles.css";
import { Link } from "react-router-dom";

const MyOrders = () => {
    const [orders, setOrders] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        const fetchOrders = async () => {
            try {
                const response = await fetch('http://localhost:8080/myOrders');
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
        return <div>Loading...</div>;
    }

    if (error) {
        return <div>Error: {error}</div>;
    }

    return (
        <div>
            <h1>My Orders</h1>
            <ul>
                {orders.map(order => (
                    <li key={order.id}>
                        <h2>Order ID: {order.id}</h2>
                        <p>Status: {order.status}</p>
                        <p>Pick-up Location: {order.pickupLocation}</p>
                        <p>Drop-off Location: {order.dropoffLocation}</p>
                        <p>Delivery Time: {order.deliveryTime}</p>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default MyOrders;
