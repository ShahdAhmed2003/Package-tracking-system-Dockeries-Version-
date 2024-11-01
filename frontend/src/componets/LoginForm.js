import React, { useState } from "react";
import "../assets/styles.css";
import { Link } from "react-router-dom";

const RegistrationForm = () => {
  const [formData, setFormData] = useState({
    email: "",
    password: "",
  });

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
        password: formData.password
    };
    try {
        const response = await fetch('http://localhost:8080/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(userPayload), 
                                           
                                             
        });

        if (response.ok) {
            // Only attempt to parse JSON if the response is successful
            alert("login successful!");
            // Log the response data
        } else {
            const errorData = await response.text(); // Read the response as text
            alert("Incorrect Email or Password. Please verify and try again");

        }
    } catch (error) {
        console.error("Error submitting the form:", error);
        alert("Something went wrong. Please try again.");
    }
    // Clear form after submission
    setFormData({ name: '', email: '', phone: '', password: '' });
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
          Don't have an account ? <Link to="/">Register</Link>
        </p>
      </form>
    </div>
  );
};

export default RegistrationForm;
