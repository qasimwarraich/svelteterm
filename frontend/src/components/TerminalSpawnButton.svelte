<script lang="ts">

    function SpawnInstance() {
        alert('Spam');
    }

</script>

<main>
    <button on:click={SpawnInstance}> Spawn a new instance </button>
</main>

<style>
    button {
        background-color: #e6adcf;
        margin: 1em;
        padding: 0.5em;
        border: none;
    }

    button:hover {
        color: #fff;
    }
</style>
