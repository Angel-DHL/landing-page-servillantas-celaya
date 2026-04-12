<!-- frontend/src/routes/privacidad/+page.svelte -->
<script>
    import { inview } from '$lib/actions/inview.js';
  </script>
  
  <svelte:head>
    <title>Aviso de Privacidad | ServiLlantas Celaya</title>
    <meta name="description" content="Aviso de Privacidad de ServiLlantas Celaya. Conoce cómo protegemos tus datos personales." />
  </svelte:head>
  
  <!-- ========== HERO ========== -->
  <section class="page-hero">
    <div class="page-hero-overlay"></div>
    <div class="page-hero-content">
      <p class="hero-subtitle">LEGAL</p>
      <h1>Aviso de <span class="text-accent">Privacidad</span></h1>
      <p class="hero-desc">Última actualización: Enero 2025</p>
      <nav class="breadcrumb">
        <a href="/">Inicio</a>
        <span>/</span>
        <span class="current">Aviso de Privacidad</span>
      </nav>
    </div>
  </section>
  
  <!-- ========== CONTENIDO ========== -->
  <section class="section section-dark">
    <div class="container">
      <div class="legal-content fade-up" use:inview>
  
        <!-- Índice -->
        <div class="legal-index">
          <h3>📋 Contenido</h3>
          <ol>
            <li><a href="#responsable">Responsable de los datos</a></li>
            <li><a href="#datos">Datos que recopilamos</a></li>
            <li><a href="#finalidad">Finalidad del tratamiento</a></li>
            <li><a href="#transferencia">Transferencia de datos</a></li>
            <li><a href="#derechos">Tus derechos ARCO</a></li>
            <li><a href="#cookies">Uso de cookies</a></li>
            <li><a href="#cambios">Cambios al aviso</a></li>
            <li><a href="#contacto-legal">Contacto</a></li>
          </ol>
        </div>
  
        <!-- Secciones -->
        <div class="legal-section" id="responsable">
          <h2>1. Responsable de los datos personales</h2>
          <p>
            <strong>ServiLlantas Celaya</strong>, con domicilio en Celaya, Guanajuato, México,
            es responsable del tratamiento de tus datos personales conforme a lo dispuesto
            por la <strong>Ley Federal de Protección de Datos Personales en Posesión de los
            Particulares</strong> (LFPDPPP).
          </p>
          <div class="legal-card">
            <p><strong>📍 Dirección:</strong> Celaya, Guanajuato, México</p>
            <p><strong>📞 Teléfono:</strong> 461 120 3488</p>
            <p><strong>📱 Teléfono 2:</strong> 720 198 7926</p>
          </div>
        </div>
  
        <div class="legal-section" id="datos">
          <h2>2. Datos personales que recopilamos</h2>
          <p>Para brindarte nuestros servicios, podemos recopilar los siguientes datos:</p>
          <div class="legal-grid">
            <div class="legal-item">
              <span class="legal-icon">👤</span>
              <div>
                <strong>Datos de identificación</strong>
                <p>Nombre completo</p>
              </div>
            </div>
            <div class="legal-item">
              <span class="legal-icon">📞</span>
              <div>
                <strong>Datos de contacto</strong>
                <p>Teléfono, correo electrónico</p>
              </div>
            </div>
            <div class="legal-item">
              <span class="legal-icon">🚗</span>
              <div>
                <strong>Datos del vehículo</strong>
                <p>Marca, modelo, año</p>
              </div>
            </div>
            <div class="legal-item">
              <span class="legal-icon">🌐</span>
              <div>
                <strong>Datos de navegación</strong>
                <p>Cookies, dirección IP, tipo de navegador</p>
              </div>
            </div>
          </div>
          <p class="legal-note">
            <strong>⚠️ Nota:</strong> No recopilamos datos personales sensibles como
            información financiera, datos de salud o creencias religiosas.
          </p>
        </div>
  
        <div class="legal-section" id="finalidad">
          <h2>3. Finalidad del tratamiento</h2>
          <p>Tus datos personales serán utilizados para las siguientes finalidades:</p>
  
          <h3>Finalidades primarias (necesarias):</h3>
          <ul>
            <li>Responder a tus solicitudes de cotización e información</li>
            <li>Agendar citas y dar seguimiento a servicios</li>
            <li>Contactarte respecto a los servicios solicitados</li>
            <li>Elaborar presupuestos y órdenes de servicio</li>
          </ul>
  
          <h3>Finalidades secundarias (opcionales):</h3>
          <ul>
            <li>Enviarte información sobre promociones y descuentos</li>
            <li>Realizar encuestas de satisfacción</li>
            <li>Mejorar nuestros servicios y experiencia del cliente</li>
            <li>Análisis estadístico de uso del sitio web</li>
          </ul>
        </div>
  
        <div class="legal-section" id="transferencia">
          <h2>4. Transferencia de datos</h2>
          <p>
            <strong>ServiLlantas Celaya</strong> se compromete a no vender, alquilar ni
            compartir tus datos personales con terceros para fines distintos a los
            descritos en este aviso.
          </p>
          <p>
            Tus datos podrán ser compartidos únicamente con proveedores de servicios
            que nos asistan en la operación del negocio (ejemplo: plataformas de
            comunicación) bajo estrictos acuerdos de confidencialidad.
          </p>
        </div>
  
        <div class="legal-section" id="derechos">
          <h2>5. Tus derechos ARCO</h2>
          <p>
            Tienes derecho a <strong>Acceder, Rectificar, Cancelar u Oponerte</strong>
            al tratamiento de tus datos personales (derechos ARCO).
          </p>
          <div class="arco-grid">
            <div class="arco-item">
              <span class="arco-letter">A</span>
              <div>
                <strong>Acceso</strong>
                <p>Conocer qué datos tenemos sobre ti</p>
              </div>
            </div>
            <div class="arco-item">
              <span class="arco-letter">R</span>
              <div>
                <strong>Rectificación</strong>
                <p>Corregir datos inexactos o incompletos</p>
              </div>
            </div>
            <div class="arco-item">
              <span class="arco-letter">C</span>
              <div>
                <strong>Cancelación</strong>
                <p>Solicitar la eliminación de tus datos</p>
              </div>
            </div>
            <div class="arco-item">
              <span class="arco-letter">O</span>
              <div>
                <strong>Oposición</strong>
                <p>Oponerte al uso de tus datos</p>
              </div>
            </div>
          </div>
          <p>
            Para ejercer estos derechos, contáctanos al teléfono
            <a href="tel:+524611203488"><strong>461 120 3488</strong></a> o envíanos
            un mensaje a través de nuestro <a href="/contacto"><strong>formulario de contacto</strong></a>.
          </p>
        </div>
  
        <div class="legal-section" id="cookies">
          <h2>6. Uso de cookies</h2>
          <p>
            Este sitio web utiliza cookies y tecnologías similares para mejorar
            tu experiencia de navegación. Las cookies son pequeños archivos de
            texto que se almacenan en tu dispositivo.
          </p>
  
          <h3>Tipos de cookies que utilizamos:</h3>
          <div class="cookies-table">
            <div class="cookie-row header">
              <span>Tipo</span>
              <span>Propósito</span>
              <span>Duración</span>
            </div>
            <div class="cookie-row">
              <span><strong>Esenciales</strong></span>
              <span>Funcionamiento básico del sitio</span>
              <span>Sesión</span>
            </div>
            <div class="cookie-row">
              <span><strong>Preferencias</strong></span>
              <span>Recordar tu configuración</span>
              <span>1 año</span>
            </div>
            <div class="cookie-row">
              <span><strong>Analíticas</strong></span>
              <span>Estadísticas de uso</span>
              <span>2 años</span>
            </div>
            <div class="cookie-row">
              <span><strong>Marketing</strong></span>
              <span>Contenido personalizado</span>
              <span>6 meses</span>
            </div>
          </div>
          <p>
            Puedes gestionar tus preferencias de cookies en cualquier momento
            desde la configuración de tu navegador.
          </p>
        </div>
  
        <div class="legal-section" id="cambios">
          <h2>7. Cambios al aviso de privacidad</h2>
          <p>
            Nos reservamos el derecho de modificar este aviso de privacidad en
            cualquier momento. Cualquier cambio será publicado en esta misma página
            con la fecha de última actualización.
          </p>
          <p>
            Te recomendamos revisar periódicamente este aviso para estar informado
            sobre cómo protegemos tus datos.
          </p>
        </div>
  
        <div class="legal-section" id="contacto-legal">
          <h2>8. Contacto</h2>
          <p>
            Si tienes dudas o comentarios sobre este aviso de privacidad o el
            tratamiento de tus datos personales, contáctanos:
          </p>
          <div class="legal-card contact-card">
            <p><strong>📞 Teléfono:</strong> <a href="tel:+524611203488">461 120 3488</a></p>
            <p><strong>📱 Teléfono 2:</strong> <a href="tel:+527201987926">720 198 7926</a></p>
            <p><strong>📍 Dirección:</strong> Celaya, Guanajuato, México</p>
            <p><strong>💬 WhatsApp:</strong> <a href="https://wa.me/524611203488" target="_blank" rel="noopener">Enviar mensaje</a></p>
            <a href="/contacto" class="btn btn-accent" style="margin-top: 1rem;">
              ✉️ Formulario de contacto
            </a>
          </div>
        </div>
  
      </div>
    </div>
  </section>
  
  <style>
    /* ====== PAGE HERO ====== */
    .page-hero {
      position: relative;
      padding: 10rem 2rem 4rem;
      text-align: center;
      background-color: var(--color-primary);
    }
    .page-hero-overlay {
      position: absolute; inset: 0;
      background: linear-gradient(180deg, rgba(22,17,47,.95), rgba(22,17,47,.9));
    }
    .page-hero-content { position: relative; z-index: 2; max-width: 700px; margin: 0 auto; }
    .page-hero .hero-subtitle { color: var(--color-accent); font-family: var(--font-heading); font-weight: 700; font-size: .85rem; letter-spacing: 4px; margin-bottom: 1rem; }
    .page-hero h1 { font-size: clamp(2rem,5vw,3rem); margin-bottom: .5rem; }
    .page-hero .hero-desc { color: var(--color-gray); font-size: 1rem; margin-bottom: 1rem; }
    .breadcrumb { display: flex; justify-content: center; gap: .5rem; font-size: .9rem; color: var(--color-gray); }
    .breadcrumb a { color: var(--color-light-3); transition: color .3s; }
    .breadcrumb a:hover { color: var(--color-accent); }
    .breadcrumb .current { color: var(--color-accent); }
  
    /* ====== LEGAL CONTENT ====== */
    .legal-content {
      max-width: 850px;
      margin: 0 auto;
    }
  
    /* Índice */
    .legal-index {
      background: var(--color-dark-3);
      border: 1px solid var(--color-dark-4);
      border-left: 4px solid var(--color-accent);
      border-radius: 0 16px 16px 0;
      padding: 2rem;
      margin-bottom: 3rem;
    }
  
    .legal-index h3 {
      font-size: 1.1rem;
      margin-bottom: 1rem;
    }
  
    .legal-index ol {
      list-style: decimal;
      padding-left: 1.5rem;
      display: flex;
      flex-direction: column;
      gap: 0.5rem;
    }
  
    .legal-index a {
      color: var(--color-secondary-light);
      font-size: 0.95rem;
      transition: color 0.3s;
    }
  
    .legal-index a:hover {
      color: var(--color-accent);
    }
  
    /* Secciones */
    .legal-section {
      margin-bottom: 3rem;
      padding-bottom: 3rem;
      border-bottom: 1px solid var(--color-dark-4);
      scroll-margin-top: 100px;
    }
  
    .legal-section:last-child {
      border-bottom: none;
    }
  
    .legal-section h2 {
      font-size: 1.4rem;
      margin-bottom: 1.2rem;
      color: var(--color-accent);
    }
  
    .legal-section h3 {
      font-size: 1.1rem;
      margin: 1.5rem 0 0.8rem;
      color: var(--color-light);
    }
  
    .legal-section p {
      color: var(--color-light-3);
      font-size: 0.95rem;
      line-height: 1.9;
      margin-bottom: 1rem;
    }
  
    .legal-section a {
      color: var(--color-accent);
      transition: color 0.3s;
    }
  
    .legal-section a:hover {
      color: var(--color-secondary-light);
    }
  
    .legal-section ul {
      margin: 0.5rem 0 1rem;
      padding-left: 1.5rem;
      display: flex;
      flex-direction: column;
      gap: 0.5rem;
    }
  
    .legal-section li {
      list-style: none;
      color: var(--color-light-3);
      font-size: 0.95rem;
      padding-left: 1.5rem;
      position: relative;
    }
  
    .legal-section li::before {
      content: '✓';
      position: absolute;
      left: 0;
      color: var(--color-accent);
      font-weight: 800;
      font-size: 0.8rem;
    }
  
    /* Cards */
    .legal-card {
      background: var(--color-dark-3);
      border: 1px solid var(--color-dark-4);
      border-radius: 12px;
      padding: 1.5rem;
      margin: 1rem 0;
    }
  
    .legal-card p {
      margin-bottom: 0.5rem;
      font-size: 0.9rem;
    }
  
    .contact-card {
      border-color: var(--color-secondary);
    }
  
    /* Legal grid */
    .legal-grid {
      display: grid;
      grid-template-columns: 1fr 1fr;
      gap: 1rem;
      margin: 1rem 0;
    }
  
    .legal-item {
      display: flex;
      gap: 1rem;
      align-items: flex-start;
      background: var(--color-dark-3);
      border: 1px solid var(--color-dark-4);
      border-radius: 12px;
      padding: 1.2rem;
    }
  
    .legal-icon {
      font-size: 1.5rem;
      flex-shrink: 0;
    }
  
    .legal-item strong {
      display: block;
      font-size: 0.9rem;
      margin-bottom: 0.2rem;
    }
  
    .legal-item p {
      color: var(--color-gray);
      font-size: 0.85rem;
      margin-bottom: 0;
    }
  
    .legal-note {
      background: rgba(243, 253, 1, 0.05);
      border: 1px solid rgba(243, 253, 1, 0.2);
      border-radius: 12px;
      padding: 1rem 1.5rem;
      margin-top: 1rem;
      font-size: 0.9rem;
    }
  
    /* ARCO */
    .arco-grid {
      display: grid;
      grid-template-columns: repeat(4, 1fr);
      gap: 1rem;
      margin: 1.5rem 0;
    }
  
    .arco-item {
      text-align: center;
      background: var(--color-dark-3);
      border: 1px solid var(--color-dark-4);
      border-radius: 16px;
      padding: 1.5rem 1rem;
      transition: all 0.3s;
    }
  
    .arco-item:hover {
      border-color: var(--color-secondary);
      transform: translateY(-4px);
    }
  
    .arco-letter {
      display: flex;
      align-items: center;
      justify-content: center;
      width: 50px;
      height: 50px;
      background: var(--color-secondary);
      border-radius: 50%;
      font-family: var(--font-heading);
      font-weight: 900;
      font-size: 1.5rem;
      margin: 0 auto 1rem;
    }
  
    .arco-item strong {
      display: block;
      margin-bottom: 0.3rem;
      font-size: 0.95rem;
    }
  
    .arco-item p {
      color: var(--color-gray);
      font-size: 0.8rem;
      margin-bottom: 0;
    }
  
    /* Cookies table */
    .cookies-table {
      margin: 1rem 0;
      border-radius: 12px;
      overflow: hidden;
      border: 1px solid var(--color-dark-4);
    }
  
    .cookie-row {
      display: grid;
      grid-template-columns: 1fr 2fr 1fr;
      gap: 1px;
    }
  
    .cookie-row span {
      padding: 0.8rem 1rem;
      background: var(--color-dark-3);
      font-size: 0.9rem;
      color: var(--color-light-3);
    }
  
    .cookie-row.header span {
      background: var(--color-dark-4);
      color: var(--color-light);
      font-weight: 700;
      font-size: 0.85rem;
      text-transform: uppercase;
      letter-spacing: 1px;
    }
  
    /* Responsive */
    @media (max-width: 768px) {
      .page-hero { padding: 8rem 1.5rem 3rem; }
      .legal-grid { grid-template-columns: 1fr; }
      .arco-grid { grid-template-columns: repeat(2, 1fr); }
      .cookie-row { grid-template-columns: 1fr; }
      .cookie-row.header { display: none; }
      .cookie-row span { padding: .6rem 1rem; }
    }
  
    @media (max-width: 480px) {
      .arco-grid { grid-template-columns: 1fr; }
    }
  </style>