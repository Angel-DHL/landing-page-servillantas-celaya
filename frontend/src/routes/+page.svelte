<!-- frontend/src/routes/+page.svelte -->
<script>
  import { onMount } from 'svelte';
  import { inview } from '$lib/actions/inview.js';

  // ===== COUNTER ANIMATION =====
  let counters = [
    { target: 15, current: 0, suffix: '+', label: 'Años de experiencia' },
    { target: 10000, current: 0, suffix: '+', label: 'Clientes satisfechos' },
    { target: 50000, current: 0, suffix: '+', label: 'Servicios realizados' },
    { target: 100, current: 0, suffix: '%', label: 'Garantía' }
  ];

  let countersStarted = false;

  function startCounters() {
    if (countersStarted) return;
    countersStarted = true;

    counters.forEach((counter, i) => {
      const duration = 2000;
      const steps = 60;
      const increment = counter.target / steps;
      let current = 0;
      const interval = setInterval(() => {
        current += increment;
        if (current >= counter.target) {
          current = counter.target;
          clearInterval(interval);
        }
        counters[i].current = Math.floor(current);
        counters = counters; // Trigger Svelte reactivity
      }, duration / steps);
    });
  }

  // ===== SERVICIOS =====
  const servicios = [
    {
      icon: '🛞',
      titulo: 'Venta de Llantas',
      desc: 'Las mejores marcas nacionales e internacionales para todo tipo de vehículo.'
    },
    {
      icon: '🎯',
      titulo: 'Alineación',
      desc: 'Alineación computarizada de última generación. Precisión garantizada.'
    },
    {
      icon: '⚖️',
      titulo: 'Balanceo',
      desc: 'Eliminamos vibraciones para un manejo suave, cómodo y seguro.'
    },
    {
      icon: '🛑',
      titulo: 'Frenos',
      desc: 'Revisión y cambio de balatas, discos y sistema completo de frenado.'
    },
    {
      icon: '🔧',
      titulo: 'Suspensión',
      desc: 'Diagnóstico y reparación completa del sistema de suspensión.'
    },
    {
      icon: '🛢️',
      titulo: 'Cambio de Aceite',
      desc: 'Aceites sintéticos y minerales premium para proteger tu motor.'
    }
  ];

  // ===== TESTIMONIOS =====
  const testimonios = [
    {
      nombre: 'Carlos Hernández',
      texto: 'Excelente servicio, muy profesionales. Me cambiaron las 4 llantas y la alineación quedó perfecta.',
      rating: 5
    },
    {
      nombre: 'María González',
      texto: 'Los mejores precios de Celaya y el servicio es rapidísimo. Siempre traigo mi carro aquí.',
      rating: 5
    },
    {
      nombre: 'Roberto Juárez',
      texto: 'Me diagnosticaron un problema en la suspensión que nadie más encontró. Muy recomendados.',
      rating: 5
    }
  ];

  // ===== MARCAS =====
  const marcas = [
    'Michelin', 'Bridgestone', 'Goodyear', 'Continental',
    'Pirelli', 'Hankook', 'General Tire', 'BFGoodrich'
  ];
</script>

<svelte:head>
  <title>ServiLlantas Celaya | Llantas, Alineación y Servicio Automotriz</title>
</svelte:head>

<!-- ============ HERO ============ -->
<section class="hero">
  <!-- Agrega tu imagen en static/images/hero-bg.jpg -->
  <div class="hero-overlay"></div>

  <div class="hero-content">
    <p class="hero-subtitle">SERVICIO AUTOMOTRIZ PROFESIONAL EN CELAYA</p>
    <h1 class="hero-title">
      Tu seguridad<br />
      <span class="text-primary">sobre ruedas</span>
    </h1>
    <p class="hero-desc">
      Más de 15 años brindando el mejor servicio en llantas, alineación,
      balanceo y mecánica automotriz. Calidad y confianza garantizada.
    </p>
    <div class="hero-buttons">
      <a href="/contacto" class="btn btn-primary btn-lg">
        🚗 Cotizar Ahora
      </a>
      <a href="tel:+524611203488" class="btn btn-outline btn-lg">
        📞 461 120 3488
      </a>
    </div>
  </div>

  <!-- Indicador de scroll -->
  <div class="scroll-indicator">
    <div class="scroll-line"></div>
  </div>
