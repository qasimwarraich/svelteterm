<script lang="ts">
    import '../assets/xterm.css';
    import './TerminalSpawnButton.svelte';
    import { ResizeContainerPTY } from './api';
    import { AttachAddon } from 'xterm-addon-attach';
    import { onMount } from 'svelte';

    export let containerid;

    onMount(async () => {
        const { Terminal } = await import('xterm');
        const term = new Terminal({
            fontSize: 12,
            rows: 100,
            cols: 180,
        });
        const socket = new WebSocket(
            `ws://0.0.0.0:6969/containers/${containerid}/attach/ws?logs=0&stream=1&stdin=1&stdout=1&stderr=1`
        );

        const attachAddon = new AttachAddon(socket);
        term.loadAddon(attachAddon);
        term.open(document.getElementById('terminal-parent'));
        let resize = { id: containerid, height: term.rows, width: term.cols };
        ResizeContainerPTY(resize);

        //Sets focus to terminal text input
        const textarea = document.getElementsByClassName('xterm-helper-textarea').item(0);
        (textarea as HTMLElement).focus();
    });
</script>

<main>
    <div class="flexparent">
        <div class="termparent" id="terminal-parent">
            <div id="terminal" />
        </div>
    </div>
</main>

<style>
    .flexparent {
        display: flex;
        justify-content: center;
        align-items: center;
    }
</style>
