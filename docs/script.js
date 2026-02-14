// ==========================================
// Mobile Menu Toggle
// ==========================================
const mobileMenuToggle = document.querySelector('.mobile-menu-toggle');
const navLinks = document.querySelector('.nav-links');

mobileMenuToggle.addEventListener('click', () => {
    navLinks.classList.toggle('active');
    mobileMenuToggle.classList.toggle('active');
});

// Close mobile menu when clicking on a link
document.querySelectorAll('.nav-links a').forEach(link => {
    link.addEventListener('click', () => {
        navLinks.classList.remove('active');
        mobileMenuToggle.classList.remove('active');
    });
});

// Close mobile menu when clicking outside
document.addEventListener('click', (e) => {
    if (!navLinks.contains(e.target) && !mobileMenuToggle.contains(e.target)) {
        navLinks.classList.remove('active');
        mobileMenuToggle.classList.remove('active');
    }
});

// ==========================================
// Code Tabs Functionality
// ==========================================
function initializeTabs() {
    const tabContainers = document.querySelectorAll('.code-tabs');
    
    tabContainers.forEach(container => {
        const tabButtons = container.querySelectorAll('.tab-btn');
        const parentSection = container.closest('.install-method, .step-card');
        
        if (!parentSection) return;
        
        const tabContents = parentSection.querySelectorAll('.tab-content');
        
        tabButtons.forEach(button => {
            button.addEventListener('click', () => {
                const targetTab = button.getAttribute('data-tab');
                
                // Remove active class from all buttons in this group
                tabButtons.forEach(btn => btn.classList.remove('active'));
                
                // Add active class to clicked button
                button.classList.add('active');
                
                // Hide all tab contents in this group
                tabContents.forEach(content => {
                    content.classList.remove('active');
                });
                
                // Show target tab content
                const targetContent = parentSection.querySelector(`.tab-content[data-tab="${targetTab}"]`);
                if (targetContent) {
                    targetContent.classList.add('active');
                }
            });
        });
    });
}

// Initialize tabs on page load
initializeTabs();

// ==========================================
// Smooth Scroll for Navigation Links
// ==========================================
document.querySelectorAll('a[href^="#"]').forEach(anchor => {
    anchor.addEventListener('click', function (e) {
        const href = this.getAttribute('href');
        
        // Skip if href is just "#"
        if (href === '#' || href === '') return;
        
        e.preventDefault();
        
        const target = document.querySelector(href);
        if (target) {
            const offsetTop = target.offsetTop - 80; // Account for fixed navbar
            
            window.scrollTo({
                top: offsetTop,
                behavior: 'smooth'
            });
        }
    });
});

// ==========================================
// Navbar Scroll Effect
// ==========================================
const navbar = document.querySelector('.navbar');
let lastScroll = 0;

window.addEventListener('scroll', () => {
    const currentScroll = window.pageYOffset;
    
    if (currentScroll > 100) {
        navbar.style.boxShadow = '0 10px 30px rgba(0, 0, 0, 0.5)';
    } else {
        navbar.style.boxShadow = 'none';
    }
    
    lastScroll = currentScroll;
});

// ==========================================
// Intersection Observer for Fade-in Animations
// ==========================================
const observerOptions = {
    threshold: 0.1,
    rootMargin: '0px 0px -50px 0px'
};

const observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
        if (entry.isIntersecting) {
            entry.target.classList.add('fade-in-up');
            observer.unobserve(entry.target);
        }
    });
}, observerOptions);

// Observe elements for animation
const animateElements = document.querySelectorAll(
    '.feature-card, .usage-card, .example-card, .prereq-card, .format-card, .build-card, .trouble-card, .step-card'
);

animateElements.forEach(element => {
    observer.observe(element);
});

