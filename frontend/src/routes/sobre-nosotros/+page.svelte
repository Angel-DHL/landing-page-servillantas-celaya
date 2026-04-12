<!-- frontend/src/routes/sobre-nosotros/+page.svelte -->
<script>
    import { inview } from '$lib/actions/inview.js';
  
    // ===== COUNTER ANIMADO =====
    let valores = [
      { target: 15, current: 0, suffix: '+', label: 'Años de experiencia' },
      { target: 10000, current: 0, suffix: '+', label: 'Clientes satisfechos' },
      { target: 6, current: 0, suffix: '', label: 'Servicios especializados' },
      { target: 50000, current: 0, suffix: '+', label: 'Servicios realizados' }
    ];
  
    let countersStarted = false;
  
    function startCounters() {
      if (countersStarted) return;
      countersStarted = true;
      valores.forEach((v, i) => {
        const duration = 2000;
        const steps = 60;
        const increment = v.target / steps;
        let current = 0;
        const interval = setInterval(() => {
          current += increment;
          if (current >= v.target) {
            current = v.target;
            clearInterval(interval);
          }
          valores[i].current = Math.floor(current);
          valores = valores;
        }, duration / steps);
      });
    }
  
    // ===== MISIÓN / VISIÓN / VALORES =====
    const pilares = [
      {
        icon: '🎯',
        titulo: 'Misión',
        texto: 'Brindar servicios automotrices de la más alta calidad, ofreciendo a nuestros clientes soluciones confiables, seguras y al mejor precio. Nos comprometemos con la seguridad de cada familia que confía en nosotros.'
      },
      {
        icon: '🔭',
        titulo: 'Visión',
        texto: 'Ser el centro de servicio automotriz líder en Celaya y la región, reconocidos por nuestra excelencia en atención, tecnología de punta y el compromiso total con la satisfacción de nuestros clientes.'
      },
      {
        icon: '💎',
        titulo: 'Valores',
        texto: 'Honestidad, responsabilidad, calidad, compromiso y trabajo en equipo. Cada servicio que realizamos refleja nuestros valores y la pasión que sentimos por lo que hacemos.'
      }
    ];
  
    // ===== TIMELINE =====
    const timeline = [
      {
        year: '2009',
        titulo: 'Nuestro inicio',
        desc: 'Abrimos nuestras puertas en Celaya con la visión de ofrecer un servicio de llantas honesto y de calidad.'
      },
      {
        year: '2013',
        titulo: 'Crecimiento y expansión',
        desc: 'Ampliamos nuestras instalaciones e incorporamos servicios de alineación computarizada y balanceo dinámico.'
      },
      {
        year: '2017',
        titulo: 'Tecnología de punta',
        desc: 'Invertimos en equipos de última generación: alineación 3D, balanceadoras digitales y herramientas especializadas.'
      },
      {
        year: '2020',
        titulo: 'Servicio integral',
        desc: 'Incorporamos servicios de frenos, suspensión y cambio de aceite para ser tu centro automotriz completo.'
      },
      {
        year: '2024',
        titulo: 'Líderes en Celaya',
        desc: 'Más de 10,000 clientes satisfechos nos respaldan. Seguimos innovando para brindarte el mejor servicio.'
      }
    ];
  

    // ===== POR QUÉ ELEGIRNOS =====
    const razones = [
      { icon: '🏆', titulo: 'Experiencia comprobada', desc: 'Más de 15 años resolviendo las necesidades de miles de clientes en Celaya.' },
      { icon: '🔬', titulo: 'Tecnología avanzada', desc: 'Equipos de diagnóstico y servicio de última generación para resultados precisos.' },
      { icon: '🤝', titulo: 'Trato personalizado', desc: 'Cada cliente es único. Te asesoramos según las necesidades específicas de tu vehículo.' },
      { icon: '✅', titulo: 'Garantía total', desc: 'Todos nuestros servicios y productos cuentan con garantía por escrito.' },
      { icon: '⏱️', titulo: 'Rapidez y puntualidad', desc: 'Valoramos tu tiempo. Servicio eficiente sin comprometer la calidad.' },
      { icon: '💰', titulo: 'Precios justos', desc: 'Presupuestos transparentes y precios competitivos. Sin sorpresas.' }
    ];
  </script>
  
  <svelte:head>
    <title>Sobre Nosotros | ServiLlantas Celaya - Nuestra Historia</title>
    <meta name="description" content="Conoce la historia de ServiLlantas Celaya. Más de 15 años de experiencia en llantas y servicio automotriz profesional." />
  </svelte:head>
  
  <!-- ========== HERO ========== -->
  <section class="page-hero">
    <div class="page-hero-overlay"></div>
    <div class="hero-particles">
      {#each Array(15) as _, i}
        <div
          class="particle"
          style="--delay: {i * 0.6}s; --x: {Math.random() * 100}%; --duration: {3 + Math.random() * 4}s"
        ></div>
      {/each}
    </div>
    <div class="page-hero-content">
      <p class="hero-subtitle">CONÓCENOS</p>
      <h1>Sobre <span class="text-accent">Nosotros</span></h1>
      <p class="hero-desc">
        Más de 15 años comprometidos con la seguridad y confianza de tu vehículo en Celaya.
      </p>
      <nav class="breadcrumb">
        <a href="/">Inicio</a>
        <span>/</span>
        <span class="current">Sobre Nosotros</span>
      </nav>
    </div>
  </section>
  
  <!-- ========== QUIÉNES SOMOS ========== -->
  <section class="section section-dark">
    <div class="container">
      <div class="about-intro">
        <div class="about-text">
          <p class="section-subtitle fade-left" use:inview>¿Quiénes somos?</p>
          <h2 class="section-title fade-left delay-1" use:inview>
            Pasión por los <span class="text-accent">automóviles</span> desde 2009
          </h2>
          <p class="about-desc fade-left delay-2" use:inview>
            <strong>ServiLlantas Celaya</strong> nació con una visión clara: ofrecer un servicio automotriz
            <span class="text-accent">honesto, profesional y de calidad</span> a cada cliente que cruza nuestras puertas.
          </p>
          <p class="about-desc fade-left delay-3" use:inview>
            Lo que comenzó como un pequeño negocio de venta de llantas, hoy es un centro de servicio
            automotriz integral con tecnología de punta, un equipo altamente capacitado y la confianza
            de más de <strong class="text-accent">10,000 clientes satisfechos</strong>.
          </p>
          <p class="about-desc fade-left delay-4" use:inview>
            Nos especializamos en llantas, alineación computarizada, balanceo, frenos, suspensión
            y cambio de aceite. Cada servicio lo realizamos con el compromiso de garantizar
            tu seguridad y la de tu familia.
          </p>
          <div class="about-signature fade-left delay-5" use:inview>
            <div class="signature-line"></div>
          </div>
        </div>
  
        <div class="about-images fade-right" use:inview>
          <div class="about-img-main">
            <div class="image-placeholder-about">
              <span>🏪</span>
              <p>Fachada del taller</p>
              <small>static/images/about-main.webp</small>
            </div>
          </div>
          <div class="about-img-secondary fade-right delay-2" use:inview>
            <div class="image-placeholder-about small">
              <span>👨‍🔧</span>
              <p>Equipo trabajando</p>
              <small>static/images/about-team.webp</small>
            </div>
          </div>
          <!-- Badge flotante -->
          <div class="about-badge scale-in delay-3" use:inview>
            <span class="badge-number">15+</span>
            <span class="badge-text">años</span>
          </div>
        </div>
      </div>
    </div>
  </section>
  
  <!-- ========== STATS ========== -->
  <section class="stats-section" use:inview on:inview={startCounters}>
    <div class="container">
      <div class="stats-grid">
        {#each valores as stat}
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
  
  <!-- ========== MISIÓN / VISIÓN / VALORES ========== -->
  <section class="section section-darker">
    <div class="container text-center">
      <p class="section-subtitle fade-up" use:inview>Lo que nos define</p>
      <h2 class="section-title fade-up" use:inview>
        Misión, Visión y <span class="text-accent">Valores</span>
      </h2>
  
      <div class="pilares-grid">
        {#each pilares as pilar, i}
          <div class="pilar-card fade-up delay-{i + 1}" use:inview>
            <div class="pilar-icon-wrapper">
              <span class="pilar-icon">{pilar.icon}</span>
            </div>
            <h3>{pilar.titulo}</h3>
            <p>{pilar.texto}</p>
          </div>
        {/each}
      </div>
    </div>
  </section>
  
  <!-- ========== TIMELINE ========== -->
  <section class="section section-dark">
    <div class="container">
      <div class="text-center">
        <p class="section-subtitle fade-up" use:inview>Nuestra trayectoria</p>
        <h2 class="section-title fade-up" use:inview>
          Una historia de <span class="text-accent">crecimiento</span>
        </h2>
        <p class="section-desc fade-up" use:inview>
          Cada año hemos dado un paso más para ser el mejor centro de servicio automotriz en Celaya.
        </p>
      </div>
  
      <div class="timeline">
        <div class="timeline-line"></div>
        {#each timeline as event, i}
          <div
            class="timeline-item"
            class:right={i % 2 !== 0}
            class:left={i % 2 === 0}
          >
            <div class="timeline-content fade-up delay-{Math.min(i + 1, 5)}" use:inview>
              <span class="timeline-year">{event.year}</span>
              <h3>{event.titulo}</h3>
              <p>{event.desc}</p>
            </div>
            <div class="timeline-dot"></div>
          </div>
        {/each}
      </div>
    </div>
  </section>
  
  <!-- ========== POR QUÉ ELEGIRNOS ========== -->
  <section class="section section-darker">
    <div class="container text-center">
      <p class="section-subtitle fade-up" use:inview>¿Por qué elegirnos?</p>
      <h2 class="section-title fade-up" use:inview>
        Lo que nos hace <span class="text-accent">diferentes</span>
      </h2>
      <p class="section-desc fade-up" use:inview>
        No solo ofrecemos servicios automotrices, ofrecemos tranquilidad y confianza.
      </p>
  
      <div class="razones-grid">
        {#each razones as razon, i}
          <div class="razon-card fade-up delay-{Math.min(i + 1, 6)}" use:inview>
            <div class="razon-icon">{razon.icon}</div>
            <h4>{razon.titulo}</h4>
            <p>{razon.desc}</p>
          </div>
        {/each}
      </div>
    </div>
  </section>
  
  <!-- ========== INSTALACIONES ========== -->
  <section class="section section-darker">
    <div class="container text-center">
      <p class="section-subtitle fade-up" use:inview>Nuestras instalaciones</p>
      <h2 class="section-title fade-up" use:inview>
        Un taller <span class="text-accent">equipado</span> para todo
      </h2>
      <p class="section-desc fade-up" use:inview>
        Instalaciones amplias y equipo de última tecnología para brindarte un servicio profesional.
      </p>
  
      <div class="gallery-grid">
        {#each ['Área de alineación', 'Zona de balanceo', 'Bahía de servicio', 'Sala de espera', 'Almacén de llantas', 'Equipo de diagnóstico'] as area, i}
          <div class="gallery-item fade-up delay-{Math.min(i + 1, 6)}" use:inview>
            <div class="gallery-placeholder">
              <span>📷</span>
              <p>{area}</p>
              <small>static/images/galeria/{i + 1}.jpg</small>
            </div>
          </div>
        {/each}
      </div>
    </div>
  </section>
  
  <!-- ========== CTA FINAL ========== -->
  <section class="cta-final">
    <div class="cta-final-overlay"></div>
    <div class="container text-center" style="position: relative; z-index: 2;">
      <h2 class="fade-up" use:inview>
        ¿Listo para conocernos <span class="text-accent">en persona</span>?
      </h2>
      <p class="cta-final-desc fade-up delay-1" use:inview>
        Visítanos en Celaya o llámanos. Será un gusto atenderte.
      </p>
      <div class="cta-final-buttons fade-up delay-2" use:inview>
        <a href="/contacto" class="btn btn-accent btn-lg">✉️ Contáctanos</a>
        <a href="tel:+524611203488" class="btn btn-primary btn-lg">📞 461 120 3488</a>
        <a href="tel:+527201987926" class="btn btn-outline btn-lg">📱 720 198 7926</a>
      </div>
    </div>
  </section>
  
  <style>
    /* ====== PAGE HERO ====== */
    .page-hero {
      position: relative;
      padding: 10rem 2rem 5rem;
      text-align: center;
      background: url('/images/hero-bg.jpg') center/cover no-repeat;
      background-color: var(--color-primary);
      overflow: hidden;
    }
  
    .page-hero-overlay {
      position: absolute;
      inset: 0;
      background: linear-gradient(
        180deg,
        rgba(22, 17, 47, 0.95) 0%,
        rgba(22, 17, 47, 0.85) 100%
      );
    }
  
    .hero-particles {
      position: absolute;
      inset: 0;
      z-index: 1;
      overflow: hidden;
    }
  
    .particle {
      position: absolute;
      width: 4px;
      height: 4px;
      background: var(--color-accent);
      border-radius: 50%;
      left: var(--x);
      bottom: -10px;
      opacity: 0;
      animation: floatUp var(--duration) var(--delay) infinite;
    }
  
    @keyframes floatUp {
      0% { opacity: 0; transform: translateY(0) scale(0); }
      10% { opacity: 0.6; }
      90% { opacity: 0.2; }
      100% { opacity: 0; transform: translateY(-100vh) scale(1); }
    }
  
    .page-hero-content {
      position: relative;
      z-index: 2;
      max-width: 700px;
      margin: 0 auto;
    }
  
    .page-hero .hero-subtitle {
      color: var(--color-accent);
      font-family: var(--font-heading);
      font-weight: 700;
      font-size: 0.85rem;
      letter-spacing: 4px;
      margin-bottom: 1rem;
    }
  
    .page-hero h1 {
      font-size: clamp(2.2rem, 5vw, 3.5rem);
      margin-bottom: 1rem;
    }
  
    .page-hero .hero-desc {
      color: var(--color-light-3);
      font-size: 1.1rem;
      margin-bottom: 1.5rem;
    }
  
    .breadcrumb {
      display: flex;
      justify-content: center;
      gap: 0.5rem;
      font-size: 0.9rem;
      color: var(--color-gray);
    }
  
    .breadcrumb a {
      color: var(--color-light-3);
      transition: color 0.3s;
    }
  
    .breadcrumb a:hover { color: var(--color-accent); }
    .breadcrumb .current { color: var(--color-accent); }
  
    /* ====== ABOUT INTRO ====== */
    .about-intro {
      display: grid;
      grid-template-columns: 1fr 1fr;
      gap: 4rem;
      align-items: center;
    }
  
    .about-desc {
      color: var(--color-light-3);
      font-size: 1rem;
      line-height: 1.9;
      margin-bottom: 1.2rem;
    }
  
    .about-desc strong {
      color: var(--color-light);
    }
  
    .about-signature {
      display: flex;
      align-items: center;
      gap: 1rem;
      margin-top: 2rem;
      padding-top: 1.5rem;
      border-top: 1px solid var(--color-dark-4);
    }
  
    .signature-line {
      width: 50px;
      height: 3px;
      background: var(--color-accent);
      border-radius: 2px;
    }
  
    .signature-name {
      display: block;
      font-family: var(--font-heading);
      font-weight: 700;
      font-size: 1rem;
    }
  
    .signature-role {
      display: block;
      color: var(--color-gray);
      font-size: 0.85rem;
      margin-top: 0.2rem;
    }
  
    /* About images */
    .about-images {
      position: relative;
    }
  
    .about-img-main {
      position: relative;
      z-index: 1;
    }
  
    .about-img-secondary {
      position: absolute;
      bottom: -30px;
      right: -20px;
      width: 55%;
      z-index: 2;
    }
  
    .image-placeholder-about {
      background: var(--color-dark-3);
      border: 2px dashed var(--color-dark-4);
      border-radius: 20px;
      height: 380px;
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
      gap: 0.8rem;
      color: var(--color-gray);
      transition: border-color 0.3s;
    }
  
    .image-placeholder-about:hover {
      border-color: var(--color-secondary);
    }
  
    .image-placeholder-about.small {
      height: 200px;
      border-color: var(--color-secondary);
      background: var(--color-dark-2);
      box-shadow: 0 10px 30px rgba(0, 0, 0, 0.4);
    }
  
    .image-placeholder-about span {
      font-size: 3rem;
    }
  
    .image-placeholder-about p {
      font-size: 0.9rem;
    }
  
    .image-placeholder-about small {
      font-size: 0.7rem;
      color: var(--color-dark-4);
    }
  
    .about-badge {
      position: absolute;
      top: -15px;
      left: -15px;
      width: 90px;
      height: 90px;
      background: var(--color-accent);
      border-radius: 50%;
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
      z-index: 3;
      box-shadow: 0 8px 25px rgba(243, 253, 1, 0.3);
    }
  
    .badge-number {
      font-family: var(--font-heading);
      font-weight: 900;
      font-size: 1.6rem;
      color: var(--color-primary);
      line-height: 1;
    }
  
    .badge-text {
      font-size: 0.75rem;
      font-weight: 700;
      color: var(--color-primary);
      text-transform: uppercase;
      letter-spacing: 1px;
    }
  
    /* ====== STATS ====== */
    .stats-section {
      background: var(--color-secondary);
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
      color: var(--color-accent);
    }
  
    .stat-label {
      font-size: 0.9rem;
      color: rgba(255, 255, 255, 0.85);
      margin-top: 0.3rem;
    }
  
    /* ====== PILARES (Misión / Visión / Valores) ====== */
    .pilares-grid {
      display: grid;
      grid-template-columns: repeat(3, 1fr);
      gap: 2rem;
      margin-top: 2rem;
    }
  
    .pilar-card {
      background: var(--color-dark-3);
      border: 1px solid var(--color-dark-4);
      border-radius: 20px;
      padding: 2.5rem 2rem;
      text-align: center;
      transition: all 0.4s ease;
      position: relative;
      overflow: hidden;
    }
  
    .pilar-card::after {
      content: '';
      position: absolute;
      bottom: 0;
      left: 0;
      width: 100%;
      height: 4px;
      background: linear-gradient(90deg, var(--color-secondary), var(--color-accent));
      transform: scaleX(0);
      transition: transform 0.4s ease;
    }
  
    .pilar-card:hover {
      transform: translateY(-8px);
      border-color: var(--color-secondary);
      box-shadow: 0 20px 40px rgba(0, 75, 188, 0.15);
    }
  
    .pilar-card:hover::after {
      transform: scaleX(1);
    }
  
    .pilar-icon-wrapper {
      width: 70px;
      height: 70px;
      background: rgba(0, 75, 188, 0.15);
      border: 2px solid var(--color-secondary);
      border-radius: 50%;
      display: flex;
      align-items: center;
      justify-content: center;
      margin: 0 auto 1.5rem;
      transition: all 0.3s ease;
    }
  
    .pilar-card:hover .pilar-icon-wrapper {
      background: var(--color-secondary);
      transform: scale(1.1);
    }
  
    .pilar-icon {
      font-size: 2rem;
    }
  
    .pilar-card h3 {
      font-size: 1.3rem;
      margin-bottom: 1rem;
    }
  
    .pilar-card p {
      color: var(--color-gray);
      font-size: 0.9rem;
      line-height: 1.8;
    }
  
    /* ====== TIMELINE ====== */
    .timeline {
      position: relative;
      max-width: 900px;
      margin: 3rem auto 0;
      padding: 2rem 0;
    }
  
    .timeline-line {
      position: absolute;
      top: 0;
      bottom: 0;
      left: 50%;
      width: 3px;
      background: linear-gradient(
        180deg,
        var(--color-secondary),
        var(--color-accent),
        var(--color-secondary)
      );
      transform: translateX(-50%);
      border-radius: 3px;
    }
  
    .timeline-item {
      position: relative;
      width: 50%;
      padding: 0 3rem 3rem;
    }
  
    .timeline-item.left {
      left: 0;
      text-align: right;
    }
  
    .timeline-item.right {
      left: 50%;
      text-align: left;
    }
  
    .timeline-dot {
      position: absolute;
      top: 5px;
      width: 18px;
      height: 18px;
      background: var(--color-accent);
      border: 3px solid var(--color-primary);
      border-radius: 50%;
      z-index: 2;
      box-shadow: 0 0 15px rgba(243, 253, 1, 0.4);
    }
  
    .timeline-item.left .timeline-dot {
      right: -9px;
    }
  
    .timeline-item.right .timeline-dot {
      left: -9px;
    }
  
    .timeline-content {
      background: var(--color-dark-3);
      border: 1px solid var(--color-dark-4);
      border-radius: 16px;
      padding: 1.8rem;
      transition: all 0.3s ease;
    }
  
    .timeline-content:hover {
      border-color: var(--color-secondary);
      transform: translateY(-3px);
      box-shadow: 0 10px 25px rgba(0, 0, 0, 0.3);
    }
  
    .timeline-year {
      display: inline-block;
      font-family: var(--font-heading);
      font-weight: 900;
      font-size: 1.8rem;
      color: var(--color-accent);
      margin-bottom: 0.5rem;
      line-height: 1;
    }
  
    .timeline-content h3 {
      font-size: 1.1rem;
      margin-bottom: 0.5rem;
    }
  
    .timeline-content p {
      color: var(--color-gray);
      font-size: 0.9rem;
      line-height: 1.7;
    }
  
    /* ====== RAZONES ====== */
    .razones-grid {
      display: grid;
      grid-template-columns: repeat(3, 1fr);
      gap: 2rem;
      margin-top: 2rem;
    }
  
    .razon-card {
      background: var(--color-dark-3);
      border: 1px solid var(--color-dark-4);
      border-radius: 16px;
      padding: 2rem;
      text-align: center;
      transition: all 0.3s ease;
    }
  
    .razon-card:hover {
      border-color: var(--color-accent);
      transform: translateY(-5px);
      box-shadow: 0 15px 35px rgba(0, 0, 0, 0.2);
    }
  
    .razon-icon {
      font-size: 2.5rem;
      margin-bottom: 1rem;
    }
  
    .razon-card h4 {
      font-size: 1.05rem;
      margin-bottom: 0.6rem;
      font-weight: 700;
    }
  
    .razon-card p {
      color: var(--color-gray);
      font-size: 0.9rem;
      line-height: 1.6;
    }
  
    /* ====== TEAM ====== */
    .team-grid {
      display: grid;
      grid-template-columns: repeat(4, 1fr);
      gap: 2rem;
      margin-top: 2rem;
    }
  
    .team-card {
      background: var(--color-dark-3);
      border: 1px solid var(--color-dark-4);
      border-radius: 20px;
      overflow: hidden;
      transition: all 0.4s ease;
    }
  
    .team-card:hover {
      border-color: var(--color-secondary);
      transform: translateY(-8px);
      box-shadow: 0 20px 40px rgba(0, 75, 188, 0.15);
    }
  
    .team-photo {
      position: relative;
      height: 240px;
      overflow: hidden;
    }
  
    .team-avatar {
      width: 100%;
      height: 100%;
      background: linear-gradient(135deg, var(--color-secondary), var(--color-primary-light));
      display: flex;
      align-items: center;
      justify-content: center;
    }
  
    .team-avatar span {
      font-family: var(--font-heading);
      font-weight: 900;
      font-size: 3rem;
      color: var(--color-accent);
    }
  
    .team-overlay {
      position: absolute;
      inset: 0;
      background: rgba(0, 75, 188, 0.9);
      display: flex;
      align-items: center;
      justify-content: center;
      padding: 1.5rem;
      opacity: 0;
      transition: opacity 0.4s ease;
    }
  
    .team-card:hover .team-overlay {
      opacity: 1;
    }
  
    .team-overlay p {
      color: var(--color-light);
      font-size: 0.9rem;
      line-height: 1.7;
      text-align: center;
    }
  
    .team-info {
      padding: 1.5rem;
      text-align: center;
    }
  
    .team-info h4 {
      font-size: 1.05rem;
      font-weight: 700;
      margin-bottom: 0.3rem;
    }
  
    .team-role {
      color: var(--color-accent);
      font-size: 0.85rem;
      font-weight: 600;
    }
  
    /* ====== GALLERY ====== */
    .gallery-grid {
      display: grid;
      grid-template-columns: repeat(3, 1fr);
      gap: 1.5rem;
      margin-top: 2rem;
    }
  
    .gallery-item {
      border-radius: 16px;
      overflow: hidden;
    }
  
    .gallery-placeholder {
      background: var(--color-dark-3);
      border: 2px dashed var(--color-dark-4);
      border-radius: 16px;
      height: 220px;
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
      gap: 0.5rem;
      color: var(--color-gray);
      transition: all 0.3s ease;
    }
  
    .gallery-placeholder:hover {
      border-color: var(--color-secondary);
      transform: scale(1.02);
    }
  
    .gallery-placeholder span {
      font-size: 2.5rem;
    }
  
    .gallery-placeholder p {
      font-size: 0.85rem;
      font-weight: 600;
    }
  
    .gallery-placeholder small {
      font-size: 0.7rem;
      color: var(--color-dark-4);
    }
  
    /* ====== CTA FINAL ====== */
    .cta-final {
      position: relative;
      padding: 5rem 0;
      background: url('/images/hero-bg.jpg') center/cover no-repeat fixed;
      background-color: var(--color-primary);
    }
  
    .cta-final-overlay {
      position: absolute;
      inset: 0;
      background: linear-gradient(135deg, rgba(0, 75, 188, 0.92), rgba(22, 17, 47, 0.95));
    }
  
    .cta-final h2 {
      font-size: clamp(1.8rem, 4vw, 2.8rem);
      margin-bottom: 1rem;
    }
  
    .cta-final-desc {
      color: rgba(255, 255, 255, 0.8);
      font-size: 1.1rem;
      margin-bottom: 2rem;
    }
  
    .cta-final-buttons {
      display: flex;
      gap: 1rem;
      justify-content: center;
      flex-wrap: wrap;
    }
  
    .btn-lg {
      padding: 1rem 2.5rem;
      font-size: 1.05rem;
    }
  
    /* ====== RESPONSIVE ====== */
    @media (max-width: 1024px) {
      .team-grid {
        grid-template-columns: repeat(2, 1fr);
      }
  
      .pilares-grid,
      .razones-grid {
        grid-template-columns: repeat(2, 1fr);
      }
    }
  
    @media (max-width: 768px) {
      .about-intro {
        grid-template-columns: 1fr;
        gap: 3rem;
      }
  
      .about-text {
        text-align: center;
      }
  
      .about-signature {
        justify-content: center;
      }
  
      .about-img-secondary {
        position: relative;
        bottom: auto;
        right: auto;
        width: 100%;
        margin-top: 1rem;
      }
  
      .about-badge {
        top: -10px;
        left: auto;
        right: -10px;
      }
  
      .stats-grid {
        grid-template-columns: repeat(2, 1fr);
      }
  
      .pilares-grid,
      .razones-grid,
      .gallery-grid {
        grid-template-columns: 1fr;
      }
  
      .team-grid {
        grid-template-columns: 1fr;
        max-width: 350px;
        margin: 2rem auto 0;
      }
  
      /* Timeline mobile */
      .timeline-line {
        left: 20px;
      }
  
      .timeline-item {
        width: 100%;
        left: 0 !important;
        padding-left: 50px;
        padding-right: 0;
        text-align: left !important;
      }
  
      .timeline-item .timeline-dot {
        left: 11px !important;
        right: auto !important;
      }
  
      .page-hero {
        padding: 8rem 1.5rem 3rem;
      }
    }
  
    @media (max-width: 480px) {
      .stat-number {
        font-size: 2rem;
      }
  
      .cta-final-buttons {
        flex-direction: column;
        align-items: center;
      }
  
      .cta-final-buttons .btn {
        width: 100%;
        justify-content: center;
      }
    }
  </style>