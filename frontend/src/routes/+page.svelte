<!-- frontend/src/routes/+page.svelte - Página principal -->
<script>
    let mensaje = "Cargando...";

    async function conectarBackend() {
        try {
            const res = await fetch("http://localhost:3000/api/saludo");
            const data = await res.json();
            mensaje = data.mensaje;
        } catch (error) {
            mensaje = "❌ Error conectando al backend";
        }
    }
</script>

<main>
    <h1>🔥 SvelteKit + Go</h1>
    <button on:click={conectarBackend}>
        Probar conexión
    </button>
    <p>{mensaje}</p>
</main>

<style>
    main {
        text-align: center;
        padding: 2rem;
        font-family: system-ui;
    }
    button {
        padding: 0.8rem 1.5rem;
        font-size: 1.1rem;
        background: #ff3e00;
        color: white;
        border: none;
        border-radius: 8px;
        cursor: pointer;
    }
    button:hover {
        background: #d63500;
    }
</style>