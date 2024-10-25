import React from 'react';
import RegistrationForm from './componets/RegisterationForm';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Login from './componets/LoginForm';

function App() {
    return (
        <BrowserRouter>
         <Routes>
            <Route path="/" element={<RegistrationForm />} />
            <Route path="/login" element={<Login />} />
        </Routes>
        </BrowserRouter>
    );
}

export default App;