// ==========================================
// Copy Code Button Functionality
// ==========================================
function addCopyButtons() {
    const codeBlocks = document.querySelectorAll('pre');
    
    codeBlocks.forEach((block, index) => {
        // Create wrapper if it doesn't exist
        if (!block.parentElement.classList.contains('code-wrapper')) {
            const wrapper = document.createElement('div');
            wrapper.className = 'code-wrapper';
            block.parentNode.insertBefore(wrapper, block);
            wrapper.appendChild(block);
            
            // Create copy button
            const copyButton = document.createElement('button');
            copyButton.className = 'copy-btn';
            copyButton.innerHTML = '📋 Copy';
            copyButton.setAttribute('aria-label', 'Copy code to clipboard');
            
            copyButton.addEventListener('click', () => {
                const code = block.querySelector('code').textContent;
                
                navigator.clipboard.writeText(code).then(() => {
                    copyButton.innerHTML = '✓ Copied!';
                    copyButton.style.background = 'linear-gradient(135deg, #27c93f, #20a838)';
                    
                    setTimeout(() => {
                        copyButton.innerHTML = '📋 Copy';
                        copyButton.style.background = '';
                    }, 2000);
                }).catch(err => {
                    console.error('Failed to copy code:', err);
                    copyButton.innerHTML = '✗ Failed';
                    
                    setTimeout(() => {
                        copyButton.innerHTML = '📋 Copy';
                    }, 2000);
                });
            });
            
            wrapper.appendChild(copyButton);
        }
    });
}

// Add copy buttons on page load
addCopyButtons();

// ==========================================
// Terminal Animation
// ==========================================
function animateTerminal() {
    const terminal = document.querySelector('.terminal-demo');
    
    if (!terminal) return;
    
    const observer = new IntersectionObserver((entries) => {
        entries.forEach(entry => {
            if (entry.isIntersecting) {
                // Trigger animation
                const progressBars = terminal.querySelectorAll('.progress-bar');
                progressBars.forEach(bar => {
                    bar.style.animation = 'none';
                    setTimeout(() => {
                        bar.style.animation = '';
                    }, 10);
                });
                
                observer.unobserve(terminal);
            }
        });
    }, { threshold: 0.5 });
    
    observer.observe(terminal);
}

animateTerminal();

// ==========================================
// Add Active State to Navigation on Scroll
// ==========================================
function updateActiveNavLink() {
    const sections = document.querySelectorAll('section[id]');
    const navLinks = document.querySelectorAll('.nav-links a[href^="#"]');
    
    let currentSection = '';
    
    sections.forEach(section => {
        const sectionTop = section.offsetTop;
        const sectionHeight = section.clientHeight;
        
        if (window.pageYOffset >= (sectionTop - 150)) {
            currentSection = section.getAttribute('id');
        }
    });
    
    navLinks.forEach(link => {
        link.classList.remove('active-link');
        
        if (link.getAttribute('href') === `#${currentSection}`) {
            link.classList.add('active-link');
        }
    });
}

window.addEventListener('scroll', updateActiveNavLink);

// ==========================================
// Add Styles for Copy Button and Active Link
// ==========================================
const style = document.createElement('style');
style.textContent = `
    .code-wrapper {
        position: relative;
    }
    
    .copy-btn {
        position: absolute;
        top: 1rem;
        right: 1rem;
        background: linear-gradient(135deg, var(--accent-cyan), var(--accent-blue));
        color: var(--bg-primary);
        border: none;
        padding: 0.5rem 1rem;
        border-radius: 0.5rem;
        cursor: pointer;
        font-size: 0.85rem;
        font-weight: 600;
        transition: all 0.3s ease;
        z-index: 10;
    }
    
    .copy-btn:hover {
        transform: translateY(-2px);
        box-shadow: 0 5px 15px rgba(0, 212, 255, 0.3);
    }
    
    .copy-btn:active {
        transform: translateY(0);
    }
    
    .active-link {
        color: var(--accent-cyan) !important;
        position: relative;
    }
    
    .active-link::after {
        content: '';
        position: absolute;
        bottom: -5px;
        left: 0;
        width: 100%;
        height: 2px;
        background: var(--accent-cyan);
    }
    
    @media (max-width: 768px) {
        .copy-btn {
            top: 0.5rem;
            right: 0.5rem;
            padding: 0.4rem 0.8rem;
            font-size: 0.75rem;
        }
        
        .active-link::after {
            display: none;
        }
    }
`;

document.head.appendChild(style);

