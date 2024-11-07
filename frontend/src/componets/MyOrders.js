import React, { useEffect, useState } from 'react';
import "../assets/styles.css";
import { Link } from "react-router-dom";
import { decodeToken } from "../utils/auth";

const MyOrders = () => {
    const [orders, setOrders] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const token = localStorage.getItem("token");
    let userID=null;
    if(token){
        const decodedToken=decodeToken(token);
        userID=decodedToken.userID;
    }

    useEffect(() => {
        if(!userID){
            setError("user is not authonticated!");
            setLoading(false);
            return
        }
        const fetchOrders = async () => {
            try {
                const response = await fetch(`http://localhost:8080/api/users/orders?userId=${userID}`, {
                    method:'GET',
                    headers:{
                        'Authorization':`Bearer ${token}`
                    }
                });
                if (!response.ok) {
                    throw new Error('Failed to fetch orders');
                }
                const data = await response.json();
                console.log(data)
                setOrders(data);
            } catch (err) {
                setError(err.message);
            } finally {
                setLoading(false);
            }
        };

        fetchOrders();
    }, [userID, token]);

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
                    <li key={order.ID} className="order-card">
                        <h2>Order ID: {order.ID}</h2>
                        <p><strong>Status:</strong> {order.status}</p>
                        <p><strong>Pick-up Location:</strong> {order.pickup_location.street_address}</p>
                        <p><strong>Drop-off Location:</strong> {order.drop_off_location.street_address}</p>
                        <p><strong>Delivery Time:</strong> {order.estimated_delivery_time}</p>
                        <Link to={`/order-details/${order.ID}`} className="details-link">
                            View Details
                        </Link>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default MyOrders;
