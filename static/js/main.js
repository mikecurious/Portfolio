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