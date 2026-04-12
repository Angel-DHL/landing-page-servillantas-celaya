<!-- frontend/src/lib/components/Preloader.svelte -->
<script>
  import { onMount } from 'svelte';

  let visible = $state(true);
  let fadeOut = $state(false);

  onMount(() => {
    const timer = setTimeout(() => {
      fadeOut = true;
    }, 1500);

    if (document.readyState === 'complete') {
      fadeOut = true;
    } else {
      window.addEventListener('load', () => {
        fadeOut = true;
      });
    }

    return () => clearTimeout(timer);
  });

  function handleEnd() {
    if (fadeOut) {
      visible = false;
    }
  }
</script>

{#if visible}
  <div class="preloader" class:fade-out={fadeOut} ontransitionend={handleEnd}>
    <div class="preloader-content">
      <div class="preloader-wheel">
        <svg viewBox="0 0 100 100" width="80" height="80">
          <circle cx="50" cy="50" r="40" fill="none" stroke="#2d2a45" stroke-width="6"/>
          <circle cx="50" cy="50" r="40" fill="none" stroke="#f3fd01" stroke-width="6"
            stroke-dasharray="80 170" stroke-linecap="round" class="spinner-circle"/>
          <circle cx="50" cy="50" r="15" fill="#f3fd01" opacity="0.15"/>
          <circle cx="50" cy="50" r="5" fill="#f3fd01"/>
        </svg>
      </div>
      <span class="preloader-text">Servi<span class="preloader-highlight">Llantas</span></span>
      <div class="preloader-bar"><div class="preloader-bar-fill"></div></div>
    </div>
  </div>
{/if}

<style>
  .preloader { position: fixed; inset: 0; background: #16112f; display: flex; align-items: center; justify-content: center; z-index: 99999; opacity: 1; transition: opacity 0.5s ease; }
  .preloader.fade-out { opacity: 0; pointer-events: none; }
  .preloader-content { text-align: center; display: flex; flex-direction: column; align-items: center; gap: 1.2rem; }
  .preloader-wheel { animation: pulse 1.5s ease infinite alternate; }
  .spinner-circle { animation: spinCircle 1.2s linear infinite; transform-origin: center; }
  .preloader-text { font-family: 'Montserrat', sans-serif; font-weight: 900; font-size: 1.8rem; color: white; }
  .preloader-highlight { color: #f3fd01; }
  .preloader-bar { width: 180px; height: 3px; background: rgba(255,255,255,0.1); border-radius: 3px; overflow: hidden; }
  .preloader-bar-fill { width: 100%; height: 100%; background: linear-gradient(90deg, #004bbc, #f3fd01); border-radius: 3px; animation: loading 1.3s ease forwards; }
  @keyframes loading { from{width:0} to{width:100%} }
  @keyframes pulse { from{transform:scale(1)} to{transform:scale(1.08)} }
  @keyframes spinCircle { to{transform:rotate(360deg)} }
</style>