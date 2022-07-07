<script lang="ts">
    import { SpawnInstance, StopContainer } from './api';


</script>

<main>
    <button on:click={SpawnInstance}> spawn a new instance </button>
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
</style>
