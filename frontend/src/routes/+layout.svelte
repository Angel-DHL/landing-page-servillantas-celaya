<!-- frontend/src/routes/+layout.svelte -->
<!-- REEMPLAZA TODO EL ARCHIVO -->
<script>
  import '../app.css';
  import Navbar from '$lib/components/Navbar.svelte';
  import Footer from '$lib/components/Footer.svelte';
  import WhatsAppButton from '$lib/components/WhatsAppButton.svelte';
  import ScrollProgress from '$lib/components/ScrollProgress.svelte';
  import ScrollToTop from '$lib/components/ScrollToTop.svelte';
  import CookieConsent from '$lib/components/CookieConsent.svelte';
  import Preloader from '$lib/components/Preloader.svelte';
</script>

<Preloader />
<ScrollProgress />
<Navbar />

<main>
  <slot />
</main>

<Footer />
<WhatsAppButton />
<ScrollToTop />
<CookieConsent />

<style>
  main {
    min-height: 100vh;
  }
</style>