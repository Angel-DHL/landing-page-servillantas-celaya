<!-- frontend/src/lib/components/PlaceholderImage.svelte -->
<!-- REEMPLAZA TODO EL ARCHIVO -->
<script>
    let { type = 'default', height = '400px', text = '' } = $props();
  
    const configs = {
      hero: {
        icon: `<svg viewBox="0 0 120 120" width="80" height="80" fill="none" xmlns="http://www.w3.org/2000/svg">
          <circle cx="60" cy="60" r="50" stroke="rgba(255,255,255,0.15)" stroke-width="4"/>
          <circle cx="60" cy="60" r="35" stroke="rgba(255,255,255,0.1)" stroke-width="3"/>
          <circle cx="60" cy="60" r="20" stroke="rgba(243,253,1,0.3)" stroke-width="3"/>
          <circle cx="60" cy="60" r="8" fill="rgba(243,253,1,0.4)"/>
          <line x1="60" y1="10" x2="60" y2="25" stroke="rgba(255,255,255,0.2)" stroke-width="3" stroke-linecap="round"/>
          <line x1="60" y1="95" x2="60" y2="110" stroke="rgba(255,255,255,0.2)" stroke-width="3" stroke-linecap="round"/>
          <line x1="10" y1="60" x2="25" y2="60" stroke="rgba(255,255,255,0.2)" stroke-width="3" stroke-linecap="round"/>
          <line x1="95" y1="60" x2="110" y2="60" stroke="rgba(255,255,255,0.2)" stroke-width="3" stroke-linecap="round"/>
        </svg>`,
        gradient: 'linear-gradient(135deg, #16112f 0%, #1a1740 30%, #004bbc 70%, #16112f 100%)'
      },
      taller: {
        icon: `<svg viewBox="0 0 100 100" width="70" height="70" fill="none" xmlns="http://www.w3.org/2000/svg">
          <rect x="10" y="40" width="80" height="45" rx="5" stroke="rgba(255,255,255,0.2)" stroke-width="3"/>
          <rect x="5" y="35" width="90" height="10" rx="3" stroke="rgba(243,253,1,0.3)" stroke-width="2"/>
          <circle cx="30" cy="85" r="8" stroke="rgba(255,255,255,0.2)" stroke-width="3"/>
          <circle cx="70" cy="85" r="8" stroke="rgba(255,255,255,0.2)" stroke-width="3"/>
          <rect x="25" y="50" width="20" height="15" rx="2" stroke="rgba(243,253,1,0.2)" stroke-width="2"/>
          <rect x="55" y="50" width="20" height="15" rx="2" stroke="rgba(243,253,1,0.2)" stroke-width="2"/>
        </svg>`,
        gradient: 'linear-gradient(135deg, #0c0a18 0%, #16112f 40%, #1a3a6b 100%)'
      },
      servicio: {
        icon: `<svg viewBox="0 0 100 100" width="60" height="60" fill="none" xmlns="http://www.w3.org/2000/svg">
          <circle cx="50" cy="50" r="40" stroke="rgba(243,253,1,0.2)" stroke-width="3"/>
          <circle cx="50" cy="50" r="25" stroke="rgba(0,75,188,0.4)" stroke-width="3"/>
          <circle cx="50" cy="50" r="10" fill="rgba(0,75,188,0.3)"/>
          <circle cx="50" cy="50" r="4" fill="rgba(243,253,1,0.5)"/>
        </svg>`,
        gradient: 'linear-gradient(135deg, #16112f 0%, #002266 50%, #16112f 100%)'
      },
      blog: {
        icon: `<svg viewBox="0 0 80 80" width="50" height="50" fill="none" xmlns="http://www.w3.org/2000/svg">
          <rect x="10" y="8" width="60" height="64" rx="5" stroke="rgba(255,255,255,0.15)" stroke-width="3"/>
          <line x1="22" y1="25" x2="58" y2="25" stroke="rgba(243,253,1,0.3)" stroke-width="3" stroke-linecap="round"/>
          <line x1="22" y1="37" x2="52" y2="37" stroke="rgba(255,255,255,0.15)" stroke-width="2" stroke-linecap="round"/>
          <line x1="22" y1="47" x2="55" y2="47" stroke="rgba(255,255,255,0.1)" stroke-width="2" stroke-linecap="round"/>
          <line x1="22" y1="57" x2="45" y2="57" stroke="rgba(255,255,255,0.1)" stroke-width="2" stroke-linecap="round"/>
        </svg>`,
        gradient: 'linear-gradient(135deg, #16112f 0%, #1a1740 50%, #003388 100%)'
      },
      galeria: {
        icon: `<svg viewBox="0 0 80 80" width="50" height="50" fill="none" xmlns="http://www.w3.org/2000/svg">
          <rect x="5" y="12" width="70" height="50" rx="5" stroke="rgba(255,255,255,0.15)" stroke-width="3"/>
          <circle cx="25" cy="30" r="7" stroke="rgba(243,253,1,0.3)" stroke-width="2"/>
          <polyline points="10,55 28,38 45,50 55,42 70,55" stroke="rgba(0,75,188,0.4)" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
        </svg>`,
        gradient: 'linear-gradient(135deg, #0c0a18 0%, #16112f 50%, #002255 100%)'
      }
    };
  
    // @ts-ignore
    let config = $derived(configs[type] || configs['servicio']);
  </script>
  
  <div class="placeholder" style="height: {height}; background: {config.gradient};">
    <div class="placeholder-pattern"></div>
    <div class="placeholder-content">
      <div class="placeholder-icon">
        {@html config.icon}
      </div>
      {#if text}
        <p class="placeholder-text">{text}</p>
      {/if}
    </div>
  </div>
  
  <style>
    .placeholder {
      position: relative;
      border-radius: 16px;
      overflow: hidden;
      display: flex;
      align-items: center;
      justify-content: center;
    }
  
    .placeholder-pattern {
      position: absolute;
      inset: 0;
      background-image:
        radial-gradient(circle at 20% 50%, rgba(0,75,188,0.15) 0%, transparent 50%),
        radial-gradient(circle at 80% 20%, rgba(243,253,1,0.08) 0%, transparent 40%),
        radial-gradient(circle at 60% 80%, rgba(0,75,188,0.1) 0%, transparent 40%);
    }
  
    .placeholder-pattern::after {
      content: '';
      position: absolute;
      inset: 0;
      background-image: radial-gradient(rgba(255,255,255,0.03) 1px, transparent 1px);
      background-size: 20px 20px;
    }
  
    .placeholder-content {
      position: relative;
      z-index: 2;
      text-align: center;
      display: flex;
      flex-direction: column;
      align-items: center;
      gap: 1rem;
    }
  
    .placeholder-icon {
      opacity: 0.8;
      animation: float 4s ease-in-out infinite;
    }
  
    .placeholder-text {
      color: rgba(255, 255, 255, 0.3);
      font-size: 0.85rem;
      font-weight: 500;
      letter-spacing: 1px;
    }
  
    @keyframes float {
      0%, 100% { transform: translateY(0); }
      50% { transform: translateY(-8px); }
    }
  </style>