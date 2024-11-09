import React, { useState } from "react";
import "../assets/styles.css";
import { Link, useNavigate } from "react-router-dom";
import {jwtDecode} from "jwt-decode"; // Install this for decoding the token

const LoginForm = ({ onLogin }) => {
  const [formData, setFormData] = useState({
    email: "",
    password: "",
  });

  const navigate = useNavigate();

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const userPayload = {
      email: formData.email,
      password: formData.password,
    };

    try {
      const response = await fetch("http://localhost:8080/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(userPayload),
      });

      if (response.ok) {
        const data = await response.json();

        // Store the token and user info
        localStorage.setItem("token", data.token);
        localStorage.setItem("user", JSON.stringify(data.user));

        alert("Login successful!");
        setFormData({ email: "", password: "" });

        if (onLogin) {
          onLogin(); // Notify App of login status change
        }

        // Decode the token to get user role
        const decodedToken = jwtDecode(data.token);
        const userRole = decodedToken.role;

        // Navigate based on the user's role
        if (userRole === "Admin") {
          navigate("/admin/manageOrders");
        } else if (userRole === "Courier") {
          navigate("/courier/assignedOrders");
        } else {
          navigate("/order");
        }
      } else {
        const errorData = await response.text();
        alert("Incorrect Email or Password. Please verify and try again.");
      }
    } catch (error) {
      console.error("Error submitting the form:", error);
      alert("Something went wrong. Please try again.");
    }
  };

  return (
    <div className="form-container">
      <form onSubmit={handleSubmit}>
        <h1>Login</h1>
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
          <label>Password:</label>
          <input
            type="password"
            name="password"
            value={formData.password}
            onChange={handleChange}
            required
          />
        </div>
        <button type="submit">Login</button>
        <p>
          Don't have an account? <Link to="/register">Register</Link>
        </p>
      </form>
    </div>
  );
};

export default LoginForm;
