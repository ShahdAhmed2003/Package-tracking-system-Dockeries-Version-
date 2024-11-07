// App.js
import React, { useState, useEffect } from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Header from './componets/Header';
import Home from './Pages/Home';
import LoginForm from './componets/LoginForm';
import RegistrationForm from './componets/RegisterationForm';
import OrderForm from './componets/OrderForm';
import MyOrders from './componets/MyOrders';
import OrderDetails from './componets/OrderDetails';

function App() {
    const [isLoggedIn, setIsLoggedIn] = useState(false);

    useEffect(() => {
        const token = localStorage.getItem("token");
        const user = localStorage.getItem("user");
        setIsLoggedIn(!!(token && user));
    }, []);

    const handleLoginStatusChange = () => {
        const token = localStorage.getItem("token");
        const user = localStorage.getItem("user");
        setIsLoggedIn(!!(token && user));
    };

    return (
        <BrowserRouter>
            <Header isLoggedIn={isLoggedIn} />
            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/login" element={<LoginForm onLogin={handleLoginStatusChange} />} />
                <Route path="/register" element={<RegistrationForm />} />
                <Route path="/order" element={<OrderForm />} />
                <Route path="/my-orders" element={<MyOrders />} />
                <Route path="/order-details/:orderId" element={<OrderDetails />} />
            </Routes>
        </BrowserRouter>
    );
}

export default App;
