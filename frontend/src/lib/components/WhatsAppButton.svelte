<!-- frontend/src/lib/components/WhatsAppButton.svelte -->
<script>
  const phone = '524611203488';
  const message = encodeURIComponent(
    'Hola, me interesa obtener información sobre sus servicios de llantas.'
  );
  const whatsappUrl = `https://wa.me/${phone}?text=${message}`;
</script>

<a
  href={whatsappUrl}
  target="_blank"
  rel="noopener noreferrer"
  class="whatsapp-float"
  aria-label="Chatea por WhatsApp"
>
  <svg viewBox="0 0 32 32" fill="currentColor" width="32" height="32">
    <path d="M16.004 0C7.165 0 .003 7.161.003 16c0 2.823.737 5.576 2.137 8.003L.007 32l8.19-2.09A15.94 15.94 0 0016.004 32C24.843 32 32 24.839 32 16S24.843 0 16.004 0zm0 29.39a13.36 13.36 0 01-6.81-1.87l-.487-.29-5.058 1.327 1.35-4.932-.316-.503A13.33 13.33 0 012.62 16c0-7.387 6.01-13.39 13.384-13.39S29.38 8.613 29.38 16s-5.99 13.39-13.376 13.39zm7.336-10.022c-.402-.201-2.378-1.174-2.747-1.308-.369-.134-.637-.201-.906.201s-1.04 1.308-1.275 1.576-.469.302-.871.101c-.402-.201-1.697-.625-3.232-1.994-1.195-1.064-2.001-2.38-2.236-2.782-.235-.402-.025-.619.177-.82.181-.18.402-.469.603-.704.201-.235.268-.402.402-.67.134-.268.067-.503-.033-.704-.101-.201-.906-2.182-1.241-2.986-.327-.784-.659-.678-.906-.69-.235-.012-.503-.014-.772-.014s-.704.101-1.073.503c-.369.402-1.409 1.376-1.409 3.357s1.442 3.893 1.643 4.161c.201.268 2.838 4.332 6.879 6.073.961.415 1.711.663 2.296.849.965.307 1.843.263 2.537.16.774-.116 2.378-.972 2.714-1.91.335-.939.335-1.744.235-1.91-.101-.168-.369-.268-.772-.469z"/>
  </svg>
  <span class="whatsapp-tooltip">¡Chatea con nosotros!</span>
</a>

<style>
  .whatsapp-float {
    position: fixed;
    bottom: 2rem;
    right: 2rem;
    width: 60px;
    height: 60px;
    background: #25d366;
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    color: white;
    box-shadow: 0 4px 20px rgba(37, 211, 102, 0.4);
    z-index: 998;
    transition: all 0.3s ease;
    animation: pulse-whatsapp 2s infinite;
  }

  .whatsapp-float:hover {
    transform: scale(1.1);
    box-shadow: 0 6px 30px rgba(37, 211, 102, 0.6);
  }

  .whatsapp-tooltip {
    position: absolute;
    right: 75px;
    background: var(--color-dark-3);
    color: white;
    padding: 0.5rem 1rem;
    border-radius: 8px;
    font-size: 0.85rem;
    white-space: nowrap;
    opacity: 0;
    pointer-events: none;
    transition: opacity 0.3s ease;
  }

  .whatsapp-float:hover .whatsapp-tooltip {
    opacity: 1;
  }

  @keyframes pulse-whatsapp {
    0%, 100% { box-shadow: 0 4px 20px rgba(37, 211, 102, 0.4); }
    50% { box-shadow: 0 4px 30px rgba(37, 211, 102, 0.7); }
  }
</style>