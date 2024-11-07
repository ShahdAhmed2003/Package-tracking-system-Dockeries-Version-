// Header.js
import React from 'react';
import { Link } from 'react-router-dom';
import '../assets/styles.css';

function Header({ isLoggedIn }) {
    return (
        <header className="header">
            <div className="logo">
                <Link to={"/"} className="login-button">Bosta</Link>
            </div>
            <nav>
                <Link to="/">Home</Link>
                {isLoggedIn ? (
                    <>
                        <Link to="/order">Create Order</Link>
                        <Link to="/my-orders">My Orders</Link>
                    </>
                ) : (
                    <>
                        <Link to="/register" className="login-button">Register</Link>
                        <Link to="/login" className="login-button">Login</Link>
                    </>
                )}
            </nav>
        </header>
    );
}

export default Header;
