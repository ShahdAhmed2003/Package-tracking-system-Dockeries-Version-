import React from 'react';
import '../assets/styles.css';

import gpsTrackingImg from '../assets/images/gps-tracking-for-telecom-company.jpg';

function Home() {
    const features = [
        { icon: gpsTrackingImg, title: 'Fast Delivery', description: 'Get your package in record time.' },
        { icon: gpsTrackingImg, title: 'Real-Time Tracking', description: 'Track your package every step of the way.' },
        { icon: gpsTrackingImg, title: 'Affordable Pricing', description: 'Cost-effective solutions for your business.' },
    ];

    return (
        <>
            <section className="hero">
                <div className="hero-content">
                    <h1>Delivering Your Packages, The Smart Way</h1>
                    <p>
                        Experience fast, reliable, and affordable delivery services tailored to your business needs.
                    </p>
                    <button className="cta-button">Get Started</button>
                </div>
                <div className="hero-image">
                    <img src={gpsTrackingImg} alt="Delivery Service" />
                </div>
            </section>

            <section className="features">
                <h2 className="section-title">Our Features</h2>
                <div className="feature-list">
                    {features.map((feature, index) => (
                        <div className="feature-item" key={index}>
                            <img src={feature.icon} alt={feature.title} className="feature-icon" />
                            <h3>{feature.title}</h3>
                            <p>{feature.description}</p>
                        </div>
                    ))}
                </div>
            </section>

            <section className="cta">
                <h2>Ready to simplify your logistics?</h2>
                <p>Join thousands of businesses who trust us for their delivery needs.</p>
                <button className="cta-button">Contact Us</button>
            </section>

            <footer className="footer">
                <div className="footer-content">
                    <p>&copy; 2025 Bosta. All rights reserved.</p>
                    <ul className="footer-links">
                        <li><a href="#">Privacy Policy</a></li>
                        <li><a href="#">Terms of Service</a></li>
                        <li><a href="#">Contact</a></li>
                    </ul>
                </div>
            </footer>
        </>
    );
}

export default Home;
