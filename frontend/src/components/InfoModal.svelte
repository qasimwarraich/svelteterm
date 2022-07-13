<script>
    let shown = false;
    export function show() {
        shown = true;
    }
    export function hide() {
        shown = false;
    }
</script>

<svelte:window
    on:keydown={(e) => {
        if (e.key == 'Escape') {
            hide();
        }
    }}
/>

{#if shown}
    <div class="modal-wrapper">
        <div class="modal">
            <span class="close" on:click={() => hide()}>&times;</span>
            <slot />
        </div>
    </div>
{/if}

<style>
    .modal-wrapper {
        background-color: rgb(0, 0, 0);
        background-color: rgba(0, 0, 0, 0.6);
        position: fixed;
        width: 100%;
        height: 100%;
        top: 0;
        left: 0;
        z-index: 10;
    }

    .modal {
        background-color: rgba(0, 0, 0, 0.5);
        color: #ebdbb2;
        border: 1px;
        border-style: solid;
        border-color: #ffb852;
        max-width: 80vw;
        padding: 1rem;
        margin: 15% auto;
    }
    .close {
        float: right;
        cursor: pointer;
        font-size: 2rem;
    }
    .close:hover {
        font-weight: bold;
    }
</style>
