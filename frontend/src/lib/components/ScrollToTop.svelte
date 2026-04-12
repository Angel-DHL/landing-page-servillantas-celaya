<!-- frontend/src/lib/components/ScrollToTop.svelte -->
<script>
  import { onMount } from 'svelte';

  let visible = $state(false);

  onMount(() => {
    const handleScroll = () => {
      visible = window.scrollY > 500;
    };
    window.addEventListener('scroll', handleScroll, { passive: true });
    return () => window.removeEventListener('scroll', handleScroll);
  });

  function scrollToTop() {
    window.scrollTo({ top: 0, behavior: 'smooth' });
  }
</script>

{#if visible}
  <button class="scroll-top" onclick={scrollToTop} aria-label="Volver arriba">
    <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round">
      <path d="M12 19V5M5 12l7-7 7 7" />
    </svg>
  </button>
{/if}

<style>
  .scroll-top {
    position: fixed; bottom: 2rem; left: 2rem; width: 48px; height: 48px;
    background: var(--color-secondary); color: white; border: none;
    border-radius: 50%; display: flex; align-items: center; justify-content: center;
    cursor: pointer; z-index: 997; transition: all 0.3s ease;
    box-shadow: 0 4px 15px rgba(0,75,188,0.4); animation: fadeInBtn 0.3s ease;
  }
  .scroll-top:hover {
    background: var(--color-accent); color: var(--color-primary);
    transform: translateY(-3px); box-shadow: 0 6px 25px rgba(243,253,1,0.3);
  }
  @keyframes fadeInBtn { from{opacity:0;transform:translateY(20px)} to{opacity:1;transform:translateY(0)} }
</style>