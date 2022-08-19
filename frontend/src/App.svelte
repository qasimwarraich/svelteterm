<script lang="ts">
    import { fade } from 'svelte/transition';
    import InfoModal from './components/InfoModal.svelte';
    import Term from './components/Terminal.svelte';
    import TerminalSpawnButton from './components/TerminalSpawnButton.svelte';
    import type { Container } from './types/types';
    import './assets/fonts.css';
    let activesession: boolean;
    let container: Container;
    let modal: InfoModal;

    function beforeUnload() {
        // Chrome requires returnValue to be set.
        if (activesession) {
            event.preventDefault();
            event.returnValue = '';
            // more compatibility
            return '';
        }
        return null;
    }
</script>

<svelte:window on:beforeunload={beforeUnload} />
<main>
    <h1>Chistole: CLI </h1>
    <body>
        <InfoModal bind:this={modal}>
            <h2>Welcome to the CLI Tutor Web Interface</h2>
            <p>
                This web interface allows you to use a terminal emulator in a sandboxed environment
                so you cannot hurt your own system. In order to initiate a new session press the
                "spawn a new instance" button. This will spawn a web terminal and connect it to a
                docker container where you will have access to a standard Linux terminal.

                <br /><br />

                When you want to exit the tool you may use the "kill this instance" button to kill
                the terminal at which point you are free to start a new session.

                <br /><br />

                Have fun!
            </p>
            <button class="infobutton" on:click={() => modal.hide()}>Close</button>
        </InfoModal>

        <div class="nav">
            <TerminalSpawnButton bind:activesession bind:container />
            <button class="infobutton" on:click={() => modal.show()}>Info</button>
        </div>
        {#if activesession}
            <div transition:fade>
                <Term containerid={container.ID} />
            </div>
        {/if}
    </body>
</main>
<footer>chistole!</footer>

<style>
    :root {
        background-color: #5c2c45;
        background-color: black;
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

    code {
        background-color: #5c7287;
        color: white;
    }

    footer {
        color: white;
        width: 100vw;
        bottom: 0;
        left: 0;
        padding: 2px;
        text-align: center;
        position: fixed;
    }

    @media (min-width: 480px) {
        h1 {
            max-width: none;
        }
    }

    button:hover {
        color: #fff;
    }

    .nav {
        display: flex;
        align-items: center;
        justify-content: center;
    }

    .infobutton {
        background-color: #ffb852;
        font: inherit;
        margin: 1em;
        padding: 0.5em;
        border: none;
    }
</style>
