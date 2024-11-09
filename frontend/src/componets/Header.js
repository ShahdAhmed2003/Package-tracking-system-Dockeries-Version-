// Header.js
import React from 'react';
import { Link, useNavigate } from 'react-router-dom';
import '../assets/styles.css';
import LogoutButton from './Log-out';

function Header({ isLoggedIn, onLogout }) {
    const navigate = useNavigate();

  

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
                        <LogoutButton onLogout={onLogout} />
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
