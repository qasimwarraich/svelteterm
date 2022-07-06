<script lang="ts">
    import '../assets/xterm.css';
    import './TerminalSpawnButton.svelte';
    import { AttachAddon } from 'xterm-addon-attach';
    import { onMount } from 'svelte';
    import TerminalSpawnButton from './TerminalSpawnButton.svelte';

    onMount(async () => {
        const { Terminal } = await import('xterm');
        const term = new Terminal({
            fontSize: 10,
            rows: 68,
            cols: 270,
        });
        const socket = new WebSocket(
            'wss://127.0.0.1:2376/containers/spamcontainer/attach/ws?logs=0&stream=1&stdin=1&stdout=1&stderr=1'
        );

        const attachAddon = new AttachAddon(socket);
        term.loadAddon(attachAddon);
        term.open(document.getElementById('terminal-parent'));

        const textarea = document.getElementsByClassName('xterm-helper-textarea').item(0);
        (textarea as HTMLElement).focus();
    });
</script>

<main>
    <div class="flexparent">
        <div class="termparent" id="terminal-parent">
            <TerminalSpawnButton />
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
