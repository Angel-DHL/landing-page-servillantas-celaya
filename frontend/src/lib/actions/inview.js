// frontend/src/lib/actions/inview.js

// Esta "acción" de Svelte detecta cuando un elemento
// entra en el viewport y le agrega la clase "visible"
/**
 * @param {Element} node
 */
export function inview(node, params = {}) {
    // @ts-ignore
    const threshold = params.threshold || 0.15;
  
    const observer = new IntersectionObserver(
      (entries) => {
        entries.forEach((entry) => {
          if (entry.isIntersecting) {
            node.classList.add('visible');
            observer.unobserve(node); // Solo anima una vez
          }
        });
      },
      { threshold }
    );
  
    observer.observe(node);
  
    return {
      destroy() {
        observer.disconnect();
      }
    };
  }