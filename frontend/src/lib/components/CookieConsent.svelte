<!-- frontend/src/lib/components/CookieConsent.svelte -->
<script>
  import { onMount } from 'svelte';

  let visible = $state(false);
  let showDetails = $state(false);

  onMount(() => {
    try {
      const consent = localStorage.getItem('cookie-consent');
      if (!consent) {
        setTimeout(() => { visible = true; }, 2500);
      }
    } catch (e) {}
  });

  function acceptAll() { saveConsent('all'); }
  function acceptEssential() { saveConsent('essential'); }

  /**
	 * @param {string} type
	 */
  function saveConsent(type) {
    try {
      localStorage.setItem('cookie-consent', type);
      localStorage.setItem('cookie-consent-date', new Date().toISOString());
    } catch (e) {}
    visible = false;
  }
</script>

{#if visible}
  <div class="cookie-overlay">
    <div class="cookie-banner" class:expanded={showDetails}>
      <div class="cookie-icon">🍪</div>
      <div class="cookie-content">
        <h3>Este sitio utiliza cookies</h3>
        <p>
          Usamos cookies para mejorar tu experiencia de navegación,
          analizar el tráfico y personalizar contenido.
          Al continuar, aceptas nuestra
          <a href="/privacidad">Política de Privacidad</a>.
        </p>
        {#if showDetails}
          <div class="cookie-details">
            <div class="cookie-type">
              <div class="cookie-type-header">
                <span class="cookie-check essential">✓</span>
                <strong>Cookies esenciales</strong>
                <span class="cookie-required">Siempre activas</span>
              </div>
              <p>Necesarias para el funcionamiento del sitio.</p>
            </div>
            <div class="cookie-type">
              <div class="cookie-type-header">
                <span class="cookie-check">✓</span>
                <strong>Cookies analíticas</strong>
              </div>
              <p>Nos ayudan a entender cómo usas el sitio.</p>
            </div>
            <div class="cookie-type">
              <div class="cookie-type-header">
                <span class="cookie-check">✓</span>
                <strong>Cookies de marketing</strong>
              </div>
              <p>Permiten mostrarte contenido relevante.</p>
            </div>
          </div>
        {/if}
      </div>
      <div class="cookie-actions">
        <button class="btn btn-accent cookie-btn" onclick={acceptAll}>Aceptar todas</button>
        <button class="btn btn-outline-accent cookie-btn" onclick={acceptEssential}>Solo esenciales</button>
        <button class="cookie-details-btn" onclick={() => showDetails = !showDetails}>
          {showDetails ? 'Ocultar detalles' : 'Ver detalles'}
        </button>
      </div>
    </div>
  </div>
{/if}

<style>
  .cookie-overlay { position: fixed; bottom: 0; left: 0; right: 0; z-index: 9999; padding: 1.5rem; animation: slideUp 0.5s ease; }
  @keyframes slideUp { from{transform:translateY(100%);opacity:0} to{transform:translateY(0);opacity:1} }
  .cookie-banner { max-width: 900px; margin: 0 auto; background: var(--color-primary); border: 1px solid var(--color-dark-4); border-radius: 20px; padding: 2rem; display: flex; align-items: flex-start; gap: 1.5rem; box-shadow: 0 -10px 40px rgba(0,0,0,0.5); }
  .cookie-banner.expanded { flex-direction: column; }
  .cookie-icon { font-size: 2.5rem; flex-shrink: 0; }
  .cookie-content { flex: 1; }
  .cookie-content h3 { font-family: var(--font-heading); font-size: 1.1rem; margin-bottom: 0.5rem; }
  .cookie-content > p { color: var(--color-light-3); font-size: 0.9rem; line-height: 1.6; }
  .cookie-content a { color: var(--color-accent); font-weight: 600; text-decoration: underline; }
  .cookie-content a:hover { color: var(--color-light); }
  .cookie-details { margin-top: 1.5rem; display: flex; flex-direction: column; gap: 1rem; animation: fadeIn 0.3s ease; }
  @keyframes fadeIn { from{opacity:0} to{opacity:1} }
  .cookie-type { background: rgba(255,255,255,0.05); border: 1px solid var(--color-dark-4); border-radius: 12px; padding: 1rem 1.2rem; }
  .cookie-type-header { display: flex; align-items: center; gap: 0.6rem; margin-bottom: 0.4rem; }
  .cookie-check { width: 22px; height: 22px; border-radius: 6px; background: var(--color-secondary); display: flex; align-items: center; justify-content: center; font-size: 0.7rem; color: white; font-weight: 800; }
  .cookie-check.essential { background: var(--color-accent); color: var(--color-primary); }
  .cookie-required { font-size: 0.7rem; color: var(--color-accent); font-weight: 600; margin-left: auto; }
  .cookie-type p { color: var(--color-gray); font-size: 0.8rem; margin-left: 1.9rem; }
  .cookie-actions { display: flex; flex-direction: column; gap: 0.6rem; flex-shrink: 0; }
  .cookie-btn { padding: 0.7rem 1.5rem; font-size: 0.85rem; white-space: nowrap; }
  .cookie-details-btn { background: none; border: none; color: var(--color-gray); font-size: 0.8rem; cursor: pointer; text-align: center; transition: color 0.3s; font-family: var(--font-body); }
  .cookie-details-btn:hover { color: var(--color-accent); }

  @media (max-width: 768px) {
    .cookie-banner { flex-direction: column; padding: 1.5rem; }
    .cookie-actions { flex-direction: row; flex-wrap: wrap; width: 100%; }
    .cookie-btn { flex: 1; min-width: 120px; }
    .cookie-details-btn { width: 100%; }
  }
</style>