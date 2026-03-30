<!-- frontend/src/routes/servicios/+page.svelte -->
<script>
  import { inview } from '$lib/actions/inview.js';

  const servicios = [
    {
      id: 'llantas',
      icon: '🛞',
      titulo: 'Venta de Llantas',
      subtitulo: 'Las mejores marcas del mercado',
      descripcion: 'Contamos con un amplio catálogo de llantas nacionales e internacionales para todo tipo de vehículo: automóvil, camioneta, SUV, carga y más. Te asesoramos para que elijas la llanta perfecta según tu tipo de manejo y presupuesto.',
      caracteristicas: [
        'Marcas premium: Michelin, Bridgestone, Goodyear, Continental',
        'Marcas económicas de alta calidad',
        'Llantas para todo terreno, carretera y uso mixto',
        'Asesoría personalizada según tu vehículo',
        'Instalación incluida sin costo extra',
        'Garantía de fábrica en todas las llantas'
      ],
      imagen: '/images/servicios/llantas.jpg'
    },
    {
      id: 'alineacion',
      icon: '🎯',
      titulo: 'Alineación Computarizada',
      subtitulo: 'Precisión milimétrica',
      descripcion: 'Utilizamos equipos de alineación computarizada de última generación para garantizar que tu vehículo se mantenga recto y seguro. Una mala alineación desgasta tus llantas prematuramente y afecta el control del vehículo.',
      caracteristicas: [
        'Equipo de alineación 3D de última generación',
        'Alineación en las 4 ruedas',
        'Ajuste de convergencia, caída y caster',
        'Diagnóstico completo del tren delantero',
        'Servicio rápido: 30-45 minutos',
        'Reporte impreso con valores antes y después'
      ],
      imagen: '/images/servicios/alineacion.jpg'
    },
    {
      id: 'balanceo',
      icon: '⚖️',
      titulo: 'Balanceo Dinámico',
      subtitulo: 'Manejo suave y sin vibraciones',
      descripcion: 'El balanceo elimina las vibraciones en el volante y asiento, mejorando la comodidad y prolongando la vida de tus llantas, amortiguadores y componentes de la suspensión.',
      caracteristicas: [
        'Balanceadora dinámica computarizada',
        'Plomos de alta precisión',
        'Detección de llantas dañadas o deformadas',
        'Incluido con la compra de llantas nuevas',
        'Servicio express disponible',
        'Balanceo de rines deportivos y de lujo'
      ],
      imagen: '/images/servicios/balanceo.jpg'
    },
    {
      id: 'frenos',
      icon: '🛑',
      titulo: 'Sistema de Frenos',
      subtitulo: 'Tu seguridad es nuestra prioridad',
      descripcion: 'Realizamos diagnóstico, mantenimiento y reparación completa del sistema de frenado. Desde el cambio de balatas hasta la rectificación de discos y sustitución de componentes hidráulicos.',
      caracteristicas: [
        'Cambio de balatas delanteras y traseras',
        'Rectificación y cambio de discos',
        'Revisión del líquido de frenos',
        'Reparación de cáliper y cilindros',
        'Purga del sistema hidráulico',
        'Diagnóstico gratuito del sistema de frenos'
      ],
      imagen: '/images/servicios/frenos.jpg'
    },
    {
      id: 'suspension',
      icon: '🔧',
      titulo: 'Suspensión',
      subtitulo: 'Estabilidad y control en cada camino',
      descripcion: 'El sistema de suspensión es clave para la seguridad y comodidad. Diagnosticamos y reparamos amortiguadores, rótulas, terminales, bujes, barras estabilizadoras y todo lo necesario.',
      caracteristicas: [
        'Cambio de amortiguadores y struts',
        'Reemplazo de rótulas y terminales',
        'Cambio de bujes y bases de amortiguador',
        'Barras estabilizadoras y ligas',
        'Diagnóstico visual y prueba en carretera',
        'Mano de obra especializada'
      ],
      imagen: '/images/servicios/suspension.jpg'
    },
    {
      id: 'aceite',
      icon: '🛢️',
      titulo: 'Cambio de Aceite',
      subtitulo: 'Protección máxima para tu motor',
      descripcion: 'El aceite es la sangre de tu motor. Realizamos cambios de aceite con las mejores marcas sintéticas y minerales. Incluimos revisión de niveles de fluidos y filtro de aceite.',
      caracteristicas: [
        'Aceites sintéticos y semi-sintéticos premium',
        'Marcas: Mobil, Castrol, Pennzoil, Valvoline',
        'Cambio de filtro de aceite incluido',
        'Revisión de 21 puntos incluida',
        'Para motores a gasolina y diésel',
        'Registro de servicio para tu historial'
      ],
      imagen: '/images/servicios/aceite.jpg'
    }
  ];

  const procesoSteps = [
    { num: '01', titulo: 'Diagnóstico', desc: 'Evaluamos tu vehículo y te damos un presupuesto transparente sin compromiso.' },
    { num: '02', titulo: 'Aprobación', desc: 'Tú decides. Te explicamos cada detalle para que tomes la mejor decisión.' },
    { num: '03', titulo: 'Servicio', desc: 'Nuestros técnicos realizan el trabajo con equipo profesional y refacciones de calidad.' },
    { num: '04', titulo: 'Entrega', desc: 'Revisamos todo, te explicamos lo realizado y sales con garantía por escrito.' }
  ];
