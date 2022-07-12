<script lang="ts">
    import { SpawnInstance, StopContainer } from './api';
    import type { Container } from '../types/types';

    export let activesession: boolean = false;

    export let container: Container = { ID: 'Default' };

    const createContainer = async () => {
        container = await SpawnInstance();
        activesession = true;
    };

    const killContainer = async () => {
        await StopContainer(container);
        activesession = false;
    };
</script>

<main>
    {#if !activesession}
        <button on:click={createContainer}> spawn a new instance </button>
    {:else}
        <button class="kill" on:click={killContainer}> kill this instance </button>
    {/if}
</main>

<style>
    button {
        background-color: #e6adcf;
        font: inherit;
        margin: 1em;
        padding: 0.5em;
        border: none;
    }

    button:hover {
        color: #fff;
    }

    .kill {
        color: #000000;
        background-color: #ff0000;
    }
</style>
