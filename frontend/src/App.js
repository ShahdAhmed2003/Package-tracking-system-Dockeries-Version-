import React from 'react';
import RegistrationForm from './componets/RegisterationForm';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Login from './componets/LoginForm';
import OrderForm from './componets/OrderForm';

function App() {
    return (
        <BrowserRouter>
         <Routes>
            <Route path="/" element={<RegistrationForm />} />
            <Route path="/login" element={<Login />} />
             <Route path="/order" element={<OrderForm />} />
        </Routes>
        </BrowserRouter>
    );
}

export default App;
