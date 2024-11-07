import React, { useState } from "react";
import "../assets/styles.css";
import { Link, useNavigate } from "react-router-dom";

const RegistrationForm = () => {
    const [formData, setFormData] = useState({
        name: "",
        email: "",
        phone: "",
        password: "",
        role: "Customer",
    });

    const navigate = useNavigate();

    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData({ ...formData, [name]: value });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        const userPayload = {
            name: formData.name,
            email: formData.email,
            phonenumber: formData.phone,
            password: formData.password,
            role: formData.role,
        };

        try {
            const response = await fetch('http://localhost:8080/signup', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Accept': 'application/json',
                },
                body: JSON.stringify(userPayload),
            });

            if (response.ok) {
                setFormData({ name: '', email: '', phone: '', password: '', role: 'Customer' });
                alert("Signup successful!");
                navigate("/");
            } else {
                const errorData = await response.text();
                throw new Error(`Error: ${errorData}`);
            }
        } catch (error) {
            console.error("Error submitting the form:", error);
            if (error.message.includes("duplicate")) {
                alert("Email already exists");
            } else {
                alert(error.message);
            }
        }
    };

    return (
        <div className="form-container">
            <form onSubmit={handleSubmit}>
                <h1>Register a new user</h1>
                <div>
                    <label>Name:</label>
                    <input
                        type="text"
                        name="name"
                        value={formData.name}
                        onChange={handleChange}
                        required
                    />
                </div>
                <div>
                    <label>Email:</label>
                    <input
                        type="email"
                        name="email"
                        value={formData.email}
                        onChange={handleChange}
                        required
                    />
                </div>
                <div>
                    <label>Phone:</label>
                    <input
                        type="tel"
                        name="phone"
                        value={formData.phone}
                        onChange={handleChange}
                        required
                    />
                </div>
                <div>
                    <label>Password:</label>
                    <input
                        type="password"
                        name="password"
                        value={formData.password}
                        onChange={handleChange}
                        required
                    />
                </div>
                <div>
                    <label>Role:</label>
                    <select
                        name="role"
                        value={formData.role}
                        onChange={handleChange}
                        required
                    >
                        <option value="Admin">Admin</option>
                        <option value="Courier">Courier</option>
                        <option value="Customer">Customer</option>
                    </select>
                </div>
                <button type="submit">Register</button>
                <p>
                    Want to login instead? <Link to="/login">Login</Link>
                </p>
            </form>
        </div>
    );
};

export default RegistrationForm;