// ==========================================
// Back to Top Button
// ==========================================
function createBackToTopButton() {
    const button = document.createElement('button');
    button.className = 'back-to-top';
    button.innerHTML = '↑';
    button.setAttribute('aria-label', 'Back to top');
    document.body.appendChild(button);
    
    // Show/hide button based on scroll position
    window.addEventListener('scroll', () => {
        if (window.pageYOffset > 500) {
            button.classList.add('visible');
        } else {
            button.classList.remove('visible');
        }
    });
    
    // Scroll to top on click
    button.addEventListener('click', () => {
        window.scrollTo({
            top: 0,
            behavior: 'smooth'
        });
    });
    
    // Add styles for back to top button
    const backToTopStyle = document.createElement('style');
    backToTopStyle.textContent = `
        .back-to-top {
            position: fixed;
            bottom: 2rem;
            right: 2rem;
            width: 50px;
            height: 50px;
            background: linear-gradient(135deg, var(--accent-cyan), var(--accent-blue));
            color: var(--bg-primary);
            border: none;
            border-radius: 50%;
            font-size: 1.5rem;
            cursor: pointer;
            opacity: 0;
            visibility: hidden;
            transition: all 0.3s ease;
            z-index: 1000;
            box-shadow: 0 5px 15px rgba(0, 0, 0, 0.3);
        }
        
        .back-to-top.visible {
            opacity: 1;
            visibility: visible;
        }
        
        .back-to-top:hover {
            transform: translateY(-5px);
            box-shadow: 0 10px 25px rgba(0, 212, 255, 0.4);
        }
        
        .back-to-top:active {
            transform: translateY(-2px);
        }
        
        @media (max-width: 768px) {
            .back-to-top {
                bottom: 1rem;
                right: 1rem;
                width: 45px;
                height: 45px;
                font-size: 1.3rem;
            }
        }
    `;
    
    document.head.appendChild(backToTopStyle);
}

createBackToTopButton();

// ==========================================
// Table Responsiveness Enhancement
// ==========================================
function enhanceTableResponsiveness() {
    const tables = document.querySelectorAll('table');
    
    tables.forEach(table => {
        const wrapper = table.closest('.table-wrapper');
        if (wrapper && window.innerWidth < 768) {
            // Add scroll hint
            const hint = document.createElement('div');
            hint.className = 'scroll-hint';
            hint.textContent = '← Scroll to see more →';
            hint.style.cssText = `
                text-align: center;
                color: var(--text-muted);
                font-size: 0.85rem;
                padding: 0.5rem;
                background: rgba(59, 130, 246, 0.1);
                border-radius: 0.5rem;
                margin-bottom: 0.5rem;
            `;
            
            if (!wrapper.querySelector('.scroll-hint')) {
                wrapper.insertBefore(hint, table.parentElement);
            }
        }
    });
}

enhanceTableResponsiveness();
window.addEventListener('resize', enhanceTableResponsiveness);

// ==========================================
// Console Easter Egg
// ==========================================
console.log('%c🖼️ img-gen', 'font-size: 24px; font-weight: bold; background: linear-gradient(135deg, #00d4ff, #a855f7); -webkit-background-clip: text; -webkit-text-fill-color: transparent;');
console.log('%cThe Ultimate CLI Image Generation Tool', 'font-size: 14px; color: #8b9cc9;');
console.log('%cGitHub: https://github.com/Parthipan-Natkunam/generate_image', 'font-size: 12px; color: #00d4ff;');

// ==========================================
// Performance Optimization
// ==========================================
// Lazy load images if any are added
if ('loading' in HTMLImageElement.prototype) {
    const images = document.querySelectorAll('img[loading="lazy"]');
    images.forEach(img => {
        img.src = img.dataset.src;
    });
} else {
    // Fallback for browsers that don't support lazy loading
    const script = document.createElement('script');
    script.src = 'https://cdnjs.cloudflare.com/ajax/libs/lazysizes/5.3.2/lazysizes.min.js';
    document.body.appendChild(script);
}

// ==========================================
// Initialize Everything on DOM Load
// ==========================================
document.addEventListener('DOMContentLoaded', () => {
    console.log('Website loaded successfully!');
    updateActiveNavLink();
});