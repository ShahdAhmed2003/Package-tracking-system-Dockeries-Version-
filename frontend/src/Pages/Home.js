import React from 'react';


function Home() {

    const features = [
        { icon: '/images/icon1.png', title: 'Fast Delivery', description: 'Get your package in record time.' },
        { icon: '/images/icon2.png', title: 'Tracking', description: 'Track your package at every step.' },
        { icon: '/images/icon3.png', title: 'Affordable', description: 'Cost-effective solutions for your business.' },
    ];
    return (<>
    <section className="hero">
            <div className="hero-content">
                <h1>Delivering Your Package, The Smart Way</h1>
                <p>Fast, reliable delivery service for your business needs.</p>
                <button className="cta-button">Get Started</button>
            </div>
            <div className="hero-image">
                {/* <img src="/images/hero-image.png" alt="Delivery" /> */}
            </div>
        </section>

        <section className="features">
            <h2>Our Features</h2>
            <div className="feature-list">
                {features.map((feature, index) => (
                    <div className="feature-item" key={index}>
                        {/* <img src={feature.icon} alt={feature.title} /> */}
                        <h3>{feature.title}</h3>
                        <p>{feature.description}</p>
                    </div>
                ))}
            </div>
        </section>
        </>
    );
}

export default Home;
