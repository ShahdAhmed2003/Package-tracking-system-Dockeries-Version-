// LogoutButton.js
import React from 'react';
import { useNavigate } from 'react-router-dom';
import "../assets/styless.css";

function LogoutButton({ onLogout }) {
    const navigate = useNavigate();

    const handleLogout = () => {
        // Clear token from localStorage
        localStorage.clear();
        
        // Notify parent component to update the state
        onLogout();

        // Redirect to login page
        navigate('/login');
    };

    return (
        <button onClick={handleLogout} className="logout-button">
            Logout
        </button>
    );
}

export default LogoutButton;
