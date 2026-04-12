<!-- frontend/src/lib/components/ScrollProgress.svelte -->
<script>
  import { onMount } from 'svelte';

  let progress = $state(0);

  onMount(() => {
    const handleScroll = () => {
      const scrollTop = window.scrollY;
      const docHeight = document.documentElement.scrollHeight - window.innerHeight;
      progress = docHeight > 0 ? (scrollTop / docHeight) * 100 : 0;
    };
    window.addEventListener('scroll', handleScroll, { passive: true });
    return () => window.removeEventListener('scroll', handleScroll);
  });
</script>

<div class="scroll-progress">
  <div class="scroll-progress-bar" style="width: {progress}%"></div>
</div>

<style>
  .scroll-progress { position: fixed; top: 0; left: 0; width: 100%; height: 3px; background: transparent; z-index: 10001; }
  .scroll-progress-bar { height: 100%; background: linear-gradient(90deg, var(--color-secondary), var(--color-accent)); transition: width 0.1s linear; border-radius: 0 2px 2px 0; box-shadow: 0 0 10px rgba(243,253,1,0.5); }
</style>