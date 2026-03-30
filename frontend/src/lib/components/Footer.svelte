<!-- frontend/src/lib/components/Footer.svelte -->
<footer class="footer">
  <div class="footer-top">
    <div class="container">
      <div class="footer-grid">
        <!-- Columna 1: Logo e info -->
        <div class="footer-col">
          <a href="/" class="footer-logo">
            <span class="logo-icon">🛞</span>
            Servi<span class="text-primary">Llantas</span>
          </a>
          <p class="footer-desc">
            Expertos en llantas y servicio automotriz en Celaya.
            Tu seguridad y confianza sobre ruedas desde hace más de 15 años.
          </p>
          <div class="footer-social">
            <a href="#" class="social-link" aria-label="Facebook">f</a>
            <a href="#" class="social-link" aria-label="Instagram">in</a>
            <a href="#" class="social-link" aria-label="TikTok">tk</a>
          </div>
        </div>

        <!-- Columna 2: Servicios -->
        <div class="footer-col">
          <h4 class="footer-title">Servicios</h4>
          <ul>
            <li><a href="/servicios">Venta de Llantas</a></li>
            <li><a href="/servicios">Alineación</a></li>
            <li><a href="/servicios">Balanceo</a></li>
            <li><a href="/servicios">Frenos</a></li>
            <li><a href="/servicios">Suspensión</a></li>
            <li><a href="/servicios">Cambio de Aceite</a></li>
          </ul>
        </div>

        <!-- Columna 3: Links -->
        <div class="footer-col">
          <h4 class="footer-title">Empresa</h4>
          <ul>
            <li><a href="/sobre-nosotros">Sobre Nosotros</a></li>
            <li><a href="/blog">Blog</a></li>
            <li><a href="/contacto">Contacto</a></li>
          </ul>
        </div>

        <!-- Columna 4: Contacto -->
        <div class="footer-col">
          <h4 class="footer-title">Contacto</h4>
          <ul class="contact-list">
            <li>
              <span>📞</span>
              <a href="tel:+524611203488">461 120 3488</a>
            </li>
            <li>
              <span>📱</span>
              <a href="tel:+527201987926">720 198 7926</a>
            </li>
            <li>
              <span>📍</span>
              <span>Celaya, Guanajuato, México</span>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>

  <!-- Footer bottom -->
  <div class="footer-bottom">
    <div class="container">
      <p>© {new Date().getFullYear()} ServiLlantas Celaya. Todos los derechos reservados.</p>
    </div>
  </div>
</footer>

<style>
  .footer {
    background: var(--color-dark-2);
    border-top: 1px solid var(--color-dark-4);
  }

  .footer-top {
    padding: 4rem 0 3rem;
  }

  .footer-grid {
    display: grid;
    grid-template-columns: 1.5fr 1fr 1fr 1.2fr;
    gap: 3rem;
  }

  .footer-logo {
    display: inline-flex;
    align-items: center;
    gap: 0.5rem;
    font-family: var(--font-heading);
    font-weight: 900;
    font-size: 1.4rem;
    margin-bottom: 1rem;
  }

  .logo-icon {
    font-size: 1.8rem;
  }

  .footer-desc {
    color: var(--color-gray);
    font-size: 0.9rem;
    margin-bottom: 1.5rem;
    line-height: 1.7;
  }

  .footer-social {
    display: flex;
    gap: 0.8rem;
  }

  .social-link {
    width: 40px;
    height: 40px;
    border-radius: 50%;
    background: var(--color-dark-4);
    display: flex;
    align-items: center;
    justify-content: center;
    font-weight: 700;
    font-size: 0.8rem;
    transition: all 0.3s ease;
    text-transform: uppercase;
  }

  .social-link:hover {
    background: var(--color-primary);
    transform: translateY(-3px);
  }

  .footer-title {
    font-size: 1.1rem;
    margin-bottom: 1.5rem;
    font-weight: 700;
    position: relative;
  }

  .footer-title::after {
    content: '';
    position: absolute;
    bottom: -8px;
    left: 0;
    width: 30px;
    height: 3px;
    background: var(--color-primary);
    border-radius: 2px;
  }

  .footer-col ul {
    display: flex;
    flex-direction: column;
    gap: 0.7rem;
  }

  .footer-col ul a,
  .footer-col ul span {
    color: var(--color-gray);
    font-size: 0.9rem;
    transition: color 0.3s ease;
  }

  .footer-col ul a:hover {
    color: var(--color-primary);
  }

  .contact-list li {
    display: flex;
    align-items: center;
    gap: 0.7rem;
  }

  .footer-bottom {
    border-top: 1px solid var(--color-dark-4);
    padding: 1.5rem 0;
    text-align: center;
  }

  .footer-bottom p {
    color: var(--color-gray);
    font-size: 0.85rem;
  }

  @media (max-width: 768px) {
    .footer-grid {
      grid-template-columns: 1fr 1fr;
      gap: 2rem;
    }
  }

  @media (max-width: 480px) {
    .footer-grid {
      grid-template-columns: 1fr;
    }
  }
</style>