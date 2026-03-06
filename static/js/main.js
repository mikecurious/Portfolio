document.addEventListener('DOMContentLoaded', function() {
    // Check for success message in URL
    const urlParams = new URLSearchParams(window.location.search);
    if (urlParams.get('message') === 'success') {
        document.getElementById('success-message').style.display = 'block';
        document.getElementById('contact-form').style.display = 'none';
        
        // Scroll to contact section
        document.getElementById('contact').scrollIntoView();
        
        // Clear success message from URL without refreshing the page
        window.history.replaceState({}, document.title, window.location.pathname);
    }
    
    // Smooth scrolling for navigation links
    document.querySelectorAll('nav a').forEach(anchor => {
        anchor.addEventListener('click', function(e) {
            e.preventDefault();
            
            const targetId = this.getAttribute('href');
            const targetElement = document.querySelector(targetId);
            
            window.scrollTo({
                top: targetElement.offsetTop - 80, // Adjust for fixed header
                behavior: 'smooth'
            });
        });
    });
    
    // Hero role typing animation
    const heroRole = document.getElementById('hero-role');
    if (heroRole) {
        const text = heroRole.textContent.trim();
        heroRole.textContent = '';
        let i = 0;
        const type = () => {
            if (i < text.length) {
                heroRole.textContent += text[i++];
                setTimeout(type, 55);
            }
        };
        setTimeout(type, 400);
    }

    // Scroll reveal via Intersection Observer
    const revealObserver = new IntersectionObserver((entries) => {
        entries.forEach(entry => {
            if (entry.isIntersecting) {
                entry.target.classList.add('visible');
                revealObserver.unobserve(entry.target);
            }
        });
    }, { threshold: 0.12 });

    document.querySelectorAll('.reveal').forEach(el => revealObserver.observe(el));

    // Header scroll effect + active nav section tracking
    const sections = document.querySelectorAll('section[id]');
    const navLinks = document.querySelectorAll('nav a');

    window.addEventListener('scroll', function() {
        const header = document.querySelector('header');
        if (window.scrollY > 100) {
            header.classList.add('scrolled');
        } else {
            header.classList.remove('scrolled');
        }

        let current = '';
        sections.forEach(section => {
            if (window.scrollY >= section.offsetTop - 120) {
                current = section.getAttribute('id');
            }
        });
        navLinks.forEach(link => {
            link.classList.remove('active');
            if (link.getAttribute('href') === '#' + current) {
                link.classList.add('active');
            }
        });
    });

    // Form validation
    const contactForm = document.getElementById('contact-form');
    contactForm.addEventListener('submit', function(e) {
        let valid = true;
        
        // Reset any previous error styling
        document.querySelectorAll('.form-group').forEach(group => {
            group.classList.remove('error');
        });
        
        // Simple validation
        const name = document.getElementById('name').value.trim();
        const email = document.getElementById('email').value.trim();
        const message = document.getElementById('message').value.trim();
        
        if (name === '') {
            document.getElementById('name').parentElement.classList.add('error');
            valid = false;
        }
        
        if (email === '') {
            document.getElementById('email').parentElement.classList.add('error');
            valid = false;
        } else if (!isValidEmail(email)) {
            document.getElementById('email').parentElement.classList.add('error');
            valid = false;
        }
        
        if (message === '') {
            document.getElementById('message').parentElement.classList.add('error');
            valid = false;
        }
        
        if (!valid) {
            e.preventDefault();
        }
    });
    
    // Email validation helper
    function isValidEmail(email) {
        const re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
        return re.test(email.toLowerCase());
    }
});