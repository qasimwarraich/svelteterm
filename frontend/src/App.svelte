<script lang="ts">
    import Term from './components/Terminal.svelte';
</script>

<main>
    <h1>Chistole: CLI Tutor</h1>
    <body>
        <Term />
    </body>
</main>

<style>
    @font-face {
        font-family: 'cozettevectorregular';
        src: url('./assets/fonts/cozette/cozettevector-webfont.woff2') format('woff2'),
            url('./assets/fonts/cozette/cozettevector-webfont.woff2') format('woff');
        font-weight: normal;
        font-style: normal;
    }
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

    p {
        max-width: 14rem;
        margin: 1rem auto;
        line-height: 1.35;
    }

    @media (min-width: 480px) {
        h1 {
            max-width: none;
        }

        p {
            max-width: none;
        }
    }
</style>