</script>

<svelte:head>
  <title>Servicios | ServiLlantas Celaya - Llantas, Alineación, Frenos y más</title>
  <meta name="description" content="Servicios automotrices profesionales: venta de llantas, alineación computarizada, balanceo, frenos, suspensión y cambio de aceite en Celaya." />
</svelte:head>

<!-- ========== HERO SERVICIOS ========== -->
<section class="page-hero">
  <div class="page-hero-overlay"></div>
  <div class="page-hero-content">
    <p class="hero-subtitle">SERVICIOS PROFESIONALES</p>
    <h1>Nuestros <span class="text-accent">Servicios</span></h1>
    <p class="hero-desc">
      Todo lo que tu vehículo necesita en un solo lugar. Calidad, experiencia y garantía.
    </p>
    <!-- Breadcrumb -->
    <nav class="breadcrumb">
      <a href="/">Inicio</a>
      <span>/</span>
      <span class="active">Servicios</span>
    </nav>
  </div>
</section>

<!-- ========== INTRO ========== -->
<section class="section section-dark">
  <div class="container text-center">
    <p class="section-subtitle fade-up" use:inview>¿Qué ofrecemos?</p>
    <h2 class="section-title fade-up" use:inview>
      Servicios automotrices <span class="text-accent">integrales</span>
    </h2>
    <p class="section-desc fade-up" use:inview>
      Desde la venta de llantas hasta el mantenimiento completo de tu vehículo.
      Contamos con equipo de última generación y técnicos certificados.
    </p>

    <!-- Quick nav de servicios -->
    <div class="services-nav fade-up delay-2" use:inview>
      {#each servicios as s}
        <a href="#{s.id}" class="service-nav-item">
          <span class="service-nav-icon">{s.icon}</span>
          <span class="service-nav-label">{s.titulo}</span>
        </a>
      {/each}
    </div>
  </div>
</section>

<!-- ========== SERVICIOS DETALLADOS ========== -->
{#each servicios as servicio, i}
  <section
    id={servicio.id}
    class="section"
    class:section-darker={i % 2 === 0}
    class:section-dark={i % 2 !== 0}
  >
    <div class="container">
      <div class="service-detail" class:reverse={i % 2 !== 0}>
        <!-- Info -->
        <div class="service-info">
          <div class="service-badge fade-left" use:inview>
            <span class="badge-icon">{servicio.icon}</span>
          </div>
          <p class="section-subtitle fade-left delay-1" use:inview>
            {servicio.subtitulo}
          </p>
          <h2 class="section-title fade-left delay-2" use:inview>
            {servicio.titulo}
          </h2>
          <p class="service-desc fade-left delay-3" use:inview>
            {servicio.descripcion}
          </p>

          <ul class="features-list">
            {#each servicio.caracteristicas as feature, fi}
              <li class="fade-left delay-{Math.min(fi + 1, 6)}" use:inview>
                <span class="feature-check">✓</span>
                <span>{feature}</span>
              </li>
            {/each}
          </ul>

          <div class="service-actions fade-left delay-4" use:inview>
            <a href="/contacto" class="btn btn-accent">
              Cotizar este servicio
            </a>
            <a href="tel:+524611203488" class="btn btn-outline-accent">
              📞 Llamar ahora
            </a>
          </div>
        </div>

        <!-- Imagen -->
        <div class="service-image fade-right" use:inview>
          <div class="image-placeholder">
            <span class="img-icon">{servicio.icon}</span>
            <p>{servicio.titulo}</p>
            <small>Coloca imagen en:<br/>static{servicio.imagen}</small>
          </div>
          <!-- Cuando tengas la imagen reemplaza con:
          <img src={servicio.imagen} alt={servicio.titulo} />
          -->
        </div>
      </div>
    </div>
  </section>
{/each}

<!-- ========== PROCESO ========== -->
<section class="section section-primary">
  <div class="container text-center">
    <p class="section-subtitle fade-up" use:inview style="color: rgba(255,255,255,0.7)">
      Nuestro Proceso
    </p>
    <h2 class="section-title fade-up" use:inview>
      Así de <span class="text-accent">fácil</span> es
    </h2>
    <p class="section-desc fade-up" use:inview style="color: rgba(255,255,255,0.6)">
      En 4 simples pasos tu vehículo estará como nuevo
    </p>

    <div class="process-grid">
      {#each procesoSteps as step, i}
        <div class="process-step fade-up delay-{i + 1}" use:inview>
          <span class="step-number">{step.num}</span>
          <h3>{step.titulo}</h3>
          <p>{step.desc}</p>
          {#if i < procesoSteps.length - 1}
            <div class="step-connector"></div>
          {/if}
        </div>
      {/each}
    </div>
  </div>
</section>

<!-- ========== CTA FINAL ========== -->
<section class="section section-dark">
  <div class="container text-center">
    <h2 class="fade-up" use:inview>
      ¿Necesitas alguno de estos servicios?<br/>
      <span class="text-accent">¡Estamos listos para ayudarte!</span>
    </h2>
    <p class="section-desc fade-up delay-1" use:inview>
      Llámanos, envíanos un WhatsApp o agenda tu cita en línea.
    </p>
    <div class="cta-buttons fade-up delay-2" use:inview>
      <a href="tel:+524611203488" class="btn btn-accent btn-lg">📞 461 120 3488</a>
      <a href="tel:+527201987926" class="btn btn-primary btn-lg">📱 720 198 7926</a>
      <a href="/contacto" class="btn btn-outline btn-lg">✉️ Agendar Cita</a>
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
    transition: color 0.3s ease;
  }

  .breadcrumb a:hover {
    color: var(--color-accent);
  }

  .breadcrumb .active {
    color: var(--color-accent);
  }

  /* ====== SERVICES NAV ====== */
  .services-nav {
    display: grid;
    grid-template-columns: repeat(6, 1fr);
    gap: 1rem;
    margin-top: 2rem;
  }

  .service-nav-item {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 0.5rem;
    padding: 1.2rem 0.5rem;
    background: var(--color-dark-3);
    border: 1px solid var(--color-dark-4);
    border-radius: 12px;
    transition: all 0.3s ease;
  }

  .service-nav-item:hover {
    border-color: var(--color-accent);
    background: rgba(243, 253, 1, 0.05);
    transform: translateY(-4px);
  }

  .service-nav-icon {
    font-size: 2rem;
  }

  .service-nav-label {
    font-size: 0.75rem;
    font-weight: 600;
    text-align: center;
    color: var(--color-light-3);
    line-height: 1.3;
  }

  .service-nav-item:hover .service-nav-label {
    color: var(--color-accent);
  }

  /* ====== SERVICE DETAIL ====== */
  .service-detail {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 4rem;
    align-items: center;
  }

  .service-detail.reverse {
    direction: rtl;
  }

  .service-detail.reverse > * {
    direction: ltr;
  }

  .service-badge {
    margin-bottom: 1rem;
  }

  .badge-icon {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    width: 70px;
    height: 70px;
    font-size: 2.5rem;
    background: rgba(0, 75, 188, 0.15);
    border: 2px solid var(--color-secondary);
    border-radius: 16px;
  }

  .service-desc {
    color: var(--color-light-3);
    font-size: 1rem;
    line-height: 1.8;
    margin-bottom: 2rem;
  }

  .features-list {
    display: flex;
    flex-direction: column;
    gap: 0.8rem;
    margin-bottom: 2rem;
  }

  .features-list li {
    display: flex;
    align-items: center;
    gap: 0.8rem;
    font-size: 0.95rem;
    color: var(--color-light-3);
  }

  .feature-check {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    width: 24px;
    height: 24px;
    background: var(--color-accent);
    color: var(--color-primary);
    border-radius: 50%;
    font-size: 0.75rem;
    font-weight: 800;
    flex-shrink: 0;
  }

  .service-actions {
    display: flex;
    gap: 1rem;
    flex-wrap: wrap;
  }

  /* Service Image */
  .service-image {
    position: relative;
  }

  .image-placeholder {
    background: var(--color-dark-3);
    border: 2px dashed var(--color-dark-4);
    border-radius: 20px;
    height: 400px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: 0.8rem;
    color: var(--color-gray);
    transition: all 0.3s ease;
  }

  .image-placeholder:hover {
    border-color: var(--color-secondary);
  }

  .img-icon {
    font-size: 4rem;
  }

  .image-placeholder small {
    font-size: 0.75rem;
    color: var(--color-dark-4);
    text-align: center;
    line-height: 1.5;
  }

  /* ====== PROCESS ====== */
  .process-grid {
    display: grid;
    grid-template-columns: repeat(4, 1fr);
    gap: 2rem;
    margin-top: 2rem;
  }

  .process-step {
    position: relative;
    padding: 2rem 1.5rem;
    background: rgba(255, 255, 255, 0.05);
    border: 1px solid rgba(255, 255, 255, 0.1);
    border-radius: 16px;
    transition: all 0.3s ease;
  }

  .process-step:hover {
    background: rgba(255, 255, 255, 0.1);
    transform: translateY(-5px);
  }

  .step-number {
    display: block;
    font-family: var(--font-heading);
    font-weight: 900;
    font-size: 2.5rem;
    color: var(--color-accent);
    margin-bottom: 0.8rem;
    line-height: 1;
  }

  .process-step h3 {
    font-size: 1.1rem;
    margin-bottom: 0.8rem;
  }

  .process-step p {
    color: rgba(255, 255, 255, 0.65);
    font-size: 0.9rem;
    line-height: 1.6;
  }

  .step-connector {
    display: none; /* Visible en desktop como línea entre pasos */
  }

  /* ====== CTA ====== */
  .cta-buttons {
    display: flex;
    gap: 1rem;
    justify-content: center;
    flex-wrap: wrap;
    margin-top: 2rem;
  }

  .btn-lg {
    padding: 1rem 2.5rem;
    font-size: 1.05rem;
  }

  /* ====== RESPONSIVE ====== */
  @media (max-width: 1024px) {
    .services-nav {
      grid-template-columns: repeat(3, 1fr);
    }

    .process-grid {
      grid-template-columns: repeat(2, 1fr);
    }
  }

  @media (max-width: 768px) {
    .service-detail {
      grid-template-columns: 1fr;
      gap: 2rem;
    }

    .service-detail.reverse {
      direction: ltr;
    }

    .services-nav {
      grid-template-columns: repeat(2, 1fr);
    }

    .process-grid {
      grid-template-columns: 1fr;
    }

    .page-hero {
      padding: 8rem 1.5rem 3rem;
    }

    .service-actions {
      flex-direction: column;
    }

    .service-actions .btn {
      width: 100%;
      justify-content: center;
    }
  }
</style>