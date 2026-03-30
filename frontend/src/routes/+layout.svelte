<!-- frontend/src/routes/+layout.svelte -->
<script>
  import '../app.css';
  import Navbar from '$lib/components/Navbar.svelte';
  import Footer from '$lib/components/Footer.svelte';
  import WhatsAppButton from '$lib/components/WhatsAppButton.svelte';
</script>

<Navbar />

<main>
  <slot />
</main>

<Footer />
<WhatsAppButton />

<style>
  main {
    min-height: 100vh;
  }
</style>