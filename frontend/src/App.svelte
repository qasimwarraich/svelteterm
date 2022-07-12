<script lang="ts">
    import Term from './components/Terminal.svelte';
    import TerminalSpawnButton from './components/TerminalSpawnButton.svelte';
    import type { Container } from './types/types';
    import './assets/fonts.css';
    let activesession: boolean;
    let container: Container;
</script>

<main>
    <h1>Chistole: CLI Tutor</h1>
    <body>
        <TerminalSpawnButton bind:activesession bind:container />
        {#if activesession}
            <Term containerid={container.ID} />
        {/if}
    </body>
</main>

<style>
    :root {
        background-color: #5c2c45;
        font-family: cozettevectorregular, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto,
            Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
    }

    main {
        text-align: center;
        padding: 1em;
        margin: 0 auto;
    }

    h1 {
        color: #e6adcf;
        text-transform: uppercase;
        font-size: 2rem;
        font-weight: 300;
        line-height: 1.1;
        margin: 2rem auto;
        max-width: 14rem;
    }

    @media (min-width: 480px) {
        h1 {
            max-width: none;
        }
    }
</style>
