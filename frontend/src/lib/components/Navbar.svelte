<!-- frontend/src/lib/components/Navbar.svelte -->
<script>
  import { onMount } from 'svelte';
  import { page } from '$app/stores';

  let scrolled = false;
  let menuOpen = false;

  const navLinks = [
    { href: '/', label: 'Inicio' },
    { href: '/servicios', label: 'Servicios' },
    { href: '/sobre-nosotros', label: 'Nosotros' },
    { href: '/blog', label: 'Blog' },
    { href: '/contacto', label: 'Contacto' }
  ];

  onMount(() => {
    const handleScroll = () => {
      scrolled = window.scrollY > 50;
    };
    window.addEventListener('scroll', handleScroll);
    return () => window.removeEventListener('scroll', handleScroll);
  });

  function closeMenu() {
    menuOpen = false;
  }
</script>

<nav class="navbar" class:scrolled>
  <div class="nav-container">
    <!-- Logo -->
    <a href="/" class="nav-logo">
      <!-- Reemplaza con tu logo: <img src="/images/logo.png" alt="ServiLlantas" /> -->
      <span class="logo-icon">🛞</span>
      <span class="logo-text">Servi<span class="text-primary">Llantas</span></span>
    </a>

    <!-- Links desktop -->
    <ul class="nav-links" class:active={menuOpen}>
      {#each navLinks as link}
        <li>
          <a
            href={link.href}
            class="nav-link"
            class:active-link={$page.url.pathname === link.href}
            on:click={closeMenu}
          >
            {link.label}
          </a>
        </li>
      {/each}
      <!-- CTA en mobile -->
      <li class="nav-cta-mobile">
        <a href="tel:+524611203488" class="btn btn-primary">
          📞 461 120 3488
        </a>
      </li>
    </ul>

    <!-- CTA desktop -->
    <a href="tel:+524611203488" class="btn btn-primary nav-cta-desktop">
      📞 Llámanos
    </a>

    <!-- Hamburger -->
    <button
      class="hamburger"
      class:open={menuOpen}
      on:click={() => (menuOpen = !menuOpen)}
      aria-label="Menú"
    >
      <span></span>
      <span></span>
      <span></span>
    </button>
  </div>
</nav>

<!-- Overlay para cerrar menú en mobile -->
{#if menuOpen}
  <div class="nav-overlay" on:click={closeMenu} on:keydown={closeMenu}></div>
{/if}

<style>
  .navbar {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    z-index: 1000;
    padding: 1rem 0;
    transition: all 0.3s ease;
    background: transparent;
  }

  .navbar.scrolled {
    background: rgba(10, 10, 10, 0.95);
    backdrop-filter: blur(20px);
    padding: 0.6rem 0;
    box-shadow: 0 4px 30px rgba(0, 0, 0, 0.5);
  }

  .nav-container {
    max-width: var(--max-width);
    margin: 0 auto;
    padding: 0 1.5rem;
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  .nav-logo {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    font-family: var(--font-heading);
    font-weight: 900;
    font-size: 1.5rem;
    z-index: 1001;
  }

  .logo-icon {
    font-size: 2rem;
  }

  .nav-links {
    display: flex;
    align-items: center;
    gap: 0.5rem;
  }

  .nav-link {
    padding: 0.5rem 1rem;
    border-radius: 8px;
    font-weight: 500;
    font-size: 0.95rem;
    transition: all 0.3s ease;
    color: var(--color-light-3);
  }

  .nav-link:hover,
  .active-link {
    color: var(--color-primary);
    background: rgba(230, 57, 70, 0.1);
  }

  .nav-cta-desktop {
    font-size: 0.85rem;
    padding: 0.7rem 1.5rem;
  }

  .nav-cta-mobile {
    display: none;
  }

  /* ====== HAMBURGER ====== */
  .hamburger {
    display: none;
    flex-direction: column;
    gap: 5px;
    background: none;
    border: none;
    cursor: pointer;
    padding: 5px;
    z-index: 1001;
  }

  .hamburger span {
    width: 28px;
    height: 3px;
    background: var(--color-light);
    border-radius: 3px;
    transition: all 0.3s ease;
  }

  .hamburger.open span:nth-child(1) {
    transform: rotate(45deg) translate(5px, 6px);
  }

  .hamburger.open span:nth-child(2) {
    opacity: 0;
  }

  .hamburger.open span:nth-child(3) {
    transform: rotate(-45deg) translate(5px, -6px);
  }

  .nav-overlay {
    position: fixed;
    inset: 0;
    background: rgba(0, 0, 0, 0.6);
    z-index: 999;
  }

  /* ====== RESPONSIVE ====== */
  @media (max-width: 768px) {
    .hamburger {
      display: flex;
    }

    .nav-cta-desktop {
      display: none;
    }

    .nav-cta-mobile {
      display: block;
      margin-top: 1rem;
    }

    .nav-links {
      position: fixed;
      top: 0;
      right: -100%;
      width: 280px;
      height: 100vh;
      flex-direction: column;
      background: var(--color-dark-2);
      padding: 5rem 2rem 2rem;
      gap: 0.5rem;
      align-items: flex-start;
      transition: right 0.3s ease;
      z-index: 1000;
      box-shadow: -10px 0 30px rgba(0, 0, 0, 0.5);
    }

    .nav-links.active {
      right: 0;
    }

    .nav-link {
      font-size: 1.1rem;
      width: 100%;
      padding: 0.8rem 1rem;
    }
  }
</style>