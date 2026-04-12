<!-- svelte-ignore non_reactive_update -->
<!-- svelte-ignore non_reactive_update -->
<!-- svelte-ignore non_reactive_update -->
<!-- frontend/src/routes/blog/+page.svelte -->
<!-- REEMPLAZA TODO EL ARCHIVO -->
<script>
// @ts-nocheck

  import { onMount } from 'svelte';
  import { inview } from '$lib/actions/inview.js';

  // ===== DATOS DE RESPALDO =====
  const fallbackPosts = [
    {
      id: 1,
      titulo: '¿Cada cuánto debo cambiar mis llantas?',
      resumen: 'Descubre las señales que indican que es momento de cambiar tus llantas y cómo prolongar su vida útil con cuidados simples.',
      imagen: '/images/blog/llantas.jpg',
      fecha: '2024-12-15',
      categoria: 'Mantenimiento',
      slug: 'cada-cuanto-cambiar-llantas',
      autor: 'Carlos Ramírez',
      autor_rol: 'Director General',
      duracion: '5 min de lectura',
      tags: ['llantas', 'mantenimiento', 'seguridad']
    },
    {
      id: 2,
      titulo: 'Importancia de la alineación y balanceo',
      resumen: 'Una mala alineación puede costarte más de lo que imaginas. Conoce por qué es vital este servicio para tu seguridad y tu bolsillo.',
      imagen: '/images/blog/alineacion.jpg',
      fecha: '2024-12-01',
      categoria: 'Seguridad',
      slug: 'importancia-alineacion-balanceo',
      autor: 'Miguel Torres',
      autor_rol: 'Jefe de Taller',
      duracion: '4 min de lectura',
      tags: ['alineación', 'balanceo', 'seguridad']
    },
    {
      id: 3,
      titulo: '5 tips para cuidar tus llantas en temporada de lluvias',
      resumen: 'La temporada de lluvias puede ser peligrosa. Estos consejos te ayudarán a mantener el control de tu vehículo.',
      imagen: '/images/blog/lluvias.jpg',
      fecha: '2024-11-20',
      categoria: 'Tips',
      slug: 'tips-llantas-lluvias',
      autor: 'Carlos Ramírez',
      autor_rol: 'Director General',
      duracion: '4 min de lectura',
      tags: ['lluvias', 'seguridad', 'consejos']
    },
    {
      id: 4,
      titulo: '¿Cómo saber si mis frenos necesitan mantenimiento?',
      resumen: 'Los frenos son el sistema de seguridad más importante de tu vehículo. Aprende a identificar las señales de alerta.',
      imagen: '/images/blog/frenos.jpg',
      fecha: '2024-11-10',
      categoria: 'Seguridad',
      slug: 'como-saber-frenos-mantenimiento',
      autor: 'Miguel Torres',
      autor_rol: 'Jefe de Taller',
      duracion: '5 min de lectura',
      tags: ['frenos', 'seguridad', 'mantenimiento']
    },
    {
      id: 5,
      titulo: 'Guía completa del cambio de aceite: todo lo que debes saber',
      resumen: 'El cambio de aceite es el mantenimiento más básico pero más importante para tu motor.',
      imagen: '/images/blog/aceite.jpg',
      fecha: '2024-10-25',
      categoria: 'Mantenimiento',
      slug: 'guia-cambio-aceite',
      autor: 'Carlos Ramírez',
      autor_rol: 'Director General',
      duracion: '6 min de lectura',
      tags: ['aceite', 'motor', 'mantenimiento']
    },
    {
      id: 6,
      titulo: 'Suspensión: señales de desgaste que no debes ignorar',
      resumen: 'Una suspensión en mal estado afecta tu seguridad, comodidad y hasta el desgaste de tus llantas.',
      imagen: '/images/blog/suspension.jpg',
      fecha: '2024-10-10',
      categoria: 'Mecánica',
      slug: 'suspension-senales-desgaste',
      autor: 'Miguel Torres',
      autor_rol: 'Jefe de Taller',
      duracion: '5 min de lectura',
      tags: ['suspensión', 'amortiguadores', 'mecánica']
    }
  ];

  /**
	 * @type {any[]}
	 */
  let posts = [];
  /**
	 * @type {any[] | null | undefined}
	 */
  let filteredPosts = [];
  let loading = true;
  let activeCategory = 'Todos';
  let searchQuery = '';

  const categorias = ['Todos', 'Mantenimiento', 'Seguridad', 'Tips', 'Mecánica'];

  onMount(async () => {
    try {
      const res = await fetch('http://localhost:3000/api/blog');
      if (res.ok) {
        posts = await res.json();
      } else {
        posts = fallbackPosts;
      }
    } catch (err) {
      // Si el backend no responde, usamos datos locales
      console.warn('Backend no disponible, usando datos locales');
      posts = fallbackPosts;
    } finally {
      filteredPosts = posts;
      loading = false;
    }
  });

  /**
	 * @param {string} cat
	 */
  function filterByCategory(cat) {
    activeCategory = cat;
    applyFilters();
  }

  function applyFilters() {
    let result = posts;

    if (activeCategory !== 'Todos') {
      result = result.filter(p => p.categoria === activeCategory);
    }

    if (searchQuery.trim()) {
      const q = searchQuery.toLowerCase();
      result = result.filter(p =>
        p.titulo.toLowerCase().includes(q) ||
        p.resumen.toLowerCase().includes(q) ||
        p.tags?.some((/** @type {string} */ t) => t.toLowerCase().includes(q))
      );
    }

    filteredPosts = result;
  }

  /**
	 * @param {string} dateStr
	 */
  function formatDate(dateStr) {
    const date = new Date(dateStr + 'T00:00:00');
    return date.toLocaleDateString('es-MX', {
      year: 'numeric',
      month: 'long',
      day: 'numeric'
    });
  }

  /**
	 * @param {string | number} cat
	 */
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
</script>

