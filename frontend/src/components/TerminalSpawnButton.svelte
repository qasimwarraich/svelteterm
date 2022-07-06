<script lang="ts">
    async function SpawnInstance() {
        const response = await fetch('http://0.0.0.0:8080/api/createcontainer');
        const container = await response.json();
        StartContainer(container);
    }

    async function StartContainer(container) {
        let options = {
            method: 'Post',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(container),
        };
        fetch('http://0.0.0.0:8080/api/startcontainer', options);
    }

    async function StopContainer(container) {
        let options = {
            method: 'Post',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(container),
        };
        fetch('http://0.0.0.0:8080/api/stopcontainer', options);
    }
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
