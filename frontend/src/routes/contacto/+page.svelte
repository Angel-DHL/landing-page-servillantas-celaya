<!-- svelte-ignore non_reactive_update -->
<!-- frontend/src/routes/contacto/+page.svelte -->
<script>
// @ts-nocheck

    import { inview } from '$lib/actions/inview.js';
  
    // ===== ESTADO DEL FORMULARIO =====
    let form = {
      nombre: '',
      email: '',
      telefono: '',
      servicio: '',
      vehiculo: '',
      mensaje: ''
    };
  
    let errors = {};
    let loading = false;
    let submitted = false;
    let submitError = '';
  
    // ===== SERVICIOS PARA SELECT =====
    const serviciosList = [
      'Venta de Llantas',
      'Alineación Computarizada',
      'Balanceo Dinámico',
      'Sistema de Frenos',
      'Suspensión',
      'Cambio de Aceite',
      'Diagnóstico General',
      'Otro'
    ];
  
    // ===== VALIDACIÓN =====
    function validate() {
      errors = {};
  
      if (!form.nombre.trim()) {
        errors.nombre = 'El nombre es obligatorio';
      } else if (form.nombre.trim().length < 3) {
        errors.nombre = 'Mínimo 3 caracteres';
      }
  
      if (!form.telefono.trim()) {
        errors.telefono = 'El teléfono es obligatorio';
      } else if (!/^[\d\s\-\+\(\)]{7,15}$/.test(form.telefono.trim())) {
        errors.telefono = 'Ingresa un teléfono válido';
      }
  
      if (form.email.trim() && !/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(form.email.trim())) {
        errors.email = 'Ingresa un correo válido';
      }
  
      if (!form.servicio) {
        errors.servicio = 'Selecciona un servicio';
      }
  
      return Object.keys(errors).length === 0;
    }
  
    // ===== ENVIAR FORMULARIO =====
    async function handleSubmit() {
      submitError = '';
  
      if (!validate()) return;
  
      loading = true;
  
      try {
        const res = await fetch('http://localhost:3000/api/contacto', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({
            nombre: form.nombre.trim(),
            email: form.email.trim(),
            telefono: form.telefono.trim(),
            servicio: form.servicio,
            mensaje: `Vehículo: ${form.vehiculo || 'No especificado'}\n${form.mensaje}`
          })
        });
  
        const data = await res.json();
  
        if (data.success) {
          submitted = true;
          // Reset
          form = { nombre: '', email: '', telefono: '', servicio: '', vehiculo: '', mensaje: '' };
        } else {
          submitError = data.error || 'Error al enviar. Intenta de nuevo.';
        }
      } catch (err) {
        submitError = 'No se pudo conectar con el servidor. Intenta llamarnos directamente.';
      } finally {
        loading = false;
      }
    }
  
    // ===== ENVIAR OTRO MENSAJE =====
    function resetForm() {
      submitted = false;
      submitError = '';
      errors = {};
    }
  
    // ===== DATOS DE CONTACTO =====
    const contactInfo = [
      {
        icon: '📞',
        titulo: 'Teléfono 1',
        dato: '461 120 3488',
        link: 'tel:+524611203488',
        action: 'Llamar ahora'
      },
      {
        icon: '📱',
        titulo: 'Teléfono 2',
        dato: '720 198 7926',
        link: 'tel:+527201987926',
        action: 'Llamar ahora'
      },
      {
        icon: '💬',
        titulo: 'WhatsApp',
        dato: '461 120 3488',
        link: 'https://wa.me/524611203488?text=Hola%2C%20me%20interesa%20información%20sobre%20sus%20servicios',
        action: 'Enviar mensaje'
      },
      {
        icon: '📍',
        titulo: 'Dirección',
        dato: 'Celaya, Guanajuato, México',
        link: 'https://maps.google.com/?q=Celaya+Guanajuato+Mexico',
        action: 'Ver en mapa'
      }
    ];
  
    // ===== HORARIOS =====
    const horarios = [
      { dia: 'Lunes a Viernes', horario: '8:00 AM - 6:00 PM' },
      { dia: 'Sábado', horario: '8:00 AM - 2:00 PM' },
      { dia: 'Domingo', horario: 'Cerrado' }
    ];
  
    // ===== FAQ =====
    let faqs = [
      {
        pregunta: '¿Necesito hacer cita?',
        respuesta: 'No es necesario para la mayoría de servicios. Sin embargo, si deseas un horario específico, te recomendamos agendar para garantizar atención inmediata.',
        open: false
      },
      {
        pregunta: '¿Cuánto tiempo toma una alineación?',
        respuesta: 'Una alineación computarizada toma aproximadamente 30-45 minutos dependiendo del vehículo. El balanceo adicional toma 15-20 minutos más.',
        open: false
      },
      {
        pregunta: '¿Tienen garantía en sus servicios?',
        respuesta: 'Sí, todos nuestros servicios cuentan con garantía por escrito. Las llantas tienen garantía de fábrica y nuestros trabajos de mano de obra tienen garantía de 6 meses.',
        open: false
      },
      {
        pregunta: '¿Aceptan tarjeta de crédito/débito?',
        respuesta: 'Sí, aceptamos todas las formas de pago: efectivo, tarjeta de crédito, débito y transferencia bancaria.',
        open: false
      },
      {
        pregunta: '¿Trabajan con todas las marcas de vehículos?',
        respuesta: 'Sí, atendemos todas las marcas y modelos de vehículos: sedán, camioneta, SUV, deportivos y vehículos de carga ligera.',
        open: false
      }
    ];
  
    function toggleFaq(index) {
      faqs = faqs.map((faq, i) => (i === index ? { ...faq, open: !faq.open } : faq));
    }
  </script>
  
  <svelte:head>
    <title>Contacto | ServiLlantas Celaya - Cotiza y Agenda tu Cita</title>
    <meta name="description" content="Contáctanos para cotizar llantas, alineación, balanceo y más. Llámanos al 461 120 3488 o envíanos un mensaje." />
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
      <p class="hero-subtitle">ESTAMOS PARA AYUDARTE</p>
      <h1><span class="text-accent">Contáctanos</span></h1>
      <p class="hero-desc">
        Cotiza sin compromiso, agenda tu cita o resuelve tus dudas. Estamos listos para atenderte.
      </p>
      <nav class="breadcrumb">
        <a href="/">Inicio</a>
        <span>/</span>
        <span class="current">Contacto</span>
      </nav>
    </div>
  </section>
  
  <!-- ========== TARJETAS DE CONTACTO RÁPIDO ========== -->
  <section class="quick-contact">
    <div class="container">
      <div class="quick-grid">
        {#each contactInfo as info, i}
          <a
            href={info.link}
            target={info.link.startsWith('http') ? '_blank' : '_self'}
            rel="noopener noreferrer"
            class="quick-card fade-up delay-{i + 1}"
            use:inview
          >
            <span class="quick-icon">{info.icon}</span>
            <span class="quick-title">{info.titulo}</span>
            <span class="quick-dato">{info.dato}</span>
            <span class="quick-action">{info.action} →</span>
          </a>
        {/each}
      </div>
    </div>
  </section>
  
  <!-- ========== FORMULARIO + INFO ========== -->
  <section class="section section-dark">
    <div class="container">
      <div class="contact-grid">
  
        <!-- COLUMNA IZQUIERDA: FORMULARIO -->
        <div class="form-column">
          <p class="section-subtitle fade-left" use:inview>Envíanos un mensaje</p>
          <h2 class="section-title fade-left delay-1" use:inview>
            Solicita tu <span class="text-accent">cotización</span>
          </h2>
          <p class="form-desc fade-left delay-2" use:inview>
            Completa el formulario y nos pondremos en contacto contigo en menos de 24 horas.
          </p>
  
          <!-- FORMULARIO -->
          {#if submitted}
            <!-- Mensaje de éxito -->
            <div class="success-card scale-in" use:inview>
              <div class="success-icon">✅</div>
              <h3>¡Mensaje enviado!</h3>
              <p>
                Gracias por contactarnos. Un asesor se comunicará contigo
                en las próximas horas.
              </p>
              <p class="success-hint">
                Si necesitas atención inmediata, llámanos:
              </p>
              <div class="success-phones">
                <a href="tel:+524611203488" class="btn btn-accent">
                  📞 461 120 3488
                </a>
                <a href="tel:+527201987926" class="btn btn-outline-accent">
                  📱 720 198 7926
                </a>
              </div>
              <button class="btn-text" on:click={resetForm}>
                ← Enviar otro mensaje
              </button>
            </div>
          {:else}
            <form
              class="contact-form fade-left delay-3"
              use:inview
              on:submit|preventDefault={handleSubmit}
              novalidate
            >
              <!-- Nombre -->
              <div class="form-group" class:error={errors.nombre}>
                <label for="nombre">
                  Nombre completo <span class="required">*</span>
                </label>
                <input
                  id="nombre"
                  type="text"
                  bind:value={form.nombre}
                  placeholder="Ej: Juan Pérez"
                  class:input-error={errors.nombre}
                />
                {#if errors.nombre}
                  <span class="error-msg">{errors.nombre}</span>
                {/if}
              </div>
  
              <!-- Teléfono + Email -->
              <div class="form-row">
                <div class="form-group" class:error={errors.telefono}>
                  <label for="telefono">
                    Teléfono <span class="required">*</span>
                  </label>
                  <input
                    id="telefono"
                    type="tel"
                    bind:value={form.telefono}
                    placeholder="Ej: 461 123 4567"
                    class:input-error={errors.telefono}
                  />
                  {#if errors.telefono}
                    <span class="error-msg">{errors.telefono}</span>
                  {/if}
                </div>
  
                <div class="form-group" class:error={errors.email}>
                  <label for="email">Correo electrónico</label>
                  <input
                    id="email"
                    type="email"
                    bind:value={form.email}
                    placeholder="correo@ejemplo.com"
                    class:input-error={errors.email}
                  />
                  {#if errors.email}
                    <span class="error-msg">{errors.email}</span>
                  {/if}
                </div>
              </div>
  
              <!-- Servicio + Vehículo -->
              <div class="form-row">
                <div class="form-group" class:error={errors.servicio}>
                  <label for="servicio">
                    Servicio de interés <span class="required">*</span>
                  </label>
                  <select
                    id="servicio"
                    bind:value={form.servicio}
                    class:input-error={errors.servicio}
                  >
                    <option value="" disabled>Selecciona un servicio</option>
                    {#each serviciosList as s}
                      <option value={s}>{s}</option>
                    {/each}
                  </select>
                  {#if errors.servicio}
                    <span class="error-msg">{errors.servicio}</span>
                  {/if}
                </div>
  
                <div class="form-group">
                  <label for="vehiculo">Vehículo (Marca / Modelo / Año)</label>
                  <input
                    id="vehiculo"
                    type="text"
                    bind:value={form.vehiculo}
                    placeholder="Ej: Honda Civic 2020"
                  />
                </div>
              </div>
  
              <!-- Mensaje -->
              <div class="form-group">
                <label for="mensaje">Mensaje o detalles adicionales</label>
                <textarea
                  id="mensaje"
                  bind:value={form.mensaje}
                  placeholder="Cuéntanos más sobre lo que necesitas..."
                  rows="4"
                ></textarea>
              </div>
  
              <!-- Error general -->
              {#if submitError}
                <div class="submit-error">
                  <span>⚠️</span> {submitError}
                </div>
              {/if}
  
              <!-- Botón -->
              <button
                type="submit"
                class="btn btn-accent btn-submit"
                disabled={loading}
              >
                {#if loading}
                  <span class="spinner"></span>
                  Enviando...
                {:else}
                  ✉️ Enviar mensaje
                {/if}
              </button>
  
              <p class="form-privacy">
                🔒 Tu información está segura. No compartimos tus datos con terceros.
              </p>
            </form>
          {/if}
        </div>
  
        <!-- COLUMNA DERECHA: INFO -->
        <div class="info-column">
  
          <!-- Card horarios -->
          <div class="info-card fade-right" use:inview>
            <div class="info-card-header">
              <span class="info-card-icon">🕐</span>
              <h3>Horario de atención</h3>
            </div>
            <div class="horarios-list">
              {#each horarios as h}
                <div class="horario-row">
                  <span class="horario-dia">{h.dia}</span>
                  <span
                    class="horario-hora"
                    class:cerrado={h.horario === 'Cerrado'}
                  >
                    {h.horario}
                  </span>
                </div>
              {/each}
            </div>
          </div>
  
          <!-- Card teléfonos -->
          <div class="info-card fade-right delay-1" use:inview>
            <div class="info-card-header">
              <span class="info-card-icon">📞</span>
              <h3>Llámanos directamente</h3>
            </div>
            <p class="info-card-desc">
              Atención inmediata de lunes a sábado en horario laboral.
            </p>
            <div class="phone-buttons">
              <a href="tel:+524611203488" class="phone-btn">
                <span class="phone-label">Línea 1</span>
                <span class="phone-number">461 120 3488</span>
              </a>
              <a href="tel:+527201987926" class="phone-btn">
                <span class="phone-label">Línea 2</span>
                <span class="phone-number">720 198 7926</span>
              </a>
            </div>
          </div>
  
          <!-- Card WhatsApp -->
          <div class="info-card whatsapp-card fade-right delay-2" use:inview>
            <div class="info-card-header">
              <span class="info-card-icon">💬</span>
              <h3>WhatsApp</h3>
            </div>
            <p class="info-card-desc">
              ¿Prefieres escribirnos? Envíanos un mensaje por WhatsApp y te respondemos al instante.
            </p>
            <a
              href="https://wa.me/524611203488?text=Hola%2C%20me%20interesa%20información%20sobre%20sus%20servicios"
              target="_blank"
              rel="noopener noreferrer"
              class="btn btn-whatsapp"
            >
              💬 Chatear por WhatsApp
            </a>
          </div>
  
          <!-- Card redes sociales -->
          <div class="info-card fade-right delay-3" use:inview>
            <div class="info-card-header">
              <span class="info-card-icon">🌐</span>
              <h3>Síguenos</h3>
            </div>
            <div class="social-buttons">
              <a href="#" class="social-btn facebook">Facebook</a>
              <a href="#" class="social-btn instagram">Instagram</a>
              <a href="#" class="social-btn tiktok">TikTok</a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
  
  <!-- ========== MAPA ========== -->
  <section class="map-section">
    <div class="container text-center">
      <p class="section-subtitle fade-up" use:inview>Ubícanos</p>
      <h2 class="section-title fade-up" use:inview>
        Visítanos en <span class="text-accent">Celaya</span>
      </h2>
      <p class="section-desc fade-up" use:inview>
        Estamos ubicados en una zona de fácil acceso. ¡Te esperamos!
      </p>
    </div>
  
    <div class="map-wrapper fade-up delay-1" use:inview>
      <div class="map-container">
        <!-- Reemplaza las coordenadas con la ubicación real -->
        <iframe
          src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d59777.77038382395!2d-100.84032795!3d20.5234884!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x842cbf24b080de93%3A0x34ab809b896e2de7!2sCelaya%2C%20Gto.!5e0!3m2!1ses-419!2smx!4v1703000000000!5m2!1ses-419!2smx"
          width="100%"
          height="100%"
          style="border:0;"
          allowfullscreen=""
          loading="lazy"
          referrerpolicy="no-referrer-when-downgrade"
          title="Ubicación ServiLlantas Celaya"
        ></iframe>
      </div>
  
      <!-- Tarjeta sobre el mapa -->
      <div class="map-info-card">
        <h3>ServiLlantas Celaya</h3>
        <div class="map-info-item">
          <span>📍</span>
          <span>Celaya, Guanajuato, México</span>
        </div>
        <div class="map-info-item">
          <span>📞</span>
          <a href="tel:+524611203488">461 120 3488</a>
        </div>
        <div class="map-info-item">
          <span>🕐</span>
          <span>Lun - Vie: 8:00 - 18:00</span>
        </div>
        <a
          href="https://maps.google.com/?q=Celaya+Guanajuato+Mexico"
          target="_blank"
          rel="noopener noreferrer"
          class="btn btn-accent btn-sm"
        >
          📍 Cómo llegar
        </a>
      </div>
    </div>
  </section>
  
  <!-- ========== FAQ ========== -->
  <section class="section section-dark">
    <div class="container">
      <div class="text-center">
        <p class="section-subtitle fade-up" use:inview>Preguntas frecuentes</p>
        <h2 class="section-title fade-up" use:inview>
          ¿Tienes <span class="text-accent">dudas</span>?
        </h2>
        <p class="section-desc fade-up" use:inview>
          Respondemos las preguntas más comunes de nuestros clientes.
        </p>
      </div>
  
      <div class="faq-list">
        {#each faqs as faq, i}
          <div class="faq-item fade-up delay-{Math.min(i + 1, 5)}" use:inview>
            <button
              class="faq-question"
              class:open={faq.open}
              on:click={() => toggleFaq(i)}
            >
              <span>{faq.pregunta}</span>
              <span class="faq-arrow" class:rotated={faq.open}>▼</span>
            </button>
            {#if faq.open}
              <div class="faq-answer">
                <p>{faq.respuesta}</p>
              </div>
            {/if}
          </div>
        {/each}
      </div>
    </div>
  </section>
  
  <!-- ========== CTA FINAL ========== -->
  <section class="cta-final">
    <div class="cta-final-overlay"></div>
    <div class="container text-center" style="position:relative; z-index:2;">
      <h2 class="fade-up" use:inview>
        ¿Prefieres que <span class="text-accent">te llamemos</span>?
      </h2>
      <p class="cta-final-desc fade-up delay-1" use:inview>
        Déjanos tu número en el formulario y un asesor se comunicará contigo.
      </p>
      <div class="cta-final-buttons fade-up delay-2" use:inview>
        <a href="#top" class="btn btn-accent btn-lg">⬆️ Ir al formulario</a>
        <a
          href="https://wa.me/524611203488?text=Hola%2C%20quiero%20agendar%20una%20cita"
          target="_blank"
          rel="noopener noreferrer"
          class="btn btn-primary btn-lg"
        >
          💬 WhatsApp directo
        </a>
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
  
    .breadcrumb a { color: var(--color-light-3); transition: color 0.3s; }
    .breadcrumb a:hover { color: var(--color-accent); }
    .breadcrumb .current { color: var(--color-accent); }
  
    /* ====== QUICK CONTACT ====== */
    .quick-contact {
      background: var(--color-dark-2);
      padding: 0;
      margin-top: -3rem;
      position: relative;
      z-index: 3;
    }
  
    .quick-grid {
      display: grid;
      grid-template-columns: repeat(4, 1fr);
      gap: 1.5rem;
    }
  
    .quick-card {
      background: var(--color-dark-3);
      border: 1px solid var(--color-dark-4);
      border-radius: 16px;
      padding: 2rem 1.5rem;
      text-align: center;
      transition: all 0.4s ease;
      display: flex;
      flex-direction: column;
      align-items: center;
      gap: 0.5rem;
      position: relative;
      overflow: hidden;
    }
  
    .quick-card::before {
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
  
    .quick-card:hover {
      transform: translateY(-8px);
      border-color: var(--color-secondary);
      box-shadow: 0 15px 35px rgba(0, 75, 188, 0.2);
    }
  
    .quick-card:hover::before { transform: scaleX(1); }
  
    .quick-icon { font-size: 2rem; }
    .quick-title { font-size: 0.8rem; color: var(--color-gray); text-transform: uppercase; letter-spacing: 1px; font-weight: 600; }
    .quick-dato { font-family: var(--font-heading); font-weight: 700; font-size: 1.05rem; }
    .quick-action { color: var(--color-accent); font-size: 0.85rem; font-weight: 600; margin-top: 0.3rem; }
  
    /* ====== CONTACT GRID ====== */
    .contact-grid {
      display: grid;
      grid-template-columns: 1.2fr 0.8fr;
      gap: 4rem;
      align-items: flex-start;
    }
  
    /* ====== FORM ====== */
    .form-desc {
      color: var(--color-gray);
      margin-bottom: 2rem;
      font-size: 0.95rem;
    }
  
    .contact-form {
      display: flex;
      flex-direction: column;
      gap: 1.5rem;
    }
  
    .form-row {
      display: grid;
      grid-template-columns: 1fr 1fr;
      gap: 1.5rem;
    }
  
    .form-group {
      display: flex;
      flex-direction: column;
      gap: 0.4rem;
    }
  
    .form-group label {
      font-size: 0.9rem;
      font-weight: 600;
      color: var(--color-light-3);
    }
  
    .required {
      color: var(--color-accent);
    }
  
    .form-group input,
    .form-group select,
    .form-group textarea {
      width: 100%;
      padding: 0.9rem 1.2rem;
      background: var(--color-dark-3);
      border: 2px solid var(--color-dark-4);
      border-radius: 12px;
      color: var(--color-light);
      font-family: var(--font-body);
      font-size: 0.95rem;
      transition: all 0.3s ease;
      outline: none;
    }
  
    .form-group select {
      cursor: pointer;
      appearance: none;
      background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='12' height='12' fill='%238888a0' viewBox='0 0 16 16'%3E%3Cpath d='M8 11L3 6h10z'/%3E%3C/svg%3E");
      background-repeat: no-repeat;
      background-position: right 1rem center;
      padding-right: 2.5rem;
    }
  
    .form-group input::placeholder,
    .form-group textarea::placeholder {
      color: var(--color-gray);
      opacity: 0.6;
    }
  
    .form-group input:focus,
    .form-group select:focus,
    .form-group textarea:focus {
      border-color: var(--color-secondary);
      box-shadow: 0 0 0 3px rgba(0, 75, 188, 0.15);
    }
  
    .input-error {
      border-color: #ff4444 !important;
    }
  
    .error-msg {
      color: #ff6b6b;
      font-size: 0.8rem;
      margin-top: 0.2rem;
    }
  
    .form-group textarea {
      resize: vertical;
      min-height: 100px;
    }
  
    .submit-error {
      background: rgba(255, 68, 68, 0.1);
      border: 1px solid rgba(255, 68, 68, 0.3);
      border-radius: 12px;
      padding: 1rem;
      color: #ff6b6b;
      font-size: 0.9rem;
      display: flex;
      align-items: center;
      gap: 0.5rem;
    }
  
    .btn-submit {
      width: 100%;
      justify-content: center;
      padding: 1.1rem;
      font-size: 1.05rem;
    }
  
    .btn-submit:disabled {
      opacity: 0.7;
      cursor: not-allowed;
    }
  
    .spinner {
      width: 20px;
      height: 20px;
      border: 3px solid rgba(22, 17, 47, 0.3);
      border-top-color: var(--color-primary);
      border-radius: 50%;
      animation: spin 0.8s linear infinite;
    }
  
    @keyframes spin {
      to { transform: rotate(360deg); }
    }
  
    .form-privacy {
      color: var(--color-gray);
      font-size: 0.8rem;
      text-align: center;
      margin-top: -0.5rem;
    }
  
    /* ====== SUCCESS CARD ====== */
    .success-card {
      background: var(--color-dark-3);
      border: 2px solid var(--color-secondary);
      border-radius: 20px;
      padding: 3rem 2rem;
      text-align: center;
    }
  
    .success-icon { font-size: 3.5rem; margin-bottom: 1rem; }
  
    .success-card h3 {
      font-size: 1.5rem;
      margin-bottom: 0.8rem;
      color: var(--color-accent);
    }
  
    .success-card > p {
      color: var(--color-light-3);
      margin-bottom: 0.5rem;
    }
  
    .success-hint {
      color: var(--color-gray);
      font-size: 0.9rem;
      margin-bottom: 1rem;
    }
  
    .success-phones {
      display: flex;
      gap: 1rem;
      justify-content: center;
      flex-wrap: wrap;
      margin-bottom: 1.5rem;
    }
  
    .btn-text {
      background: none;
      border: none;
      color: var(--color-accent);
      font-weight: 600;
      cursor: pointer;
      font-size: 0.9rem;
      transition: color 0.3s;
    }
  
    .btn-text:hover { color: var(--color-light); }
  
    /* ====== INFO CARDS ====== */
    .info-column {
      display: flex;
      flex-direction: column;
      gap: 1.5rem;
    }
  
    .info-card {
      background: var(--color-dark-3);
      border: 1px solid var(--color-dark-4);
      border-radius: 16px;
      padding: 1.8rem;
      transition: all 0.3s ease;
    }
  
    .info-card:hover {
      border-color: var(--color-secondary);
    }
  
    .info-card-header {
      display: flex;
      align-items: center;
      gap: 0.8rem;
      margin-bottom: 1rem;
    }
  
    .info-card-icon { font-size: 1.5rem; }
  
    .info-card-header h3 {
      font-size: 1.05rem;
      font-weight: 700;
    }
  
    .info-card-desc {
      color: var(--color-gray);
      font-size: 0.9rem;
      margin-bottom: 1rem;
      line-height: 1.6;
    }
  
    /* Horarios */
    .horarios-list {
      display: flex;
      flex-direction: column;
      gap: 0.8rem;
    }
  
    .horario-row {
      display: flex;
      justify-content: space-between;
      align-items: center;
      padding-bottom: 0.8rem;
      border-bottom: 1px solid var(--color-dark-4);
    }
  
    .horario-row:last-child { border-bottom: none; padding-bottom: 0; }
  
    .horario-dia {
      font-weight: 600;
      font-size: 0.9rem;
    }
  
    .horario-hora {
      color: var(--color-accent);
      font-weight: 600;
      font-size: 0.9rem;
    }
  
    .horario-hora.cerrado {
      color: #ff6b6b;
    }
  
    /* Phone buttons */
    .phone-buttons {
      display: flex;
      flex-direction: column;
      gap: 0.8rem;
    }
  
    .phone-btn {
      display: flex;
      justify-content: space-between;
      align-items: center;
      padding: 1rem 1.2rem;
      background: rgba(0, 75, 188, 0.1);
      border: 1px solid rgba(0, 75, 188, 0.3);
      border-radius: 12px;
      transition: all 0.3s ease;
    }
  
    .phone-btn:hover {
      background: rgba(0, 75, 188, 0.2);
      border-color: var(--color-secondary);
      transform: translateX(5px);
    }
  
    .phone-label {
      color: var(--color-gray);
      font-size: 0.8rem;
      text-transform: uppercase;
      letter-spacing: 1px;
    }
  
    .phone-number {
      font-family: var(--font-heading);
      font-weight: 700;
      font-size: 1.1rem;
      color: var(--color-accent);
    }
  
    /* WhatsApp */
    .whatsapp-card { border-color: rgba(37, 211, 102, 0.3); }
    .whatsapp-card:hover { border-color: #25d366; }
  
    .btn-whatsapp {
      display: inline-flex;
      align-items: center;
      gap: 0.5rem;
      padding: 0.9rem 2rem;
      background: #25d366;
      color: white;
      border: none;
      border-radius: 50px;
      font-family: var(--font-heading);
      font-weight: 700;
      font-size: 0.95rem;
      cursor: pointer;
      transition: all 0.3s ease;
      width: 100%;
      justify-content: center;
    }
  
    .btn-whatsapp:hover {
      background: #1ea952;
      transform: translateY(-2px);
      box-shadow: 0 8px 25px rgba(37, 211, 102, 0.3);
    }
  
    /* Social */
    .social-buttons {
      display: flex;
      gap: 0.8rem;
    }
  
    .social-btn {
      flex: 1;
      padding: 0.8rem;
      text-align: center;
      border-radius: 10px;
      font-weight: 600;
      font-size: 0.85rem;
      transition: all 0.3s ease;
    }
  
    .social-btn.facebook {
      background: rgba(59, 89, 152, 0.15);
      border: 1px solid rgba(59, 89, 152, 0.3);
      color: #7e9bd5;
    }
    .social-btn.facebook:hover { background: #3b5998; color: white; }
  
    .social-btn.instagram {
      background: rgba(225, 48, 108, 0.15);
      border: 1px solid rgba(225, 48, 108, 0.3);
      color: #e1306c;
    }
    .social-btn.instagram:hover { background: #e1306c; color: white; }
  
    .social-btn.tiktok {
      background: rgba(255, 255, 255, 0.05);
      border: 1px solid rgba(255, 255, 255, 0.15);
      color: var(--color-light-3);
    }
    .social-btn.tiktok:hover { background: var(--color-light); color: var(--color-dark); }
  
    /* ====== MAP ====== */
    .map-section {
      background: var(--color-dark);
      padding: 6rem 0 0;
    }
  
    .map-wrapper {
      position: relative;
      margin-top: 3rem;
    }
  
    .map-container {
      width: 100%;
      height: 450px;
      border-top: 3px solid var(--color-secondary);
    }
  
    .map-container iframe {
      width: 100%;
      height: 100%;
      filter: brightness(0.8) contrast(1.1) saturate(0.3);
      transition: filter 0.3s ease;
    }
  
    .map-container:hover iframe {
      filter: brightness(0.9) contrast(1.1) saturate(0.6);
    }
  
    .map-info-card {
      position: absolute;
      top: 2rem;
      left: 2rem;
      background: var(--color-dark-2);
      border: 1px solid var(--color-dark-4);
      border-radius: 16px;
      padding: 2rem;
      max-width: 320px;
      box-shadow: 0 10px 30px rgba(0, 0, 0, 0.5);
      z-index: 2;
    }
  
    .map-info-card h3 {
      font-size: 1.2rem;
      margin-bottom: 1rem;
      color: var(--color-accent);
    }
  
    .map-info-item {
      display: flex;
      align-items: center;
      gap: 0.6rem;
      margin-bottom: 0.8rem;
      font-size: 0.9rem;
      color: var(--color-light-3);
    }
  
    .map-info-item a {
      color: var(--color-light-3);
      transition: color 0.3s;
    }
  
    .map-info-item a:hover { color: var(--color-accent); }
  
    .btn-sm {
      padding: 0.6rem 1.2rem;
      font-size: 0.85rem;
      margin-top: 0.5rem;
    }
  
    /* ====== FAQ ====== */
    .faq-list {
      max-width: 800px;
      margin: 2rem auto 0;
      display: flex;
      flex-direction: column;
      gap: 1rem;
    }
  
    .faq-item {
      background: var(--color-dark-3);
      border: 1px solid var(--color-dark-4);
      border-radius: 14px;
      overflow: hidden;
      transition: border-color 0.3s ease;
    }
  
    .faq-item:hover {
      border-color: var(--color-secondary);
    }
  
    .faq-question {
      width: 100%;
      display: flex;
      justify-content: space-between;
      align-items: center;
      padding: 1.3rem 1.5rem;
      background: none;
      border: none;
      color: var(--color-light);
      font-family: var(--font-heading);
      font-weight: 700;
      font-size: 1rem;
      cursor: pointer;
      text-align: left;
      transition: color 0.3s ease;
    }
  
    .faq-question:hover { color: var(--color-accent); }
  
    .faq-question.open { color: var(--color-accent); }
  
    .faq-arrow {
      font-size: 0.7rem;
      transition: transform 0.3s ease;
      color: var(--color-gray);
    }
  
    .faq-arrow.rotated {
      transform: rotate(180deg);
      color: var(--color-accent);
    }
  
    .faq-answer {
      padding: 0 1.5rem 1.3rem;
      animation: slideDown 0.3s ease;
    }
  
    .faq-answer p {
      color: var(--color-gray);
      font-size: 0.95rem;
      line-height: 1.8;
    }
  
    @keyframes slideDown {
      from { opacity: 0; transform: translateY(-10px); }
      to { opacity: 1; transform: translateY(0); }
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
      .quick-grid {
        grid-template-columns: repeat(2, 1fr);
      }
  
      .contact-grid {
        grid-template-columns: 1fr;
        gap: 3rem;
      }
    }
  
    @media (max-width: 768px) {
      .quick-grid {
        grid-template-columns: 1fr 1fr;
      }
  
      .form-row {
        grid-template-columns: 1fr;
      }
  
      .page-hero {
        padding: 8rem 1.5rem 3rem;
      }
  
      .quick-contact {
        margin-top: -2rem;
      }
  
      .map-info-card {
        position: relative;
        top: auto;
        left: auto;
        max-width: 100%;
        margin: 0 1.5rem;
        transform: translateY(-2rem);
      }
  
      .map-container {
        height: 300px;
      }
  
      .success-phones {
        flex-direction: column;
      }
  
      .social-buttons {
        flex-wrap: wrap;
      }
  
      .social-btn {
        flex: 1 1 calc(50% - 0.4rem);
      }
    }
  
    @media (max-width: 480px) {
      .quick-grid {
        grid-template-columns: 1fr;
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