<svelte:head>
  <title>Blog | ServiLlantas Celaya - Tips y Consejos Automotrices</title>
  <meta name="description" content="Consejos, tips y guías sobre llantas, mantenimiento automotriz y seguridad vial." />
</svelte:head>

<!-- ========== HERO ========== -->
<section class="page-hero">
  <div class="page-hero-overlay"></div>
  <div class="hero-particles">
    {#each Array(15) as _, i}
      <div class="particle" style="--delay: {i * 0.6}s; --x: {Math.random() * 100}%; --duration: {3 + Math.random() * 4}s"></div>
    {/each}
  </div>
  <div class="page-hero-content">
    <p class="hero-subtitle">BLOG AUTOMOTRIZ</p>
    <h1>Tips y <span class="text-accent">Consejos</span></h1>
    <p class="hero-desc">
      Artículos, guías y consejos para mantener tu vehículo en las mejores condiciones.
    </p>
    <nav class="breadcrumb">
      <a href="/">Inicio</a>
      <span>/</span>
      <span class="current">Blog</span>
    </nav>
  </div>
</section>

<!-- ========== FILTROS ========== -->
<section class="filters-section">
  <div class="container">
    <div class="filters-bar fade-up" use:inview>
      <div class="category-filters">
        {#each categorias as cat}
          <button
            class="category-btn"
            class:active={activeCategory === cat}
            on:click={() => filterByCategory(cat)}
          >
            {cat}
            {#if cat !== 'Todos'}
              <span class="cat-count">
                {posts.filter(p => p.categoria === cat).length}
              </span>
            {/if}
          </button>
        {/each}
      </div>

      <div class="search-box">
        <span class="search-icon">🔍</span>
        <input
          type="text"
          placeholder="Buscar artículos..."
          bind:value={searchQuery}
          on:input={applyFilters}
        />
        {#if searchQuery}
          <button class="search-clear" on:click={() => { searchQuery = ''; applyFilters(); }}>✕</button>
        {/if}
      </div>
    </div>

    <!-- Contador de resultados -->
    {#if !loading}
      <p class="results-count">
        {filteredPosts.length} {filteredPosts.length === 1 ? 'artículo' : 'artículos'}
        {activeCategory !== 'Todos' ? `en ${activeCategory}` : ''}
        {searchQuery ? `con "${searchQuery}"` : ''}
      </p>
    {/if}
  </div>
</section>

<!-- ========== ARTÍCULOS ========== -->
<section class="section section-dark">
  <div class="container">

    {#if loading}
      <div class="blog-grid">
        {#each Array(6) as _}
          <div class="skeleton-card">
            <div class="skeleton-img"></div>
            <div class="skeleton-body">
              <div class="skeleton-line short"></div>
              <div class="skeleton-line"></div>
              <div class="skeleton-line"></div>
              <div class="skeleton-line medium"></div>
            </div>
          </div>
        {/each}
      </div>

    {:else if filteredPosts.length === 0}
      <div class="empty-state fade-up" use:inview>
        <span class="empty-icon">📝</span>
        <h3>No se encontraron artículos</h3>
        <p>Intenta con otra categoría o término de búsqueda.</p>
        <button class="btn btn-outline-accent" on:click={() => { activeCategory = 'Todos'; searchQuery = ''; applyFilters(); }}>
          Ver todos los artículos
        </button>
      </div>

    {:else}
      <!-- Artículo destacado -->
      {#if activeCategory === 'Todos' && !searchQuery}
        <div class="featured-post fade-up" use:inview>
          <a href="/blog/{filteredPosts[0].slug}" class="featured-link">
            <div class="featured-image">
              <div class="featured-placeholder">
                <div class="featured-icon-bg">📰</div>
              </div>
              <div class="featured-badge">
                <span style="background: {getCategoryColor(filteredPosts[0].categoria)}">
                  {filteredPosts[0].categoria}
                </span>
              </div>
            </div>
            <div class="featured-content">
              <div class="featured-meta">
                <span>📅 {formatDate(filteredPosts[0].fecha)}</span>
                <span>⏱️ {filteredPosts[0].duracion}</span>
              </div>
              <h2>{filteredPosts[0].titulo}</h2>
              <p>{filteredPosts[0].resumen}</p>
              <div class="featured-tags">
                {#each filteredPosts[0].tags?.slice(0, 3) || [] as tag}
                  <span class="mini-tag">#{tag}</span>
                {/each}
              </div>
              <div class="featured-footer">
                <div class="author-info">
                  <div class="author-avatar-sm">{filteredPosts[0].autor.charAt(0)}</div>
                  <div>
                    <span class="author-name">{filteredPosts[0].autor}</span>
                    <span class="author-role">{filteredPosts[0].autor_rol}</span>
                  </div>
                </div>
                <span class="read-more">Leer artículo →</span>
              </div>
            </div>
          </a>
        </div>
      {/if}

      <!-- Grid -->
      <div class="blog-grid">
        {#each (activeCategory === 'Todos' && !searchQuery ? filteredPosts.slice(1) : filteredPosts) as post, i}
          <a href="/blog/{post.slug}" class="blog-card fade-up delay-{Math.min(i + 1, 6)}" use:inview>
            <div class="card-image">
              <div class="card-img-placeholder">
                <span>📄</span>
              </div>
              <div class="card-category" style="background: {getCategoryColor(post.categoria)}">
                {post.categoria}
              </div>
              <div class="card-duration">⏱️ {post.duracion}</div>
            </div>

            <div class="card-body">
              <div class="card-meta">
                <span>📅 {formatDate(post.fecha)}</span>
              </div>
              <h3>{post.titulo}</h3>
              <p>{post.resumen}</p>
              <div class="card-tags">
                {#each post.tags?.slice(0, 2) || [] as tag}
                  <span class="mini-tag">#{tag}</span>
                {/each}
              </div>
            </div>

            <div class="card-footer">
              <div class="card-author">
                <div class="author-avatar-xs">{post.autor.charAt(0)}</div>
                <span>{post.autor}</span>
              </div>
              <span class="card-read-more">Leer →</span>
            </div>
          </a>
        {/each}
      </div>
    {/if}
  </div>
</section>

<!-- ========== NEWSLETTER CTA ========== -->
<section class="cta-newsletter">
  <div class="cta-newsletter-overlay"></div>
  <div class="container text-center" style="position:relative; z-index:2;">
    <span class="newsletter-icon fade-up" use:inview>📬</span>
    <h2 class="fade-up delay-1" use:inview>
      ¿Te gustaría recibir <span class="text-accent">tips automotrices</span>?
    </h2>
    <p class="newsletter-desc fade-up delay-2" use:inview>
      Síguenos en redes sociales para recibir consejos, promociones y novedades.
    </p>
    <div class="newsletter-buttons fade-up delay-3" use:inview>
      <a href="#" class="btn btn-accent btn-lg">📱 Facebook</a>
      <a href="#" class="btn btn-primary btn-lg">📸 Instagram</a>
    </div>
  </div>
</section>

<!-- ========== CTA CONTACTO ========== -->
<section class="section section-dark">
  <div class="container text-center">
    <h2 class="fade-up" use:inview>
      ¿Necesitas alguno de estos <span class="text-accent">servicios</span>?
    </h2>
    <p class="section-desc fade-up delay-1" use:inview>
      Contáctanos para cotizar sin compromiso.
    </p>
    <div class="cta-buttons fade-up delay-2" use:inview>
      <a href="/contacto" class="btn btn-accent btn-lg">✉️ Cotizar ahora</a>
      <a href="tel:+524611203488" class="btn btn-primary btn-lg">📞 461 120 3488</a>
    </div>
  </div>
</section>

<style>
  .page-hero {
    position: relative; padding: 10rem 2rem 5rem; text-align: center;
    background: url('/images/hero-bg.jpg') center/cover no-repeat;
    background-color: var(--color-primary); overflow: hidden;
  }
  .page-hero-overlay { position: absolute; inset: 0; background: linear-gradient(180deg, rgba(22,17,47,0.95), rgba(22,17,47,0.85)); }
  .hero-particles { position: absolute; inset: 0; z-index: 1; overflow: hidden; }
  .particle { position: absolute; width: 4px; height: 4px; background: var(--color-accent); border-radius: 50%; left: var(--x); bottom: -10px; opacity: 0; animation: floatUp var(--duration) var(--delay) infinite; }
  @keyframes floatUp { 0%{opacity:0;transform:translateY(0) scale(0)} 10%{opacity:.6} 90%{opacity:.2} 100%{opacity:0;transform:translateY(-100vh) scale(1)} }
  .page-hero-content { position: relative; z-index: 2; max-width: 700px; margin: 0 auto; }
  .page-hero .hero-subtitle { color: var(--color-accent); font-family: var(--font-heading); font-weight: 700; font-size: .85rem; letter-spacing: 4px; margin-bottom: 1rem; }
  .page-hero h1 { font-size: clamp(2.2rem,5vw,3.5rem); margin-bottom: 1rem; }
  .page-hero .hero-desc { color: var(--color-light-3); font-size: 1.1rem; margin-bottom: 1.5rem; }
  .breadcrumb { display: flex; justify-content: center; gap: .5rem; font-size: .9rem; color: var(--color-gray); }
  .breadcrumb a { color: var(--color-light-3); transition: color .3s; }
  .breadcrumb a:hover { color: var(--color-accent); }
  .breadcrumb .current { color: var(--color-accent); }

  .filters-section { background: var(--color-dark-2); padding: 2rem 0; border-bottom: 1px solid var(--color-dark-4); position: sticky; top: 60px; z-index: 50; }
  .filters-bar { display: flex; justify-content: space-between; align-items: center; gap: 2rem; flex-wrap: wrap; }
  .category-filters { display: flex; gap: .5rem; flex-wrap: wrap; }
  .category-btn { padding: .6rem 1.2rem; background: var(--color-dark-3); border: 1px solid var(--color-dark-4); border-radius: 50px; color: var(--color-light-3); font-family: var(--font-body); font-size: .85rem; font-weight: 600; cursor: pointer; transition: all .3s; display: inline-flex; align-items: center; gap: .4rem; }
  .category-btn:hover { border-color: var(--color-secondary); color: var(--color-light); }
  .category-btn.active { background: var(--color-secondary); border-color: var(--color-secondary); color: var(--color-light); }
  .cat-count { background: rgba(255,255,255,.15); padding: .1rem .45rem; border-radius: 50px; font-size: .7rem; }
  .category-btn.active .cat-count { background: rgba(255,255,255,.25); }

  .search-box { display: flex; align-items: center; gap: .5rem; background: var(--color-dark-3); border: 1px solid var(--color-dark-4); border-radius: 50px; padding: .5rem 1rem; min-width: 250px; transition: border-color .3s; }
  .search-box:focus-within { border-color: var(--color-secondary); }
  .search-icon { font-size: .9rem; }
  .search-box input { background: none; border: none; color: var(--color-light); font-size: .9rem; outline: none; width: 100%; font-family: var(--font-body); }
  .search-box input::placeholder { color: var(--color-gray); }
  .search-clear { background: none; border: none; color: var(--color-gray); cursor: pointer; font-size: .9rem; padding: .2rem; transition: color .3s; }
  .search-clear:hover { color: var(--color-accent); }

  .results-count { color: var(--color-gray); font-size: .85rem; margin-top: 1rem; }

  .featured-post { margin-bottom: 3rem; }
  .featured-link { display: grid; grid-template-columns: 1.2fr 1fr; background: var(--color-dark-3); border: 1px solid var(--color-dark-4); border-radius: 20px; overflow: hidden; transition: all .4s; }
  .featured-link:hover { border-color: var(--color-secondary); transform: translateY(-5px); box-shadow: 0 20px 40px rgba(0,75,188,.15); }
  .featured-image { position: relative; min-height: 350px; }
  .featured-placeholder { width: 100%; height: 100%; background: linear-gradient(135deg, var(--color-primary-light), var(--color-secondary)); display: flex; align-items: center; justify-content: center; }
  .featured-icon-bg { font-size: 5rem; opacity: .3; }
  .featured-badge { position: absolute; top: 1.5rem; left: 1.5rem; }
  .featured-badge span { padding: .4rem 1rem; border-radius: 50px; font-size: .8rem; font-weight: 700; color: white; }
  .featured-content { padding: 2.5rem; display: flex; flex-direction: column; justify-content: center; }
  .featured-meta { display: flex; gap: 1.5rem; color: var(--color-gray); font-size: .85rem; margin-bottom: 1rem; }
  .featured-content h2 { font-size: 1.6rem; margin-bottom: 1rem; line-height: 1.3; transition: color .3s; }
  .featured-link:hover .featured-content h2 { color: var(--color-accent); }
  .featured-content > p { color: var(--color-light-3); line-height: 1.8; margin-bottom: 1rem; font-size: .95rem; }
  .featured-tags { display: flex; gap: .5rem; flex-wrap: wrap; margin-bottom: 1.5rem; }
  .mini-tag { font-size: .75rem; color: var(--color-secondary-light); font-weight: 600; }
  .featured-footer { display: flex; justify-content: space-between; align-items: center; }
  .author-info { display: flex; align-items: center; gap: .7rem; }
  .author-avatar-sm { width: 40px; height: 40px; border-radius: 50%; background: var(--color-secondary); display: flex; align-items: center; justify-content: center; font-family: var(--font-heading); font-weight: 800; font-size: 1rem; }
  .author-name { display: block; font-weight: 600; font-size: .9rem; }
  .author-role { display: block; color: var(--color-gray); font-size: .8rem; }
  .read-more { color: var(--color-accent); font-weight: 700; font-size: .9rem; transition: letter-spacing .3s; }
  .featured-link:hover .read-more { letter-spacing: 2px; }

  .blog-grid { display: grid; grid-template-columns: repeat(3,1fr); gap: 2rem; }
  .blog-card { background: var(--color-dark-3); border: 1px solid var(--color-dark-4); border-radius: 16px; overflow: hidden; transition: all .4s; display: flex; flex-direction: column; }
  .blog-card:hover { border-color: var(--color-secondary); transform: translateY(-8px); box-shadow: 0 20px 40px rgba(0,75,188,.15); }
  .card-image { position: relative; height: 200px; overflow: hidden; }
  .card-img-placeholder { width: 100%; height: 100%; background: linear-gradient(135deg, var(--color-primary-light), var(--color-secondary)); display: flex; align-items: center; justify-content: center; font-size: 3rem; transition: transform .4s; }
  .blog-card:hover .card-img-placeholder { transform: scale(1.05); }
  .card-category { position: absolute; top: 1rem; left: 1rem; padding: .3rem .8rem; border-radius: 50px; font-size: .75rem; font-weight: 700; color: white; }
  .card-duration { position: absolute; top: 1rem; right: 1rem; padding: .3rem .6rem; border-radius: 50px; font-size: .7rem; background: rgba(0,0,0,.6); color: white; backdrop-filter: blur(5px); }
  .card-body { padding: 1.5rem; flex: 1; }
  .card-meta { display: flex; gap: 1rem; color: var(--color-gray); font-size: .8rem; margin-bottom: .8rem; }
  .card-body h3 { font-size: 1.1rem; margin-bottom: .8rem; line-height: 1.4; transition: color .3s; }
  .blog-card:hover .card-body h3 { color: var(--color-accent); }
  .card-body > p { color: var(--color-gray); font-size: .9rem; line-height: 1.6; display: -webkit-box; -webkit-line-clamp: 3; -webkit-box-orient: vertical; overflow: hidden; }
  .card-tags { display: flex; gap: .5rem; margin-top: .8rem; }
  .card-footer { padding: 1rem 1.5rem; border-top: 1px solid var(--color-dark-4); display: flex; justify-content: space-between; align-items: center; }
  .card-author { display: flex; align-items: center; gap: .5rem; font-size: .85rem; color: var(--color-light-3); }
  .author-avatar-xs { width: 28px; height: 28px; border-radius: 50%; background: var(--color-secondary); display: flex; align-items: center; justify-content: center; font-weight: 800; font-size: .75rem; }
  .card-read-more { color: var(--color-accent); font-weight: 700; font-size: .85rem; }

  .skeleton-card { background: var(--color-dark-3); border-radius: 16px; overflow: hidden; }
  .skeleton-img { height: 200px; background: linear-gradient(90deg, var(--color-dark-4) 25%, var(--color-dark-3) 50%, var(--color-dark-4) 75%); background-size: 200% 100%; animation: shimmer 1.5s infinite; }
  .skeleton-body { padding: 1.5rem; display: flex; flex-direction: column; gap: .8rem; }
  .skeleton-line { height: 14px; background: linear-gradient(90deg, var(--color-dark-4) 25%, var(--color-dark-3) 50%, var(--color-dark-4) 75%); background-size: 200% 100%; border-radius: 4px; animation: shimmer 1.5s infinite; }
  .skeleton-line.short { width: 40%; } .skeleton-line.medium { width: 70%; }
  @keyframes shimmer { 0%{background-position:-200% 0} 100%{background-position:200% 0} }

  .empty-state { text-align: center; padding: 4rem 2rem; }
  .empty-icon { font-size: 4rem; display: block; margin-bottom: 1rem; }
  .empty-state h3 { margin-bottom: .5rem; }
  .empty-state p { color: var(--color-gray); margin-bottom: 1.5rem; }

  .cta-newsletter { position: relative; padding: 5rem 0; background: url('/images/hero-bg.jpg') center/cover no-repeat fixed; background-color: var(--color-primary); }
  .cta-newsletter-overlay { position: absolute; inset: 0; background: linear-gradient(135deg, rgba(0,75,188,.92), rgba(22,17,47,.95)); }
  .newsletter-icon { font-size: 3rem; display: block; margin-bottom: 1rem; }
  .cta-newsletter h2 { font-size: clamp(1.8rem,4vw,2.5rem); margin-bottom: 1rem; }
  .newsletter-desc { color: rgba(255,255,255,.7); font-size: 1.05rem; margin-bottom: 2rem; max-width: 500px; margin-left: auto; margin-right: auto; }
  .newsletter-buttons { display: flex; gap: 1rem; justify-content: center; flex-wrap: wrap; }

  .cta-buttons { display: flex; gap: 1rem; justify-content: center; flex-wrap: wrap; margin-top: 1rem; }
  .btn-lg { padding: 1rem 2.5rem; font-size: 1.05rem; }

  @media (max-width: 1024px) { .blog-grid { grid-template-columns: repeat(2,1fr); } }
  @media (max-width: 768px) {
    .page-hero { padding: 8rem 1.5rem 3rem; }
    .featured-link { grid-template-columns: 1fr; }
    .featured-image { min-height: 220px; }
    .blog-grid { grid-template-columns: 1fr; }
    .filters-bar { flex-direction: column; gap: 1rem; }
    .search-box { min-width: 100%; }
    .category-filters { justify-content: center; }
    .filters-section { position: relative; top: 0; }
  }
  @media (max-width: 480px) {
    .newsletter-buttons, .cta-buttons { flex-direction: column; align-items: center; }
    .newsletter-buttons .btn, .cta-buttons .btn { width: 100%; justify-content: center; }
  }
</style>