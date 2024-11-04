import React, { useState } from "react";
import "../assets/styles.css";
import { Link } from "react-router-dom";

const OrderForm = () => {
    const [formData, setFormData] = useState({
        pickupLocation: "",
        dropoffLocation: "",
        packageDetails: "",
        deliveryTime: "",
    });

    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData({ ...formData, [name]: value });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        const orderPayload = {
            pickupLocation: formData.pickupLocation,
            dropoffLocation: formData.dropoffLocation,
            packageDetails: formData.packageDetails,
            deliveryTime: formData.deliveryTime,
        };

        try {
            const response = await fetch('http://localhost:8080/createOrder', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(orderPayload),
            });

            if (response.ok) {
                alert("Order created successfully ");
            } else {
                const errorData = await response.text();
                alert("Failed to create order: " + errorData);
            }
        } catch (error) {
            console.error("Error submitting the form:", error);
            alert("Something went wrong! Please try again.");
        }

        setFormData({
            pickupLocation: "",
            dropoffLocation: "",
            packageDetails: "",
            deliveryTime: "",
        });
    };

    return (
        <div className="form-container">
            <form onSubmit={handleSubmit}>
                <h1>Create Order</h1>
                <div>
                    <label>Pick-up Location:</label>
                    <input
                        type="text"
                        name="pickupLocation"
                        value={formData.pickupLocation}
                        onChange={handleChange}
                        required
                    />
                </div>
                <div>
                    <label>Drop-off Location:</label>
                    <input
                        type="text"
                        name="dropoffLocation"
                        value={formData.dropoffLocation}
                        onChange={handleChange}
                        required
                    />
                </div>
                <div>
                    <label>Package Details:</label>
                    <textarea
                        name="packageDetails"
                        value={formData.packageDetails}
                        onChange={handleChange}
                        required
                    />
                </div>
                <div>
                    <label>Delivery Time:</label>
                    <input
                        type="datetime-local"
                        name="deliveryTime"
                        value={formData.deliveryTime}
                        onChange={handleChange}
                        required
                    />
                </div>
                <button type="submit">Create Order</button>
            </form>
        </div>
    );
};

export default OrderForm;
