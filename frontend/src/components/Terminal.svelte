<script lang="ts">
    import { onMount } from 'svelte';
    import './TerminalSpawnButton.svelte';
    import { AttachAddon } from 'xterm-addon-attach';
    import { ResizeContainerPTY } from './api';
    import type { ResizeRequest } from '../types/types';
    import '../assets/xterm.css';

    export let containerid: string;

    onMount(async () => {
        const { Terminal } = await import('xterm');
        const term = new Terminal({
            fontSize: 14,
            fontFamily: 'monospace',
            rows: 40,
            cols: 160,
        });

        const socket = new WebSocket(
            import.meta.env.VITE_DOCKERENDPOINT + containerid + import.meta.env.VITE_WEBSOCKETATTACH
        );

        const attachAddon = new AttachAddon(socket, {bidirectional: true} );
        term.loadAddon(attachAddon);
        term.open(document.getElementById('terminal-parent'));
        let resize: ResizeRequest = { ID: containerid, Height: term.rows, Width: term.cols };
        ResizeContainerPTY(resize);

        //Sets focus to terminal text input
        const textarea: Element = document.getElementsByClassName('xterm-helper-textarea').item(0);
        (textarea as HTMLElement).focus();
        const terminal: Element = document.querySelector(".terminal")
        terminal.scrollIntoView({ behavior: 'smooth', block: 'center' })
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