</section>

<!-- ============ STATS ============ -->
<section class="stats-section" use:inview on:inview={startCounters}>
  <div class="container">
    <div class="stats-grid">
      {#each counters as stat}
        <div class="stat-item">
          <span class="stat-number">
            {stat.current.toLocaleString()}{stat.suffix}
          </span>
          <span class="stat-label">{stat.label}</span>
        </div>
      {/each}
    </div>
  </div>
</section>

<!-- ============ SERVICIOS ============ -->
<section class="section section-dark">
  <div class="container text-center">
    <p class="section-subtitle" use:inview>Nuestros Servicios</p>
    <h2 class="section-title fade-up" use:inview>
      Todo lo que tu vehículo <span class="text-primary">necesita</span>
    </h2>
    <p class="section-desc fade-up" use:inview>
      Contamos con el equipo y la experiencia para mantener tu vehículo
      en las mejores condiciones.
    </p>

    <div class="services-grid">
      {#each servicios as servicio, i}
        <div class="service-card fade-up delay-{i + 1}" use:inview>
          <span class="service-icon">{servicio.icon}</span>
          <h3>{servicio.titulo}</h3>
          <p>{servicio.desc}</p>
          <a href="/servicios" class="service-link">
            Ver más →
          </a>
        </div>
      {/each}
    </div>
  </div>
</section>

<!-- ============ POR QUÉ ELEGIRNOS ============ -->
<section class="section section-darker">
  <div class="container">
    <div class="why-grid">
      <div class="why-content">
        <p class="section-subtitle" use:inview>¿Por qué elegirnos?</p>
        <h2 class="section-title fade-up" use:inview>
          La confianza de <span class="text-primary">miles de clientes</span>
        </h2>

        <div class="why-list">
          <div class="why-item fade-left delay-1" use:inview>
            <div class="why-icon">✅</div>
            <div>
              <h4>Equipo de última generación</h4>
              <p>Tecnología computarizada para resultados precisos en cada servicio.</p>
            </div>
          </div>

          <div class="why-item fade-left delay-2" use:inview>
            <div class="why-icon">⚡</div>
            <div>
              <h4>Servicio rápido y eficiente</h4>
              <p>Respetamos tu tiempo. Servicio ágil sin sacrificar calidad.</p>
            </div>
          </div>

          <div class="why-item fade-left delay-3" use:inview>
            <div class="why-icon">🛡️</div>
            <div>
              <h4>Garantía en todos los trabajos</h4>
              <p>Respaldamos cada servicio con garantía escrita.</p>
            </div>
          </div>

          <div class="why-item fade-left delay-4" use:inview>
            <div class="why-icon">💰</div>
            <div>
              <h4>Precios competitivos</h4>
              <p>La mejor relación calidad-precio del mercado en Celaya.</p>
            </div>
          </div>
        </div>
      </div>

      <div class="why-image fade-right" use:inview>
        <!-- Reemplaza con tu imagen -->
        <div class="why-image-placeholder">
          <span>🔧</span>
          <p>Imagen del taller</p>
          <small>Coloca tu foto en static/images/taller.jpg</small>
        </div>
      </div>
    </div>
  </div>
</section>

<!-- ============ CTA BANNER ============ -->
<section class="cta-banner">
  <div class="cta-overlay"></div>
  <div class="container text-center" style="position:relative; z-index:2;">
    <h2 class="fade-up" use:inview>
      ¿Necesitas un servicio?<br />
      <span class="text-secondary">¡Contáctanos ahora!</span>
    </h2>
    <p class="cta-desc fade-up delay-1" use:inview>
      Llámanos o envíanos un WhatsApp para agendar tu cita
    </p>
    <div class="cta-buttons fade-up delay-2" use:inview>
      <a href="tel:+524611203488" class="btn btn-primary btn-lg">
        📞 461 120 3488
      </a>
      <a href="tel:+527201987926" class="btn btn-secondary btn-lg">
        📱 720 198 7926
      </a>
      <a href="/contacto" class="btn btn-outline btn-lg">
        ✉️ Formulario de Contacto
      </a>
    </div>
  </div>
</section>

<!-- ============ TESTIMONIOS ============ -->
<section class="section section-dark">
  <div class="container text-center">
    <p class="section-subtitle" use:inview>Testimonios</p>
    <h2 class="section-title fade-up" use:inview>
      Lo que dicen nuestros <span class="text-primary">clientes</span>
    </h2>

    <div class="testimonials-grid">
      {#each testimonios as testimonio, i}
        <div class="testimonial-card fade-up delay-{i + 1}" use:inview>
          <div class="testimonial-stars">
            {#each Array(testimonio.rating) as _}
              <span>⭐</span>
            {/each}
          </div>
          <p class="testimonial-text">"{testimonio.texto}"</p>
          <div class="testimonial-author">
            <div class="author-avatar">
              {testimonio.nombre.charAt(0)}
            </div>
            <span class="author-name">{testimonio.nombre}</span>
          </div>
        </div>
      {/each}
    </div>
  </div>
</section>

<!-- ============ MARCAS ============ -->
<section class="section section-darker">
  <div class="container text-center">
    <p class="section-subtitle" use:inview>Marcas que manejamos</p>
    <h2 class="section-title fade-up" use:inview>
      Las <span class="text-primary">mejores marcas</span> del mercado
    </h2>
    <div class="brands-grid fade-up delay-1" use:inview>
      {#each marcas as marca}
        <div class="brand-item">
          <span>{marca}</span>
        </div>
      {/each}
    </div>
  </div>
</section>

<!-- ============ CONTACT PREVIEW ============ -->
<section class="section section-dark">
  <div class="container text-center">
    <p class="section-subtitle" use:inview>Contáctanos</p>
    <h2 class="section-title fade-up" use:inview>
      ¿Listo para darle el mejor <span class="text-primary">servicio</span> a tu auto?
    </h2>
    <p class="section-desc fade-up" use:inview>
      Déjanos tus datos y nos pondremos en contacto contigo en menos de 24 horas.
    </p>
    <div class="fade-up delay-2" use:inview>
      <a href="/contacto" class="btn btn-primary btn-lg">
        Ir al formulario de contacto →
      </a>
    </div>
  </div>
</section>

<!-- ============ ESTILOS ============ -->
<style>
  /* ====== HERO ====== */
  .hero {
    position: relative;
    min-height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    /* Si tienes imagen de fondo: */
    background: url('/images/hero-bg.jpg') center/cover no-repeat;
    /* Si no tienes imagen aún, usa este gradiente: */
    background-color: #0a0a0a;
  }

  .hero-overlay {
    position: absolute;
    inset: 0;
    background: linear-gradient(
      135deg,
      rgba(0, 0, 0, 0.85) 0%,
      rgba(0, 0, 0, 0.6) 50%,
      rgba(230, 57, 70, 0.2) 100%
    );
    z-index: 1;
  }

  .hero-content {
    position: relative;
    z-index: 2;
    text-align: center;
    padding: 2rem;
    max-width: 800px;
    animation: fadeInUp 1s ease;
  }

  .hero-subtitle {
    color: var(--color-primary);
    font-family: var(--font-heading);
    font-weight: 700;
    font-size: 0.9rem;
    letter-spacing: 4px;
    margin-bottom: 1rem;
    animation: fadeInUp 1s ease 0.2s both;
  }

  .hero-title {
    font-size: clamp(2.5rem, 6vw, 4.5rem);
    margin-bottom: 1.5rem;
    animation: fadeInUp 1s ease 0.4s both;
  }

  .hero-desc {
    color: var(--color-light-3);
    font-size: 1.15rem;
    max-width: 600px;
    margin: 0 auto 2.5rem;
    line-height: 1.8;
    animation: fadeInUp 1s ease 0.6s both;
  }

  .hero-buttons {
    display: flex;
    gap: 1rem;
    justify-content: center;
    flex-wrap: wrap;
    animation: fadeInUp 1s ease 0.8s both;
  }

  .btn-lg {
    padding: 1rem 2.5rem;
    font-size: 1.05rem;
  }

  /* Scroll indicator */
  .scroll-indicator {
    position: absolute;
    bottom: 2rem;
    left: 50%;
    transform: translateX(-50%);
    z-index: 2;
  }

  .scroll-line {
    width: 2px;
    height: 40px;
    background: var(--color-primary);
    animation: scrollPulse 2s infinite;
    border-radius: 2px;
  }

  @keyframes scrollPulse {
    0%, 100% { opacity: 0; transform: scaleY(0.5); transform-origin: top; }
    50% { opacity: 1; transform: scaleY(1); }
  }

  @keyframes fadeInUp {
    from { opacity: 0; transform: translateY(30px); }
    to { opacity: 1; transform: translateY(0); }
  }

  /* ====== STATS ====== */
  .stats-section {
    background: var(--color-primary);
    padding: 3rem 0;
  }

  .stats-grid {
    display: grid;
    grid-template-columns: repeat(4, 1fr);
    gap: 2rem;
    text-align: center;
  }

  .stat-number {
    display: block;
    font-family: var(--font-heading);
    font-weight: 900;
    font-size: 2.5rem;
    color: var(--color-light);
  }

  .stat-label {
    font-size: 0.9rem;
    color: rgba(255, 255, 255, 0.8);
    margin-top: 0.3rem;
  }

  /* ====== SERVICES GRID ====== */
  .services-grid {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 2rem;
    margin-top: 1rem;
  }

  .service-card {
    background: var(--color-dark-3);
    border: 1px solid var(--color-dark-4);
    border-radius: 16px;
    padding: 2.5rem 2rem;
    text-align: center;
    transition: all 0.4s ease;
    position: relative;
    overflow: hidden;
  }

  .service-card::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 4px;
    background: var(--color-primary);
    transform: scaleX(0);
    transition: transform 0.4s ease;
  }

  .service-card:hover {
    transform: translateY(-8px);
    border-color: var(--color-primary);
    box-shadow: 0 20px 40px rgba(230, 57, 70, 0.15);
  }

  .service-card:hover::before {
    transform: scaleX(1);
  }

  .service-icon {
    font-size: 3rem;
    display: block;
    margin-bottom: 1rem;
  }

  .service-card h3 {
    font-size: 1.2rem;
    margin-bottom: 0.8rem;
    font-weight: 700;
  }

  .service-card p {
    color: var(--color-gray);
    font-size: 0.9rem;
    line-height: 1.6;
    margin-bottom: 1.5rem;
  }

  .service-link {
    color: var(--color-primary);
    font-weight: 600;
    font-size: 0.9rem;
    transition: all 0.3s ease;
  }

  .service-link:hover {
    color: var(--color-primary-light);
    letter-spacing: 1px;
  }

  /* ====== WHY SECTION ====== */
  .why-grid {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 4rem;
    align-items: center;
  }

  .why-list {
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
    margin-top: 2rem;
  }

  .why-item {
    display: flex;
    gap: 1rem;
    align-items: flex-start;
  }

  .why-icon {
    font-size: 1.5rem;
    width: 50px;
    height: 50px;
    background: rgba(230, 57, 70, 0.1);
    border-radius: 12px;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-shrink: 0;
  }

  .why-item h4 {
    font-size: 1.05rem;
    margin-bottom: 0.3rem;
    font-weight: 700;
  }

  .why-item p {
    color: var(--color-gray);
    font-size: 0.9rem;
  }

  .why-image-placeholder {
    background: var(--color-dark-3);
    border: 2px dashed var(--color-dark-4);
    border-radius: 16px;
    height: 450px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: 1rem;
    color: var(--color-gray);
  }

  .why-image-placeholder span {
    font-size: 4rem;
  }

  .why-image-placeholder small {
    font-size: 0.8rem;
    color: var(--color-dark-4);
  }

  /* ====== CTA BANNER ====== */
  .cta-banner {
    position: relative;
    padding: 5rem 0;
    background: url('/images/hero-bg.jpg') center/cover no-repeat fixed;
    background-color: var(--color-dark);
  }

  .cta-overlay {
    position: absolute;
    inset: 0;
    background: linear-gradient(135deg, rgba(230, 57, 70, 0.9), rgba(0, 0, 0, 0.85));
  }

  .cta-banner h2 {
    font-size: clamp(1.8rem, 4vw, 2.8rem);
    margin-bottom: 1rem;
  }

  .cta-desc {
    color: rgba(255, 255, 255, 0.8);
    font-size: 1.1rem;
    margin-bottom: 2rem;
  }

  .cta-buttons {
    display: flex;
    gap: 1rem;
    justify-content: center;
    flex-wrap: wrap;
  }

  /* ====== TESTIMONIALS ====== */
  .testimonials-grid {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 2rem;
    margin-top: 1rem;
  }

  .testimonial-card {
    background: var(--color-dark-3);
    border: 1px solid var(--color-dark-4);
    border-radius: 16px;
    padding: 2rem;
    text-align: left;
    transition: all 0.3s ease;
  }

  .testimonial-card:hover {
    border-color: var(--color-primary);
    transform: translateY(-4px);
  }

  .testimonial-stars {
    margin-bottom: 1rem;
    font-size: 1.1rem;
  }

  .testimonial-text {
    color: var(--color-light-3);
    font-size: 0.95rem;
    line-height: 1.7;
    font-style: italic;
    margin-bottom: 1.5rem;
  }

  .testimonial-author {
    display: flex;
    align-items: center;
    gap: 0.8rem;
  }

  .author-avatar {
    width: 45px;
    height: 45px;
    border-radius: 50%;
    background: var(--color-primary);
    display: flex;
    align-items: center;
    justify-content: center;
    font-family: var(--font-heading);
    font-weight: 800;
    font-size: 1.2rem;
  }

  .author-name {
    font-weight: 600;
    font-size: 0.95rem;
  }

  /* ====== BRANDS ====== */
  .brands-grid {
    display: grid;
    grid-template-columns: repeat(4, 1fr);
    gap: 1.5rem;
    margin-top: 1rem;
  }

  .brand-item {
    background: var(--color-dark-3);
    border: 1px solid var(--color-dark-4);
    border-radius: 12px;
    padding: 1.5rem;
    display: flex;
    align-items: center;
    justify-content: center;
    font-family: var(--font-heading);
    font-weight: 700;
    font-size: 1rem;
    color: var(--color-light-3);
    transition: all 0.3s ease;
  }

  .brand-item:hover {
    border-color: var(--color-primary);
    color: var(--color-primary);
    transform: translateY(-3px);
  }

  /* ====== RESPONSIVE ====== */
  @media (max-width: 1024px) {
    .services-grid,
    .testimonials-grid {
      grid-template-columns: repeat(2, 1fr);
    }
  }

  @media (max-width: 768px) {
    .stats-grid {
      grid-template-columns: repeat(2, 1fr);
    }

    .services-grid,
    .testimonials-grid {
      grid-template-columns: 1fr;
    }

    .why-grid {
      grid-template-columns: 1fr;
    }

    .brands-grid {
      grid-template-columns: repeat(2, 1fr);
    }

    .hero-title {
      font-size: clamp(2rem, 7vw, 3rem);
    }

    .btn-lg {
      padding: 0.9rem 1.8rem;
      font-size: 0.95rem;
    }
  }

  @media (max-width: 480px) {
    .stats-grid {
      grid-template-columns: 1fr 1fr;
      gap: 1.5rem;
    }

    .stat-number {
      font-size: 2rem;
    }
  }
</style>