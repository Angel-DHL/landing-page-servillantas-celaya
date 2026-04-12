<!-- frontend/src/routes/blog/[slug]/+page.svelte -->
<script>
// @ts-nocheck

    import { onMount } from 'svelte';
    import { page } from '$app/stores';
    import { inview } from '$lib/actions/inview.js';
  
    // @ts-ignore
    /**
	 * @type {{ titulo: any; resumen: any; categoria: any; autor: string; autor_rol: any; fecha: any; duracion: any; imagen: any; contenido: any; tags: any; } | null}
	 */
    let post = null;
    // @ts-ignore
    /**
	 * @type {string | any[] | null | undefined}
	 */
    let relatedPosts = [];
    let loading = true;
    let error = false;
  
    let slug = $derived(() => $page.params.slug);
  
    onMount(async () => {
      await loadPost();
    });
  
    async function loadPost() {
      loading = true;
      error = false;
  
      try {
        // Cargar artículo
        const res = await fetch(`http://localhost:3000/api/blog/${slug}`);
  
        if (!res.ok) {
          error = true;
          return;
        }
  
        post = await res.json();
  
        // Cargar artículos relacionados
        const relRes = await fetch(`http://localhost:3000/api/blog/${slug}/related`);
        if (relRes.ok) {
          relatedPosts = await relRes.json();
        }
      } catch (err) {
        error = true;
        console.error('Error:', err);
      } finally {
        loading = false;
      }
    }
  
    // @ts-ignore
    function formatDate(dateStr) {
      const date = new Date(dateStr + 'T00:00:00');
      return date.toLocaleDateString('es-MX', {
        year: 'numeric',
        month: 'long',
        day: 'numeric'
      });
    }
  
    // @ts-ignore
    function getCategoryColor(cat) {
      const colors = {
        'Mantenimiento': '#004bbc',
        'Seguridad': '#e63946',
        'Tips': '#25d366',
        'Mecánica': '#ff8c00'
      };
      // @ts-ignore
      return colors[cat] || '#004bbc';
    }
  
    // Convertir markdown básico a HTML
    // @ts-ignore
    function renderContent(text) {
      if (!text) return '';
      return text
        // Headers
        .replace(/^### (.+)$/gm, '<h3>$1</h3>')
        .replace(/^## (.+)$/gm, '<h2>$1</h2>')
        // Bold
        .replace(/\*\*(.+?)\*\*/g, '<strong>$1</strong>')
        // Lists
        .replace(/^- (.+)$/gm, '<li>$1</li>')
        // Tables (básico)
        // @ts-ignore
        .replace(/\|(.+)\|/g, (match) => {
          // @ts-ignore
          const cells = match.split('|').filter(c => c.trim());
          // @ts-ignore
          if (cells.some(c => c.includes('---'))) return '';
          // @ts-ignore
          const cellHtml = cells.map(c => `<td>${c.trim()}</td>`).join('');
          return `<tr>${cellHtml}</tr>`;
        })
        // Paragraphs
        .replace(/\n\n/g, '</p><p>')
        // Line breaks
        .replace(/\n/g, '<br/>');
    }
  
    function shareOnWhatsApp() {
      const url = encodeURIComponent(window.location.href);
      // @ts-ignore
      const text = encodeURIComponent(`Mira este artículo: ${post.titulo}`);
      window.open(`https://wa.me/?text=${text}%20${url}`, '_blank');
    }
  
    function copyLink() {
      navigator.clipboard.writeText(window.location.href);
      alert('¡Link copiado!');
    }
  </script>
  
  <svelte:head>
    {#if post}
      <title>{post.titulo} | Blog ServiLlantas Celaya</title>
      <meta name="description" content={post.resumen} />
    {:else}
      <title>Blog | ServiLlantas Celaya</title>
    {/if}
  </svelte:head>
  
  {#if loading}
    <!-- ========== LOADING ========== -->
    <section class="loading-section">
      <div class="loading-spinner"></div>
      <p>Cargando artículo...</p>
    </section>
  
  {:else if error || !post}
    <!-- ========== ERROR ========== -->
    <section class="error-section">
      <div class="container text-center">
        <span class="error-icon">😕</span>
        <h1>Artículo no encontrado</h1>
        <p>El artículo que buscas no existe o fue eliminado.</p>
        <a href="/blog" class="btn btn-accent">← Volver al blog</a>
      </div>
    </section>
  
  {:else}
    <!-- ========== HERO ARTÍCULO ========== -->
    <section class="article-hero">
      <div class="article-hero-overlay"></div>
      <div class="article-hero-content">
        <nav class="breadcrumb">
          <a href="/">Inicio</a>
          <span>/</span>
          <a href="/blog">Blog</a>
          <span>/</span>
          <span class="current">{post.categoria}</span>
        </nav>
  
        <span class="article-category" style="background: {getCategoryColor(post.categoria)}">
          {post.categoria}
        </span>
  
        <h1>{post.titulo}</h1>
  
        <div class="article-meta">
          <div class="meta-author">
            <div class="meta-avatar">{post.autor.charAt(0)}</div>
            <div>
              <span class="meta-name">{post.autor}</span>
              <span class="meta-role">{post.autor_rol}</span>
            </div>
          </div>
          <div class="meta-details">
            <span>📅 {formatDate(post.fecha)}</span>
            <span>⏱️ {post.duracion}</span>
          </div>
        </div>
      </div>
    </section>
  
    <!-- ========== CONTENIDO ========== -->
    <section class="article-section">
      <div class="container">
        <div class="article-layout">
  
          <!-- Contenido principal -->
          <article class="article-content fade-up" use:inview>
  
            <!-- Imagen principal -->
            <div class="article-image">
              <div class="article-img-placeholder">
                <span>📷</span>
                <p>Imagen: {post.titulo}</p>
                <small>static{post.imagen}</small>
              </div>
            </div>
  
            <!-- Resumen destacado -->
            <div class="article-summary">
              <p>{post.resumen}</p>
            </div>
  
            <!-- Contenido del artículo -->
            <div class="article-body">
              {@html renderContent(post.contenido)}
            </div>
  
            <!-- Tags -->
            <div class="article-tags">
              <span class="tags-label">Etiquetas:</span>
              {#each post.tags as tag}
                <span class="tag">#{tag}</span>
              {/each}
            </div>
  
            <!-- Compartir -->
            <div class="article-share">
              <span class="share-label">Compartir:</span>
              <div class="share-buttons">
                <button class="share-btn whatsapp" on:click={shareOnWhatsApp}>
                  💬 WhatsApp
                </button>
                <button class="share-btn copy" on:click={copyLink}>
                  🔗 Copiar link
                </button>
              </div>
            </div>
          </article>
  
          <!-- Sidebar -->
          <aside class="article-sidebar">
  
            <!-- CTA Cotizar -->
            <div class="sidebar-card cta-card fade-right" use:inview>
              <h3>¿Necesitas este servicio?</h3>
              <p>Cotiza sin compromiso y agenda tu cita hoy mismo.</p>
              <a href="/contacto" class="btn btn-accent btn-block">
                ✉️ Cotizar ahora
              </a>
              <a href="tel:+524611203488" class="btn btn-outline-accent btn-block">
                📞 461 120 3488
              </a>
            </div>
  
            <!-- Artículos relacionados -->
            {#if relatedPosts.length > 0}
              <div class="sidebar-card fade-right delay-1" use:inview>
                <h3>Artículos relacionados</h3>
                <div class="related-list">
                  {#each relatedPosts as related}
                    <a href="/blog/{related.slug}" class="related-item">
                      <div class="related-thumb">
                        <span>📄</span>
                      </div>
                      <div class="related-info">
                        <span class="related-title">{related.titulo}</span>
                        <span class="related-date">{formatDate(related.fecha)}</span>
                      </div>
                    </a>
                  {/each}
                </div>
              </div>
            {/if}
  
            <!-- Card WhatsApp -->
            <div class="sidebar-card whatsapp-sidebar fade-right delay-2" use:inview>
              <span class="sidebar-emoji">💬</span>
              <h3>¿Tienes dudas?</h3>
              <p>Escríbenos por WhatsApp y resolvemos tus preguntas.</p>
              <a
                href="https://wa.me/524611203488?text=Hola%2C%20leí%20su%20artículo%20y%20tengo%20una%20duda"
                target="_blank"
                rel="noopener noreferrer"
                class="btn btn-whatsapp btn-block"
              >
                Chatear ahora
              </a>
            </div>
          </aside>
        </div>
      </div>
    </section>
  
    <!-- ========== MÁS ARTÍCULOS ========== -->
    {#if relatedPosts.length > 0}
      <section class="section section-dark">
        <div class="container text-center">
          <p class="section-subtitle fade-up" use:inview>Sigue leyendo</p>
          <h2 class="section-title fade-up" use:inview>
            Más artículos que podrían <span class="text-accent">interesarte</span>
          </h2>
  
          <div class="more-grid">
            {#each relatedPosts as related, i}
              <a href="/blog/{related.slug}" class="more-card fade-up delay-{i + 1}" use:inview>
                <div class="more-img">
                  <span>📄</span>
                </div>
                <div class="more-body">
                  <span class="more-cat" style="color: {getCategoryColor(related.categoria)}">
                    {related.categoria}
                  </span>
                  <h3>{related.titulo}</h3>
                  <p>{related.resumen}</p>
                  <span class="more-read">Leer artículo →</span>
                </div>
              </a>
            {/each}
          </div>
        </div>
      </section>
    {/if}
  
    <!-- ========== CTA FINAL ========== -->
    <section class="cta-final">
      <div class="cta-final-overlay"></div>
      <div class="container text-center" style="position:relative; z-index:2;">
        <h2 class="fade-up" use:inview>
          ¿Tu auto necesita <span class="text-accent">atención</span>?
        </h2>
        <p class="cta-desc fade-up delay-1" use:inview>
          No esperes a que sea tarde. Agenda tu cita hoy.
        </p>
        <div class="cta-buttons fade-up delay-2" use:inview>
          <a href="/contacto" class="btn btn-accent btn-lg">✉️ Agendar cita</a>
          <a href="tel:+524611203488" class="btn btn-primary btn-lg">📞 461 120 3488</a>
        </div>
      </div>
    </section>
  {/if}
  
  <style>
    /* ====== LOADING / ERROR ====== */
    .loading-section, .error-section {
      min-height: 80vh;
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
      gap: 1rem;
      padding-top: 5rem;
    }
  
    .loading-spinner {
      width: 50px;
      height: 50px;
      border: 4px solid var(--color-dark-4);
      border-top-color: var(--color-accent);
      border-radius: 50%;
      animation: spin 0.8s linear infinite;
    }
  
    @keyframes spin { to { transform: rotate(360deg); } }
  
    .loading-section p { color: var(--color-gray); }
  
    .error-icon { font-size: 4rem; display: block; margin-bottom: 1rem; }
    .error-section h1 { margin-bottom: 0.5rem; }
    .error-section p { color: var(--color-gray); margin-bottom: 1.5rem; }
  
    /* ====== ARTICLE HERO ====== */
    .article-hero {
      position: relative;
      padding: 10rem 2rem 4rem;
      text-align: center;
      background: url('/images/hero-bg.jpg') center/cover no-repeat;
      background-color: var(--color-primary);
    }
  
    .article-hero-overlay {
      position: absolute;
      inset: 0;
      background: linear-gradient(180deg, rgba(22, 17, 47, 0.95), rgba(22, 17, 47, 0.88));
    }
  
    .article-hero-content {
      position: relative;
      z-index: 2;
      max-width: 800px;
      margin: 0 auto;
    }
  
    .article-category {
      display: inline-block;
      padding: 0.4rem 1.2rem;
      border-radius: 50px;
      font-size: 0.8rem;
      font-weight: 700;
      color: white;
      margin: 1.5rem 0 1rem;
    }
  
    .article-hero h1 {
      font-size: clamp(1.8rem, 4vw, 2.8rem);
      margin-bottom: 1.5rem;
      line-height: 1.3;
    }
  
    .article-meta {
      display: flex;
      justify-content: center;
      align-items: center;
      gap: 2rem;
      flex-wrap: wrap;
    }
  
    .meta-author {
      display: flex;
      align-items: center;
      gap: 0.7rem;
    }
  
    .meta-avatar {
      width: 45px;
      height: 45px;
      border-radius: 50%;
      background: var(--color-secondary);
      display: flex;
      align-items: center;
      justify-content: center;
      font-family: var(--font-heading);
      font-weight: 800;
      font-size: 1.1rem;
    }
  
    .meta-name { display: block; font-weight: 600; }
    .meta-role { display: block; color: var(--color-gray); font-size: 0.8rem; }
  
    .meta-details {
      display: flex;
      gap: 1.5rem;
      color: var(--color-gray);
      font-size: 0.9rem;
    }
  
    .breadcrumb { display: flex; justify-content: center; gap: 0.5rem; font-size: 0.85rem; color: var(--color-gray); margin-bottom: 0.5rem; }
    .breadcrumb a { color: var(--color-light-3); transition: color 0.3s; }
    .breadcrumb a:hover { color: var(--color-accent); }
    .breadcrumb .current { color: var(--color-accent); }
  
    /* ====== ARTICLE LAYOUT ====== */
    .article-section {
      background: var(--color-dark);
      padding: 4rem 0;
    }
  
    .article-layout {
      display: grid;
      grid-template-columns: 1fr 350px;
      gap: 3rem;
      align-items: flex-start;
    }
  
    /* ====== ARTICLE CONTENT ====== */
    .article-image {
      margin-bottom: 2rem;
      border-radius: 16px;
      overflow: hidden;
    }
  
    .article-img-placeholder {
      height: 350px;
      background: linear-gradient(135deg, var(--color-primary-light), var(--color-secondary));
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
      gap: 0.5rem;
      color: rgba(255, 255, 255, 0.5);
      border-radius: 16px;
    }
  
    .article-img-placeholder span { font-size: 3rem; }
    .article-img-placeholder small { font-size: 0.75rem; }
  
    .article-summary {
      background: var(--color-dark-3);
      border-left: 4px solid var(--color-accent);
      padding: 1.5rem 2rem;
      border-radius: 0 12px 12px 0;
      margin-bottom: 2rem;
    }
  
    .article-summary p {
      color: var(--color-light-3);
      font-size: 1.05rem;
      line-height: 1.8;
      font-style: italic;
    }
  
    /* Article body */
    .article-body {
      color: var(--color-light-3);
      font-size: 1rem;
      line-height: 1.9;
      margin-bottom: 2rem;
    }
  
    .article-body :global(h2) {
      font-family: var(--font-heading);
      font-size: 1.5rem;
      color: var(--color-light);
      margin: 2.5rem 0 1rem;
      padding-bottom: 0.5rem;
      border-bottom: 2px solid var(--color-dark-4);
    }
  
    .article-body :global(h3) {
      font-family: var(--font-heading);
      font-size: 1.2rem;
      color: var(--color-accent);
      margin: 2rem 0 0.8rem;
    }
  
    .article-body :global(strong) {
      color: var(--color-light);
    }
  
    .article-body :global(li) {
      padding-left: 1.5rem;
      position: relative;
      margin-bottom: 0.5rem;
    }
  
    .article-body :global(li)::before {
      content: '▸';
      position: absolute;
      left: 0;
      color: var(--color-accent);
      font-weight: bold;
    }
  
    .article-body :global(p) { margin-bottom: 1rem; }
  
    .article-body :global(tr) {
      display: grid;
      grid-template-columns: 1fr 1fr;
      gap: 1px;
      margin-bottom: 1px;
    }
  
    .article-body :global(td) {
      padding: 0.8rem 1rem;
      background: var(--color-dark-3);
      font-size: 0.9rem;
    }
  
    /* Tags */
    .article-tags {
      display: flex;
      align-items: center;
      gap: 0.8rem;
      flex-wrap: wrap;
      padding: 1.5rem 0;
      border-top: 1px solid var(--color-dark-4);
      border-bottom: 1px solid var(--color-dark-4);
      margin-bottom: 1.5rem;
    }
  
    .tags-label { color: var(--color-gray); font-size: 0.9rem; font-weight: 600; }
  
    .tag {
      padding: 0.3rem 0.8rem;
      background: rgba(0, 75, 188, 0.15);
      border: 1px solid rgba(0, 75, 188, 0.3);
      border-radius: 50px;
      font-size: 0.8rem;
      color: var(--color-secondary-light);
      font-weight: 600;
    }
  
    /* Share */
    .article-share {
      display: flex;
      align-items: center;
      gap: 1rem;
      flex-wrap: wrap;
    }
  
    .share-label { color: var(--color-gray); font-weight: 600; font-size: 0.9rem; }
  
    .share-buttons { display: flex; gap: 0.8rem; }
  
    .share-btn {
      padding: 0.6rem 1.2rem;
      border-radius: 50px;
      font-size: 0.85rem;
      font-weight: 600;
      cursor: pointer;
      transition: all 0.3s ease;
      border: none;
      font-family: var(--font-body);
    }
  
    .share-btn.whatsapp { background: #25d366; color: white; }
    .share-btn.whatsapp:hover { background: #1ea952; transform: translateY(-2px); }
  
    .share-btn.copy { background: var(--color-dark-3); border: 1px solid var(--color-dark-4); color: var(--color-light-3); }
    .share-btn.copy:hover { border-color: var(--color-accent); color: var(--color-accent); }
  
    /* ====== SIDEBAR ====== */
    .article-sidebar {
      display: flex;
      flex-direction: column;
      gap: 1.5rem;
      position: sticky;
      top: 100px;
    }
  
    .sidebar-card {
      background: var(--color-dark-3);
      border: 1px solid var(--color-dark-4);
      border-radius: 16px;
      padding: 1.8rem;
      transition: border-color 0.3s;
    }
  
    .sidebar-card:hover { border-color: var(--color-secondary); }
  
    .sidebar-card h3 { font-size: 1.05rem; margin-bottom: 0.8rem; }
    .sidebar-card p { color: var(--color-gray); font-size: 0.9rem; line-height: 1.6; margin-bottom: 1rem; }
  
    .sidebar-emoji { font-size: 2rem; display: block; margin-bottom: 0.5rem; }
  
    .btn-block { width: 100%; justify-content: center; margin-bottom: 0.8rem; display: inline-flex; }
    .btn-block:last-child { margin-bottom: 0; }
  
    .cta-card { border-color: var(--color-accent); background: rgba(243, 253, 1, 0.03); }
  
    /* Related */
    .related-list { display: flex; flex-direction: column; gap: 1rem; }
  
    .related-item {
      display: flex;
      gap: 1rem;
      align-items: center;
      padding: 0.8rem;
      border-radius: 10px;
      transition: background 0.3s;
    }
  
    .related-item:hover { background: var(--color-dark-4); }
  
    .related-thumb {
      width: 55px;
      height: 55px;
      background: linear-gradient(135deg, var(--color-primary-light), var(--color-secondary));
      border-radius: 10px;
      display: flex;
      align-items: center;
      justify-content: center;
      font-size: 1.3rem;
      flex-shrink: 0;
    }
  
    .related-title {
      display: block;
      font-weight: 600;
      font-size: 0.85rem;
      line-height: 1.4;
      margin-bottom: 0.2rem;
      transition: color 0.3s;
    }
  
    .related-item:hover .related-title { color: var(--color-accent); }
  
    .related-date { color: var(--color-gray); font-size: 0.75rem; }
  
    .whatsapp-sidebar { border-color: rgba(37, 211, 102, 0.3); }
    .whatsapp-sidebar:hover { border-color: #25d366; }
  
    .btn-whatsapp {
      padding: 0.9rem 2rem;
      background: #25d366;
      color: white;
      border: none;
      border-radius: 50px;
      font-family: var(--font-heading);
      font-weight: 700;
      cursor: pointer;
      transition: all 0.3s;
    }
  
    .btn-whatsapp:hover { background: #1ea952; transform: translateY(-2px); }
  
    /* ====== MORE ARTICLES ====== */
    .more-grid {
      display: grid;
      grid-template-columns: repeat(3, 1fr);
      gap: 2rem;
      margin-top: 2rem;
    }
  
    .more-card {
      background: var(--color-dark-3);
      border: 1px solid var(--color-dark-4);
      border-radius: 16px;
      overflow: hidden;
      transition: all 0.4s ease;
      text-align: left;
    }
  
    .more-card:hover {
      border-color: var(--color-secondary);
      transform: translateY(-5px);
    }
  
    .more-img {
      height: 160px;
      background: linear-gradient(135deg, var(--color-primary-light), var(--color-secondary));
      display: flex;
      align-items: center;
      justify-content: center;
      font-size: 2.5rem;
    }
  
    .more-body { padding: 1.5rem; }
  
    .more-cat { font-size: 0.8rem; font-weight: 700; text-transform: uppercase; letter-spacing: 1px; }
    .more-body h3 { font-size: 1.05rem; margin: 0.5rem 0; line-height: 1.4; transition: color 0.3s; }
    .more-card:hover h3 { color: var(--color-accent); }
    .more-body p { color: var(--color-gray); font-size: 0.85rem; line-height: 1.6; display: -webkit-box; -webkit-line-clamp: 2; -webkit-box-orient: vertical; overflow: hidden; margin-bottom: 1rem; }
    .more-read { color: var(--color-accent); font-weight: 700; font-size: 0.85rem; }
  
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
  
    .cta-final h2 { font-size: clamp(1.8rem, 4vw, 2.8rem); margin-bottom: 1rem; }
    .cta-desc { color: rgba(255, 255, 255, 0.8); font-size: 1.1rem; margin-bottom: 2rem; }
    .cta-buttons { display: flex; gap: 1rem; justify-content: center; flex-wrap: wrap; }
    .btn-lg { padding: 1rem 2.5rem; font-size: 1.05rem; }
  
    /* ====== RESPONSIVE ====== */
    @media (max-width: 1024px) {
      .article-layout {
        grid-template-columns: 1fr;
      }
  
      .article-sidebar {
        position: relative;
        top: 0;
        display: grid;
        grid-template-columns: 1fr 1fr;
        gap: 1.5rem;
      }
  
      .more-grid { grid-template-columns: repeat(2, 1fr); }
    }
  
    @media (max-width: 768px) {
      .article-hero { padding: 8rem 1.5rem 3rem; }
  
      .article-sidebar {
        grid-template-columns: 1fr;
      }
  
      .article-meta {
        flex-direction: column;
        gap: 1rem;
      }
  
      .more-grid { grid-template-columns: 1fr; }
  
      .article-img-placeholder { height: 220px; }
  
      .share-buttons { flex-wrap: wrap; }
    }
  
    @media (max-width: 480px) {
      .cta-buttons {
        flex-direction: column;
        align-items: center;
      }
      .cta-buttons .btn {
        width: 100%;
        justify-content: center;
      }
    }
  </style>